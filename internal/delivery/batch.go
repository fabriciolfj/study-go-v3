package delivery

func IDs(ds []*Delivery) []string {
	value := make([]string, 0, len(ds))

	return value
}

func GroupByMerchant(ds []*Delivery) map[string][]*Delivery {
	deliveries := make(map[string][]*Delivery)

	for _, d := range ds {
		deliveries[d.MerchantId] = append(deliveries[d.MerchantId], d)
	}

	return deliveries
}

func Preview(d *Delivery, n int) []byte {
	value := make([]byte, n, len(d.ID))

	return value
}
