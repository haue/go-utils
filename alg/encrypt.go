package alg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io/ioutil"
)

func EncryptBytes(key string, plain []byte) ([]byte, error) {
	bkey := []byte(key)
	block, err := aes.NewCipher(bkey)
	if err != nil {
		return make([]byte, 0), err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plain))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return make([]byte, 0), err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plain)
	return ciphertext, nil
}

func EncryptFile(key string, inputFile string, outputFile string) error {
	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	cipher, err := EncryptBytes(key, plaintext)
	return ioutil.WriteFile(outputFile, cipher, 0644)
}
