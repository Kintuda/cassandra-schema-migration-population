package migrator

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"path/filepath"

	"github.com/gocql/gocql"
)

type Migrator struct {
	Session *gocql.Session
}

func NewMigrator(session *gocql.Session) *Migrator {
	return &Migrator{Session: session}
}

func GetTLSConfig(path string) gocql.SslOptions {
	certPath, _ := filepath.Abs(path + "/cert.pfx")
	keyPath, _ := filepath.Abs(path + "/key")
	caPath, _ := filepath.Abs(path + "/ca.crt")

	cert, _ := tls.LoadX509KeyPair(certPath, keyPath)
	caCert, _ := ioutil.ReadFile(caPath)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return gocql.SslOptions{
		Config:                 tlsConfig,
		EnableHostVerification: false,
	}
}

func CreateConnection(contactPoints string, tlsPath *string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(contactPoints)
	cluster.Hosts = []string{contactPoints + ":" + "9042"}
	session, err := cluster.CreateSession()

	if tlsPath != nil {
		options := GetTLSConfig(*tlsPath)
		cluster.SslOpts = &options
	}

	return session, err
}

func (m *Migrator) CheckSchema(keyspace string, table string) *gocql.Query {
	return m.Session.Query("SELECT * FROM system_schema.columns WHERE keyspace_name = $1 AND table_name = $2", keyspace, table)
}
