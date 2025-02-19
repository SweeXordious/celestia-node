package p2p

import (
	"context"
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	routinghelpers "github.com/libp2p/go-libp2p-routing-helpers"
	"go.uber.org/fx"

	"github.com/celestiaorg/celestia-node/node/fxutil"
)

// ContentRouting constructs nil content routing,
// as for our use-case existing ContentRouting mechanisms, e.g DHT, are unsuitable
func ContentRouting() routing.ContentRouting {
	return &routinghelpers.Null{}
}

// PeerRouting provides constructor for PeerRouting over DHT.
// Basically, this provides a way to discover peer addresses by respecting public keys.
func PeerRouting(cfg Config) func(routingParams) (routing.PeerRouting, error) {
	return func(params routingParams) (routing.PeerRouting, error) {
		bpeers, err := cfg.bootstrapPeers()
		if err != nil {
			return nil, err
		}

		prefix := fmt.Sprintf("/celestia/%s", cfg.Network)
		mode := dht.ModeAuto
		if cfg.Bootstrapper {
			mode = dht.ModeServer
		}

		d, err := dht.New(
			fxutil.WithLifecycle(params.Ctx, params.Lc),
			params.Host,
			dht.Mode(mode),
			dht.BootstrapPeers(bpeers...),
			dht.ProtocolPrefix(protocol.ID(prefix)),
			dht.Datastore(params.DataStore),
			dht.QueryFilter(dht.PublicQueryFilter),
			dht.RoutingTableFilter(dht.PublicRoutingTableFilter),
			// disable DHT for everything besides peer routing
			dht.DisableValues(),
			dht.DisableProviders(),
		)
		if err != nil {
			return nil, err
		}
		params.Lc.Append(fx.Hook{OnStop: func(context.Context) error {
			return d.Close()
		}})

		return d, nil
	}
}

type routingParams struct {
	fx.In

	Ctx       context.Context
	Lc        fx.Lifecycle
	Host      hostBase
	DataStore datastore.Batching
}
