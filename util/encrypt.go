package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
func decrypt(data []byte, passphrase string)[]byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}
func encryptFile(filename string, data []byte, passphrase string){
	f,_:=os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data,passphrase))
}
func decryptFile(filename string, passphrase string)[]byte{
	data,_:=ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	//fmt.Println("Encryption Program v0.01")

	text := []byte("My Super Secret Code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	ciphertext := encrypt([]byte(text), string(key))
	fmt.Println(string(ciphertext))
	plaintext := decrypt(ciphertext, string(key))
	fmt.Println(string(plaintext))


	//// generate a new aes cipher using our 32 byte long key
	//c, err := aes.NewCipher(key)
	//// if there are any errors, handle them
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// gcm or Galois/Counter Mode, is a mode of operation
	//// for symmetric key cryptographic block ciphers
	//// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//gcm, err := cipher.NewGCM(c)
	//// if any error generating new GCM
	//// handle them
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//// creates a new byte array the size of the nonce
	//// which must be passed to Seal
	//nonce := make([]byte, gcm.NonceSize())
	//// populates our nonce with a cryptographically secure
	//// random sequence
	//if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
	//	fmt.Println(err)
	//}
	//
	//// here we encrypt our text using the Seal function
	//// Seal encrypts and authenticates plaintext, authenticates the
	//// additional data and appends the result to dst, returning the updated
	//// slice. The nonce must be NonceSize() bytes long and unique for all
	//// time, for a given key.
	//
	//sHash:=gcm.Seal(nonce, nonce, text, nil)
	//
	//err = ioutil.WriteFile("encrypt.data", sHash,0777)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
