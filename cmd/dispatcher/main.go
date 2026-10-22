package main

import (
	"fmt"

	"fabriciolfj.github/study/internal/delivery"
)

type Money int64

func main() {
	var deliveries = make([]*delivery.Delivery, 0, 3)
	for i := 0; i < 3; i++ {
		value := delivery.New("001", "http://localhost:8080/test", nil)
		incrementar(value)

		deliveries = append(deliveries, value)
	}

	for i := 0; i < 2; i++ {
		value := delivery.New("002", "http://localhost:8080/test", nil)
		incrementar(value)

		deliveries = append(deliveries, value)
	}

	result := delivery.IDs(deliveries)
	println(result)

	println("=============")
	resultMap := delivery.GroupByMerchant(deliveries)
	fmt.Printf("%v\n", resultMap)

	money := delivery.Money(100).Split(3)
	fmt.Printf("%v\n", somar(money))
}

func incrementar(d *delivery.Delivery) {
	d.Attempts++
}

func somar(ds []delivery.Money) delivery.Money {
	var sum delivery.Money
	for _, d := range ds {
		sum += d
	}

	return sum
}
