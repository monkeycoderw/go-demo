package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Printf("for if/else switch")

	/**
	* for 循环
	*
	*/
	// for 1
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for 2
	for j := 7; j <=9; j++ {
		fmt.Println(i)
	}

	// for 3
	for {
		fmt.Printf("loop")
		break;
	}

	/**

	* if/else
	*
	 */
	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println(" 8 is divisible by 4")
	}

	// 在条件语句之前可以有一个语句；
	// 任何在这里声明的变量, 都可以在所有的条件分支中使用。
	if num := 11; num < 0 {
		fmt.Println(num, "is negatives")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/**
	 switch
	 */
	value := 2
	fmt.Println("write ", value, "as")
	switch value {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 在一个 `case` 语句中，你可以使用逗号来分隔多个表达式。
	// 在这个例子中，我们很好的使用了可选的`default` 分支。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	vtime := time.Now()
	switch {
	case vtime.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

}