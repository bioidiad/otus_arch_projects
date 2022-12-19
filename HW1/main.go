package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

type QuadEq struct {
	A        float64
	B        float64
	C        float64
	Solution struct {
		isSolvable bool
		roots      []float64
	}
}

func (qe *QuadEq) solve() {
	a, b, c := decimal.NewFromFloat(qe.A), decimal.NewFromFloat(qe.B), decimal.NewFromFloat(qe.C)
	d := b.Mul(b).Sub(a.Mul(c.Mul(decimal.NewFromFloat(4))))
	if d.LessThan(decimal.NewFromFloat(0)) {
		qe.Solution.isSolvable = false
	} else if d.Equals(decimal.NewFromFloat(0)) {
		tmp := b.Div(a.Mul(decimal.NewFromFloat(-2)))
		res, _ := tmp.Float64()
		qe.Solution.isSolvable = true
		qe.Solution.roots = append(qe.Solution.roots, res)
	} else {
		floatD, _ := d.Float64()
		decSqRoot := decimal.NewFromFloat(math.Sqrt(floatD))
		solution1, _ := (b.Mul(decimal.NewFromFloat(-1)).Add(decSqRoot).Div(a.Mul(decimal.NewFromFloat(2)))).Float64()
		solution2, _ := (b.Mul(decimal.NewFromFloat(-1)).Sub(decSqRoot).Div(a.Mul(decimal.NewFromFloat(2)))).Float64()
		qe.Solution.isSolvable = true
		qe.Solution.roots = append(qe.Solution.roots, solution1, solution2)
	}
}

func main() {
	task := QuadEq{A: -4, B: 28, C: -49}
	task.solve()
	if task.Solution.isSolvable {
		fmt.Println(task.Solution.roots)
	} else {
		fmt.Println("No solution")
	}
}

func solveQuadratic(a float64, b float64, c float64) (float64, float64) {
	// calculate the discriminant
	discriminant := b*b - 4*a*c

	// calculate the two solutions
	solution1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	solution2 := (-b - math.Sqrt(discriminant)) / (2 * a)

	return solution1, solution2
}
