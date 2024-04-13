package main

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log"
	"zk-snarks-example-project/circuit"
	"zk-snarks-example-project/polynomials"
)

func main() {
	coefficients := []int{1, 2, 3}
	x := 5

	polynomial := polynomials.NewPolynomial(coefficients)

	var polyCircuit circuit.PolynomialEvaluationCircuit
	polyCircuit.InitCircuit(polynomial)

	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &polyCircuit)

	// Setup:
	// Given the constraint system `cs`,
	// generate the quadratic arithmetic program (QAP) internally,
	// and output the proving key and the verification key.
	provingKey, verificationKey, err := groth16.Setup(cs)
	if err != nil {
		log.Fatalf("Error happened during the groth16 setup: %e", err)
	}

	// Prover:

	// Finds a secret, i.e. evaluates a polynomials at the given value x
	y := polynomial.ValueAt(x)

	// Generates a witness, i.e. a representation of the
	// public input: [a0, a1, ..., al],
	// and the secret input known to the prover only: [a(l + 1) ... am]
	witness, err := frontend.NewWitness(
		&circuit.PolynomialEvaluationCircuit{
			/* public */ X: x,
			/* secret */ Y: y,
		},
		ecc.BN254.ScalarField(),
	)
	if err != nil {
		log.Fatalf("Error happened during generating the witness: %e", err)
	}

	// Generates proof of the secret for the given constraint system.
	// Uses the generated proving key and the witness
	proof, err := groth16.Prove(cs, provingKey, witness)
	if err != nil {
		log.Fatalf("Error happened on the side of the prover: %e", err)
	}

	fmt.Println("Prover: generated the proof")

	// Verifier:

	// Has the public input only
	publicWitness, _ := witness.Public()

	// Verifies the proof received from the prover.
	// Uses the generated verification key and the public input [a0, a1, ..., al]
	err = groth16.Verify(proof, verificationKey, publicWitness)
	if err != nil {
		log.Fatalf("Error happened on the side of the verifier: %e", err)
	}

	fmt.Println("Verifier: successfully verified the proof")
}
