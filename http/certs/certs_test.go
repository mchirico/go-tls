package certs

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCertFile_KeyFile(t *testing.T) {

	cf := NewCertFile()
	cf.AddPath("./fixtures").AddPath("./http/certs/fixtures")

	file, err := cf.KeyFile()
	if err != nil {
		t.Fatalf("Can't get file")
	}
	dat, err := ioutil.ReadFile(file)
	if !strings.Contains(string(dat), "dummy tls.key") {
		t.Fatalf("Bad info in file")
	}

	cf.ClearPath()
	file, err = cf.KeyFile()
	if err == nil {
		t.Fatalf("not working")
	}
}

func TestCertFile_CertFile(t *testing.T) {

	cf := NewCertFile()
	cf.AddPath("./fixtures").AddPath("./http/certs/fixtures")

	file, err := cf.CertFile()
	if err != nil {
		t.Fatalf("Can't get file")
	}
	dat, err := ioutil.ReadFile(file)
	if !strings.Contains(string(dat), "dummy tls.cert") {
		t.Fatalf("Bad info in file")
	}

	cf.ClearPath()
	file, err = cf.KeyFile()
	if err == nil {
		t.Fatalf("not working")
	}
}
