package frequence

import (
	"testing"
)

/*
Benchmark av FrequenceForBench metoden
Tre benchmark funksjoner for å teste tre forskjellige filer av varierende størrelse.
*/
//Tester tekst fil pg100.txt
func BenchmarkFrequencepg100(b *testing.B) {
	benchFrequencePG100(1, b)

}

//test tekstfil pg200.txt
func BenchmarkFrequencepg200(b *testing.B) {
	benchmarkFrequencepg200(1, b)
}

//Tester tekstfil bible.txt
func BenchmarkFrequence1Bible(b *testing.B) {
	benchFrequenceBible(1, b)
}

//tester tekstfil bible2.txt
func BenchmarkFrequenceBible2(b *testing.B) {
	benchmarkFrequenceBible2(1, b)
}
func BenchmarkFrequenceFile(b *testing.B) {
	benchmarkFrequenceFile(1, b)
}

/*
Funksjonene under setter absolutt path til filen for benchtesting av hver fil.
Funksjonene blir kalt inn i testene over.
 */
func benchFrequencePG100(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/pg100.txt")
	}
}
func benchmarkFrequencepg200(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/pg200.txt")
	}
}
func benchFrequenceBible(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/bible.txt")
	}
}

func benchmarkFrequenceBible2(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/bible2.txt")
	}

}
func benchmarkFrequenceFile(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		b.StartTimer()
		FrequenceForBench("../files/file.txt")

	}
}
