package cookie

import (
	"net/http"
	"net/url"
)

func CreateCookie(name string, keys []string, values []string, maxAge int, path, domain string, sameSite http.SameSite,
	security, httpOnly bool) string {
	return NewCookie(name, NewParam(keys, values).Encode(), maxAge, path, domain, sameSite, security, httpOnly).String()

}

func NewParam(keys []string, values []string) url.Values {
	params := url.Values{}
	l := len(values)
	for i, v := range keys {
		if i < l {
			params.Add(v, values[i])
		} else {
			params.Add(v, "")
		}
	}
	return params
}

func NewCookie(name string, value string, maxAge int, path, domain string, sameSite http.SameSite,
	security, httpOnly bool) *http.Cookie {
	// add to header
	return &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge, // 30 day
		Path:     path,
		Domain:   domain,
		SameSite: sameSite,
		Secure:   security,
		HttpOnly: httpOnly,
	}
}