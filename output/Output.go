package output

import (
	ds "url/dss"
	"os"
	"strings"
	"bufio"
)

func Put_URL() (string,bool) {
	ss:=bufio.NewReader(os.Stdin)
	s,_:=ss.ReadString('\n')
	s=strings.TrimSpace(s)
	for i,j:=range ds.Ma{
		if j==s{
			return i,true
		}
	}
	return "NO URL found",false
}