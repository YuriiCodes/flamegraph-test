package main

import (
	"fmt"
	"time"
)

// fibRecursive - неоптимізована рекурсивна функція для обчислення чисел Фібоначчі
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibDynamic(n int) int {
	if n <= 1 {
		return n
	}
	fibCache := make([]int, n+1)
	fibCache[1] = 1
	for i := 2; i <= n; i++ {
		fibCache[i] = fibCache[i-1] + fibCache[i-2]
	}
	return fibCache[n]
}

// unoptimised.main.go
//func main() {
//	timeBefore := time.Now()
//	n := 40 // Досить велике число для демонстрації неефективності
//	fmt.Printf("Фібоначчі(%d) = %d\n", n, fibRecursive(n))
//	timeAfter := time.Now()
//	fmt.Printf("Час виконання (повільний): %v\n", timeAfter.Sub(timeBefore))
//}

// optimised.main.go
func main() {
	n := 100 // Досить велике число для демонстрації неефективності

	timeBefore := time.Now()
	fmt.Printf("Фібоначчі оптимізоване(%d) = %d\n", n, fibDynamic(n))
	timeAfter := time.Now()
	fmt.Printf("Час виконання (оптимізований): %v\n", timeAfter.Sub(timeBefore))
}
