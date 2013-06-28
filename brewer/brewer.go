// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Pallete type comments ©2002 Cynthia Brewer.

// Package brewer provides Brewer Palletes for informative graphics.
//
// The colors defined here are from http://www.ColorBrewer.org/ by Cynthia A. Brewer,
// Geography, Pennsylvania State University.
//
// For more information see:
// http://www.personal.psu.edu/cab38/ColorBrewer/ColorBrewer_learnMore.html
//
package brewer

import (
	"image/color"
)

// Color represents a Brewer Pallete color.
type Color struct {
	Letter byte
	color.Color
}

// DivergingPallete represents a diverging color scheme.
type DivergingPallete []Color

// CriticalValue returns the indexish of the lightest (median) color in the DivergingPallete.
func (d DivergingPallete) CriticalValue() float64 { return float64(len(d)+1)/2 - 1 }

// DivergingPallete represents a sequential or qualitative color schemes.
type Pallete []Color

// Diverging schemes put equal emphasis on mid-range critical values and extremes
// at both ends of the data range. The critical class or break in the middle of the
// legend is emphasized with light colors and low and high extremes are emphasized
// with dark colors that have contrasting hues.
type Diverging map[int]DivergingPallete

// Qualitative schemes do not imply magnitude differences between legend classes,
// and hues are used to create the primary visual differences between classes.
// Qualitative schemes are best suited to representing nominal or categorical data.
type Qualitative map[int]Pallete

// Sequential schemes are suited to ordered data that progress from low to high.
// Lightness steps dominate the look of these schemes, with light colors for low
// data values to dark colors for high data values.
type Sequential map[int]Pallete
