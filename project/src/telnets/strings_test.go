package main


import (
	"testing"
)



func TestExtractUsernameAndHost(t *testing.T) {

	tests := []struct{
		String string

		ExpectedUsername string
		ExpectedHost     string
	}{
		{
			String:            "username@host",

			ExpectedUsername: "username",
			ExpectedHost:     "host",
		},



		{
			String:            "host",

			ExpectedUsername: "",
			ExpectedHost:     "host",
		},
	}


	for testNumber, test := range tests {

		username, host := extractUsernameAndHost(test.String)

		if expected, actual := test.ExpectedUsername, username; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", expected, actual)
			continue
		}

		if expected, actual := test.ExpectedHost, host; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
