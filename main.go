package main

import "fmt"

// materi variabel
func variabel() {
	var nama string = "Bagas"
	var angka int = 20
	var pecahan float32 = 0.5
	var boolean bool = true

	const alamat string = "Bandung"

	fmt.Println("Tipe data string", nama)
	fmt.Println("Tipe data integer", angka)
	fmt.Println("Tipe data float", pecahan)
	fmt.Println("Tipe data boolean", boolean)
	fmt.Println("Konstanta", alamat)
}

// materi operator
func operator() {
	angka1 := 20
	angka2 := 8

	fmt.Println("Hasil penjumlahan", angka1+angka2)
	fmt.Println("Hasil pengurangan", angka1-angka2)
	fmt.Println("Hasil perkalian", angka1*angka2)
	fmt.Println("Hasil pembagian", float32(angka1)/float32(angka2))
	fmt.Println("Hasil sisa pembagian", angka1%angka2)
}

// materi if else
func ifElse() {
	nilai := 78

	if nilai > 86 {
		fmt.Println("Nilai A")
	} else if nilai > 75 {
		fmt.Println("Nilai AB")
	} else if nilai > 66 {
		fmt.Println("Nilai B")
	} else if nilai > 55 {
		fmt.Println("Nilai C")
	} else {
		fmt.Println("Tidak lulus")
	}
}

// materi switch case
func switchCase() {
	warnaLampu := "Merah"

	switch warnaLampu {
	case "Merah":
		fmt.Println("Berhenti")
	case "Kuning":
		fmt.Println("Siap-siap")
	case "Hijau":
		fmt.Println("Jalan")
	default:
		fmt.Println("Error")
	}

	angka := 10

	switch {
	case angka == 4:
		fmt.Println("Angka 4")
	case angka > 4:
		fmt.Println("Lebih dari 4")
	default:
		fmt.Println("Kurang dari 4")
	}

}

// materi array
func array() {
	arr := []int{1, 2, 3, 4, 5}
	arr = append(arr, 6, 7, 8, 9)
	fmt.Println(arr)
}

// materi looping
func loop() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Ini angka", arr[i])
	}

	for i := 0; i < len(arr); i++ {
		if arr[i]%3 == 0 {
			continue
		}
		fmt.Println("Ini bukan kelipatan 3 :", arr[i])
	}
}

// materi slice
func slice() {
	var arr1 = [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	arr2 := make([]int, 4)
	arr2 = append(arr2, 1, 2, 3, 4, 5)

	fmt.Println(arr2)
}

// materi map
func mapType() {
	ibukota := map[string]string{
		"Indonesia": "Jakarta",
		"Thailand":  "Bangkok",
		"Filipina":  "Manila",
	}

	fmt.Println(ibukota)

	// delete data map
	delete(ibukota, "Filipina")

	fmt.Println(ibukota)

	makanan := make(map[string]string)
	makanan["Indonesia"] = "Rendang"
	makanan["Thailand"] = "Tom Yum"
	makanan["Filipina"] = "Tapsilog"

	fmt.Println(makanan)

	// cek key dan value map
	value, hasKey := ibukota["Indonesia"]
	fmt.Println(hasKey, value)
}

// materi range loop
func rangeLoop() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}
}

// materi function
func functionReturn(operator string, angka1 int, angka2 int) int {
	if operator == "+" {
		return angka1 + angka2
	} else {
		return angka1 + angka2
	}

	// function yang huruf depannya besar = public, kalo kecil = private
}

// materi multiple return function
func multipleReturnFunction(angka1 int, angka2 int) (int, int) {
	penjumlahan := angka1 + angka2
	pengurangan := angka1 - angka2
	return penjumlahan, pengurangan
}

// materi variadic function
func variadicFunction(angka ...int) int {
	sum := 0
	for i := 0; i < len(angka); i++ {
		sum += angka[i]
	}

	return sum
}

// materi closure function
func closureFunction() {
	nama := func(name string) string {
		return name
	}

	functionNama := nama("Bagas")
	fmt.Println(functionNama)

	func(name string) {
		fmt.Println(name)
	}("Adrianus Bagas")

	functionNamaLengkap := functionDalamFunction()
	fmt.Println(functionNamaLengkap("Adrianus Bagas Tantyo Dananjaya"))

	angka := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fungsiGanjilGenap := GanjilGenap()
	ganjil, genap := fungsiGanjilGenap(angka...)
	fmt.Println("Angka Ganjil :", ganjil)
	fmt.Println("Angka Genap :", genap)

}

func functionDalamFunction() func(string) string {
	return func(nama string) string {
		return nama
	}
}

func GanjilGenap() func(...int) ([]int, []int) {
	return func(angka ...int) ([]int, []int) {
		ganjil, genap := []int{}, []int{}

		for _, value := range angka {
			if value%2 == 0 {
				genap = append(genap, value)
			} else {
				ganjil = append(ganjil, value)
			}
		}

		return ganjil, genap
	}
}

