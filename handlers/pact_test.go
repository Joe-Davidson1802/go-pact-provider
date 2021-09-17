package handlers

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var version = os.Getenv("VERSION")
var pactbroker = os.Getenv("PACT_BROKER_BASE_URL")
var pact *dsl.Pact

func startServer() (port int) {
	handler := http.HandlerFunc(GetTimeHandler)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port = listener.Addr().(*net.TCPAddr).Port
	go func() {
		panic(http.Serve(listener, handler))
	}()
	return
}

func TestProvider(t *testing.T) {
	pact = &dsl.Pact{
		Provider: "time-back-end",
	}
	port := startServer()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		PactURLs:                   []string{pactbroker},
		PublishVerificationResults: true,
		ProviderVersion:            version,
	})
	if err != nil {
		t.Error(err)
	}
}
