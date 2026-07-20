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

func SpacePresent(s string) error{
	for _,ch:=range s{
		if ch==' '{
			return errors.New("Blank space found in URL")
		}
	}
	return nil
}

func LengthShortURL(s string) error{
	if len(s)-9!=ds.Size{
		return errors.New("Short URL must be exactly test.com/ and "+string(ds.Size+48)+" characters long.")
	}
	return nil
}