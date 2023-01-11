package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name    string
		handler func(*testing.T)
	}{
		{
			name: "no roots",
			handler: func(t *testing.T) {
				var expected []float64
				var a float64 = 1
				var b float64 = 0
				var c float64 = 1
				roots, err := Solve(a, b, c)
				require.Equal(t, expected, roots)
				require.NoError(t, err)
			},
		},
		{
			name: "[-1, 1] roots",
			handler: func(t *testing.T) {
				var a float64 = 1
				var b float64 = 0
				var c float64 = -1
				result, err := Solve(a, b, c)
				fmt.Println(result)
				require.Equal(t, result[0], -result[1])
				require.NoError(t, err)
			},
		},
		{
			name: "equal roots",
			handler: func(t *testing.T) {
				var a float64 = 0.0001
				var b float64 = 0.001
				var c float64 = 0.000001
				result, err := Solve(a, b, c)
				require.Equal(t, result[1], result[0])
				require.NoError(t, err)
			},
		},
		{
			name: "a < epsilon",
			handler: func(t *testing.T) {
				var a float64 = 10e-6
				var b float64 = 0
				var c float64 = 0
				result, err := Solve(a, b, c)
				require.EqualError(t, err, ErrNotQuadratic.Error())
				require.Nil(t, result)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.handler(t)
		})
	}
}
