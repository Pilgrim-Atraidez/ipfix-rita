package yaml

import (
	"github.com/activecm/ipfix-rita/converter/config"
	"github.com/activecm/mgosec"
	"github.com/pkg/errors"
)

//mongoDB implements config.MongoDB
type mongoDB struct {
	ConnectionString string `yaml:"ConnectionString"`
	AuthMechanism    string `yaml:"AuthenticationMechanism"`
	TLS              tls    `yaml:"TLS"`
	Database         string `yaml:"Database"`
	Collection       string `yaml:"Collection"`
}

func (m *mongoDB) GetConnectionString() string {
	return m.ConnectionString
}

func (m *mongoDB) GetAuthMechanism() (mgosec.AuthMechanism, error) {
	mechanism, err := mgosec.ParseAuthMechanism(m.AuthMechanism)
	return mechanism, errors.Wrapf(err, "could not parse MongoDB AuthMechanism: %s", m.AuthMechanism)
}

func (m *mongoDB) GetTLS() config.TLS {
	return &m.TLS
}

//tls implements config.TLS
type tls struct {
	Enabled           bool   `yaml:"Enable"`
	VerifyCertificate bool   `yaml:"VerifyCertificate"`
	CAFile            string `yaml:"CAFile"`
}

func (t *tls) IsEnabled() bool {
	return t.Enabled
}

func (t *tls) ShouldVerifyCertificate() bool {
	return t.VerifyCertificate
}

func (t *tls) GetCAFile() string {
	return t.CAFile
}

func (m *mongoDB) GetDatabase() string {
	return m.Database
}

func (m *mongoDB) GetCollection() string {
	return m.Collection
}
