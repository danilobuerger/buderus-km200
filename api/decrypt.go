package api

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

func decode(body []byte) ([]byte, error) {
	data := make([]byte, base64.StdEncoding.DecodedLen(len(body)))
	n, err := base64.StdEncoding.Decode(data, body)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func decrypt(body, key []byte) ([]byte, error) {
	data, err := decode(body)
	if err != nil {
		return nil, err
	}

	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return bytes.Trim(decrypted, "\x00"), nil
}
