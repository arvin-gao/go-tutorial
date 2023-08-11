package mytesting

import (
	"fmt"
	"math"
	"testing"
)

func TestRoundFloat(t *testing.T) {
	p_num := 5.4
	n_num := -5.4
	// 與 java 一樣算法；而 python3 是銀行家拾入法(Banker's rounding)
	// round: 往最接近數字方向的拾入操作
	// ceil: 	往正方向拾入
	// floor: 往負方向拾入
	fmt.Printf("%f\n", math.Round(p_num)) //  5.000000(Round)
	fmt.Printf("%f\n", math.Ceil(p_num))  //  6.000000(Ceil)
	fmt.Printf("%f\n", math.Floor(p_num)) //  5.000000(Floor)
	fmt.Printf("%f\n", math.Round(n_num)) // -5.000000(Round)
	fmt.Printf("%f\n", math.Ceil(n_num))  // -5.000000(Ceil)
	fmt.Printf("%f\n", math.Floor(n_num)) // -6.000000(Floor)
}
