package coord

import (
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"math"
)

type point struct {
	X, Y, Z float64
}

var CalcedLinearPTermId *hashset.Set = hashset.New()

func CalcLinearCoord(twoNodeDistArray [][]float64, respectiveDistance [][]float64) (x, y, z float64, err error) {
	// todo 线性上的距离
	// 只保证至少有两个已知节点
	// 直接计算沿着线的距离

	//S.S.Logger.Infof("进行线性坐标位置计算中，直接按比例取已知两点的某个比例中间点")

	if len(twoNodeDistArray) != 2 || len(respectiveDistance) != 2 {
		return 0, 0, 0, errors.New("入参有误")
	}

	// 已知两点的坐标
	A := point{twoNodeDistArray[0][0], twoNodeDistArray[0][1], twoNodeDistArray[0][2]}
	B := point{twoNodeDistArray[1][0], twoNodeDistArray[1][1], twoNodeDistArray[1][2]}

	// 已知点P到A和B的距离
	d1 := respectiveDistance[0][0]
	d2 := respectiveDistance[1][0]

	PPrime := calcLinearCoord(A, d1, B, d2)

	//fmt.Printf("投影点P'的坐标: (%.2f, %.2f, %.2f)\n", PPrime.X, PPrime.Y, PPrime.Z)
	return PPrime.X, PPrime.Y, PPrime.Z, nil
}

func calcLinearCoord(A point, d1 float64, B point, d2 float64) point {
	// 计算AB向量
	AB := point{B.X - A.X, B.Y - A.Y, B.Z - A.Z}

	// 计算AB向量的长度
	lengthAB := distance(point{0, 0, 0}, AB)

	// 计算投影点P'到A的距离d1'
	d1Prime := (d1*d1 - d2*d2 + lengthAB*lengthAB) / (2 * lengthAB)

	// 计算单位化的AB向量
	AB_unit := point{AB.X / lengthAB, AB.Y / lengthAB, AB.Z / lengthAB}

	// 计算投影点P'的坐标
	PPrime := point{A.X + AB_unit.X*d1Prime, A.Y + AB_unit.Y*d1Prime, A.Z + AB_unit.Z*d1Prime}
	return PPrime
}

// 计算两点之间的距离
func distance(p1, p2 point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return sqrt(dx*dx + dy*dy + dz*dz)
}

// 计算平方根
func sqrt(x float64) float64 {
	return math.Sqrt(x)
}
