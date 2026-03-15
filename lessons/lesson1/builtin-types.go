package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

type Empty struct{}

type Bad struct {
	a bool  // 1 байт
	b int64 // 8 байт, выравнивание 8 → +7 байт padding
	c bool  // 1 байт → +7 байт padding
}

type Good struct {
	b int64 // 8 байт
	a bool  // 1 байт
	c bool  // 1 байт → +6 байт padding до кратного 8
}

func main() {

	// integers
	// uint8, uint16, uint32, uint64 — беззнаковые
	// int8, int16, int32, int64 — знаковые
	// uint, int — размер зависит от платформы (32 или 64 бита)
	// uintptr — для хранения указателей
	// byte — алиас для uint8
	// rune — алиас для int32, используется для символов Unicode

	var ui8 uint8 = 5
	fmt.Println(ui8)

	// overflow without panic
	var u uint8 = 255
	u++
	fmt.Println(u)

	//type casting
	var i int = 11
	var i32 int32 = int32(i)
	fmt.Println(i32)

	// floats
	// float32, float64

	// bool — true или false

	// string — неизменяемая последовательность байт (обычно UTF-8)

	// structs — пустая структура занимает 0 байт
	type Empty struct{}
	var e Empty
	fmt.Println(unsafe.Sizeof(e))

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
	fmt.Println(unsafe.Sizeof(p))

	// Выравнивание
	fmt.Println(unsafe.Sizeof(Bad{}))  // 24 на 64-bit
	fmt.Println(unsafe.Alignof(Bad{})) // 8

	fmt.Println(unsafe.Sizeof(Good{}))  // 16 на 64-bit
	fmt.Println(unsafe.Alignof(Good{})) // 8

	// Тип	Размер	Выравнивание
	// bool		1	1
	// int8		1	1
	// int16	2	2
	// int32	4	4
	// int64	8	8
	// float64	8	8
	// string	16	8
	// slice	24	8

	// slice
	// uintptr (8), len (8), cap (8)

	// Литерал
	s := []int{1, 2, 3}
	// make — длина и ёмкость
	s = make([]int, 5)     // len=5, cap=5
	s = make([]int, 3, 10) // len=3, cap=10
	// Срез из массива
	arr := [5]int{1, 2, 3, 4, 5}
	s = arr[1:4] // [2, 3, 4], len=3, cap=4
	fmt.Println(s)
}
