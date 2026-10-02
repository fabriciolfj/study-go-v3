package httpsender

import (
	"context"

	"fabriciolfj.github/study/internal/delivery"
)

type Client struct {
}

func (c *Client) Send(ctx context.Context, d *delivery.Delivery) error {
	println("sending delivery")
	return nil
}
