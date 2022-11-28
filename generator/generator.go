package generator

import (
	"math/rand"
	"strings"

	"genpwd/cryptosource"
)

type generator struct {
	rnd *rand.Rand
}

func New(csrc *cryptosource.CryptoSource) *generator {
	return &generator{
		rnd: rand.New(csrc),
	}
}

func (g *generator) GenPassword(partLen, partsCount int) (pass string) {
	alphaLower := "abcdefghijklmnopqrstuvwxyz"
	alphaUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	all := alphaLower + alphaUpper + numbers
	symbols := strings.Split(all, "")
	separator := "-"

	for i := 0; i < partsCount; i++ {
		pass += g.genPart(symbols, partLen)
		if i != partsCount-1 {
			pass += separator
		}
	}
	return
}

func (g *generator) genPart(a []string, length int) (part string) {
	g.rnd.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for i := 0; i < length; i++ {
		part += a[g.rnd.Intn(len(a))]
	}
	return
}
