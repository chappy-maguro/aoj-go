// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var totalSeconds int
// 	fmt.Scan(&totalSeconds)
// 	// 秒から time.Duration を作成
// 	d := time.Duration(totalSeconds) * time.Second

// 	// 時間、分、秒を計算
// 	h := int(d.Hours())
// 	m := int(d.Minutes()) % 60
// 	s := int(d.Seconds()) % 60

//		// h:m:s 形式でフォーマット
//		result := fmt.Sprintf("%d:%02d:%02d\n", h, m, s)
//		fmt.Println(result) // 出力例: 1:01:20
//	}
package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	h := s / 3600
	m := (s % 3600) / 60
	sec := s % 60

	// 0埋めせず、普通の %d でコロン区切りにし、最後に改行を入れる
	fmt.Printf("%d:%d:%d\n", h, m, sec)
}
