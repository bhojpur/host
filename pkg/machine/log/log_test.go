package log

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripSecrets(t *testing.T) {
	testCases := []struct {
		description string
		input       []string
		expected    []string
	}{
		{
			description: "Log that does contain certs should have them stripped",
			input: []string{
				"Some mundane log lines",
				"IP is foo.bar",
				`Secret here: printf '%s' '-----BEGIN CERTIFICATE-----
MIIC4DCCAcigAwIBAgIRAMMHbb4WFRVYsCOIrfM3dqkwDQYJKoZIhvcNAQELBQAw
GTEXMBUGA1UEChMObmF0aGFubGVjbGFpcmUwHhcNMTUxMDEwMDE1MDAwWhcNMTgw
OTI0MDE1MDAwWjAZMRcwFQYDVQQKEw5uYXRoYW5sZWNsYWlyZTCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBANLMyaAZPThE6lXtXYfUMZeF0pEfO4BQ7Rv8
Q9/aIKwm8SlKNm+g+6+RoexsiaPXmAkqk04kg+f9WRgtUKC3nhaiUwTqx2HtxowY
Kp7VVW9QyzwCP1r04WTNTdICzhwM5GfaCMKLmibVUfh9GqIYg4Z6eFly7t0PaN1P
uaLClow1e4sWgAgkpIx7ko9ZtW+73knAnp9PPoH4KPBLS+sIPNGh62WsDlvQrOnq
KDiBPIAAMxu2UefIPeGe6xxFuCG89RoJYYsB627IaR8R8iGJMwjJsiAiObGu6z8M
lcWxT4dC+cEIDRu+XQmavJlAydBeHY6/gtJXzsyRExHTyDwi8xkCAwEAAaMjMCEw
DgYDVR0PAQH/BAQDAgKsMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQAD
ggEBAA5CBXPgjvxfY5bR+f6YfcDcKBWxOQ5zN+OH6jWpVzJMEUWp/ZvTQ1GcV1CT
J4HDMRUOL6lQigZDKR6OJ0g/pD4cDGEQlCuPDXx0O8eenxj9TQ+X+qdtxQNkgjId
QWj3k3JDHCh4BQ7h1ZJIg4SnGCUsrQQ+M8TS4Z0YZ/bZ6ZTktJgQgWMn9Uum1GN9
hXJ/fa/E9OJuRxTXou7J0WwrV9aX9sEM9syOANR88PcA1fSE7+wNSdj5ZCfY6mQn
II9e8NZEf5ktPXCNi0LKI6R1berejwQI3KKHEFbdZ8SKn93HgDh/Ip/dFctj+zBt
CAlTWS3abehlCERn6Ze9IfZBtpI=
-----END CERTIFICATE-----' | sudo tee /etc/bhojpur/ca.pem`,
			},
			expected: []string{
				"Some mundane log lines",
				"IP is foo.bar",
				`Secret here: printf '%s' '<REDACTED>' | sudo tee /etc/bhojpur/ca.pem`,
			},
		},
		{
			description: "Log that does contain private keys should have them stripped",
			input: []string{
				"Some mundane log lines",
				"IP is foo.bar",
				`Secret here: printf '%s' '-----BEGIN RSA PRIVATE KEY-----
MIIC4DCCAcigAwIBAgIRAMMHbb4WFRVYsCOIrfM3dqkwDQYJKoZIhvcNAQELBQAw
GTEXMBUGA1UEChMObmF0aGFubGVjbGFpcmUwHhcNMTUxMDEwMDE1MDAwWhcNMTgw
OTI0MDE1MDAwWjAZMRcwFQYDVQQKEw5uYXRoYW5sZWNsYWlyZTCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBANLMyaAZPThE6lXtXYfUMZeF0pEfO4BQ7Rv8
Q9/aIKwm8SlKNm+g+6+RoexsiaPXmAkqk04kg+f9WRgtUKC3nhaiUwTqx2HtxowY
Kp7VVW9QyzwCP1r04WTNTdICzhwM5GfaCMKLmibVUfh9GqIYg4Z6eFly7t0PaN1P
uaLClow1e4sWgAgkpIx7ko9ZtW+73knAnp9PPoH4KPBLS+sIPNGh62WsDlvQrOnq
KDiBPIAAMxu2UefIPeGe6xxFuCG89RoJYYsB627IaR8R8iGJMwjJsiAiObGu6z8M
lcWxT4dC+cEIDRu+XQmavJlAydBeHY6/gtJXzsyRExHTyDwi8xkCAwEAAaMjMCEw
DgYDVR0PAQH/BAQDAgKsMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQAD
ggEBAA5CBXPgjvxfY5bR+f6YfcDcKBWxOQ5zN+OH6jWpVzJMEUWp/ZvTQ1GcV1CT
J4HDMRUOL6lQigZDKR6OJ0g/pD4cDGEQlCuPDXx0O8eenxj9TQ+X+qdtxQNkgjId
QWj3k3JDHCh4BQ7h1ZJIg4SnGCUsrQQ+M8TS4Z0YZ/bZ6ZTktJgQgWMn9Uum1GN9
hXJ/fa/E9OJuRxTXou7J0WwrV9aX9sEM9syOANR88PcA1fSE7+wNSdj5ZCfY6mQn
II9e8NZEf5ktPXCNi0LKI6R1berejwQI3KKHEFbdZ8SKn93HgDh/Ip/dFctj+zBt
CAlTWS3abehlCERn6Ze9IfZBtpI=
-----END RSA PRIVATE KEY-----' | sudo tee /etc/bhojpur/server-key.pem`,
			},
			expected: []string{
				"Some mundane log lines",
				"IP is foo.bar",
				`Secret here: printf '%s' '<REDACTED>' | sudo tee /etc/bhojpur/server-key.pem`,
			},
		},
		{
			description: "Log that does not contain secrets should not change",
			input: []string{
				"Some mundane log lines",
				"IP is foo.bar",
			},
			expected: []string{
				"Some mundane log lines",
				"IP is foo.bar",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, stripSecrets(tc.input))
	}
}
