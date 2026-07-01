package input

import (
	"fmt"
	ds "url/dss"
	hash "url/Hashing"
	"strings"
	"bufio"
	"os"
)

func Get_URL(){
	fmt.Println("Enter the URL:")
	ss:=bufio.NewReader(os.Stdin)
	ss.ReadString('\n')
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	if _,ok:=ds.Ma[s];!ok{
		ds.Ma[s]=hash.Hash(s)[:7]
	}
	fmt.Println("Shortened URL is:",ds.Ma[s])
}