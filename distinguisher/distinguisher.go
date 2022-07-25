package distinguisher

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

func pillingLemmaApproximation(u, v []byte) float64 {
	return -1
}
