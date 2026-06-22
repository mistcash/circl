// Package math provides exported types and functions for Kyber512.
package math

import (
	"github.com/cloudflare/circl/pke/kyber/internal/common"
	"github.com/cloudflare/circl/pke/kyber/kyber512/internal"
)

type PK = internal.PublicKey
type SK = internal.PrivateKey
type Vec = internal.Vec
type Mat = internal.Mat
type Poly = common.Poly

// common params

const Q = common.Q
const N = common.N
const PolySize = common.PolySize
const PlaintextSize = common.PlaintextSize
const Eta2 = common.Eta2

// params

const K = internal.K
const Eta1 = internal.Eta1
const DU = internal.DU
const DV = internal.DV
const PublicKeySize int = internal.PublicKeySize
const PrivateKeySize int = internal.PrivateKeySize
const SeedSize int = internal.SeedSize
const CiphertextSize int = internal.CiphertextSize

// methods

var NewKeyPair = internal.NewKeyFromSeed
var PolyDotHat = internal.PolyDotHat
var CompressedPolySize = internal.CompressedPolySize
