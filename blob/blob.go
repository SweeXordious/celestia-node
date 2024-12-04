package blob

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tendermint/tendermint/crypto/merkle"
	coretypes "github.com/tendermint/tendermint/types"

	"github.com/celestiaorg/celestia-app/v2/pkg/appconsts"
	v2 "github.com/celestiaorg/celestia-app/v2/pkg/appconsts/v2"
	"github.com/celestiaorg/go-square/blob"
	"github.com/celestiaorg/go-square/inclusion"
	"github.com/celestiaorg/go-square/shares"
	"github.com/celestiaorg/nmt"

	"github.com/celestiaorg/celestia-node/share"
)

// appVersion is the current application version of celestia-app.
const appVersion = v2.Version

var errEmptyShares = errors.New("empty shares")

// Proof constructs the proof of a blob to the data root.
type Proof struct {
	// SubtreeRoots are the subtree roots of the blob's data that are
	// used to create the commitment.
	SubtreeRoots [][]byte `json:"subtree_roots"`
	// SubtreeRootProofs the proofs of the subtree roots to the row roots they belong to.
	// If the blob spans across multiple rows, then this will contain multiple proofs.
	SubtreeRootProofs []*nmt.Proof `json:"share_to_row_root_proofs"`
	// RowToDataRootProof the proofs of the row roots containing the blob shares
	// to the data root.
	RowToDataRootProof coretypes.RowProof `json:"row_to_data_root_proof"`
}

// namespaceToRowRootProof a proof of a set of namespace shares to the row
// roots they belong to.
type namespaceToRowRootProof []*nmt.Proof

// Commitment is a Merkle Root of the subtree built from shares of the Blob.
// It is computed by splitting the blob into shares and building the Merkle subtree to be included
// after Submit.
type Commitment []byte

