package main

import "fmt"

// func main() {
// 	user := map[string]any{
// 		"name": "Andrey",
// 		"age":  25,
// 		"role": "Backend Developer",
// 	}

// 	fmt.Println("Nama:", user["name"])
// 	fmt.Println("Umur:", user["age"])
// 	fmt.Println("Role:", user["role"])
// }

/* Latihan 02
Lanjutkan kode kamu sebelumnya dengan ketentuan:
1. Ubah name menjadi "Andi".
2. Ubah age menjadi 26.
3. Tambahkan key email dengan value "andi@example.com".
4. Hapus key role.
5. Tampilkan Map setelah semua perubahan.

Jawaban

func main() {
	user := map[string]any{
		"name": "Andrey",
		"age":  25,
		"role": "Backend Developer",
	}

	user["name"] = "Andi"
	user["age"] = 26
	user["email"] = "andi@example.com"

	delete(user, "role")

	fmt.Println("Nama:", user["name"])
	fmt.Println("Umur:", user["age"])
	fmt.Println("Email:", user["email"])
	fmt.Println("Role:", user["role"])
}

*/

/*
Latihan 03
Coba buat program untuk menampilkan seluruh isi Map menggunakan for range

Jawaban

func main() {
	user := map[string]any{
		"name": "Andrey",
		"age":  25,
		"role": "Backend Developer",
	}

	for key, value := range user {
		fmt.Println(key, ":", value)
	}
}
*/

/*
Buat Map
scores := map[string]int{
    "Andrey": 80,
    "Agik":   90,
    "Andi":   85,
}

Tugas:
1. Tampilkan seluruh nama dan nilai menggunakan for range.
2. Hitung total seluruh nilai.
3. Tampilkan jumlah siswa menggunakan len(scores).

Andrey: 80
Agik: 90
Andi: 85

Total Score: 255
Total Students: 3

Urutan nama tidak harus sama.
Jawaban:

func main() {
	scores := map[string]int{
		"Andrey": 80,
		"Agik":   90,
		"Andi":   85,
	}

	var totalScore int

	for key, value := range scores {
		fmt.Println(key, ":", value)
		totalScore += value
	}

	fmt.Println("Total Score:", totalScore)
	fmt.Println("Total Student: ", len(scores))

}
*/

/*
Nested Map

*/

func main() {
	users := map[string]map[string]any{
		"user_1": {
			"name": "Andrey",
			"age":  25,
		},
		"user_2": {
			"name": "Agik",
			"age":  26,
		},
		"user_3": {
			"name": "Andi",
			"age":  24,
		},
	}

	for key, value := range users {
		fmt.Println("ID:", key)
		fmt.Println("Name:", value["name"])
		fmt.Println("Age:", value["age"])
		fmt.Println("")
	}

	searchID := "user_2"
	value, exists := users[searchID]
	if exists {
		fmt.Println("Nama: ", value["name"])
		fmt.Println("Umur: ", value["age"])
	} else {
		fmt.Println("User not found")
	}

}
