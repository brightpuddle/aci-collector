package ssh

import (
	"testing"

	"github.com/brightpuddle/aci-collector/pkg/req"
	"github.com/stretchr/testify/assert"
)

func TestParseHost(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"bare IP", "192.168.1.1", "192.168.1.1"},
		{"https scheme", "https://192.168.1.1", "192.168.1.1"},
		{"http scheme", "http://192.168.1.1", "192.168.1.1"},
		{"IP with port", "192.168.1.1:443", "192.168.1.1"},
		{"https with port", "https://192.168.1.1:443", "192.168.1.1"},
		{"hostname", "apic.example.com", "apic.example.com"},
		{"https hostname", "https://apic.example.com", "apic.example.com"},
		{"hostname with trailing slash", "https://apic.example.com/", "apic.example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseHost(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCLIRequestsValid(t *testing.T) {
	a := assert.New(t)

	for _, r := range req.CLIRequests {
		a.NotEmpty(r.Command, "CLIRequest.Command must not be empty")
		a.NotEmpty(r.Filename, "CLIRequest.Filename must not be empty")
		a.True(
			len(r.Filename) > 4 && r.Filename[len(r.Filename)-4:] == ".txt",
			"CLIRequest.Filename must end in .txt: %s", r.Filename,
		)
	}
}
