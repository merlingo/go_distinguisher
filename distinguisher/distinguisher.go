package distinguisher

import (
	"container/list"
	"github.com/merlingo/go_distinguisher/LAT"
	"math"
)

// random plainText olustur.
// crypto algoritması ile cipher çıkar
// cipher plainText içinden "distinguisher" turune gore secim yap
// secimi n defa uygula ve matris olustur
// matristen istatistiksel hesap yap  - ex: chi square

//The purpose of this method is to obtain a linear approximate expression of a given
//cipher algorithm. For this purpose, we begin by constructing a statistical linear path
//between input and output bits of each Sbox. Then we extend this path to the entire
//algorithm, and finally reach a linear approximate expression without any intermediate
//value

func pillingLemmaApproximation(biases ...float64) float64 {
	// calculate pilling-up bias value from all biases
	var total float64 = 0
	for _, bias := range biases {
		total *= bias
	}
	total = total*math.Pow(2, float64(len(biases)-1)) + 0.5
	return total
}

func matsui_alg1(cipherPairs *list.List, tuple LAT.MaskTuple) int {
	t0 := 0
	t1 := 0
	var res byte = 0

	for e := cipherPairs.Front(); e != nil; e = e.Next() {
		pair := e.Value.(LAT.CipherPair)
		res = (tuple.A & pair.Plain) ^ (tuple.B & pair.Cipher)
		if res == 0 {
			t0 += 1
		} else if res == 1 {
			t1 += 1
		}
	}
	if t0 > t1 {
		return 0
	} else {
		return 1
	}
}

func matsui_alg2_for_Kr(cipherPairs *list.List, tuple LAT.MaskTuple) bool {
	t0 := 0
	t1 := 0
	var res byte = 0

	for e := cipherPairs.Front(); e != nil; e = e.Next() {
		pair := e.Value.(LAT.CipherPair)
		res = (tuple.A & pair.Plain) ^ (tuple.B & pair.Cipher)
		if res == 0 {
			t0 += 1
		} else {
			t1 += 1
		}
	}
	return t0 > t1*4
}
