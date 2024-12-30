package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"
)

func SignMessage(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	hash := sha256.Sum256([]byte(message))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return hex.EncodeToString(signature), nil
}

func VerifySignature(publicKey *ecdsa.PublicKey, message, signatureHex string) (bool, error) {
	sigBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}

	if len(sigBytes) != 64 {
		return false, errors.New("invalid signature length")
	}

	r := big.NewInt(0).SetBytes(sigBytes[:32])
	s := big.NewInt(0).SetBytes(sigBytes[32:])
	hash := sha256.Sum256([]byte(message))
	return ecdsa.Verify(publicKey, hash[:], r, s), nil
}
