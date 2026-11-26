package main

import "fmt"

/*
Tugas:
1. Buat struct bernama Employee.
2. Tambahkan field:
   - Name bertipe string
   - Position bertipe string
   - Salary bertipe int
   - IsActive bertipe bool
3. Buat satu variabel employee menggunakan struct tersebut.
4. Isi semua field dengan data.
5. Tampilkan seluruh data employee.
6. Buat satu variabel employee2 menggunakan zero value struct, kemudian tampilkan seluruh field-nya

Jawaban:

type Employee struct {
	Name     string
	Position string
	Salary   int
	IsActive bool
}

func main() {
	employee := Employee{
		Name:     "Andrey",
		Position: "Backend Developer",
		Salary:   1000000,
		IsActive: true,
	}

	fmt.Println("==== EMPLOYEE =====")
	fmt.Println("Name:", employee.Name)
	fmt.Println("Position:", employee.Position)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Is Active:", employee.IsActive)

	employee2 := Employee{}

	fmt.Println()
	fmt.Println("==== EMPLOYEE 2 =====")
	fmt.Println("Name:", employee2.Name)
	fmt.Println("Position:", employee2.Position)
	fmt.Println("Salary:", employee2.Salary)
	fmt.Println("Is Active:", employee2.IsActive)
}

*/

/*
Praktik 10.4
Tambahkan method pada struct Employee yang sudah kamu buat sebelumnya.

Tugas:
1. Buat method displayInfo() untuk menampilkan seluruh informasi employee.
2. Buat method getAnnualSalary() yang mengembalikan gaji tahunan (Salary × 12).
3. Panggil kedua method dari main().
4. Tampilkan hasil gaji tahunan.

Jawaban:

*/

type Employee struct {
	Name     string
	Position string
	Salary   int
	IsActive bool
}

func (e *Employee) displayInfo() {
	fmt.Println("==== EMPLOYEE =====")
	fmt.Println("Name:", e.Name)
	fmt.Println("Position:", e.Position)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Is Active:", e.IsActive)
}

func (e *Employee) getAnnualSalary() int {
	return e.Salary * 12
}

/*
Step 10.6 — Praktik Value vs Pointer Receiver
Sekarang buat dua method pada Employee:
1. increaseSalary() menggunakan value receiver, menaikkan salary sebesar 500.000.
2. increaseSalaryByPointer() menggunakan pointer receiver, menaikkan salary sebesar 500.000.
Panggil kedua method secara berurutan dan tampilkan salary setelah masing-masing pemanggilan. */

func (e Employee) increaseSalary() {
	e.Salary += 500000
}

func (e *Employee) increaseSalaryByPointer() {
	e.Salary += 500000
}

func main() {
	employee := &Employee{
		Name:     "Andrey",
		Position: "Backend Developer",
		Salary:   1000000,
		IsActive: true,
	}

	employee.displayInfo()
	fmt.Println("Gaji Tahunan:", employee.getAnnualSalary())

	employee.increaseSalary()
	employee.increaseSalaryByPointer()

	fmt.Println("Gaji Terakhir:", employee.Salary)

}
