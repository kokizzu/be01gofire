package utils

import (
	"crypto/sha512"
	"encoding/base64"
)

const Salt = `.837*(&(o8g8`

func HashGenerator(str string) string {
	s := sha512.New()
	s.Write([]byte(str+Salt))
	res := s.Sum(nil)
	return base64.StdEncoding.EncodeToString(res)
}
func SamePassword(hashedByte string, pass string) bool {
	return hashedByte == HashGenerator(pass)
}
