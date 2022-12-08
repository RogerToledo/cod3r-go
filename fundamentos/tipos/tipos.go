package main

import (
	"fmt"
	"math"
	"reflect"
)

func main(){
	// numeros inteiros
	fmt.Println(1, 10, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal, apenas positivo - uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("Byte é", reflect.TypeOf(b))

	var ui8 uint8 = math.MaxUint8
	var ui16 uint16 = math.MaxUint16
	var ui32 uint32 = math.MaxUint32
	var ui64 uint64 = math.MaxUint64
	var uint uint = math.MaxUint
	fmt.Println("ui8:", ui8, "| ui16:", ui16, "| ui32:", ui32, "| uint64:", ui64, "| uint:", uint)

	// com sinal - int8, int16, int32, int64
	i8 := math.MaxInt8
	i16 := math.MaxInt16
	i32 := math.MaxInt32
	i64 := math.MaxInt64
	int := math.MaxInt
	fmt.Println("i8:", i8, "| i16:", i16, "| i32:", i32, "| i64:", i64, "| int:", int)
	fmt.Println("i64 é", reflect.TypeOf(i64))

	// Não existe tipo char. Representa mapeamento da tabela Unicode (int32)
	var v1 rune = 'a'
	fmt.Println("rune é", reflect.TypeOf(v1))
	fmt.Println(v1)

	// Numeros reais float32, float64
	var x float32 = 49.99
	fmt.Println("x é", reflect.TypeOf(x))
	fmt.Println("literal de 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Hello World"
	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho de s1 é", len(s1))

	s2 := `Hello
	World`
	fmt.Println(s2)
	fmt.Println("O tamanho de s2 é", len(s2))

	s3 := `Hello
World`
	fmt.Println(s3)
	fmt.Println("O tamanho de s2 é", len(s3))



}