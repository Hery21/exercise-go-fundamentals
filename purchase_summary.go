package main

type PurchasedItem struct {
	Barcode  string  `json:"barcode"`
	Name     string  `json:"name"`
	Subtotal float32 `json:"subtotal"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type PurchaseSummary struct {
	PurchasedItems []*PurchasedItem `json:"purchasedItems"`
	TotalPrice     float32          `json:"totalPrice"`
}
