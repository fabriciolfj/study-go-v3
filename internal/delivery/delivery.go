package delivery

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
