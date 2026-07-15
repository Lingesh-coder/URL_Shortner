package errors
import (
	"errors"
	ds "url/dss"
)
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