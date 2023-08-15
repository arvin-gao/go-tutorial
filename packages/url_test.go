package packages

import (
	"net"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	ptr(u.Scheme)

	ptr(u.User)
	ptr(u.User.Username())
	p, _ := u.User.Password()
	ptr(p)

	ptr(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	ptr(host)
	ptr(port)

	ptr(u.Path)
	ptr(u.Fragment)

	ptr(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	ptr(m)
	ptr(m["k"][0])
}
