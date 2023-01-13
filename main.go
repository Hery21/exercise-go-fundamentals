package main

import (
	"encoding/json"
	"fmt"
)

func debugToJson(purchase *PurchaseSummary) {
	data, _ := json.Marshal(purchase)
	fmt.Println(string(data))
}

func runPosApp() {
	barcodes := []string{"00001", "00001", "00002", "00003"}
	purchaseSummary := CalculatePurchase(barcodes)
	debugToJson(purchaseSummary)

	receipt := GenerateReceiptText(purchaseSummary)
	fmt.Println(receipt)
}

func runParkBilling() {
	price := CalculateParkingBill(Car, 2)
	fmt.Println("price", price)
}

func main() {
	runPosApp()
	//runParkBilling()
}
