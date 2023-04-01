// Package check主要提供了检测矩阵牌型是否为和牌的情况，并且暴露一定的api供外部使用
package check

import (
	"Richii_cal/models"
)

// sum 返回牌的总数
func sum(Pai models.PaiArr) int {
	total := 0
	for i := 0; i < 10; i++ {
		total += Pai.S[i]
		total += Pai.M[i]
		total += Pai.P[i]
		if i < 7 {
			total += Pai.Z[i]
		}
	}
	return total
}
