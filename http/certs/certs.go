package certs

type CertFile struct {
	certFile string
	keyFile  string
	path     []string
}

func NewCertFile() *CertFile {
	cf := &CertFile{"tls.cert", "tls.key", []string{}}
	return cf
}

func (cf *CertFile) AddPath(path string) *CertFile {
	cf.path = append(cf.path, path)
	return cf
}
