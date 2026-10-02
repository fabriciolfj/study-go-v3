package dispatcher

import (
	"context"

	"fabriciolfj.github/study/internal/delivery"
)

type Dispatcher struct {
	send delivery.Sender
}

func New(s delivery.Sender) *Dispatcher {
	return &Dispatcher{send: s}
}

func (dp *Dispatcher) Process(ctx context.Context, d *delivery.Delivery) error {
	if err := dp.send.Send(ctx, d); err != nil {
		d.Fail(err)
		return err
	}

	d.MarkSend()
	return nil
}
