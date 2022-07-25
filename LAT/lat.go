package LAT

import (
	"container/list"
	"errors"
	"fmt"
	"math"
)

type CipherPair struct {
	plain   []byte
	cipher  []byte
	bytelen int
}
type MaskTuple struct {
	a, b int
}

func makeCipherPair(plaintext []byte, cipher func([]byte) []byte) CipherPair {
	cp := CipherPair{
		plain:   make([]byte, len(plaintext)),
		cipher:  nil,
		bytelen: len(plaintext),
	}
	copy(cp.plain, plaintext)

	cp.cipher = cipher(plaintext)
	return cp
}
func CipherPairList(bytelen int, cipher func([]byte) []byte) *list.List {
	cipherPairs := new(list.List)
	plaintexts := make([]byte, bytelen)
	pcount := int(math.Pow(2, float64(8))) * bytelen

	for a_i := 0; a_i < int(pcount); a_i++ {
		byte_index := a_i / int(math.Pow(2, float64(8)))
		plaintexts[byte_index] += 1
		cipherPairs.PushBack(makeCipherPair(plaintexts, cipher))
	}
	return cipherPairs
}

func MaskingZeroFound(cipherPairs *list.List, a_mask, b_mask []byte, bytelen int) (int, error) {
	var res = make([]byte, bytelen)
	var index = 0
	for e := cipherPairs.Front(); e != nil; e = e.Next() {
		b := false
		pair := e.Value.(CipherPair)
		fmt.Println("cipher pair: ", e.Value)
		if bytelen != len(pair.plain) && bytelen != len(pair.cipher) {
			return 0, errors.New("plaintext and cipher text length error!!")
		}
		for i := 0; i < bytelen; i++ {
			res[i] = (a_mask[i] & pair.plain[i]) ^ (b_mask[i] & pair.cipher[i])
			if res[i] != 0 {
				b = false
				break
			}
			b = true
		}
		if b {
			index += 1
		}
	}
	return index, nil
}

func AllMaskingResult(cipherPairs *list.List, bytelen int) (map[MaskTuple]int, error) {
	//2^8*bytelen kadarlık tum alfa beta masklari cikarilir. plaintext ve cipher'a uygulanır.
	mask_count := int(math.Pow(2, float64(8))) * bytelen
	a_mask := make([]byte, bytelen)
	b_mask := make([]byte, bytelen)
	m := make(map[MaskTuple]int)
	for a_i := 0; a_i < int(mask_count); a_i++ {
		a_mask[a_i/int(math.Pow(2, float64(8)))] += 1
		for b_i := 0; b_i < int(mask_count); b_i++ {
			b_mask[b_i/int(math.Pow(2, float64(8)))] += 1
			//a-b icin kac tane plain cipher ikilisi 0 verir
			count, err := MaskingZeroFound(cipherPairs, a_mask, b_mask, bytelen)
			if err != nil {
				print(err)
				return nil, err
			}
			m[MaskTuple{a: a_i, b: b_i}] = count - cipherPairs.Len()/2
		}
	}
	return m, nil
}
