package service

import "testing"

func TestReadFiles(t *testing.T) {
	rst, err := readFileToLine("/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/gitRepoCmd1.txt")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(rst)
	}
}
