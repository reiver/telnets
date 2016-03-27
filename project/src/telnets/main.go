package main


import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"

	"bufio"
	"bytes"
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



	var telnetsClient *telnet.Client
	var err error

	switch len(os.Args) {
	case 1:
		telnetsClient, err = telnet.DialTLS(tlsConfig)
	case 2:
		addr := fmt.Sprintf("%s:telnets", os.Args[1])

		telnetsClient, err = telnet.DialToTLS(addr, tlsConfig)
	case 3:
		addr := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

		telnetsClient, err = telnet.DialToTLS(addr, tlsConfig)
	}
	if nil != err {
		fmt.Fprintf(os.Stderr, "telnets: Unable to connect to remote host: %v", err)
		os.Exit(1)
		return
	}
	defer telnetsClient.Close()


	fmt.Fprintf(os.Stdout, "Connected to TELNETS (secure TELNET) server at %q.\n", telnetsClient.RemoteAddr())

	go func(writer io.Writer, reader io.Reader) {
//@TODO: What about the TELNET control characters? (What about the bell, for example.)
//       The client should deal with them too.

		var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
		p := buffer[:]

		for {
			// Read 1 byte.
			n, err := reader.Read(p)
			if n <= 0 && nil == err {
				continue
			} else if n <= 0 && nil != err {
				break
			}

			oi.LongWrite(writer, p)
		}
	}(os.Stdout, telnetsClient)



	var buffer bytes.Buffer
	var p []byte

	var crlfBuffer [2]byte = [2]byte{'\r','\n'}
	crlf := crlfBuffer[:]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.Write(crlf)

		p = buffer.Bytes()

		n, err := oi.LongWrite(telnetsClient, p)
		if nil != err {
			break
		}
		if expected, actual := int64(len(p)), n; expected != actual {
			fmt.Fprintf(os.Stderr, "Transmission problem: tried sending %d bytes, but actually only sent %d bytes.", expected, actual)
			os.Exit(1)
			return
		}


		buffer.Reset()
	}



	os.Exit(0)
	return
}



func help(w io.Writer) {
	fmt.Fprint(w, "Usage: telnets host [port]\n")
}
