package main

import "fmt"

/* LEARN 1
Tugas
Buat tiga function:
1. sayHello() menampilkan "Hello, Go!".
2. sayGoodbye() menampilkan "Goodbye, Go!".
3. showInfo() menampilkan "Learning Go Functions".

Panggil ketiga function tersebut dari dalam main().
*/

// func sayHello() {
// 	fmt.Println("Hello, Go!")
// }

// func sayGoodbye() {
// 	fmt.Println("Goodbye, Go!")
// }

// func showInfo() {
// 	fmt.Println("Learning Go Functions")
// }

// func main() {
// 	sayHello()
// 	sayGoodbye()
// 	showInfo()
// }

// END LEARN 1

/* Learn 2
Latihan 02
Buat function bernama greetUser() dengan parameter name bertipe string.

Ketentuannya:
1. Function menerima nama user.
2. Menampilkan "Welcome, [nama]!".
3. Panggil function untuk tiga nama berbeda.

func introduceUser(name string, age int, role string) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Role:", role)
}

func main() {
	introduceUser("Andrey", 25, "Backend Developer")
}

End Learn 2 */

/* Learn 3
Latihan 04
Buat function calculateArea() dengan ketentuan:
1. Menerima parameter length bertipe int.
2. Menerima parameter width bertipe int.
3. Mengembalikan hasil perkalian keduanya.
4. Simpan hasil function ke variabel area.
5. Tampilkan hasilnya.


func calculateArea(length int, width int) int {
	return length * width
}

func main() {
	area := calculateArea(10, 5)

	fmt.Println("Area:", area)
}

*/

/* Learn 4
Latihan 05
Buat function calculateNumbers() dengan ketentuan:
1. Menerima dua parameter integer: a dan b.
2. Mengembalikan hasil penjumlahan.
3. Mengembalikan hasil perkalian.
4. Terima kedua hasil tersebut di dalam main().
5. Tampilkan hasilnya


func calculateNumbers(a int, b int) (sum int, multiply int) {
	return a + b, a * b
}

func main() {
	sum, multiply := calculateNumbers(10, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Multiply:", multiply)
}
End Learn 4
*/

/* Learn 5
Functions 07 — Praktik Variadic Function
Sekarang kita praktik.

Buat function calculateTotal() dengan ketentuan:
1. Menerima argument integer dalam jumlah bebas menggunakan ...int.
2. Menghitung total seluruh angka.
3. Mengembalikan hasil total dalam bentuk int.
4. Tampilkan hasilnya dari main().



func calculateTotal(numbers ...int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := calculateTotal(10, 20, 30, 40, 50)
	fmt.Println("Total:", total)
}
End Learn 5*/

/*
3. Latihan
Buat program dengan ketentuan:
1. Buat fungsi processData().
2. Cetak "Memulai proses".
3. Gunakan tiga defer untuk mencetak:
   - "Membersihkan resource"
   - "Menutup koneksi"
   - "Proses selesai"
4. Sebelum fungsi berakhir, cetak "Memproses data".


func processData() {
	fmt.Println("Memulai Proses")

	defer fmt.Println("Membersihkan resource")
	defer fmt.Println("Menutup Koneksi")
	defer fmt.Println("Proses Selesai")

	fmt.Println("Memproses data")
}

func main() {
	processData()
}

*/

/*Latihan 08.10
Coba buat fungsi berikut:

getUserAge() dengan ketentuan:
1. Mengembalikan nilai 25 bertipe int.
2. Sebelum return, cetak "Mengambil data user".
3. Gunakan defer untuk mencetak "Selesai mengambil data".
4. Di main(), panggil fungsi tersebut dan tampilkan hasilnya dengan format "User Age: 25".

func getUserAge() int {
	age := 25
	fmt.Println("Mengambil data user")

	defer fmt.Println("Selesai mengambil data")

	return age
}

func main() {
	age := getUserAge()
	fmt.Println("User Age:", age)
}


func calculateScore() (score int) {
	score = 80

	defer func() {
		score += 20
	}()

	return score
}

func main() {
	score := calculateScore()
	fmt.Println("Final Score:", score)
}
*/

/*
3. Latihan 08.12
Buat program dengan ketentuan:
1. Buat anonymous function tanpa nama.
2. Simpan fungsi tersebut ke variabel calculate.
3. Fungsi menerima dua parameter bertipe int.
4. Fungsi mengembalikan hasil penjumlahan.
5. Panggil fungsi dengan nilai 10 dan 20.
6. Cetak hasilnya.

func calculate(a int, b int, operation func(a int, b int) int) int {
	return operation(a, b)
}

func main() {
	perkalian := func(a int, b int) int {
		return a * b
	}

	penjumlahan := func(a int, b int) int {
		return a + b
	}

	fmt.Println("Addition:", calculate(10, 20, penjumlahan))
	fmt.Println("Multiplication:", calculate(10, 20, perkalian))
}

func getOperation() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	add := getOperation()

	result := add(10, 20)

	fmt.Println("Result:", result)
} */

/*
3. Latihan 08.14
Buat fungsi getOperation() dengan ketentuan:
1. Menerima parameter operation string.
2. Mengembalikan fungsi dengan signature func(int, int) int.
3. Jika operation == "add", kembalikan fungsi penjumlahan.
4. Jika operation == "multiply", kembalikan fungsi perkalian.
5. Di main(), gunakan keduanya.

func getOperation(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(i1, i2 int) int {
			return i1 + i2
		}
	case "multiply":
		return func(i1, i2 int) int {
			return i1 * i2
		}
	default:
		return func(i1, i2 int) int {
			return 0
		}
	}
}

func main() {
	add := getOperation("add")
	multiply := getOperation("multiply")

	fmt.Println("Addition:", add(10, 20))
	fmt.Println("Multiplication:", multiply(10, 20))
}
*/

/*
08.15 — Function Review (Latihan Integrasi)
Sebelum menutup materi Functions, kita lakukan satu latihan integrasi.

Buat program kalkulator sederhana dengan ketentuan:
1. Buat fungsi getOperation() yang menerima operator (string).
2. Fungsi mengembalikan function func(int, int) int.
3. Dukung operator:
   - add
   - subtract
   - multiply
4. Gunakan anonymous function untuk setiap operasi.
5. Panggil ketiga operasi dari main(). */

func getOperation(operator string) func(int, int) int {
	switch operator {
	case "add":
		return func(i1, i2 int) int {
			return i1 + i2
		}
	case "subtract":
		return func(i1, i2 int) int {
			return i2 - i1
		}
	case "multiply":
		return func(i1, i2 int) int {
			return i1 * i2
		}
	default:
		return func(i1, i2 int) int {
			return 0
		}
	}
}

func main() {
	add := getOperation("add")
	subtract := getOperation("subtract")
	multiply := getOperation("multiply")

	fmt.Println("Addition:", add(10, 20))
	fmt.Println("Subtraction:", subtract(10, 20))
	fmt.Println("Multiplication:", multiply(10, 20))
}
