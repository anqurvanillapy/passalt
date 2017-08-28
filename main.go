package passalt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
)

// New a password hash.
func New(passwd string) string {
	salt := generateSalt(16)
	hash := hmac.New(sha512.New, salt)
	hash.Write([]byte(passwd))

	saltStr := hex.EncodeToString(salt)
	hashStr := hex.EncodeToString(hash.Sum(nil))

	return fmt.Sprintf("sha512$%s$%s", saltStr, hashStr)
}

// Check if the password is valid.
func Check(token string, passwd string) bool {
	strArr := strings.Split(token, "$")
	tokenHashStr := strArr[2]
	saltByteArr, err := hex.DecodeString(strArr[1])

	if err != nil {
		log.Fatal(err)
	}

	hash := hmac.New(sha512.New, saltByteArr)
	hash.Write([]byte(passwd))
	hashStr := hex.EncodeToString(hash.Sum(nil))

	return tokenHashStr == hashStr
}

func generateSalt(len int) []byte {
	salt := make([]byte, int(math.Ceil(float64(len)/2.0)))
	_, err := rand.Read(salt)

	if err != nil {
		log.Fatal(err)
	}

	return salt
}