// materi pointer
func pointer() {
	a := 10
	fmt.Println("Value a :", a)

	// mencetak alamat memori a
	fmt.Println("Alamat memori a :", &a)
	// deklarasi variabel b sebagai pointer
	var b *int
	// assign variabel b dengan alamat memori a
	b = &a
	fmt.Println("Alamat memori b :", b)
	// mencetak nilai b
	fmt.Println("Value b :", *b)
	// pointer dalam function (yang dikirim alamat memori parameter)
	angka1, angka2 := 10, 9
	pointerInFunction(&angka1, &angka2)
	// nested assign variabel dengan alamat memori
	bilangan1 := 10
	bilangan2 := &bilangan1
	bilangan3 := &bilangan2

	fmt.Println("Value bilangan 1 :", bilangan1)
	fmt.Println("Value bilangan 2 (alamat memori bilangan 1):", bilangan2)
	fmt.Println("Value bilangan 2 :", *bilangan2)
	fmt.Println("Value bilangan 3 (alamat memori bilangan 2):", bilangan3)
	fmt.Println("Value bilangan 3 (alamat memori bilangan 1):", *bilangan3)
	fmt.Println("Value bilangan 3 :", **bilangan3)
}

func pointerInFunction(a *int, b *int) {
	fmt.Println("Hasil penjumlahan function dengan parameter pointer :", *a+*b)
}

// materi struct
func structLesson() {
	type User struct {
		FirstName string
		LastName  string
		Address   string
		Age       int
	}

	user := User{
		FirstName: "Adrianus",
		LastName:  "Bagas",
		Address:   "Bandung",
		Age:       22,
	}

	fmt.Println(user)
	fmt.Println("Alamat user :", user.Address)

	userPointer := &user
	// mencetak alamat user
	fmt.Println(userPointer)
	// mencetak value user
	fmt.Println(*userPointer)
	// mencetak element userPointer
	fmt.Println(userPointer.Address)

	type Customer struct {
		JenisKelamin string
		User         User
	}

	customer := Customer{
		JenisKelamin: "Pria",
		User:         user,
	}

	fmt.Println(customer)
}

// materi method
type Nilai struct {
	a, b int
}

// method return
func (n Nilai) Penjumlahan() int {
	return n.a + n.b
}

// method tanpa return
func (n Nilai) Perkalian() {
	fmt.Println("Hasil perkalian :", n.a*n.b)
}

// method dengan pointer (untuk ubah value variabel)
func (n *Nilai) UbahNilaiA(a int) {
	n.a = a
}

func main() {
	fmt.Println("Halo Guys")
	fmt.Println("=====VARIABEL=====")
	variabel()
	fmt.Println()
	fmt.Println("=====OPERATOR=====")
	operator()
	fmt.Println()
	fmt.Println("=====IF ELSE=====")
	ifElse()
	fmt.Println()
	fmt.Println("=====SWITCH CASE=====")
	switchCase()
	fmt.Println()
	fmt.Println("=====ARRAY=====")
	array()
	fmt.Println()
	fmt.Println("=====LOOP FOR=====")
	loop()
	fmt.Println()
	fmt.Println("=====SLICE=====")
	slice()
	fmt.Println()
	fmt.Println("=====MAP=====")
	mapType()
	fmt.Println()
	fmt.Println("=====RANGE=====")
	rangeLoop()
	fmt.Println()
	fmt.Println("=====FUNCTION=====")
	fmt.Println(functionReturn("+", 10, 10))
	fmt.Println()
	fmt.Println("=====MULTIPLE RETURN FUNCTION=====")
	penjumlahan, pengurangan := multipleReturnFunction(20, 10)
	fmt.Println("Hasil penjumlahan :", penjumlahan)
	fmt.Println("Hasil pengurangan :", pengurangan)
	fmt.Println()
	fmt.Println("=====VARIADIC FUNCTION=====")
	fmt.Println("Hasil penjumlahan :", variadicFunction(1, 2, 3, 4, 5))
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println("Hasil penjumlahan slice :", variadicFunction(slice...))
	fmt.Println()
	fmt.Println("=====CLOSURE FUNCTION=====")
	closureFunction()
	fmt.Println()
	fmt.Println("=====POINTER=====")
	pointer()
	fmt.Println()
	fmt.Println("=====STRUCT=====")
	structLesson()
	fmt.Println()
	fmt.Println("=====METHOD=====")
	nilai := Nilai{2, 4}
	fmt.Println("Nilai a sebelum diubah :", nilai.a)
	fmt.Println("Hasil penjumlahan sebelum ubah a :", nilai.Penjumlahan())
	nilai.Perkalian()
	nilai.UbahNilaiA(10)
	fmt.Println("Nilai a sesudah diubah :", nilai.a)
	fmt.Println("Hasil penjumlahan sesudah ubah a :", nilai.Penjumlahan())
	nilai.Perkalian()
	fmt.Println()
}
