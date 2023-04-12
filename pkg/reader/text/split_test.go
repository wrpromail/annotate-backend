package text

import "testing"

// /Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/gitRepoCmd1.txt
func TestSplitLocalFile(t *testing.T) {
	lf := LocalFileSplit{
		path: "/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/gitRepoCmd1.txt",
		out:  "/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/out2",
	}
	lf.Split()
}
