// Package check主要提供了检测矩阵牌型是否为和牌的情况，并且暴露一定的api供外部使用
package check

import (
	"richii_go/models"
)

// getSum 返回牌的总数
func getSum(pai models.PaiArr) int {
	total := 0
	for i := 0; i < 10; i++ {
		total += pai.S[i]
		total += pai.M[i]
		total += pai.P[i]
		if i < 7 {
			total += pai.Z[i]
		}
	}
	return total
}

// If7D 对七对的判断
func If7D(pai models.PaiArr) bool {
	//合并数组，方便处理
	arrs := append(append(append(pai.S[:], pai.M[:]...), pai.Z[:]...), pai.P[:]...)
	sum := 0
	for _, arr := range arrs {
		if arr%2 != 0 {
			return false
		}
		sum += arr
	}
	return sum == 14
}
