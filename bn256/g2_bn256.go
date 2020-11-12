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

// ClearCofactor maps a point in E'(Fp2) to E'(Fp2)[r]
// cf http://cacr.uwaterloo.ca/techreports/2011/cacr2011-26.pdf, 6.1
func (p *g2Jac) ClearCofactor(a *g2Jac) *g2Jac {

	var points [4]g2Jac

	points[0].ScalarMultiplication(a, &xGen)

	points[1].Double(&points[0]).
		AddAssign(&points[0]).
		psi(&points[1])

	points[2].psi(&points[0]).
		psi(&points[2])

	points[3].psi(a).psi(&points[3]).psi(&points[3])

	var res g2Jac
	res.Set(&g2Infinity)
	for i := 0; i < 4; i++ {
		res.AddAssign(&points[i])
	}
	p.Set(&res)
	return p

}

// ClearCofactor maps a point in E(Fp) to E(Fp)[r]
// cf https://eprint.iacr.org/2019/403.pdf, 5
func (p *G2Affine) ClearCofactor(a *G2Affine) *G2Affine {
	var _p g2Jac
	_p.FromAffine(a)
	_p.ClearCofactor(&_p)
	p.FromJacobian(&_p)
	return p
}
