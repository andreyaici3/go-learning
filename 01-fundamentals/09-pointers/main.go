package main

import "fmt"

/*
Latihan 09.4
Buat fungsi changeScore() dengan ketentuan:
1. Menerima parameter pointer *int.
2. Mengubah nilai yang ditunjuk menjadi 100.
3. Di main(), buat variabel score dengan nilai 80.
4. Kirim alamat score ke fungsi changeScore().
5. Cetak nilai score sebelum dan sesudah pemanggilan fungsi.
Target output:
Before: 80
After: 100

Jawaban

func changeScore(score *int) {
	*score = 100
}

func main() {
	score := 80

	fmt.Println("Before:", score)
	changeScore(&score)
	fmt.Println("After:", score)
}

*/

/*
Latihan 09.5
Buat program dengan ketentuan:
1. Deklarasikan pointer scorePointer bertipe *int tanpa nilai awal.
2. Periksa apakah pointer bernilai nil.
3. Jika nil, cetak "Pointer belum memiliki alamat".
4. Buat variabel score dengan nilai 90.
5. Arahkan scorePointer ke alamat score.
6. Cetak nilai melalui pointer.
Target output:
Pointer belum memiliki alamat
Score: 90

Jawaban:

func main() {
	var scorePointer *int

	if scorePointer == nil {
		fmt.Println("Pointer belum memiliki alamat")
	}

	score := 90

	scorePointer = &score
	fmt.Println("Score:", *scorePointer)
}
*/

/*
Latihan 09.6
Buat struct Product dengan field:
- Name string
- Price int
Kemudian:
1. Buat variabel product dengan nama "Laptop" dan harga 10000000.
2. Buat pointer productPointer yang menunjuk ke product.
3. Cetak nama dan harga melalui pointer.
4. Ubah harga menjadi 12000000 melalui pointer.
5. Cetak kembali harga dari variabel product.
Target output:
Product: Laptop
Price: 10000000
Updated Price: 12000000

Jawaban:

type Product struct {
	Name  string
	Price int
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 10000000,
	}

	productPointer := &product

	fmt.Println("Product:", productPointer.Name)
	fmt.Println("Price:", productPointer.Price)

	productPointer.Price = 12000000

	fmt.Println("Updated Price:", productPointer.Price)
}

*/

/*
Latihan Integrasi — Pointers + Struct + Functions
Buat program sederhana untuk mengelola data produk.
Ketentuan
1. Buat struct Product

Memiliki field:
- Name bertipe string
- Price bertipe int
- Stock bertipe int
2. Buat fungsi updateProduct()

Ketentuan:
- Menerima parameter pointer *Product.
- Mengubah Name menjadi "MacBook Pro".
- Mengubah Price menjadi 25000000.
- Mengubah Stock menjadi 10.
3. Implementasikan di main()
- Buat variabel product dengan data awal:
  - Name: "Laptop"
  - Price: 10000000
  - Stock: 5
- Tampilkan data sebelum perubahan.
- Panggil updateProduct() dengan mengirim alamat variabel product.
- Tampilkan data setelah perubahan.
Target output
=== Before Update ===
Name: Laptop
Price: 10000000
Stock: 5

=== After Update ===
Name: MacBook Pro
Price: 25000000
Stock: 10


Ketentuan tambahan:
- Jangan mengembalikan Product menggunakan return.
- Perubahan harus dilakukan melalui pointer.
- Pisahkan fungsi updateProduct() dari main().
*/

type Product struct {
	Name  string
	Price int
	Stock int
}

func updateProduct(product *Product) {
	product.Name = "MacBook Pro"
	product.Price = 25000000
	product.Stock = 10
}

func main() {
	product := Product{
		Name:  "Laptop",
		Price: 10000000,
		Stock: 5,
	}

	fmt.Println("=== Before Update ===")
	fmt.Println("Name:", product.Name)
	fmt.Println("Price:", product.Price)
	fmt.Println("Stock:", product.Stock)
	fmt.Println()
	updateProduct(&product)
	fmt.Println("=== After Update ===")
	fmt.Println("Name:", product.Name)
	fmt.Println("Price:", product.Price)
	fmt.Println("Stock:", product.Stock)

}
