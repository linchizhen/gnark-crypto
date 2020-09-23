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

// Code generated by gurvy DO NOT EDIT

package bn256

func addE2(z, x, y *e2) {
	z.A0.Add(&x.A0, &y.A0)
	z.A1.Add(&x.A1, &y.A1)
}

func subE2(z, x, y *e2) {
	z.A0.Sub(&x.A0, &y.A0)
	z.A1.Sub(&x.A1, &y.A1)
}

func doubleE2(z, x *e2) {
	z.A0.Double(&x.A0)
	z.A1.Double(&x.A1)
}

func negE2(z, x *e2) {
	z.A0.Neg(&x.A0)
	z.A1.Neg(&x.A1)
}

func squareAdxE2(z, x *e2) {
	panic("not implemented")
}

func mulAdxE2(z, x, y *e2) {
	panic("not implemented")
}
