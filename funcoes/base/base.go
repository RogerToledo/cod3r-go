package main

import "fmt"

func main() {
	f1()
	f2("Hello", "World")
	fmt.Println(f3())
	fmt.Println(f4("Roger", "Toledo"))

	r3, r4 := f3(), f4("Parametro 1", "Parametro 2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	fmt.Printf("F5: %s, %s\n", r51, r52)

	r53, _ := f5()
	fmt.Println(r53)

	_, r54 := f5()
	fmt.Println(r54)
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s!\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "Parametro 1", "Parametro 2"
}
