// Copyright 2012 Dorival de Moraes Pedroso. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package utl implements functions for simplifying calculations and allocation of structures
// such as slices and slices of slices. It also contains functions for sorting quantities.
package utl

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/cpmech/gosl/io"
)

// IntFill fills a slice of integers
func IntFill(s []int, val int) {
	for i := 0; i < len(s); i++ {
		s[i] = val
	}
}

// IntVals allocates a slice of integers with size==n, filled with val
func IntVals(n int, val int) (s []int) {
	s = make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = val
	}
	return
}

// StrVals allocates a slice of strings with size==n, filled with val
func StrVals(n int, val string) (s []string) {
	s = make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = val
	}
	return
}

// DblsAlloc allocates a matrix of float64
func DblsAlloc(m, n int) (mat [][]float64) {
	mat = make([][]float64, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]float64, n)
	}
	return
}

// StrsAlloc allocates a matrix of strings
func StrsAlloc(m, n int) (mat [][]string) {
	mat = make([][]string, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]string, n)
	}
	return
}

// IntsAlloc allocates a matrix of integers
func IntsAlloc(m, n int) (mat [][]int) {
	mat = make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	return
}

// IntRange generates a slice of integers from 0 to n-1
func IntRange(n int) (res []int) {
	if n <= 0 {
		return
	}
	res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return
}

// IntRange2 generates slice of integers from start to stop (but not stop)
func IntRange2(start, stop int) []int {
	return IntRange3(start, stop, 1)
}

// IntRange3 generates a slice of integers from start to stop (but not stop), afer each 'step'
func IntRange3(start, stop, step int) (res []int) {
	switch {
	case stop == start:
		return
	case stop > start:
		n := (stop - start) / step
		res = make([]int, n)
		for i, v := 0, start; v < stop; i, v = i+1, v+step {
			res[i] = v
		}
	case stop < start:
		if step > 0 {
			return
		}
		n := (stop - start) / step
		res = make([]int, n)
		for i, v := 0, start; v > stop; i, v = i+1, v+step {
			res[i] = v
		}
	}
	return
}

// IntAddScalar adds a scalar to all values in a slice of integers
func IntAddScalar(a []int, s int) (res []int) {
	res = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] + s
	}
	return
}

// IntUnique returns a unique and sorted slice of integers
func IntUnique(slices ...[]int) (res []int) {
	if len(slices) == 0 {
		return
	}
	nn := 0
	for i := 0; i < len(slices); i++ {
		nn += len(slices[i])
	}
	res = make([]int, 0, nn)
	for i := 0; i < len(slices); i++ {
		a := make([]int, len(slices[i]))
		copy(a, slices[i])
		sort.Ints(a)
		for j := 0; j < len(a); j++ {
			idx := sort.SearchInts(res, a[j])
			if idx < len(res) && res[idx] == a[j] {
				continue // found
			} else {
				if idx == len(res) { // append
					res = append(res, a[j])
				} else { // insert
					res = append(res[:idx], append([]int{a[j]}, res[idx:]...)...)
				}
			}
		}
	}
	return
}

// IntPy returns a Python string representing a slice of integers
func IntPy(a []int) (res string) {
	res = "["
	for i := 0; i < len(a); i++ {
		res += strconv.Itoa(a[i])
		if i < len(a)-1 {
			res += ", "
		}
	}
	res += "]"
	return
}

// DblOnes generates a slice of double precision '1s'
func DblOnes(n int) (res []float64) {
	res = make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = 1.0
	}
	return
}

// DblVals generates a slice of double precision values
func DblVals(n int, v float64) (res []float64) {
	res = make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = v
	}
	return
}

// DblCopy gets a copy of slice of doubles
func DblCopy(in []float64) (out []float64) {
	out = make([]float64, len(in))
	copy(out, in)
	return
}

// LinSpace returns evenly spaced numbers over a specified closed interval.
func LinSpace(start, stop float64, num int) (res []float64) {
	if num <= 0 {
		return []float64{}
	}
	if num == 1 {
		return []float64{start}
	}
	step := (stop - start) / float64(num-1)
	res = make([]float64, num)
	for i := 0; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	res[num-1] = stop
	return
}

// LinSpaceOpen returns evenly spaced numbers over a specified open interval.
func LinSpaceOpen(start, stop float64, num int) (res []float64) {
	if num <= 0 {
		return []float64{}
	}
	step := (stop - start) / float64(num)
	res = make([]float64, num)
	for i := 0; i < num; i++ {
		res[i] = start + float64(i)*step
	}
	return
}

// Dbl2Str converts a slice of doubles (float64) to a slice of strings
func Dbl2Str(v []float64, format string) (s []string) {
	s = make([]string, len(v))
	for i := 0; i < len(v); i++ {
		s[i] = io.Sf(format, v[i])
	}
	return
}

// Str2Dbl converts a slice of strings to a slice of doubles (float64)
func Str2Dbl(s []string) (v []float64) {
	v = make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		v[i] = io.Atof(s[i])
	}
	return
}

// DblSplit splits a string into floats
func DblSplit(s string) (r []float64) {
	ss := strings.Fields(s)
	r = make([]float64, len(ss))
	for i, v := range ss {
		r[i] = io.Atof(v)
	}
	return
}

// Digits returns the nubmer of digits
func Digits(maxint int) (ndigits int, format string) {
	ndigits = int(math.Log10(float64(maxint))) + 1
	format = io.Sf("%%%dd", ndigits)
	return
}

// Expon returns the exponent
func Expon(val float64) (ndigits int) {
	if val == 0.0 {
		return
	}
	ndigits = int(math.Log10(math.Abs(val)))
	return
}

// MeshGrid2D creates a grid with x-y coordinates
//  x -- [ny][nx]
//  y -- [ny][nx]
func MeshGrid2D(xmin, xmax, ymin, ymax float64, nx, ny int) (x, y [][]float64) {
	if nx < 2 {
		return
	}
	if ny < 2 {
		return
	}
	dx := (xmax - xmin) / float64(nx-1)
	dy := (ymax - ymin) / float64(ny-1)
	x = DblsAlloc(ny, nx)
	y = DblsAlloc(ny, nx)
	for i := 0; i < ny; i++ {
		for j := 0; j < nx; j++ {
			x[i][j] = xmin + float64(j)*dx
			y[i][j] = ymin + float64(i)*dy
		}
	}
	return
}
