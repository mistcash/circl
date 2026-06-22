// Package internal provides internal Kyber512 implementation details.
package internal

// Kyber512 exports

var CompressedPolySize = compressedPolySize

// Kyber512 key components

// Sh returns sk.sh NTT(s), normalized
func (sk PrivateKey) Sh() Vec {
	return sk.sh
}

// Rho returns pk.rho ρ, the seed for the matrix A
func (pk PublicKey) Rho() [32]byte {
	return pk.rho
}

// Th returns pk.th NTT(t), normalized
func (pk PublicKey) Th() Vec {
	return pk.th
}

// AT returns pk.aT the matrix Aᵀ
func (pk PublicKey) AT() Mat {
	return pk.aT
}
