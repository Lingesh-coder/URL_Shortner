package input

import (
	"fmt"
	ds "url/dss"
	hash "url/Hashing"
	"strings"
	"bufio"
	"os"
	//"regexp"
)
// Yet to use regexp
func format(s string) string{
	return s
}

func Get_URL(){
	fmt.Println("Enter the URL:")
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	s=format(s)
	if _,ok:=ds.Ma[s];!ok{
		er:=""
		for {
			er=hash.Hash()
			if _,okk:=ds.Am[er];!okk{
				break
			}
		}
		ds.Ma[s]=er
		ds.Am[er]=s
	}
	fmt.Println("Shortened URL is:",ds.Ma[s])
}
