// +build !amd64

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

package bn256

import "github.com/consensys/gurvy/bn256/fp"

// MulByNonResidue multiplies a e2 by (9,1)
func (z *e2) MulByNonResidue(x *e2) *e2 {
	var a, b fp.Element
	a.Double(&x.A0).Double(&a).Double(&a).Add(&a, &x.A0).Sub(&a, &x.A1)
	b.Double(&x.A1).Double(&b).Double(&b).Add(&b, &x.A1).Add(&b, &x.A0)
	z.A0.Set(&a)
	z.A1.Set(&b)
	return z
}

// Mul sets z to the e2-product of x,y, returns z
func (z *e2) Mul(x, y *e2) *e2 {
	mulGenericE2(z, x, y)
	return z
}

// Square sets z to the e2-product of x,x returns z
func (z *e2) Square(x *e2) *e2 {
	squareGenericE2(z, x)
	return z
}
