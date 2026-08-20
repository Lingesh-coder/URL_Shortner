package hashing

import (
	"math/rand"
	ds "url/dss"
)

func Hash() string {
	d := ""
	for i := 0; i < ds.Size; i++ {
		switch i % 3 {
		case 1:
			d += string(rune(rand.Intn(26) + 'A'))
		case 2:
			d += string(rune(rand.Intn(26) + 'a'))
		default:
			d += string(rune(rand.Intn(10) + '0'))
		}
	}
	return "test.com/" + d
}
