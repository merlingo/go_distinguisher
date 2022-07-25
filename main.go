package main

import (
	"crypto/md5"
	"github.com/merlingo/go_distinguisher/LAT"
)

func simple_cipher_for_test(p []byte) []byte {

	c := md5.Sum(p)

	return c[0:2]
}

func main() {
	bytelen := 2
	cipherPairs := LAT.CipherPairList(bytelen, simple_cipher_for_test)
	res, err := LAT.AllMaskingResult(cipherPairs, bytelen)
	if err != nil {
		print(err)
	}
	print(res)
}
