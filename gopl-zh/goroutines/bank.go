package main

import (
	"fmt"
	"log"
	"time"
)

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) {
	if b := Balance(); b > amount {
		withdraw <- amount
	} else {
		log.Fatalln("not enough money ...")
	}
}

func teller() {
	var balance int
	cnt := 0
	for i := 0; i < 10; i++ {
		select {
		case amount := <-deposits:
			balance += amount
		case da := <-withdraw:
			balance -= da
		case balances <- balance:
			cnt++
			fmt.Println("balances receive once ", cnt)
		}
	}
}

func init() {
	go teller()
}

func main() {
	fmt.Println(Balance())
	go Deposit(10)
	go Deposit(15)
	time.Sleep(time.Second)

	fmt.Println(Balance())
	Withdraw(20)
	fmt.Println(Balance())
	Withdraw(30)
}
