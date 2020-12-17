package certs

import "testing"

func TestCertFile_AddPath(t *testing.T) {

	cf := NewCertFile()
	cf.AddPath("a").AddPath("b")

}
