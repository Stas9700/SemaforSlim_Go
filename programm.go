package main

import (
	"fmt"
	"time"
)

type calc func(a, b int)

var counter int = 0

func main() {
	SemaforTest()
	fmt.Scanln()
}

func SemaforTest() {
	semaforChanel := make(chan bool)
	for i := 0; i < 30; i++ {
		fmt.Print("Operation ")
		fmt.Print(i)
		fmt.Println(" go")
		go Start(i, semaforChanel, Calc, i, 1)
	}
	semaforChanel <- true
}

func Start(i int, chanel chan bool, funcToExecute calc, a, b int) {
	for {
		var res = <-chanel
		if res == true {
			chanel <- false
			fmt.Println("Cap -", len(chanel), " задача -", i, " начата")
			Calc(a, b)
			chanel <- true
			fmt.Println("Cap -", len(chanel), " задача -", i, " закончена")
			return
		}
		fmt.Println("задача -", i, "ожидает")
		time.Sleep(50 * time.Millisecond)
	}
}

func Calc(a, b int) {
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
