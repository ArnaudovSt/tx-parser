// The streamProvider is designed as a future improvement to utilize a data stream,
// such as WebSocket, which is more efficient than polling for new blocks.
package streamprovider

import (
	"context"

	"github.com/ArnaudovSt/tx-parser/provider"
)

var _ provider.IProvider = (*streamProvider)(nil)

type streamProvider struct{}

func NewWebsocketProvider() *streamProvider {
	return &streamProvider{}
}

func (w *streamProvider) Start(_ context.Context) error {
	panic("unimplemented")
}
