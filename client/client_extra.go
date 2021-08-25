package client

import (
	"crypto/tls"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHTTPClientWithConfig creates a new device42 HTTP client,
// using a customizable transport config and auth.
func NewHTTPClientWithConfigAndAuth(formats strfmt.Registry, cfg *TransportConfig, username, password, useragent string, verify bool) *Device42 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: verify},
	}
	httpClient := &http.Client{Transport: tr}

	// create transport and client
	transport := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, httpClient)
	transport.DefaultAuthentication = httptransport.BasicAuth(username, password)
	transport.Transport = SetUserAgent(transport.Transport, useragent)
	return New(transport, formats)
}
