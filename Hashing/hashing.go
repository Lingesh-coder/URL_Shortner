package hashing

import (
	"math/rand"
	ds "url/dss"
)

func big(s string) string{
	return s+string(rand.Intn(26)+65)
}

func small(s string) string{
	return s+string(rand.Intn(26)+97)
}

func num(s string) string{
	return s+string(rand.Intn(10)+48)
}

func Hash() string{
	d:=""
	for i:=0;i<ds.Size;i++{
		e:=rand.Intn(3)
		if e==1{
			d=big(d)
		} else if e==2{
			d=small(d)
		} else{
			d=num(d)
		}
	}
	return "test.com/"+d
}