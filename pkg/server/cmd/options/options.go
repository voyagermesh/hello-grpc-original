package options

import (
	"github.com/spf13/pflag"
)

type Options struct {
	SecureAddr               string
	PlaintextAddr            string
	EnableJavaClient         bool
	APIDomain                string
	CACertFile               string
	CertFile                 string
	KeyFile                  string
	EnableCORS               bool
	CORSOriginHost           string
	CORSOriginAllowSubdomain bool
	OpsAddress               string
}

func New() *Options {
	return &Options{
		SecureAddr:    ":8443",
		PlaintextAddr: ":8080",
		OpsAddress:    ":56790",
	}
}

func (opt *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&opt.SecureAddr, "secure-addr", opt.SecureAddr, "host:port used to server secure apis")
	fs.StringVar(&opt.PlaintextAddr, "plaintext-addr", opt.PlaintextAddr, "host:port used to server plaintext apis")

	fs.BoolVar(&opt.EnableJavaClient, "enable-java-client", opt.EnableJavaClient, "Set true to send SETTINGS frame from the server. Default set to false")
	fs.StringVar(&opt.APIDomain, "api-domain", opt.APIDomain, "Domain used to server hello api")
	fs.StringVar(&opt.CACertFile, "caCertFile", opt.CACertFile, "File containing CA certificate")
	fs.StringVar(&opt.CertFile, "certFile", opt.CertFile, "File container server TLS certificate")
	fs.StringVar(&opt.KeyFile, "keyFile", opt.KeyFile, "File containing server TLS private key")

	fs.BoolVar(&opt.EnableCORS, "enable-cors", opt.EnableCORS, "Enable CORS support")
	fs.StringVar(&opt.CORSOriginHost, "cors-origin-host", opt.CORSOriginHost, `Allowed CORS origin host e.g, domain[:port]`)
	fs.BoolVar(&opt.CORSOriginAllowSubdomain, "cors-origin-allow-subdomain", opt.CORSOriginAllowSubdomain, "Allow CORS request from subdomains of origin")

	fs.StringVar(&opt.OpsAddress, "ops-addr", opt.OpsAddress, "Address to listen on for web interface and telemetry.")
}
