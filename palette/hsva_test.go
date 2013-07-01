// Copyright ©2011-2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette

import (
	"image/color"
	check "launchpad.net/gocheck"
)

var e uint32 = 1

// Checkers
type float32WithinRange struct {
	*check.CheckerInfo
}

var floatWithinRange check.Checker = &float32WithinRange{
	&check.CheckerInfo{Name: "WithinRange", Params: []string{"obtained", "min", "max"}},
}

func (checker *float32WithinRange) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0].(float64) >= params[1].(float64) && params[0].(float64) <= params[2].(float64), ""
}

type uint32WithinRange struct {
	*check.CheckerInfo
}

var uintWithinRange check.Checker = &uint32WithinRange{
	&check.CheckerInfo{Name: "WithinRange", Params: []string{"obtained", "min", "max"}},
}

func (checker *uint32WithinRange) Check(params []interface{}, names []string) (result bool, error string) {
	return params[0].(uint32) >= params[1].(uint32) && params[0].(uint32) <= params[2].(uint32), ""
}

type uint32EpsilonChecker struct {
	*check.CheckerInfo
}

var withinEpsilon check.Checker = &uint32EpsilonChecker{
	&check.CheckerInfo{Name: "EpsilonLessThan", Params: []string{"obtained", "expected", "error"}},
}

func (checker *uint32EpsilonChecker) Check(params []interface{}, names []string) (result bool, error string) {
	d := int64(params[0].(uint32)) - int64(params[1].(uint32))
	if d < 0 {
		if d == -d {
			panic("color: weird number overflow")
		}
		d = -d
	}
	return uint32(d) <= params[2].(uint32), ""
}

// Tests

func (s *S) TestColor(c *check.C) {
	for r := 0; r < 256; r += 5 {
		for g := 0; g < 256; g += 5 {
			for b := 0; b < 256; b += 5 {
				col := color.RGBA{uint8(r), uint8(g), uint8(b), 0}
				cDirectR, cDirectG, cDirectB, cDirectA := col.RGBA()
				hsva := rgbaToHsva(col.RGBA())
				c.Check(hsva.H, floatWithinRange, float64(0), float64(1))
				c.Check(hsva.S, floatWithinRange, float64(0), float64(1))
				c.Check(hsva.V, floatWithinRange, float64(0), float64(1))
				cFromHSVR, cFromHSVG, cFromHSVB, cFromHSVA := hsva.RGBA()
				c.Check(cFromHSVR, uintWithinRange, uint32(0), uint32(0xFFFF))
				c.Check(cFromHSVG, uintWithinRange, uint32(0), uint32(0xFFFF))
				c.Check(cFromHSVB, uintWithinRange, uint32(0), uint32(0xFFFF))
				back := rgbaToHsva(color.RGBA{uint8(cFromHSVR >> 8), uint8(cFromHSVG >> 8), uint8(cFromHSVB >> 8), uint8(cFromHSVA)}.RGBA())
				c.Check(hsva, check.Equals, back)
				c.Check(cFromHSVR, withinEpsilon, cDirectR, e)
				c.Check(cFromHSVG, withinEpsilon, cDirectG, e)
				c.Check(cFromHSVB, withinEpsilon, cDirectB, e)
				c.Check(cFromHSVA, check.Equals, cDirectA)
				if c.Failed() {
					return
				}
			}
		}
	}
}
