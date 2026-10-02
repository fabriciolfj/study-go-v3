package main

import "fabriciolfj.github/study/internal/delivery"

func main() {
	var value = delivery.New("001", "http://localhost:8080/test", nil)

	incrementar(value)
	println(value.MerchantId)
	println(value.Attempts)

	incrementar(value)
	println(value.Attempts)
}

func incrementar(d *delivery.Delivery) {
	d.Attempts++
}
