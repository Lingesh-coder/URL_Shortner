package errors
import (
	"errors"
	ds "url/dss"
	"os"
	"strconv"
	"bufio"
	"strings"
)

func TextToInteger() (int,error){
	scan:=bufio.NewReader(os.Stdin)
	scanner,_:=scan.ReadString('\n')
	scanner=strings.TrimSpace(scanner)
	r,err:=strconv.Atoi(scanner)
	if(err!=nil){
		return 0,errors.New("Enter numbers only")
	}
	return r,nil
}

func Format(s string) error {
	if ds.DomainCheck.MatchString(s){
		return nil
	} else{
		return errors.New("Format error")
	}
}

func LengthURL(s string) error{
	if len(s)>ds.MaxURLLen{
		return errors.New("URL too long.")
	}
	return nil
}

func SpacePresent(s string) error{
	for _,ch := range s{
		if ch==' '{
			return errors.New("Blank space found in URL")
		}
	}
	return nil
}
