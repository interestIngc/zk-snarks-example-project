package circuit

import (
	"github.com/consensys/gnark/frontend"
	"zk-snarks-example-project/polynomials"
)

// PolynomialEvaluationCircuit is a circuit representing the check of
// `polynomial` evaluation at the specific value `X`.
// Its Define method defines a series of mathematical operations to prove that
// Y = polynomial(X).
type PolynomialEvaluationCircuit struct {
	polynomial *polynomials.Polynomial
	X          frontend.Variable `gnark:",public"` /* public variable */
	Y          frontend.Variable /* secret variable */
}

// InitCircuit initializes the circuit with the polynomial to be evaluated.
func (circuit *PolynomialEvaluationCircuit) InitCircuit(polynomial *polynomials.Polynomial) {
	circuit.polynomial = polynomial
}

// Define is the base method provided by gnark to construct a circuit.
// It defines a series of mathematical operations on both public
// and secret variables to prove that the secret is valid.
// This series of operations is then compiled into r1cs serving as an input to the
// zk-snarks protocol.
// In this case, we define a series of additions and multiplications to check that
// Y = polynomial(X).
func (circuit *PolynomialEvaluationCircuit) Define(api frontend.API) error {
	var result frontend.Variable = 0
	var xPower frontend.Variable = 1
	for i := 0; i <= circuit.polynomial.Degree; i++ {
		result = api.Add(result, api.Mul(circuit.polynomial.CoefficientAt(i), xPower))
		xPower = api.Mul(xPower, circuit.X)
	}

	api.AssertIsEqual(result, circuit.Y)

	return nil
}
