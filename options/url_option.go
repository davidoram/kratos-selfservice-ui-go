package options

import (
	"net/url"
)

type URLValue struct {
	URL     *url.URL
	Default string // Default URL value if none provided
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if s == "" {
		s = v.Default
	}
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}
