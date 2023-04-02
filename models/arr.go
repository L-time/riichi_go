package models

// MPSZ 牌面后缀
var MPSZ = [4]string{"m", "p", "s", "z"}

// HaiArr 是对牌的数组存储，剥离了Dora的存在
// 从上到下：万M、筒P、索S、字Z
// 合理的数据应该是每个位置只有0-4
type HaiArr [4][9]int

// GetSum 这个方法我不说了……
func (a HaiArr) GetSum() int {
	sum := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {
			sum += a[i][j]
		}
	}
	return sum
}
