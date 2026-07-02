package hashing

import (
	"math/rand"
	"fmt"
)

func Hash(s string) string{
	var r int64
	for _,ch:=range s{
		r+=int64(ch)+int64(rand.Intn(1000000000000000000))
	}
	return fmt.Sprintf("%x",r);
}