// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build math_big_pure_go

package big
import "math/big"


func mulWW(x, y big.Word) (z1, z0 big.Word) {
	return mulWW_g(x, y)
}

func divWW(x1, x0, y big.Word) (q, r big.Word) {
	return divWW_g(x1, x0, y)
}

func addVV(z, x, y []big.Word) (c big.Word) {
	return addVV_g(z, x, y)
}

func subVV(z, x, y []big.Word) (c big.Word) {
	return subVV_g(z, x, y)
}

func addVW(z, x []big.Word, y big.Word) (c big.Word) {
	return addVW_g(z, x, y)
}

func subVW(z, x []big.Word, y big.Word) (c big.Word) {
	return subVW_g(z, x, y)
}

func shlVU(z, x []big.Word, s uint) (c big.Word) {
	return shlVU_g(z, x, s)
}

func shrVU(z, x []big.Word, s uint) (c big.Word) {
	return shrVU_g(z, x, s)
}

func mulAddVWW(z, x []big.Word, y, r big.Word) (c big.Word) {
	return mulAddVWW_g(z, x, y, r)
}

func addMulVVW(z, x []big.Word, y big.Word) (c big.Word) {
	return addMulVVW_g(z, x, y)
}

func divWVW(z []big.Word, xn big.Word, x []big.Word, y big.Word) (r big.Word) {
	return divWVW_g(z, xn, x, y)
}