// Verify takes a data root and verifies if the
// provided proof's subtree roots were committed to the given data root.
func (p *Proof) Verify(dataRoot []byte, subtreeRootThreshold int) (bool, error) {
	if len(dataRoot) == 0 {
		return false, errors.New("root must be non-empty")
	}

	if subtreeRootThreshold <= 0 {
		return false, errors.New("subtreeRootThreshold must be > 0")
	}

	// this check is < instead of != because we can have two subtree roots
	// at the same height, depending on the subtree root threshold,
	// and they can be used to create the above inner node without needing a proof inner node.
	if len(p.SubtreeRoots) < len(p.SubtreeRootProofs) {
		return false, fmt.Errorf(
			"the number of subtree roots %d should be bigger than the number of subtree root proofs %d",
			len(p.SubtreeRoots),
			len(p.SubtreeRootProofs),
		)
	}

	// for each row, one or more subtree roots' inclusion is verified against
	// their corresponding row root. then, these row roots' inclusion is verified
	// against the data root. so their number should be the same.
	if len(p.SubtreeRootProofs) != len(p.RowToDataRootProof.Proofs) {
		return false, fmt.Errorf(
			"the number of subtree root proofs %d should be equal to the number of row root proofs %d",
			len(p.SubtreeRootProofs),
			len(p.RowToDataRootProof.Proofs),
		)
	}

	// the row root proofs' ranges are defined as [startRow, endRow].
	if int(p.RowToDataRootProof.EndRow-p.RowToDataRootProof.StartRow+1) != len(p.RowToDataRootProof.RowRoots) {
		return false, fmt.Errorf(
			"the number of rows %d must equal the number of row roots %d",
			int(p.RowToDataRootProof.EndRow-p.RowToDataRootProof.StartRow+1),
			len(p.RowToDataRootProof.RowRoots),
		)
	}
	if len(p.RowToDataRootProof.Proofs) != len(p.RowToDataRootProof.RowRoots) {
		return false, fmt.Errorf(
			"the number of proofs %d must equal the number of row roots %d",
			len(p.RowToDataRootProof.Proofs),
			len(p.RowToDataRootProof.RowRoots),
		)
	}

	// verify the inclusion of the rows to the data root
	if err := p.RowToDataRootProof.Validate(dataRoot); err != nil {
		return false, err
	}

	// computes the total number of shares proven given that each subtree root
	// references a specific set of leaves.
	numberOfShares := 0
	for _, proof := range p.SubtreeRootProofs {
		numberOfShares += proof.End() - proof.Start()
	}

	// use the computed total number of shares to calculate the subtree roots
	// width.
	// the subtree roots width is defined in ADR-013:
	//
	//https://github.com/celestiaorg/celestia-app/blob/main/docs/architecture/adr-013-non-interactive-default-rules-for-zero-padding.md
	subtreeRootsWidth := inclusion.SubTreeWidth(numberOfShares, subtreeRootThreshold)

	nmtHasher := nmt.NewNmtHasher(appconsts.NewBaseHashFunc(), share.NamespaceSize, true)
	// verify the proof of the subtree roots
	subtreeRootsCursor := 0
	for i, subtreeRootProof := range p.SubtreeRootProofs {
		// calculate the share range that each subtree root commits to.
		ranges, err := nmt.ToLeafRanges(subtreeRootProof.Start(), subtreeRootProof.End(), subtreeRootsWidth)
		if err != nil {
			return false, err
		}

		if len(p.SubtreeRoots) < subtreeRootsCursor {
			return false, fmt.Errorf("len(commitmentProof.SubtreeRoots)=%d < subtreeRootsCursor=%d",
				len(p.SubtreeRoots), subtreeRootsCursor)
		}
		if len(p.SubtreeRoots) < subtreeRootsCursor+len(ranges) {
			return false, fmt.Errorf("len(commitmentProof.SubtreeRoots)=%d < subtreeRootsCursor+len(ranges)=%d",
				len(p.SubtreeRoots), subtreeRootsCursor+len(ranges))
		}
		valid, err := subtreeRootProof.VerifySubtreeRootInclusion(
			nmtHasher,
			p.SubtreeRoots[subtreeRootsCursor:subtreeRootsCursor+len(ranges)],
			subtreeRootsWidth,
			p.RowToDataRootProof.RowRoots[i],
		)
		if err != nil {
			return false, err
		}
		if !valid {
			return false,
				fmt.Errorf(
					"subtree root proof for range [%d, %d) is invalid",
					subtreeRootProof.Start(),
					subtreeRootProof.End(),
				)
		}
		subtreeRootsCursor += len(ranges)
	}

	return true, nil
}

// GenerateCommitment generates the share commitment corresponding
// to the proof's subtree roots
func (p *Proof) GenerateCommitment() Commitment {
	return merkle.HashFromByteSlices(p.SubtreeRoots)
}

func (com Commitment) String() string {
	return string(com)
}

// Equal ensures that commitments are the same
func (com Commitment) Equal(c Commitment) bool {
	return bytes.Equal(com, c)
}

// Blob represents any application-specific binary data that anyone can submit to Celestia.
type Blob struct {
	*blob.Blob `json:"blob"`

	Commitment Commitment `json:"commitment"`

	// the celestia-node's namespace type
	// this is to avoid converting to and from app's type
	namespace share.Namespace

	// index represents the index of the blob's first share in the EDS.
	// Only retrieved, on-chain blobs will have the index set. Default is -1.
	index int
}

// NewBlobV0 constructs a new blob from the provided Namespace and data.
// The blob will be formatted as v0 shares.
func NewBlobV0(namespace share.Namespace, data []byte) (*Blob, error) {
	return NewBlob(appconsts.ShareVersionZero, namespace, data)
}

