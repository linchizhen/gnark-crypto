// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goff (v0.3.3) DO NOT EDIT

// Package fr contains field arithmetic operations for modulus 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177
package fr

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"math/bits"
	"strconv"
	"sync"
)

// Element represents a field element stored on 6 words (uint64)
// Element are assumed to be in Montgomery form in all methods
// field modulus q =
//
// 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177
type Element [6]uint64

// Limbs number of 64 bits words needed to represent Element
const Limbs = 6

// Bits number bits needed to represent Element
const Bits = 377

// field modulus stored as big.Int
var _modulus big.Int
var onceModulus sync.Once

// Modulus returns q as a big.Int
// q =
//
// 258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177
func Modulus() *big.Int {
	onceModulus.Do(func() {
		_modulus.SetString("258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177", 10)
	})
	return new(big.Int).Set(&_modulus)
}

// q (modulus)
var qElement = Element{
	9586122913090633729,
	1660523435060625408,
	2230234197602682880,
	1883307231910630287,
	14284016967150029115,
	121098312706494698,
}

// rSquare
var rSquare = Element{
	13224372171368877346,
	227991066186625457,
	2496666625421784173,
	13825906835078366124,
	9475172226622360569,
	30958721782860680,
}

// Bytes returns the regular (non montgomery) value
// of z as a big-endian byte slice.
func (z *Element) Bytes() []byte {
	_z := z.ToRegular()
	var res [Limbs * 8]byte
	binary.BigEndian.PutUint64(res[40:48], _z[0])
	binary.BigEndian.PutUint64(res[32:40], _z[1])
	binary.BigEndian.PutUint64(res[24:32], _z[2])
	binary.BigEndian.PutUint64(res[16:24], _z[3])
	binary.BigEndian.PutUint64(res[8:16], _z[4])
	binary.BigEndian.PutUint64(res[0:8], _z[5])

	return res[:]
}

// SetBytes interprets e as the bytes of a big-endian unsigned integer,
// sets z to that value (in Montgomery form), and returns z.
func (z *Element) SetBytes(e []byte) *Element {
	var tmp big.Int
	tmp.SetBytes(e)
	z.SetBigInt(&tmp)
	return z
}

// SetUint64 z = v, sets z LSB to v (non-Montgomery form) and convert z to Montgomery form
func (z *Element) SetUint64(v uint64) *Element {
	*z = Element{v}
	return z.Mul(z, &rSquare) // z.ToMont()
}

// Set z = x
func (z *Element) Set(x *Element) *Element {
	z[0] = x[0]
	z[1] = x[1]
	z[2] = x[2]
	z[3] = x[3]
	z[4] = x[4]
	z[5] = x[5]
	return z
}

// SetInterface converts i1 from uint64, int, string, or Element, big.Int into Element
// panic if provided type is not supported
func (z *Element) SetInterface(i1 interface{}) *Element {
	switch c1 := i1.(type) {
	case Element:
		return z.Set(&c1)
	case *Element:
		return z.Set(c1)
	case uint64:
		return z.SetUint64(c1)
	case int:
		return z.SetString(strconv.Itoa(c1))
	case string:
		return z.SetString(c1)
	case *big.Int:
		return z.SetBigInt(c1)
	case big.Int:
		return z.SetBigInt(&c1)
	case []byte:
		return z.SetBytes(c1)
	default:
		panic("invalid type")
	}
}

// SetZero z = 0
func (z *Element) SetZero() *Element {
	z[0] = 0
	z[1] = 0
	z[2] = 0
	z[3] = 0
	z[4] = 0
	z[5] = 0
	return z
}

// SetOne z = 1 (in Montgomery form)
func (z *Element) SetOne() *Element {
	z[0] = 202099033278250856
	z[1] = 5854854902718660529
	z[2] = 11492539364873682930
	z[3] = 8885205928937022213
	z[4] = 5545221690922665192
	z[5] = 39800542322357402
	return z
}

// Div z = x*y^-1 mod q
func (z *Element) Div(x, y *Element) *Element {
	var yInv Element
	yInv.Inverse(y)
	z.Mul(x, &yInv)
	return z
}

// Equal returns z == x
func (z *Element) Equal(x *Element) bool {
	return (z[5] == x[5]) && (z[4] == x[4]) && (z[3] == x[3]) && (z[2] == x[2]) && (z[1] == x[1]) && (z[0] == x[0])
}

// IsZero returns z == 0
func (z *Element) IsZero() bool {
	return (z[5] | z[4] | z[3] | z[2] | z[1] | z[0]) == 0
}

