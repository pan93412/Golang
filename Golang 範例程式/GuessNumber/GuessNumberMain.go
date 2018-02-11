package main

import (
	"fmt"
	"time"
)

var errcnt int = 0
var randomInt int = random(1, 10)

func main() {
	game()
}

func again() {
	time.Sleep(5 * time.Second)
	errcnt++
	game()
}

func game() {
	var number int
	fmt.Print("輸入 1~10 以內的數字： ")
	fmt.Scan(&number)
	switch {
	case randomInt == number:
		fmt.Println("你答對了！ :DDDDDDDD")
		fmt.Printf("錯誤次數： %v 次！\n", errcnt)
		fmt.Println("正確答案是", randomInt)
	case number > 10:
		fmt.Println("你在輸入三小？？？")
		time.Sleep(10 * time.Second)
		again()
	case number < 0:
		fmt.Println("你在輸入三小？？？")
		time.Sleep(10 * time.Second)
		again()
	case randomInt > number:
		fmt.Println("你輸入的數字太小了！ :(((\n重新開始！")
		again()
	case randomInt < number:
		fmt.Println("你輸入的數字太小了！ :(((\n重新開始！")
		again()
	default:
		panic("")
	}
}
