package main

import (
	"fmt"
)

/*
Buat main.go dengan ketentuan berikut:
1. Buat interface bernama PaymentMethod.
2. Interface memiliki method pay(amount int).
3. Buat struct BankTransfer dengan field BankName string.
4. Buat struct EWallet dengan field Provider string.
5. Kedua struct harus memiliki method pay(amount int).
6. Buat function processPayment(payment PaymentMethod, amount int) yang menjalankan method pay().
7. Panggil function tersebut menggunakan BankTransfer dan EWallet.
Target output:
Processing payment using BCA
Amount: 100000

Processing payment using GoPay
Amount: 50000

Jawaban:

type PaymentMethod interface {
	pay(amount int)
}

type BankTransfer struct {
	BankName string
}

func (e *BankTransfer) pay(amount int) {
	fmt.Println("Processing payment using", e.BankName)
	fmt.Println("Amount:", amount)
}

type EWallet struct {
	Provider string
}

func (e *EWallet) pay(amount int) {
	fmt.Println("Processing payment using", e.Provider)
	fmt.Println("Amount:", amount)
}

func processPayment(payment PaymentMethod, amount int) {
	payment.pay(amount)
}

func main() {
	bank := &BankTransfer{
		BankName: "BCA",
	}

	ewallet := &EWallet{
		Provider: "GoPay",
	}

	processPayment(bank, 100000)
	fmt.Println()
	processPayment(ewallet, 50000)

	data := any(100)

	value, ok := data.(int)
	fmt.Println(value)
	fmt.Println(ok)
} */

func checkData(data any) {
	switch data.(type) {
	case string:
		fmt.Println("Data berupa string")
	case int:
		fmt.Println("Data berupa integer")
	case float64:
		fmt.Println("Data berupa float")
	case bool:
		fmt.Println("Data berupa boolean")
	default:
		fmt.Println("Tipe data tidak dikenali")
	}
}

func main() {
	checkData("STR")
	checkData(111)
	checkData(10.5)
	checkData(true)
}
