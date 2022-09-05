package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, 
// складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

// add two numbers a + b
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// sub two numbers a + b
func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

// multiply two numbers a * b
func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

// dvide two numbers a / b
func divide(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {

	bigIntA := big.NewInt(int64(1 << 29))
	bigIntB := big.NewInt(int64(1 << 22))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Adding is equal: ", add(bigIntA, bigIntB))

	fmt.Println("Subtracting is equal: ", sub(bigIntA, bigIntB))

	fmt.Println("Multiplying is equal: ", multiply(bigIntA, bigIntB))

	fmt.Println("Devision is equal: ", divide(bigIntA, bigIntB))

}
