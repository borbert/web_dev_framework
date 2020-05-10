package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CreateHash(key string) string {
	hasher := sha512.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
//func Encrypt2(data []byte, passphrase string)[]byte{
//	cipher.text:=bcrypt.GenerateFromPassword(passphrase)
//}

func Decrypt(data []byte, passphrase string)[]byte {
	key := []byte(CreateHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}
func EncryptFile(filename string, data []byte, passphrase string){
	f,_:=os.Create(filename)
	defer f.Close()
	f.Write(Encrypt(data,passphrase))
}
func DecryptFile(filename string, passphrase string)[]byte{
	data,_:=ioutil.ReadFile(filename)
	return Decrypt(data, passphrase)
}

func main(){
	key:="test@test.test"
	hashed:=Encrypt([]byte("1234"),key)
	fmt.Println(hashed)
	dehashed:=Decrypt(hashed,key)
	fmt.Println(dehashed)

}


