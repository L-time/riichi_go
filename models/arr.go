package models

// HaiArr 是对牌的数组存储，剥离了Dora的存在
// 从上到下：万M、筒P、索S、字Z
type HaiArr [4][9]int

func (a HaiArr) GetSum() int {
	sum := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			sum += a[i][j]
		}
	}
	return sum
}
