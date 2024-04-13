package circuit

import (
	"github.com/consensys/gnark/test"
	"testing"
	"zk-snarks-example-project/polynomials"
)

const x = 5
const y = 10

var basePolynomial = []int{1, 2, 3}
var constantPolynomial = []int{y}

func TestCircuit_basePolynomial_validWitness_proverSucceeded(t *testing.T) {
	assert := test.NewAssert(t)
	polynomial := polynomials.NewPolynomial(basePolynomial)

	var circuit PolynomialEvaluationCircuit
	circuit.InitCircuit(polynomial)

	assert.ProverSucceeded(
		&circuit,
		&PolynomialEvaluationCircuit{
			X: x,
			Y: polynomial.ValueAt(x),
		},
	)
}

func TestCircuit_basePolynomial_invalidWitness_proverFailed(t *testing.T) {
	assert := test.NewAssert(t)
	polynomial := polynomials.NewPolynomial(basePolynomial)

	var circuit PolynomialEvaluationCircuit
	circuit.InitCircuit(polynomial)

	assert.ProverFailed(
		&circuit,
		&PolynomialEvaluationCircuit{
			X: x,
			Y: polynomial.ValueAt(x) + 10,
		},
	)
}

func TestCircuit_constantPolynomial_validWitness_proverSucceeded(t *testing.T) {
	assert := test.NewAssert(t)
	polynomial := polynomials.NewPolynomial(constantPolynomial)

	var circuit PolynomialEvaluationCircuit
	circuit.InitCircuit(polynomial)

	assert.ProverSucceeded(
		&circuit,
		&PolynomialEvaluationCircuit{
			X: x,
			Y: y,
		},
	)
}

func TestCircuit_constantPolynomial_invalidWitness_proverFailed(t *testing.T) {
	assert := test.NewAssert(t)
	polynomial := polynomials.NewPolynomial(constantPolynomial)

	var circuit PolynomialEvaluationCircuit
	circuit.InitCircuit(polynomial)

	assert.ProverFailed(
		&circuit,
		&PolynomialEvaluationCircuit{
			X: x,
			Y: y * 5,
		},
	)
}
