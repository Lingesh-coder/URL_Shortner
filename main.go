package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/rand"
)
	var ma=make(map[string]string)

func hash(s string) string{
	var r int64
	for _,ch:=range s{
		r+=int64(ch)+int64(rand.Intn(1000000000000000000))
	}
	return fmt.Sprintf("%x",r);
}
func main(){
	fmt.Println("Enter the URL:")
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	if _,ok:=ma[s];!ok{
		ma[s]=hash(s)[:7]
	}
	fmt.Println("Shortened URL is:",ma[s])
}