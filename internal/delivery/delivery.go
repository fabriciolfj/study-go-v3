package delivery

import (
	"fmt"

	"github.com/google/uuid"
)

type Money int64

type Delivery struct {
	ID         string
	MerchantId string
	Endpoint   string
	Payload    []byte
	Attempts   int
	Status     string
	Error      error
}

func New(merchantId string, endpoint string, payload []byte) *Delivery {
	return &Delivery{
		MerchantId: merchantId,
		Endpoint:   endpoint,
		Payload:    payload,
		Attempts:   0,
		ID:         uuid.New().String(),
	}
}

func (d *Delivery) MarkSend() {
	d.Status = "SEND"
}

func (d *Delivery) Fail(message error) {
	d.Attempts++
	d.Error = message
}

func (d *Delivery) IsRetryable() bool {
	return d.Attempts < 5
}

func (d *Delivery) String() string {
	return fmt.Sprintf("%+v", *d)
}

func (m Money) Split(partes int) []Money {
	if partes < 1 {
		return nil
	}

	out := make([]Money, partes)

	passo := Money(1)

	base := m / Money(partes)
	resto := m % Money(partes)

	if resto < 0 {
		passo, resto = -1, -resto
	}

	for i := range out {
		out[i] = base
		if Money(i) < resto {
			out[i] += passo
		}
	}

	return out
}
