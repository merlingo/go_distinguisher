package main

import (
	"github.com/merlingo/go_distinguisher/LAT"
)

func simple_cipher_for_test(key, p byte) byte {

	return key ^ p
}

func main() {
	//bytelen := 2,
	var key byte = 45
	cipherPairs := LAT.CipherPairList(key, simple_cipher_for_test)
	res, err := LAT.AllMaskingResult(cipherPairs)
	if err != nil {
		print(err)
	}
	print(res)
}
