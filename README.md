## zk-snarks-example-project


### Description
This project demonstrates the concept of zk-snarks 
applied to polynomial evaluations at arbitrary points. 

The zk-snarks procedure is implemented using [gnark](https://github.com/Consensys/gnark) 
and the [groth16](https://eprint.iacr.org/2016/260.pdf) protocol.

Workflow:
* Setup
  * A trusted party generates a pair of (proving key, verification key) according to the scheme specified in [circuit.go](circuit/circuit.go).
* Prover
  * Evaluates the polynomial at the given point `x`, and generates a proof of the result.
* Verifier
  * Verifies correctness of the proof, received from the prover, using public input only, i.e. `x`.