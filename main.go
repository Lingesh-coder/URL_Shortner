package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
	ds "url/dss"
)

func hash(s string) string{
	var r int64
	for _,ch:=range s{
		r+=int64(ch)+int64(rand.Intn(1000000000000000000))
	}
	return fmt.Sprintf("%x",r);
}
func main(){
	outer:
	for{
		var ch rune
		fmt.Println("1. Enter Y/y to get a shortened URL\n2. Enter N/n to exit")
		fmt.Scanf("%c",&ch)
		switch(ch){
		case 'Y','y':
			fmt.Println("Enter the URL:")
			ss:=bufio.NewReader(os.Stdin)
			ss.ReadString('\n')
			s,_:=ss.ReadString('\n')
			s=strings.TrimSpace(s)
			if _,ok:=ds.Ma[s];!ok{
				ds.Ma[s]=hash(s)[:7]
			}
			fmt.Println("Shortened URL is:",ds.Ma[s])
		case 'N','n':
			fmt.Println("Exiting...")
			break outer
		default:
			fmt.Println("Enter character again")
		}
	}
}