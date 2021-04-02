package options

import (
	"log"
	"net/url"
)

// URLValue is a custom data type for parsing URL values with the 'flag' package
type URLValue struct {
	URL *url.URL
}

func MustMakeURLValue(s string) URLValue {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatalf("Error parsing URL '%s', error; %v", s, err)
	}
	return URLValue{u}

}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}
