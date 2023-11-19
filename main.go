package main

import (
	"fmt"
	"time"

	"github.com/distatus/battery"
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
func main() {
	// Get initial battery info
	initialBatteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get initial battery info!")
		return
	}
	initialBattery := initialBatteries[0] // Assuming single battery system
	timeBefore := time.Now()
	n := 30 // Досить велике число для демонстрації неефективності
	fmt.Printf("Фібоначчі(%d) = %d\n", n, fibRecursive(n))
	timeAfter := time.Now()
	fmt.Printf("Час виконання (повільний): %v\n", timeAfter.Sub(timeBefore))
	finalBatteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get final battery info!")
		return
	}
	finalBattery := finalBatteries[0]

	// Calculate battery usage
	usedCapacity := initialBattery.Current - finalBattery.Current
	fmt.Printf("Used Battery Capacity: %f mWh\n", usedCapacity)

}

// optimised.main.go
//func main() {
//	// Get initial battery info
//	initialBatteries, err := battery.GetAll()
//	if err != nil {
//		fmt.Println("Could not get initial battery info!")
//		return
//	}
//	initialBattery := initialBatteries[0] // Assuming single battery system
//
//	// Run Fibonacci function
//	n := 100
//	timeBefore := time.Now()
//	fibResult := fibDynamic(n)
//	timeAfter := time.Now()
//
//	// Get battery info after running Fibonacci function
//	finalBatteries, err := battery.GetAll()
//	if err != nil {
//		fmt.Println("Could not get final battery info!")
//		return
//	}
//	finalBattery := finalBatteries[0]
//
//	// Calculate battery usage
//	usedCapacity := initialBattery.Current - finalBattery.Current
//
//	// Output results
//	fmt.Printf("Фібоначчі оптимізоване(%d) = %d\n", n, fibResult)
//	fmt.Printf("Час виконання (оптимізований): %v\n", timeAfter.Sub(timeBefore))
//	fmt.Printf("Used Battery Capacity: %f mWh\n", usedCapacity)
//}
