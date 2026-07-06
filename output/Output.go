package output

import (
	ds "url/dss"
	"os"
	"strings"
	"bufio"
	"fmt"
)

func Put_URL() string {
	fmt.Println("Enter shortened URL:")
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	if y,ok:=ds.Am[s];ok{
		return "Original URL is: "+y
	}
	return "NO URL found"
}