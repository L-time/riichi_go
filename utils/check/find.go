package check

import (
	"richii_go/models"
	"strconv"
)

// find 主要定义了一些能返回牌型基本信息的函数
// 私有函数都是组件

// findKotSu 返回找到的刻子
func findKotSu(pai models.HaiArr) []string {
	var KotSu []string
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			if pai[i][j] >= 3 {
				pai[i][j] -= 3
			}
			//校验成功说明找到了一个新刻子
			if IsNormal(pai) {
				str := strconv.Itoa(j+1) + models.MPSZ[i]
				KotSu = append(KotSu, str)
			} else {
				pai[i][j] += 3
			}
		}
	}
	return KotSu
}

// findJyuntsu 返回找到的顺子
func findJyuntsu(pai models.HaiArr) [][]string {
	var Jyuntsu [][]string
	for i := 0; i < 4; i++ {
		if i == 3 {
			break
		}
		//只需要判断到j=6的情况就行，j>6后，无法再成顺子
		for j := 0; j < 7; j++ {
			for {
				//当j,j+1,j+2牌内有一张为0，则绝对不可能成为顺子
				if pai[i][j] < 1 || pai[i][j+1] < 1 || pai[i][j+2] < 1 {
					break
				}
				pai[i][j]--
				pai[i][j+1]--
				pai[i][j+2]--
				if IsNormal(pai) {
					str := []string{strconv.Itoa(j+1) + models.MPSZ[i], strconv.Itoa(j+2) + models.MPSZ[i], strconv.Itoa(j+3) + models.MPSZ[i]}
					Jyuntsu = append(Jyuntsu, str)
				} else {
					pai[i][j]++
					pai[i][j+1]++
					pai[i][j+2]++
					break
				}
			}
		}
	}
	return Jyuntsu
}

// findJyanto 返回雀头
// 由于雀头只能有一个（七对子特判），因此只返回一个string
func findJyanto(pai models.HaiArr) string {
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			if pai[i][j] >= 2 {
				return strconv.Itoa(j+1) + models.MPSZ[i]
			}
		}
	}
	//找不到就是空字符串了……但是绝对不可能出现这种情况啦
	return ""
}
