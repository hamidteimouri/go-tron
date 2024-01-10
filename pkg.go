package gotron

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/mr-tron/base58"
)

func S256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}
func HexToBase58(str string) (string, error) {
	addb, err := hex.DecodeString(str)
	if err != nil {
		return "", nil
	}
	hash1 := S256(S256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}
	return base58.Encode(addb), nil
}

func Base58ToHex(str string) (string, error) {
	// Decode base58 string
	decoded, err := base58.Decode(str)
	if err != nil {
		return "", err
	}

	// Remove the last 4 bytes (checksum) from the decoded data
	address := decoded[:len(decoded)-4]

	// Convert the remaining data to hexadecimal
	hexStr := hex.EncodeToString(address)

	return hexStr, nil
}
