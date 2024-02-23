package packages

import (
	"net"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	s := "postgres://myusername:mypassword@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	ptr()
	ptr(
		// "postgres://myusername:mypassword@host.com:5432/path?k=v#f"
		u,
		// postgres
		u.Scheme,
		// myusername:mypassword
		u.User,
		// myusername
		u.User.Username(),
		// /path
		u.Path,
		// f
		u.Fragment,
		// k=v
		u.RawQuery,
		// host.com:5432
		u.Host,
	)

	p, _ := u.User.Password()
	// mypassword
	ptr(p)

	host, port, _ := net.SplitHostPort(u.Host)
	ptr(
		// host.com
		host,
		// 5432
		port,
	)

	m, _ := url.ParseQuery(u.RawQuery)
	ptr(
		// map[k:[v]]
		m,
		// v
		m["k"][0],
	)
}
