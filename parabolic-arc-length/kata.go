package kata

import (
	"math"
)

type point struct {
	x, y float64
}

func LenCurve(n int) float64 {
	points := []point{
		{x: 0, y: 0},
	}
	path := 0.0
	for i := 1; i <= n; i++ {
		x := float64(i) / float64(n)
		y := x * x
		points = append(points, point{
			x: x,
			y: y,
		})
	}

	for i := 1; i < len(points); i++ {
		x := points[i].x - points[i-1].x
		y := points[i].y - points[i-1].y
		path += math.Hypot(x, y)
	}

	return path
}

func LenCurveWithHypot(n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += math.Hypot(float64(2*i+1), float64(n))
	}

	return sum / float64(n*n)
}
