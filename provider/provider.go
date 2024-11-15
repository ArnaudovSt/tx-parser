package provider

import "context"

type IProvider interface {
	Start(ctx context.Context) error
}