// NewBlob constructs a new blob from the provided Namespace, data and share version.
func NewBlob(shareVersion uint8, namespace share.Namespace, data []byte) (*Blob, error) {
	if len(data) == 0 || len(data) > appconsts.DefaultMaxBytes {
		return nil, fmt.Errorf("blob data must be > 0 && <= %d, but it was %d bytes", appconsts.DefaultMaxBytes, len(data))
	}
	if err := namespace.ValidateForBlob(); err != nil {
		return nil, err
	}

	blob := blob.Blob{
		NamespaceId:      namespace.ID(),
		Data:             data,
		ShareVersion:     uint32(shareVersion),
		NamespaceVersion: uint32(namespace.Version()),
	}

	com, err := inclusion.CreateCommitment(&blob, merkle.HashFromByteSlices, appconsts.SubtreeRootThreshold(appVersion))
	if err != nil {
		return nil, err
	}
	return &Blob{Blob: &blob, Commitment: com, namespace: namespace, index: -1}, nil
}

// Namespace returns blob's namespace.
func (b *Blob) Namespace() share.Namespace {
	return b.namespace
}

// Index returns the blob's first share index in the EDS.
// Only retrieved, on-chain blobs will have the index set. Default is -1.
func (b *Blob) Index() int {
	return b.index
}

// Length returns the number of shares in the blob.
func (b *Blob) Length() (int, error) {
	s, err := BlobsToShares(b)
	if err != nil {
		return 0, err
	}

	if len(s) == 0 {
		return 0, errors.New("blob with zero shares received")
	}

	appShare, err := shares.NewShare(s[0])
	if err != nil {
		return 0, err
	}

	seqLength, err := appShare.SequenceLen()
	if err != nil {
		return 0, err
	}

	return shares.SparseSharesNeeded(seqLength), nil
}

func (b *Blob) compareCommitments(com Commitment) bool {
	return bytes.Equal(b.Commitment, com)
}

type jsonBlob struct {
	Namespace    share.Namespace `json:"namespace"`
	Data         []byte          `json:"data"`
	ShareVersion uint32          `json:"share_version"`
	Commitment   Commitment      `json:"commitment"`
	Index        int             `json:"index"`
}

func (b *Blob) MarshalJSON() ([]byte, error) {
	blob := &jsonBlob{
		Namespace:    b.Namespace(),
		Data:         b.Data,
		ShareVersion: b.ShareVersion,
		Commitment:   b.Commitment,
		Index:        b.index,
	}
	return json.Marshal(blob)
}

func (b *Blob) UnmarshalJSON(data []byte) error {
	var jsonBlob jsonBlob
	err := json.Unmarshal(data, &jsonBlob)
	if err != nil {
		return err
	}

	if len(jsonBlob.Namespace) == 0 {
		return errors.New("expected a non-empty namespace")
	}

	b.Blob = &blob.Blob{}
	b.Blob.NamespaceVersion = uint32(jsonBlob.Namespace.Version())
	b.Blob.NamespaceId = jsonBlob.Namespace.ID()
	b.Blob.Data = jsonBlob.Data
	b.Blob.ShareVersion = jsonBlob.ShareVersion
	b.Commitment = jsonBlob.Commitment
	b.namespace = jsonBlob.Namespace
	b.index = jsonBlob.Index
	return nil
}

func (b *Blob) ComputeSubtreeRoots() ([][]byte, error) {
	gsBlob := blob.Blob{
		NamespaceId:      b.NamespaceId,
		Data:             b.Data,
		ShareVersion:     b.ShareVersion,
		NamespaceVersion: b.NamespaceVersion,
	}
	return inclusion.GenerateSubtreeRoots(&gsBlob, appconsts.SubtreeRootThreshold(appVersion))
}

// proveRowRootsToDataRoot creates a set of binary merkle proofs for all the
// roots defined by the range [start, end).
func proveRowRootsToDataRoot(roots [][]byte, start, end int) []*merkle.Proof {
	_, proofs := merkle.ProofsFromByteSlices(roots)
	return proofs[start:end]
}
