package config

import (
	"math/rand"
	"strconv"
)

var (
	jwtSignature []byte
	expiredTime  = 1000
)

func GenerateJwtSignature() {
	jwtSignature = []byte(strconv.FormatUint(rand.Uint64(), 10))
}

func GetJwtExpiredTime() int {
	return expiredTime
}

func GetJwtSignature() []byte {
	return jwtSignature
}
