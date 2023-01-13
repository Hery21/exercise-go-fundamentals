package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func loadProducts() []*Product {
	file, _ := ioutil.ReadFile("products.json")
	var data []*Product
	if err := json.Unmarshal(file, &data); err != nil {
		panic("failed load json product")
	}
	return data
}

func CalculatePurchase(barcodes []string) *PurchaseSummary {
	purchaseBarcodes := make(map[string]int)

	for _, barcode := range barcodes {
		if len(purchaseBarcodes) == 0 {
			purchaseBarcodes[barcode] = 1
			continue
		}
		if _, ok := purchaseBarcodes[barcode]; ok {
			purchaseBarcodes[barcode] += 1
			continue
		}
		purchaseBarcodes[barcode] = 1
	}

	var bill PurchaseSummary = PurchaseSummary{[]*PurchasedItem{}, 0}

	for _, product := range loadProducts() {
		if val, ok := purchaseBarcodes[product.Barcode]; ok {
			tempItems := PurchasedItem{product.Barcode, product.Name, product.Price * float32(val), product.Price, val}
			tempSlice := []*PurchasedItem{&tempItems}

			if bill.TotalPrice == 0 {
				bill = PurchaseSummary{tempSlice, tempItems.Subtotal}
			} else {
				fmt.Println(tempItems)
				bill.PurchasedItems = append(bill.PurchasedItems, &tempItems)
				bill.TotalPrice += tempItems.Subtotal
			}
		}
	}
	return &bill
}

func GenerateReceiptText(purchaseSummary *PurchaseSummary) string {
	var receiptText []string = []string{""}

	for _, item := range purchaseSummary.PurchasedItems {
		receiptText = append(receiptText, fmt.Sprintf("%v %v x %v = %v", item.Name, item.Price, item.Quantity, item.Subtotal))
	}

	receiptText = append(receiptText, fmt.Sprintf("--------------------"))
	receiptText = append(receiptText, fmt.Sprintf("Total = %v", purchaseSummary.TotalPrice))

	return fmt.Sprintf("%s", strings.Join(receiptText, "\n"))
}
