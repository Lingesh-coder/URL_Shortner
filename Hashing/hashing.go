package hashing
import (
	"math/rand"
	ds "url/dss"
)
func Hash() string {
	shortUrlString := ""
	for i := 0; i < ds.Size; i++ {
		switch i % 3 {
		case 1:
			shortUrlString += string(rune(rand.Intn(26) + 'A'))
		case 2:
			shortUrlString += string(rune(rand.Intn(26) + 'a'))
		default:
			shortUrlString += string(rune(rand.Intn(10) + '0'))
		}
	}
	shortUrlString = "test.com/" + shortUrlString
	if _,kk := ds.ShortUrl[shortUrlString];kk{
		Hash()
		return ""
	}
	return shortUrlString
}
