// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bezier implements 2D Bézier curve calculation.
package bezier

// Point represents a 2 dimensional value.
type Point struct {
	X, Y float64
}

type point struct {
	Point, Control Point
}

// Curve implements Bezier curve calculation according to the algorithm of Robert D. Miller.
//
// Graphics Gems 5, 'Quick and Simple Bézier Curve Drawing', pages 206-209.
type Curve []point

// NewCurve returns a Curve initialized with the control points in cp.
func New(cp ...Point) Curve {
	if len(cp) == 0 {
		return nil
	}
	c := make(Curve, len(cp))
	for i, p := range cp {
		c[i].Point = p
	}

	var w float64
	for i, p := range c {
		if i == 0 {
			w = 1
		} else if i == 1 {
			w = float64(len(c)) - 1
		} else {
			w *= float64(len(c)-i) / float64(i)
		}
		c[i].Control.X = p.Point.X * w
		c[i].Control.Y = p.Point.Y * w
	}

	return c
}

// Point returns the point at t along the curve, where 0 ≤ t ≤ 1.
func (c Curve) Point(t float64) Point {
	c[0].Point = c[0].Control
	u := t
	for i, p := range c[1:] {
		c[i+1].Point = Point{p.Control.X * u, p.Control.Y * u}
		u *= t
	}

	var (
		t1 = 1 - t
		tt = t1
	)
	p := c[len(c)-1].Point
	for i := len(c) - 2; i >= 0; i-- {
		p.X += c[i].Point.X * tt
		p.Y += c[i].Point.Y * tt
		tt *= t1
	}

	return p
}

// Curve returns a slice of Point, p, filled with equally spaced points along the Bézier curve
// described by c. If the length of p is less than 2, the curve points are undefined. The length of
// p is not altered by the call.
func (c Curve) Curve(p []Point) []Point {
	for i, nf := 0, float64(len(p)-1); i < len(p); i++ {
		p[i] = c.Point(float64(i) / nf)
	}
	return p
}
