package LAT

import (
	"container/list"
	"fmt"
	"math"
)

type CipherPair struct {
	Plain   byte
	Cipher  byte
	bytelen int
}
type MaskTuple struct {
	A, B byte
}

func makeCipherPair(plaintext byte, key byte, cipher func(byte, byte) byte) CipherPair {
	cp := CipherPair{
		Plain:  plaintext,
		Cipher: 0,
	}

	cp.Cipher = cipher(plaintext, key)
	return cp
}
func CipherPairList(key byte, cipher func(byte, byte) byte) *list.List {
	cipherPairs := new(list.List)
	var plaintexts byte = 0
	pcount := int(math.Pow(2, float64(8))) //* bytelen

	for a_i := 0; a_i < int(pcount); a_i++ {
		//byte_index := a_i / int(math.Pow(2, float64(8)))
		//plaintexts[byte_index] += 1
		plaintexts += 1
		cipherPairs.PushBack(makeCipherPair(plaintexts, key, cipher))
	}
	return cipherPairs
}

func MaskingZeroFound(cipherPairs *list.List, a_mask, b_mask byte) (int, error) {
	var res byte = 0
	var index = 0
	for e := cipherPairs.Front(); e != nil; e = e.Next() {
		b := false
		pair := e.Value.(CipherPair)

		//if bytelen != len(pair.Plain) && bytelen != len(pair.Cipher) {
		//	return 0, errors.New("plaintext and cipher text length error!!")
		//}
		res = (a_mask & pair.Plain) ^ (b_mask & pair.Cipher)
		if res != 0 {
			b = false
		} else {
			b = true
		}

		if b {
			index += 1
		}
	}

	return index, nil
}

func AllMaskingResult(cipherPairs *list.List) (map[MaskTuple]int, error) {
	//2^8*bytelen kadarlık tum alfa beta masklari cikarilir. plaintext ve cipher'a uygulanır.
	mask_count := int(math.Pow(2, float64(8))) //* bytelen
	var a_mask byte = 0
	var b_mask byte = 0
	m := make(map[MaskTuple]int)
	for a_i := 0; a_i < int(mask_count); a_i++ {
		a_mask += 1
		for b_i := 0; b_i < int(mask_count); b_i++ {
			b_mask += 1
			//a-b icin kac tane plain cipher ikilisi 0 verir
			count, err := MaskingZeroFound(cipherPairs, a_mask, b_mask)
			if err != nil {
				print(err)
				return nil, err
			}
			r := count - cipherPairs.Len()/2
			m[MaskTuple{A: a_mask, B: b_mask}] = r
			fmt.Println("mask a: ", a_mask, "mask b:", b_mask, " --- index: ", r)
		}
	}
	return m, nil
}
