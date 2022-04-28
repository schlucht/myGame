package main

import "github.com/schlucht/myGame/pkg/payment"

func main() {
	payment.ReadFile("./daten/pay.csv")
}
