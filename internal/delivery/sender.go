package delivery

import (
	"context"
)

type Sender interface {
	Send(ctx context.Context, d *Delivery) error
}
