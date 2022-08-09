package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	cnt := 0
	bool := 0
	for bool < 1 {
		for i, v := range a {
			// 2で割り切れない場合、ループから抜ける
			if v%2 != 0 {
				bool++
				break
			}
			// 割り算の結果を再代入
			a[i] = v / 2

			// 最後まで2で割り切れたら、カウントアップ
			if i == n-1 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
