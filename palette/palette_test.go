// Copyright ©2011-2012 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette

import (
	"image/color"
	"testing"

	"gopkg.in/check.v1"
)

// Tests
func Test(t *testing.T) { check.TestingT(t) }

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestRainbow(c *check.C) {
	c.Check(Rainbow(10, 0, 1, 1, 1, 1), check.DeepEquals, palette{
		color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // "#FF0000FF"
		color.NRGBA{R: 0xff, G: 0xaa, B: 0x00, A: 0xff}, // "#FFAA00FF"
		color.NRGBA{R: 0xaa, G: 0xff, B: 0x00, A: 0xff}, // "#AAFF00FF"
		color.NRGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // "#00FF00FF"
		color.NRGBA{R: 0x00, G: 0xff, B: 0xaa, A: 0xff}, // "#00FFAAFF"
		color.NRGBA{R: 0x00, G: 0xaa, B: 0xff, A: 0xff}, // "#00AAFFFF"
		color.NRGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // "#0000FFFF"
		color.NRGBA{R: 0xaa, G: 0x00, B: 0xff, A: 0xff}, // "#AA00FFFF"
		color.NRGBA{R: 0xff, G: 0x00, B: 0xaa, A: 0xff}, // "#FF00AAFF"
		color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // "#FF0000FF"
	})
}

func (s *S) TestHeat(c *check.C) {
	c.Check(Heat(10, 1), check.DeepEquals, palette{
		color.NRGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // "#FF0000FF"
		color.NRGBA{R: 0xff, G: 0x24, B: 0x00, A: 0xff}, // "#FF2400FF"
		color.NRGBA{R: 0xff, G: 0x49, B: 0x00, A: 0xff}, // "#FF4900FF"
		color.NRGBA{R: 0xff, G: 0x6d, B: 0x00, A: 0xff}, // "#FF6D00FF"
		color.NRGBA{R: 0xff, G: 0x92, B: 0x00, A: 0xff}, // "#FF9200FF"
		color.NRGBA{R: 0xff, G: 0xb6, B: 0x00, A: 0xff}, // "#FFB600FF"
		color.NRGBA{R: 0xff, G: 0xdb, B: 0x00, A: 0xff}, // "#FFDB00FF"
		color.NRGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}, // "#FFFF00FF"
		color.NRGBA{R: 0xff, G: 0xff, B: 0x3f, A: 0xff}, // "#FFFF40FF" Off by one compared to R.
		color.NRGBA{R: 0xff, G: 0xff, B: 0xbF, A: 0xff}, // "#FFFFBFFF"
	})
}

func (s *S) TestRadial(c *check.C) {
	c.Check(Radial(10, Cyan, Magenta, 1), check.DeepEquals, palette{
		color.NRGBA{R: 0x7f, G: 0xff, B: 0xff, A: 0xff}, // "#80FFFFFF" Off by one compared to R.
		color.NRGBA{R: 0x99, G: 0xff, B: 0xff, A: 0xff}, // "#99FFFFFF"
		color.NRGBA{R: 0xb3, G: 0xff, B: 0xff, A: 0xff}, // "#B3FFFFFF"
		color.NRGBA{R: 0xcc, G: 0xff, B: 0xff, A: 0xff}, // "#CCFFFFFF"
		color.NRGBA{R: 0xe6, G: 0xff, B: 0xff, A: 0xff}, // "#E6FFFFFF"
		color.NRGBA{R: 0xff, G: 0xe6, B: 0xff, A: 0xff}, // "#FFE6FFFF"
		color.NRGBA{R: 0xff, G: 0xcc, B: 0xff, A: 0xff}, // "#FFCCFFFF"
		color.NRGBA{R: 0xff, G: 0xb3, B: 0xff, A: 0xff}, // "#FFB3FFFF"
		color.NRGBA{R: 0xff, G: 0x99, B: 0xff, A: 0xff}, // "#FF99FFFF"
		color.NRGBA{R: 0xff, G: 0x7f, B: 0xff, A: 0xff}, // "#FF80FFFF" Off by one compared to R.
	})
}
