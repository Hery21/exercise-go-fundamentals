package main

type Product struct {
	Barcode string  `json:"barcode"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}
