package ner_spliter

import (
	"regexp"
	"testing"
)

func TestSplitText(t *testing.T) {
	Run()
}

func TestRegexReplace(t *testing.T) {
	f := func(line string) string {
		re := regexp.MustCompile(`^\d+.\s+`)
		return re.ReplaceAllString(line, "")
	}
	t.Log(f("36. Kibana 是 Elasticsearch 的数据可视化工具。"))
	t.Log(f("GitLab CI/CD 允许用户自定义 CI/CD 流程以满足特定需求。"))
}
