package polynomials

import "log"

// Polynomial represents a polynomial of the form:
// f(x) = a0 + a1*x + a2*x^2 + ... + an*x^n
// `Degree` of the polynomial is an index of its last non-zero coefficient.
// Example: f(x) = 5 + 3*x + 6*x^2
type Polynomial struct {
	coefficients []int
	Degree       int
}

// NewPolynomial creates a new polynomial with the given coefficients.
func NewPolynomial(coefficients []int) *Polynomial {
	if len(coefficients) == 0 {
		log.Fatal(
			"Cannot instantiate a polynomials: at least one coefficient must be provided",
		)
	}

	polynomial := new(Polynomial)

	polynomial.coefficients = coefficients

	degree := len(coefficients) - 1
	for degree > 0 {
		if coefficients[degree] != 0 {
			break
		}
		degree--
	}

	polynomial.Degree = degree

	return polynomial
}

// CoefficientAt finds the polynomial coefficient matching the specific index:
// e.g given i -> ai
func (polynomial *Polynomial) CoefficientAt(i int) int {
	if i < len(polynomial.coefficients) {
		return polynomial.coefficients[i]
	}
	return 0
}

// ValueAt evaluates the polynomial at the specific value x.
// Returns f(x).
func (polynomial *Polynomial) ValueAt(x int) int {
	result := 0
	xPower := 1

	for i := 0; i <= polynomial.Degree; i++ {
		result += polynomial.CoefficientAt(i) * xPower
		xPower *= x
	}

	return result
}
