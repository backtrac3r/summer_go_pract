package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var b bool
	var s string
	var p *int

	fmt.Println("Размер базовых типов в байтах на данной архитектуре:")
	fmt.Printf("int: \t\t%d\n", unsafe.Sizeof(i))
	fmt.Printf("int8: \t\t%d\n", unsafe.Sizeof(i8))
	fmt.Printf("int16: \t\t%d\n", unsafe.Sizeof(i16))
	fmt.Printf("int32 (rune): \t%d\n", unsafe.Sizeof(i32))
	fmt.Printf("int64: \t\t%d\n", unsafe.Sizeof(i64))
	fmt.Printf("float32: \t%d\n", unsafe.Sizeof(f32))
	fmt.Printf("float64: \t%d\n", unsafe.Sizeof(f64))
	fmt.Printf("complex64: \t%d\n", unsafe.Sizeof(c64))
	fmt.Printf("complex128: \t%d\n", unsafe.Sizeof(c128))
	fmt.Printf("bool: \t\t%d\n", unsafe.Sizeof(b))
	fmt.Printf("string: \t%d (размер дескриптора, а не данных)\n", unsafe.Sizeof(s))
	fmt.Printf("pointer: \t%d\n", unsafe.Sizeof(p))
}
