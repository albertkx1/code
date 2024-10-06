package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f_(info [][]int, x, y int) int {
	ans := -1
	for i, v := range info {
		x1 := v[0] + v[2]
		y1 := v[1] + v[3]
		if x <= x1 && x >= v[0] && y >= v[1] && y <= y1 {
			ans = i + 1
		}
	}
	return ans
}
func main() {
	var n, x, y int
	fmt.Scanln(&n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			fmt.Scan(&matrix[i][j])

		}
	}
	fmt.Scanf("%d %d", &x, &y)

	fmt.Println(f_(matrix, x, y))
}
func maxProfit(prices []int) int {
	ans := 0
	var min1, max1 int
	min1 = prices[0]
	flag := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			if flag == 1 {
				ans += max1 - min1
			}

			min1 = prices[i]
			fmt.Println("min", min1)
		} else {
			flag = 1
			max1 = prices[i]
			fmt.Println("max", max1)
			if i == len(prices)-1 {
				ans += max1 - min1
			}
		}

	}
	return ans
}

//	func canJump(nums []int) bool {
//		ans := nums[0]
//
//		for i := 0; i < len(nums); {
//			max1:=0
//			for i<nums[k]
//		}
//
//		return ans > len(nums)
//	}
func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	var reverse func([]int)
	reverse = func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}
func discountPrices(sentence string, discount int) string {
	split := strings.Split(sentence, " ")
	dis := 1 - 0.01*float64(discount)
	str := ""
	for i := 0; i < len(split); i++ {
		if len(split[i]) > 0 && split[i][0] == '$' {
			cnt := 0.00
			for j := 1; j < len(split[i]); j++ {
				if split[i][j] > '0' && split[i][j] <= '9' {
					cnt = cnt*10 + float64(int(split[i][j]-'0'))
				} else {
					cnt = -1
					break
				}
			}

			if cnt == -1 {
				str += split[i]
			} else {
				cnt *= dis
				str += "$" + strconv.FormatFloat(dis, 'f', 2, 64)
			}
		}
		if i < len(split)-1 {
			str += " "
		}

	}
	return str

}
func removeElement(nums []int, val int) int {
	cnt := 0

	for i := 0; i+cnt <= len(nums); {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1-cnt]

			cnt++
		} else {
			i++
		}

	}
	return cnt
}
func removeTrailingZeros(num string) string {
	i := len(num) - 1
	for ; i >= 0; i-- {
		if num[i] != '0' {
			break
		}
	}
	return num[0:i]
}
