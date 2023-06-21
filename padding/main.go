package main

import (
	"fmt"
	"unsafe"
)

type ZeroStruct struct {
}

type A24 struct {
	flag     uint8
	length   int64
	boolFlag bool
}

type A16 struct {
	length   int64
	boolFlag bool
	flag     uint8
}

type ASup struct {
	b        bool
	length   int64
	boolFlag bool
	flag     uint8
}

type ASupOptimized struct {
	length   int64
	b        bool
	boolFlag bool
	flag     uint8
}

type A struct {
	a bool
	b int64
}

type ComposedStruct struct {
	a bool
	b A
	c bool
}

type MinA struct {
	a int8
	b int8
}

func main() {
	z := ZeroStruct{}
	fmt.Println("z size", unsafe.Sizeof(z))   // 0
	fmt.Println("z align", unsafe.Alignof(z)) // 1
	fmt.Println()

	// not optimized struct
	a24 := A24{}
	fmt.Println("a24 size", unsafe.Sizeof(a24))   // 24
	fmt.Println("a24 align", unsafe.Alignof(a24)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct A24 {
			flag     uint8 %d
			length   int64 %d
			boolFlag bool  %d
		}
	`,
		unsafe.Offsetof(a24.flag),
		unsafe.Offsetof(a24.length),
		unsafe.Offsetof(a24.boolFlag)))
	fmt.Println()

	// optimized
	a16 := A16{}
	fmt.Println("a16 size", unsafe.Sizeof(a16))   // 16
	fmt.Println("a16 align", unsafe.Alignof(a16)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct A16 {
			length   int64 %d
			boolFlag bool  %d
			flag     uint8 %d
		}
	`,
		unsafe.Offsetof(a16.length),
		unsafe.Offsetof(a16.boolFlag),
		unsafe.Offsetof(a16.flag)))
	fmt.Println()

	// not optimized struct
	as := ASup{}
	fmt.Println("asup size", unsafe.Sizeof(as))   // 24
	fmt.Println("asup align", unsafe.Alignof(as)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct ASup {
			b        bool  %d
			length   int64 %d
			boolFlag bool  %d
			flag     uint8 %d
		}
	`,
		unsafe.Offsetof(as.b),
		unsafe.Offsetof(as.length),
		unsafe.Offsetof(as.boolFlag),
		unsafe.Offsetof(as.flag)))
	fmt.Println()

	// optimized struct
	ao := ASupOptimized{}
	fmt.Println("ao size", unsafe.Sizeof(ao))   // 16
	fmt.Println("ao align", unsafe.Alignof(ao)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct ASupOptimized {
			length   int64 %d
			b        bool  %d
			boolFlag bool  %d
			flag     uint8 %d
		}
	`,
		unsafe.Offsetof(ao.length),
		unsafe.Offsetof(ao.b),
		unsafe.Offsetof(ao.boolFlag),
		unsafe.Offsetof(ao.flag)))
	fmt.Println()

	// non optimized struct
	ac := ComposedStruct{}
	fmt.Println("ac size", unsafe.Sizeof(ac))   // 32
	fmt.Println("ac align", unsafe.Alignof(ac)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct ComposedStruct {
			a   bool %d
			b   A    %d
			c   bool %d
		}
	`,
		unsafe.Offsetof(ac.a),
		unsafe.Offsetof(ac.b),
		unsafe.Offsetof(ac.c)))
	fmt.Println()

	m := MinA{}
	fmt.Println("m size", unsafe.Sizeof(m))   // 32
	fmt.Println("m align", unsafe.Alignof(m)) // 8
	fmt.Println("offsets", fmt.Sprintf(`
		struct MinA {
			a   int8 %d
			b   int8 %d
		}
	`,
		unsafe.Offsetof(m.a),
		unsafe.Offsetof(m.b)))
	fmt.Println()
}
