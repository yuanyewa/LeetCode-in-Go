package problem0909

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]int
	ans   int
}{

	{
		[][]int{
			{1, 1, -1},
			{1, 1, 1},
			{-1, 1, 1},
		},
		-1,
	},

	{
		[][]int{
			{1, -1, -1, -1, 48, 5, -1},
			{12, 29, 13, 9, -1, 2, 32},
			{-1, -1, 21, 7, -1, 12, 49},
			{42, 37, 21, 40, -1, 22, 12},
			{42, -1, 2, -1, -1, -1, 6},
			{39, -1, 35, -1, -1, 39, -1},
			{-1, 36, -1, -1, -1, -1, 5},
		},
		3,
	},

	{
		[][]int{
			{-1, -1, 27, 13, -1, 25, -1},
			{-1, -1, -1, -1, -1, -1, -1},
			{44, -1, 8, -1, -1, 2, -1},
			{-1, 30, -1, -1, -1, -1, -1},
			{3, -1, 20, -1, 46, 6, -1},
			{-1, -1, -1, -1, -1, -1, 29},
			{-1, 29, 21, 33, -1, -1, -1},
		},
		4,
	},

	{
		[][]int{
			{-1, 1, 2, -1},
			{2, 13, 15, -1},
			{-1, 10, -1, -1},
			{-1, 6, 2, 8},
		},
		2,
	},

	{
		[][]int{
			{-1, -1, -1, -1, 162, -1, -1, -1, -1, -1, -1, -1, 256, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, 396, -1, 184, 120, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, 384, -1, -1, -1, -1, -1, 76, -1, -1, -1, 196, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, 152, -1, -1, -1, -1, -1, 40, -1, -1, -1, 30, 159},
			{-1, 86, -1, -1, 25, 307, 280, -1, -1, -1, -1, -1, -1, -1, 6, -1, -1, -1, 330, -1},
			{-1, -1, -1, -1, -1, 266, -1, -1, -1, -1, -1, -1, -1, -1, -1, 213, 57, 36, -1, 399},
			{-1, -1, -1, -1, -1, -1, -1, -1, 149, -1, -1, 27, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, 366, -1, -1, -1, -1, -1, 140, 247, -1, -1, -1, -1, 396, -1, -1, 347, -1, -1, -1},
			{-1, -1, 73, -1, -1, -1, 120, 178, -1, -1, -1, -1, -1, -1, -1, 24, -1, -1, 22, -1},
			{-1, 394, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, 208, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 32, -1, -1, -1, 370, 345, -1},
			{-1, -1, -1, -1, 356, -1, 28, -1, -1, -1, -1, -1, -1, 102, -1, 312, -1, 242, -1, -1},
			{-1, -1, -1, 235, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, 81, -1, -1, 177, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, 306, -1, -1, -1, -1, -1, 62, -1, -1, -1, -1, -1, 236, 395},
			{-1, -1, -1, -1, -1, -1, -1, 291, -1, -1, -1, -1, -1, -1, 40, 387, -1, -1, -1, -1},
			{-1, -1, -1, 151, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 222, 230, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 186, -1, -1, -1, -1, -1, 236, -1, -1, -1},
		},
		6,
	},

	{
		[][]int{
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 35, -1, -1, 13, -1},
			{-1, -1, -1, -1, -1, -1},
			{-1, 15, -1, -1, -1, -1}},
		4,
	},

	// 可以有多个 testcase
}

func Test_snakesAndLadders(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, snakesAndLadders(tc.board), "输入:%v", tc)
	}
}

func Benchmark_snakesAndLadders(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			snakesAndLadders(tc.board)
		}
	}
}
func Test_location(t *testing.T) {
	ast := assert.New(t)
	//
	const (
		n = 3
	)
	//
	expected := [n][n]int{
		{7, 8, 9},
		{6, 5, 4},
		{1, 2, 3},
	}
	actual := [n][n]int{}
	for i := 1; i <= 9; i++ {
		x, y := location(n, i)
		actual[x][y] = i
	}
	ast.Equal(expected, actual)
}
