package main

import (
	"errors"
	"math"
)

const (
	Epsilon = 10e-5
)

var (
	ErrNotQuadratic        = errors.New("a == 0")
	ErrInvalidCoefficients = errors.New("invalid coefficients")
)

type QuadEq struct {
	A     float64
	B     float64
	C     float64
	roots []float64
}

func (qe *QuadEq) solve() error {
	if qe.A < Epsilon {
		return ErrNotQuadratic
	}

	var discriminant float64
	discriminant = (qe.B * qe.B) - (4 * qe.A * qe.C)

	if discriminant > Epsilon {
		solution1 := -qe.B + math.Sqrt(discriminant)/(2*qe.A)
		solution2 := -qe.B - math.Sqrt(discriminant)/(2*qe.A)
		var roots []float64
		roots = append(qe.roots, solution1, solution2)
		qe.roots = roots
	} else if discriminant < Epsilon && discriminant > -Epsilon {
		solution := -qe.B / (2 * qe.A)
		var roots []float64
		roots = append(qe.roots, solution, solution)
		qe.roots = roots
	} else if discriminant < 0 {
	} else {
		return ErrInvalidCoefficients
	}
	return nil
}

func Solve(a, b, c float64) ([]float64, error) {
	task := QuadEq{A: a, B: b, C: c}
	err := task.solve()
	return task.roots, err
}
