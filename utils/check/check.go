// Package check主要提供了检测矩阵牌型是否为和牌的情况，并且暴露一定的api供外部使用
package check

import (
	"richii_go/models"
)

func sum(p [9]int) int {
	s := 0
	for i := 0; i < 9; i++ {
		s += p[i]
	}
	return s
}

// MPSZ 牌面后缀
var MPSZ = [4]string{"m", "p", "s", "z"}

// Is7D 对七对的判断
func Is7D(pai models.HaiArr) bool {
	sum := pai.GetSum()
	//七对一定14张
	if sum != 14 {
		return false
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			if pai[i][j] == 0 {
				continue
			}
			if pai[i][j] != 2 {
				return false
			}
		}
	}
	return true
}

// Is13G 对国士无双的判断
func Is13G(pai models.HaiArr) bool {
	arrs := []int{pai[0][0], pai[0][8], pai[1][0], pai[1][8], pai[2][0], pai[2][8]}
	arrs = append(arrs, pai[3][0:7]...)
	for _, arr := range arrs {
		if arr == 0 {
			return false
		}
	}
	return pai.GetSum() == 14
}

// isRowNormal 判断单种牌型是否合法
func isRowNormal(pai [9]int, isJiHai bool) bool {
	arr := pai
	s := sum(arr)
	if s == 0 {
		return true
	}
	//摘去雀头再进入判断
	if s%3 == 2 {
		for i := 0; i < 9; i++ {
			if arr[i] >= 2 {
				arr[i] -= 2
			} else {
				continue
			}
			if !isRowNormal(arr, isJiHai) {
				arr[i] += 2
			} else {
				return true
			}
		}
		return false
	}
	//检测面子
	for i := 0; i < 9; i++ {
		if arr[i] == 0 {
			continue
		} else if arr[i] == 3 {
			//去除刻子
			arr[i] = 0
			continue
		} else {
			//如果是字牌或者数到8了，那无论如何都不可能组成顺子
			if isJiHai || i >= 7 {
				return false
			}
			//去除如111123s牌型的影响，将其转换为111s 123s
			if arr[i] == 4 {
				arr[i] -= 3
			}
			arr[i+1] -= arr[i]
			arr[i+2] -= arr[i]
			//组不出来顺子就返回false
			if arr[i+1] < 0 || arr[i+2] < 0 {
				return false
			}
			arr[i] = 0
		}
	}
	return true
}

// IsNormal 对一般牌型的和牌判断
func IsNormal(pai models.HaiArr) bool {
	j := 0
	for i := 0; i < 4; i++ {
		//对于每一种牌，如果多出来一张，那么无论如何这张牌都无法被合并，因此绝对不可能是合法和牌型
		if sum(pai[i])%3 == 1 {
			return false
		}
		//一般牌型只可能有一个雀头，所以如果找到若干个，绝对不可能和牌
		if sum(pai[i])%3 == 2 {
			j++
		}
	}
	return j == 1 &&
		isRowNormal(pai[3], true) &&
		isRowNormal(pai[0], false) &&
		isRowNormal(pai[1], false) &&
		isRowNormal(pai[2], false)
}

// IsAll 对所有验证方法的封装
func IsAll(pai models.HaiArr) bool {
	return IsNormal(pai) || Is7D(pai) || Is13G(pai)
}
