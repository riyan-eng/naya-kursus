package main

import "fmt"

func main() {
	// aritmatika
	fmt.Println("--- aritmatika ---")
	var a int = 10
	var b int = 4
	fmt.Println("variabel a =", a)
	fmt.Println("variabel b =", b)
	fmt.Println("penjumlahan: a + b =", a+b)
	fmt.Println("pengurangan: a - b =", a-b)
	fmt.Println("perkalian: a * b =", a*b)
	fmt.Println("pembagian: a / b =", a/b)
	fmt.Println("modulus: a % b =", a%b)
	a++
	fmt.Println("increment: a++ =", a)
	b--
	fmt.Println("decrement: b-- =", b)

	fmt.Println()
	// assignment
	fmt.Println("--- assignment ---")
	var c int = 6
	fmt.Println("variabel =", c)
	c += 1
	fmt.Println("asignment: +=", c)
	c -= 2
	fmt.Println("asignment: -=", c)
	c *= 3
	fmt.Println("asignment: *=", c)
	c /= 4
	fmt.Println("asignment: /=", c)
	c %= 1
	fmt.Println("asignment: %=", c)

	fmt.Println()
	fmt.Println("--- perbandingan ---")
	var d int = 5
	var e int = 6
	fmt.Println("perbandingan: d == e =", d == e)
	fmt.Println("perbandingan: d != e =", d != e)
	fmt.Println("perbandingan: d > e =", d > e)
	fmt.Println("perbandingan: d < e =", d < e)
	fmt.Println("perbandingan: d >= e =", d >= e)
	fmt.Println("perbandingan: d <= e =", d <= e)

	fmt.Println()
	fmt.Println("--- logika ---")
	var f bool = true
	var g bool = false
	fmt.Println("logika: f && g =", f && g)
	fmt.Println("logika: f || g =", f || g)
	fmt.Println("logika: !f =", !f)
}