// SetRandom sets z to a random element < q
func (z *Element) SetRandom() *Element {
	bytes := make([]byte, 48)
	io.ReadFull(rand.Reader, bytes)
	z[0] = binary.BigEndian.Uint64(bytes[0:8])
	z[1] = binary.BigEndian.Uint64(bytes[8:16])
	z[2] = binary.BigEndian.Uint64(bytes[16:24])
	z[3] = binary.BigEndian.Uint64(bytes[24:32])
	z[4] = binary.BigEndian.Uint64(bytes[32:40])
	z[5] = binary.BigEndian.Uint64(bytes[40:48])
	z[5] %= 121098312706494698

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 121098312706494698 || (z[5] == 121098312706494698 && (z[4] < 14284016967150029115 || (z[4] == 14284016967150029115 && (z[3] < 1883307231910630287 || (z[3] == 1883307231910630287 && (z[2] < 2230234197602682880 || (z[2] == 2230234197602682880 && (z[1] < 1660523435060625408 || (z[1] == 1660523435060625408 && (z[0] < 9586122913090633729))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 9586122913090633729, 0)
		z[1], b = bits.Sub64(z[1], 1660523435060625408, b)
		z[2], b = bits.Sub64(z[2], 2230234197602682880, b)
		z[3], b = bits.Sub64(z[3], 1883307231910630287, b)
		z[4], b = bits.Sub64(z[4], 14284016967150029115, b)
		z[5], _ = bits.Sub64(z[5], 121098312706494698, b)
	}

	return z
}

// One returns 1 (in montgommery form)
func One() Element {
	var one Element
	one.SetOne()
	return one
}

// MulAssign is deprecated
// Deprecated: use Mul instead
func (z *Element) MulAssign(x *Element) *Element {
	return z.Mul(z, x)
}

// AddAssign is deprecated
// Deprecated: use Add instead
func (z *Element) AddAssign(x *Element) *Element {
	return z.Add(z, x)
}

// SubAssign is deprecated
// Deprecated: use Sub instead
func (z *Element) SubAssign(x *Element) *Element {
	return z.Sub(z, x)
}

// API with assembly impl

// Mul z = x * y mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Mul(x, y *Element) *Element {
	mul(z, x, y)
	return z
}

// Square z = x * x mod q
// see https://hackmd.io/@zkteam/modular_multiplication
func (z *Element) Square(x *Element) *Element {
	square(z, x)
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element) FromMont() *Element {
	fromMont(z)
	return z
}

// Add z = x + y mod q
func (z *Element) Add(x, y *Element) *Element {
	add(z, x, y)
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element) Double(x *Element) *Element {
	double(z, x)
	return z
}

// Sub  z = x - y mod q
func (z *Element) Sub(x, y *Element) *Element {
	sub(z, x, y)
	return z
}

// Neg z = q - x
func (z *Element) Neg(x *Element) *Element {
	neg(z, x)
	return z
}

// Generic (no ADX instructions, no AMD64) versions of multiplication and squaring algorithms

func _mulGeneric(z, x, y *Element) {

	var t [6]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd1(v, y[4], c[1])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd1(v, y[5], c[1])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 4
		v := x[4]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 5
		v := x[5]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		c[2], z[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, y[4], c[1], t[4])
		c[2], z[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, y[5], c[1], t[5])
		z[5], z[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 121098312706494698 || (z[5] == 121098312706494698 && (z[4] < 14284016967150029115 || (z[4] == 14284016967150029115 && (z[3] < 1883307231910630287 || (z[3] == 1883307231910630287 && (z[2] < 2230234197602682880 || (z[2] == 2230234197602682880 && (z[1] < 1660523435060625408 || (z[1] == 1660523435060625408 && (z[0] < 9586122913090633729))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 9586122913090633729, 0)
		z[1], b = bits.Sub64(z[1], 1660523435060625408, b)
		z[2], b = bits.Sub64(z[2], 2230234197602682880, b)
		z[3], b = bits.Sub64(z[3], 1883307231910630287, b)
		z[4], b = bits.Sub64(z[4], 14284016967150029115, b)
		z[5], _ = bits.Sub64(z[5], 121098312706494698, b)
	}
}

func _squareGeneric(z, x *Element) {

	var t [6]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, x[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd1(v, x[1], c[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd1(v, x[2], c[1])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd1(v, x[3], c[1])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd1(v, x[4], c[1])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd1(v, x[5], c[1])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 4
		v := x[4]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], t[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], t[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		t[5], t[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}
	{
		// round 5
		v := x[5]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 9586122913090633727
		c[2] = madd0(m, 9586122913090633729, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], z[0] = madd2(m, 1660523435060625408, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], z[1] = madd2(m, 2230234197602682880, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		c[2], z[2] = madd2(m, 1883307231910630287, c[2], c[0])
		c[1], c[0] = madd2(v, x[4], c[1], t[4])
		c[2], z[3] = madd2(m, 14284016967150029115, c[2], c[0])
		c[1], c[0] = madd2(v, x[5], c[1], t[5])
		z[5], z[4] = madd3(m, 121098312706494698, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 121098312706494698 || (z[5] == 121098312706494698 && (z[4] < 14284016967150029115 || (z[4] == 14284016967150029115 && (z[3] < 1883307231910630287 || (z[3] == 1883307231910630287 && (z[2] < 2230234197602682880 || (z[2] == 2230234197602682880 && (z[1] < 1660523435060625408 || (z[1] == 1660523435060625408 && (z[0] < 9586122913090633729))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 9586122913090633729, 0)
		z[1], b = bits.Sub64(z[1], 1660523435060625408, b)
		z[2], b = bits.Sub64(z[2], 2230234197602682880, b)
		z[3], b = bits.Sub64(z[3], 1883307231910630287, b)
		z[4], b = bits.Sub64(z[4], 14284016967150029115, b)
		z[5], _ = bits.Sub64(z[5], 121098312706494698, b)
	}
}

func _fromMontGeneric(z *Element) {
	// the following lines implement z = z * 1
	// with a modified CIOS montgomery multiplication
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 9586122913090633727
		C := madd0(m, 9586122913090633729, z[0])
		C, z[0] = madd2(m, 1660523435060625408, z[1], C)
		C, z[1] = madd2(m, 2230234197602682880, z[2], C)
		C, z[2] = madd2(m, 1883307231910630287, z[3], C)
		C, z[3] = madd2(m, 14284016967150029115, z[4], C)
		C, z[4] = madd2(m, 121098312706494698, z[5], C)
		z[5] = C
	}

	// if z > q --> z -= q
	// note: this is NOT constant time
	if !(z[5] < 121098312706494698 || (z[5] == 121098312706494698 && (z[4] < 14284016967150029115 || (z[4] == 14284016967150029115 && (z[3] < 1883307231910630287 || (z[3] == 1883307231910630287 && (z[2] < 2230234197602682880 || (z[2] == 2230234197602682880 && (z[1] < 1660523435060625408 || (z[1] == 1660523435060625408 && (z[0] < 9586122913090633729))))))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 9586122913090633729, 0)
		z[1], b = bits.Sub64(z[1], 1660523435060625408, b)
		z[2], b = bits.Sub64(z[2], 2230234197602682880, b)
		z[3], b = bits.Sub64(z[3], 1883307231910630287, b)
		z[4], b = bits.Sub64(z[4], 14284016967150029115, b)
		z[5], _ = bits.Sub64(z[5], 121098312706494698, b)
	}
}

// Exp z = x^exponent mod q
func (z *Element) Exp(x Element, exponent *big.Int) *Element {
	var bZero big.Int
	if exponent.Cmp(&bZero) == 0 {
		return z.SetOne()
	}

	z.Set(&x)

	for i := exponent.BitLen() - 2; i >= 0; i-- {
		z.Square(z)
		if exponent.Bit(i) == 1 {
			z.Mul(z, &x)
		}
	}

	return z
}

// ToMont converts z to Montgomery form
// sets and returns z = z * r^2
func (z *Element) ToMont() *Element {
	return z.Mul(z, &rSquare)
}

// ToRegular returns z in regular form (doesn't mutate z)
func (z Element) ToRegular() Element {
	return *z.FromMont()
}

// String returns the string form of an Element in Montgomery form
func (z *Element) String() string {
	var _z big.Int
	return z.ToBigIntRegular(&_z).String()
}

// ToBigInt returns z as a big.Int in Montgomery form
func (z *Element) ToBigInt(res *big.Int) *big.Int {
	var b [Limbs * 8]byte
	binary.BigEndian.PutUint64(b[40:48], z[0])
	binary.BigEndian.PutUint64(b[32:40], z[1])
	binary.BigEndian.PutUint64(b[24:32], z[2])
	binary.BigEndian.PutUint64(b[16:24], z[3])
	binary.BigEndian.PutUint64(b[8:16], z[4])
	binary.BigEndian.PutUint64(b[0:8], z[5])

	return res.SetBytes(b[:])
}

// ToBigIntRegular returns z as a big.Int in regular form
func (z Element) ToBigIntRegular(res *big.Int) *big.Int {
	z.FromMont()
	return z.ToBigInt(res)
}

// SetBigInt sets z to v (regular form) and returns z in Montgomery form
func (z *Element) SetBigInt(v *big.Int) *Element {
	z.SetZero()

	var zero big.Int
	q := Modulus()

	// fast path
	c := v.Cmp(q)
	if c == 0 {
		// v == 0
		return z
	} else if c != 1 && v.Cmp(&zero) != -1 {
		// 0 < v < q
		return z.setBigInt(v)
	}

	// copy input + modular reduction
	vv := new(big.Int).Set(v)
	vv.Mod(v, q)

	return z.setBigInt(vv)
}

// setBigInt assumes 0 <= v < q
func (z *Element) setBigInt(v *big.Int) *Element {
	vBits := v.Bits()

	if bits.UintSize == 64 {
		for i := 0; i < len(vBits); i++ {
			z[i] = uint64(vBits[i])
		}
	} else {
		for i := 0; i < len(vBits); i++ {
			if i%2 == 0 {
				z[i/2] = uint64(vBits[i])
			} else {
				z[i/2] |= uint64(vBits[i]) << 32
			}
		}
	}

	return z.ToMont()
}

// SetString creates a big.Int with s (in base 10) and calls SetBigInt on z
func (z *Element) SetString(s string) *Element {
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Element.SetString failed -> can't parse number in base10 into a big.Int")
	}
	return z.SetBigInt(x)
}

var (
	_bLegendreExponentElement *big.Int
	_bSqrtExponentElement     *big.Int
)

func init() {
	_bLegendreExponentElement, _ = new(big.Int).SetString("d71d230be28875631d82e03650a49d8d116cf9807a89c78f79b117dd04a4000b85aea2180000004284600000000000", 16)
	const sqrtExponentElement = "35c748c2f8a21d58c760b80d94292763445b3e601ea271e3de6c45f741290002e16ba88600000010a11"
	_bSqrtExponentElement, _ = new(big.Int).SetString(sqrtExponentElement, 16)
}

// Legendre returns the Legendre symbol of z (either +1, -1, or 0.)
func (z *Element) Legendre() int {
	var l Element
	// z^((q-1)/2)
	l.Exp(*z, _bLegendreExponentElement)

	if l.IsZero() {
		return 0
	}

	// if l == 1
	if (l[5] == 39800542322357402) && (l[4] == 5545221690922665192) && (l[3] == 8885205928937022213) && (l[2] == 11492539364873682930) && (l[1] == 5854854902718660529) && (l[0] == 202099033278250856) {
		return 1
	}
	return -1
}

// Sqrt z = √x mod q
// if the square root doesn't exist (x is not a square mod q)
// Sqrt leaves z unchanged and returns nil
func (z *Element) Sqrt(x *Element) *Element {
	// q ≡ 1 (mod 4)
	// see modSqrtTonelliShanks in math/big/int.go
	// using https://www.maa.org/sites/default/files/pdf/upload_library/22/Polya/07468342.di020786.02p0470a.pdf

	var y, b, t, w Element
	// w = x^((s-1)/2))
	w.Exp(*x, _bSqrtExponentElement)

	// y = x^((s+1)/2)) = w * x
	y.Mul(x, &w)

	// b = x^s = w * w * x = y * x
	b.Mul(&w, &y)

	// g = nonResidue ^ s
	var g = Element{
		7563926049028936178,
		2688164645460651601,
		12112688591437172399,
		3177973240564633687,
		14764383749841851163,
		52487407124055189,
	}
	r := uint64(46)

	// compute legendre symbol
	// t = x^((q-1)/2) = r-1 squaring of x^s
	t = b
	for i := uint64(0); i < r-1; i++ {
		t.Square(&t)
	}
	if t.IsZero() {
		return z.SetZero()
	}
	if !((t[5] == 39800542322357402) && (t[4] == 5545221690922665192) && (t[3] == 8885205928937022213) && (t[2] == 11492539364873682930) && (t[1] == 5854854902718660529) && (t[0] == 202099033278250856)) {
		// t != 1, we don't have a square root
		return nil
	}
	for {
		var m uint64
		t = b

		// for t != 1
		for !((t[5] == 39800542322357402) && (t[4] == 5545221690922665192) && (t[3] == 8885205928937022213) && (t[2] == 11492539364873682930) && (t[1] == 5854854902718660529) && (t[0] == 202099033278250856)) {
			t.Square(&t)
			m++
		}

		if m == 0 {
			return z.Set(&y)
		}
		// t = g^(2^(r-m-1)) mod q
		ge := int(r - m - 1)
		t = g
		for ge > 0 {
			t.Square(&t)
			ge--
		}

		g.Square(&t)
		y.Mul(&y, &t)
		b.Mul(&b, &g)
		r = m
	}
}

// Inverse z = x^-1 mod q
// Algorithm 16 in "Efficient Software-Implementation of Finite Fields with Applications to Cryptography"
// if x == 0, sets and returns z = x
func (z *Element) Inverse(x *Element) *Element {
	if x.IsZero() {
		return z.Set(x)
	}

	// initialize u = q
	var u = Element{
		9586122913090633729,
		1660523435060625408,
		2230234197602682880,
		1883307231910630287,
		14284016967150029115,
		121098312706494698,
	}

	// initialize s = r^2
	var s = Element{
		13224372171368877346,
		227991066186625457,
		2496666625421784173,
		13825906835078366124,
		9475172226622360569,
		30958721782860680,
	}

	// r = 0
	r := Element{}

	v := *x

	var carry, borrow, t, t2 uint64
	var bigger, uIsOne, vIsOne bool

	for !uIsOne && !vIsOne {
		for v[0]&1 == 0 {

			// v = v >> 1
			t2 = v[5] << 63
			v[5] >>= 1
			t = t2
			t2 = v[4] << 63
			v[4] = (v[4] >> 1) | t
			t = t2
			t2 = v[3] << 63
			v[3] = (v[3] >> 1) | t
			t = t2
			t2 = v[2] << 63
			v[2] = (v[2] >> 1) | t
			t = t2
			t2 = v[1] << 63
			v[1] = (v[1] >> 1) | t
			t = t2
			v[0] = (v[0] >> 1) | t

			if s[0]&1 == 1 {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 9586122913090633729, 0)
				s[1], carry = bits.Add64(s[1], 1660523435060625408, carry)
				s[2], carry = bits.Add64(s[2], 2230234197602682880, carry)
				s[3], carry = bits.Add64(s[3], 1883307231910630287, carry)
				s[4], carry = bits.Add64(s[4], 14284016967150029115, carry)
				s[5], _ = bits.Add64(s[5], 121098312706494698, carry)

			}

			// s = s >> 1
			t2 = s[5] << 63
			s[5] >>= 1
			t = t2
			t2 = s[4] << 63
			s[4] = (s[4] >> 1) | t
			t = t2
			t2 = s[3] << 63
			s[3] = (s[3] >> 1) | t
			t = t2
			t2 = s[2] << 63
			s[2] = (s[2] >> 1) | t
			t = t2
			t2 = s[1] << 63
			s[1] = (s[1] >> 1) | t
			t = t2
			s[0] = (s[0] >> 1) | t

		}
		for u[0]&1 == 0 {

			// u = u >> 1
			t2 = u[5] << 63
			u[5] >>= 1
			t = t2
			t2 = u[4] << 63
			u[4] = (u[4] >> 1) | t
			t = t2
			t2 = u[3] << 63
			u[3] = (u[3] >> 1) | t
			t = t2
			t2 = u[2] << 63
			u[2] = (u[2] >> 1) | t
			t = t2
			t2 = u[1] << 63
			u[1] = (u[1] >> 1) | t
			t = t2
			u[0] = (u[0] >> 1) | t

			if r[0]&1 == 1 {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 9586122913090633729, 0)
				r[1], carry = bits.Add64(r[1], 1660523435060625408, carry)
				r[2], carry = bits.Add64(r[2], 2230234197602682880, carry)
				r[3], carry = bits.Add64(r[3], 1883307231910630287, carry)
				r[4], carry = bits.Add64(r[4], 14284016967150029115, carry)
				r[5], _ = bits.Add64(r[5], 121098312706494698, carry)

			}

			// r = r >> 1
			t2 = r[5] << 63
			r[5] >>= 1
			t = t2
			t2 = r[4] << 63
			r[4] = (r[4] >> 1) | t
			t = t2
			t2 = r[3] << 63
			r[3] = (r[3] >> 1) | t
			t = t2
			t2 = r[2] << 63
			r[2] = (r[2] >> 1) | t
			t = t2
			t2 = r[1] << 63
			r[1] = (r[1] >> 1) | t
			t = t2
			r[0] = (r[0] >> 1) | t

		}

		// v >= u
		bigger = !(v[5] < u[5] || (v[5] == u[5] && (v[4] < u[4] || (v[4] == u[4] && (v[3] < u[3] || (v[3] == u[3] && (v[2] < u[2] || (v[2] == u[2] && (v[1] < u[1] || (v[1] == u[1] && (v[0] < u[0])))))))))))

		if bigger {

			// v = v - u
			v[0], borrow = bits.Sub64(v[0], u[0], 0)
			v[1], borrow = bits.Sub64(v[1], u[1], borrow)
			v[2], borrow = bits.Sub64(v[2], u[2], borrow)
			v[3], borrow = bits.Sub64(v[3], u[3], borrow)
			v[4], borrow = bits.Sub64(v[4], u[4], borrow)
			v[5], _ = bits.Sub64(v[5], u[5], borrow)

			// r >= s
			bigger = !(r[5] < s[5] || (r[5] == s[5] && (r[4] < s[4] || (r[4] == s[4] && (r[3] < s[3] || (r[3] == s[3] && (r[2] < s[2] || (r[2] == s[2] && (r[1] < s[1] || (r[1] == s[1] && (r[0] < s[0])))))))))))

			if bigger {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 9586122913090633729, 0)
				s[1], carry = bits.Add64(s[1], 1660523435060625408, carry)
				s[2], carry = bits.Add64(s[2], 2230234197602682880, carry)
				s[3], carry = bits.Add64(s[3], 1883307231910630287, carry)
				s[4], carry = bits.Add64(s[4], 14284016967150029115, carry)
				s[5], _ = bits.Add64(s[5], 121098312706494698, carry)

			}

			// s = s - r
			s[0], borrow = bits.Sub64(s[0], r[0], 0)
			s[1], borrow = bits.Sub64(s[1], r[1], borrow)
			s[2], borrow = bits.Sub64(s[2], r[2], borrow)
			s[3], borrow = bits.Sub64(s[3], r[3], borrow)
			s[4], borrow = bits.Sub64(s[4], r[4], borrow)
			s[5], _ = bits.Sub64(s[5], r[5], borrow)

		} else {

			// u = u - v
			u[0], borrow = bits.Sub64(u[0], v[0], 0)
			u[1], borrow = bits.Sub64(u[1], v[1], borrow)
			u[2], borrow = bits.Sub64(u[2], v[2], borrow)
			u[3], borrow = bits.Sub64(u[3], v[3], borrow)
			u[4], borrow = bits.Sub64(u[4], v[4], borrow)
			u[5], _ = bits.Sub64(u[5], v[5], borrow)

			// s >= r
			bigger = !(s[5] < r[5] || (s[5] == r[5] && (s[4] < r[4] || (s[4] == r[4] && (s[3] < r[3] || (s[3] == r[3] && (s[2] < r[2] || (s[2] == r[2] && (s[1] < r[1] || (s[1] == r[1] && (s[0] < r[0])))))))))))

			if bigger {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 9586122913090633729, 0)
				r[1], carry = bits.Add64(r[1], 1660523435060625408, carry)
				r[2], carry = bits.Add64(r[2], 2230234197602682880, carry)
				r[3], carry = bits.Add64(r[3], 1883307231910630287, carry)
				r[4], carry = bits.Add64(r[4], 14284016967150029115, carry)
				r[5], _ = bits.Add64(r[5], 121098312706494698, carry)

			}

			// r = r - s
			r[0], borrow = bits.Sub64(r[0], s[0], 0)
			r[1], borrow = bits.Sub64(r[1], s[1], borrow)
			r[2], borrow = bits.Sub64(r[2], s[2], borrow)
			r[3], borrow = bits.Sub64(r[3], s[3], borrow)
			r[4], borrow = bits.Sub64(r[4], s[4], borrow)
			r[5], _ = bits.Sub64(r[5], s[5], borrow)

		}
		uIsOne = (u[0] == 1) && (u[5]|u[4]|u[3]|u[2]|u[1]) == 0
		vIsOne = (v[0] == 1) && (v[5]|v[4]|v[3]|v[2]|v[1]) == 0
	}

	if uIsOne {
		z.Set(&r)
	} else {
		z.Set(&s)
	}

	return z
}
