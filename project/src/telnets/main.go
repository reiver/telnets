package main


import (
	"github.com/reiver/go-telnet"

	"crypto/tls"
	"fmt"
	"io"
	"os"
)


func main() {

	if len(os.Args) <= 1 {
		help(os.Stderr)
		os.Exit(1)
		return
	}



	tlsConfig := &tls.Config{
//@TODO: What does this actually do (in a deep sense)? Is this a security issue?
		InsecureSkipVerify:true,
	}



	var conn *telnet.Conn
	var err error
	var username string
	var host string

	switch len(os.Args) {
	case 1:
		conn, err = telnet.DialTLS(tlsConfig)
	case 2:
		username, host = extractUsernameAndHost(os.Args[1])

		addr := fmt.Sprintf("%s:telnets", host)

		conn, err = telnet.DialToTLS(addr, tlsConfig)
	case 3:
		username, host = extractUsernameAndHost(os.Args[1])

		addr := fmt.Sprintf("%s:%s", host, os.Args[2])

		conn, err = telnet.DialToTLS(addr, tlsConfig)
	}
	if nil != err {
		fmt.Fprintf(os.Stderr, "telnets: Unable to connect to remote host: %v\n", err)
		os.Exit(1)
		return
	}


	fmt.Fprintf(os.Stdout, "Connected to TELNETS (secure TELNET) server at %q.\n", conn.RemoteAddr())


	client := &telnet.Client{}
	client.SetAuth(username)

	if err := client.Call(conn); nil != err {
		fmt.Fprintf(os.Stderr, "Problem calling TELNETS (secure TELNET) server: %v\n", err)
		os.Exit(1)
		return
	}



	os.Exit(0)
	return
}



func help(w io.Writer) {
	fmt.Fprint(w, "Usage: telnets [user@]host [port]\n")
}
