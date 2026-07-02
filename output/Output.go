package output

import (
	ds "url/dss"
	"os"
	"strings"
	"bufio"
)

func Put_URL() string {
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	if y,ok:=ds.Am[s];ok{
		return "Original URL is: "+y
	}
	return "NO URL found"
}