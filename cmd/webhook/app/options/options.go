/*
Copyright 2020 The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	"strings"

	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"

	config "github.com/jetstack/cert-manager/internal/apis/config/webhook"
	configscheme "github.com/jetstack/cert-manager/internal/apis/config/webhook/scheme"
	configv1alpha1 "github.com/jetstack/cert-manager/pkg/apis/config/webhook/v1alpha1"
)

// WebhookFlags defines options that can only be configured via flags.
type WebhookFlags struct {
	// Path to a file containing a WebhookConfiguration resource
	Config string
}

func NewWebhookFlags() *WebhookFlags {
	return &WebhookFlags{}
}

func (f *WebhookFlags) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&f.Config, "config", "", "Path to a file containing a WebhookConfiguration object used to configure the webhook")
}

func ValidateWebhookFlags(f *WebhookFlags) error {
	// No validation needed today
	return nil
}

func NewWebhookConfiguration() (*config.WebhookConfiguration, error) {
	scheme, _, err := configscheme.NewSchemeAndCodecs()
	if err != nil {
		return nil, err
	}
	versioned := &configv1alpha1.WebhookConfiguration{}
	scheme.Default(versioned)
	config := &config.WebhookConfiguration{}
	if err := scheme.Convert(versioned, config, nil); err != nil {
		return nil, err
	}
	return config, nil
}

func AddConfigFlags(fs *pflag.FlagSet, c *config.WebhookConfiguration) {
	fs.IntVar(c.SecurePort, "secure-port", *c.SecurePort, "port number to listen on for secure TLS connections")
	fs.IntVar(c.HealthzPort, "healthz-port", *c.HealthzPort, "port number to listen on for insecure healthz connections")

	fs.StringVar(&c.TLSConfig.Filesystem.CertFile, "tls-cert-file", c.TLSConfig.Filesystem.CertFile, "path to the file containing the TLS certificate to serve with")
	fs.StringVar(&c.TLSConfig.Filesystem.KeyFile, "tls-private-key-file", c.TLSConfig.Filesystem.KeyFile, "path to the file containing the TLS private key to serve with")

	fs.StringVar(&c.TLSConfig.Dynamic.SecretNamespace, "dynamic-serving-ca-secret-namespace", c.TLSConfig.Dynamic.SecretNamespace, "namespace of the secret used to store the CA that signs serving certificates")
	fs.StringVar(&c.TLSConfig.Dynamic.SecretName, "dynamic-serving-ca-secret-name", c.TLSConfig.Dynamic.SecretName, "name of the secret used to store the CA that signs serving certificates certificates")
	fs.StringSliceVar(&c.TLSConfig.Dynamic.DNSNames, "dynamic-serving-dns-names", c.TLSConfig.Dynamic.DNSNames, "DNS names that should be present on certificates generated by the dynamic serving CA")

	fs.StringVar(&c.KubeConfig, "kubeconfig", c.KubeConfig, "optional path to the kubeconfig used to connect to the apiserver. If not specified, in-cluster-config will be used")
	fs.StringVar(&c.APIServerHost, "api-server-host", c.APIServerHost, ""+
		"Optional apiserver host address to connect to. If not specified, autoconfiguration "+
		"will be attempted.")
	fs.BoolVar(&c.EnablePprof, "enable-profiling", c.EnablePprof, ""+
		"Enable profiling for controller.")
	fs.StringVar(&c.PprofAddress, "profiler-address", c.PprofAddress,
		"Address of the Go profiler (pprof). This should never be exposed on a public interface. If this flag is not set, the profiler is not run.")
	tlsCipherPossibleValues := cliflag.TLSCipherPossibleValues()
	fs.StringSliceVar(&c.TLSConfig.CipherSuites, "tls-cipher-suites", c.TLSConfig.CipherSuites,
		"Comma-separated list of cipher suites for the server. "+
			"If omitted, the default Go cipher suites will be use.  "+
			"Possible values: "+strings.Join(tlsCipherPossibleValues, ","))
	tlsPossibleVersions := cliflag.TLSPossibleVersions()
	fs.StringVar(&c.TLSConfig.MinTLSVersion, "tls-min-version", c.TLSConfig.MinTLSVersion,
		"Minimum TLS version supported. "+
			"Possible values: "+strings.Join(tlsPossibleVersions, ", "))

}
