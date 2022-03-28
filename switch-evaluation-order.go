package main
// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
// 没有条件的 switch 同 switch true 一样。
// 这种形式能将一长串 if-then-else 写得更加清晰。

import (
	"fmt"
	"time"
)

func switchEvaluationOrder_study() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today+2:
		fmt.Println("Aftertomorrow")
	case today+1:
		fmt.Println("Tomorrow")
	case today+0:
		fmt.Println("Today")
	default:
		fmt.Println("far away")
	}

	switch  {
	case time.Now().Hour() < 12:
		fmt.Println("Now is noon")
	case time.Now().Hour() < 18:
		fmt.Println("Now is afternoon")
	default:
		fmt.Println("Now is night")
	}
}
