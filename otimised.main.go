package main

import (
	"fmt"
	"time"
)

func main_optimised() {
	timeBefore := time.Now()
	n := 30 // Досить велике число для демонстрації неефективності
	//fmt.Printf("Фібоначчі(%d) = %d\n", n, fibRecursive(n))
	timeAfter := time.Now()
	//fmt.Printf("Час виконання (повільний): %v\n", timeAfter.Sub(timeBefore))

	timeBefore = time.Now()
	fmt.Printf("Фібоначчі оптимізоване(%d) = %d\n", n, fibDynamic(n))
	timeAfter = time.Now()
	fmt.Printf("Час виконання (оптимізований): %v\n", timeAfter.Sub(timeBefore))
}
