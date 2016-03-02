package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random number from [1, p).
func PrivateKey(p *big.Int) *big.Int {
	pSubThree := new(big.Int).Sub(p, big.NewInt(3))
	n, _ := rand.Int(rand.Reader, pSubThree)
	return n.Add(n, big.NewInt(2))
}

// PublicKey generates a public key from a private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair generates a public/private key pair.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey uses a private and a public key to generate a shared secret.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
