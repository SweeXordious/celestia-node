package node

import (
	"encoding/json"
	"net/http"
)

type StatusMessage struct {
	ListenAddresses []string `json:"listenAddresses"`
	Network         string   `json:"network"`
}

func NewStatusMessage(listenAddrs []string, network string) *StatusMessage {
	return &StatusMessage{
		ListenAddresses: listenAddrs,
		Network:         network,
	}
}

func (s StatusMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonStatus, err := json.Marshal(s)
	if err != nil {
		log.Errorf("could not marshal the status: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, err = w.Write(jsonStatus)
		if err != nil {
			log.Errorf("could not write status response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
