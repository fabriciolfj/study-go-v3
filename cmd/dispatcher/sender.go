package main

import (
	"context"

	"fabriciolfj.github/study/internal/delivery"
)

type Sender interface {
	Send(ctx context.Context, d *delivery.Delivery) error
}
