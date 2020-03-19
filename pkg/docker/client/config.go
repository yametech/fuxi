package client

//DockerConfig  config the docker
type DockerConfig struct {
	Host               string
	TLS                bool
	CertDir            string
	CertPassword       string
	InsecureSkipVerify bool
}
