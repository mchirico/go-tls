package certs

import (
	"os"
)

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

func (cf *CertFile) ClearPath() {
	cf.path = cf.path[:0]
}

func (cf *CertFile) KeyFile() (string, error) {
	for _, path := range cf.path {
		filename := path + "/" + cf.keyFile
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		}
		return filename, nil
	}

	return "", os.ErrNotExist
}

func (cf *CertFile) CertFile() (string, error) {
	for _, path := range cf.path {
		filename := path + "/" + cf.certFile
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		}
		return filename, nil
	}

	return "", os.ErrNotExist
}
