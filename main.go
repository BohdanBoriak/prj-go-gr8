package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalPoints = 50

func main() {
	fmt.Println("Вітаємо у грі MATHCORE!")
	time.Sleep(1 * time.Second)

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	start := time.Now()
	myPoints := 0
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		res := x + y
		fmt.Println(x, "+", y, "=")
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)
		if ans == res {

		}
	}
}
