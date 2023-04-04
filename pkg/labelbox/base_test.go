package labelbox

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestWriteIntentionResult(t *testing.T) {
	var rst []IntentionLabeledPair

	data, err := os.ReadFile("export-2023-04-03T16_51_10.660Z.json")
	if err != nil {
		panic(err)
	}

	// unmarshal
	var result []EntityRecord
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	for _, entity := range result {
		text, e := entity.ReadLabeledDataContent()
		if e != nil {
			log.Error(e)
			continue
		}

		rst = append(rst, IntentionLabeledPair{
			Text:   text,
			Intent: entity.GetClassificationValue(),
		})
	}

	t.Log(len(rst))
	jsonBytes, err := json.Marshal(rst)
	if err != nil {
		panic(err)
	}

	// 将JSON写入文件
	err = os.WriteFile("intent.json", jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func TestReadLabelBoxExportFile(t *testing.T) {
	data, err := os.ReadFile("export-2023-04-03T16_51_10.660Z.json")
	if err != nil {
		panic(err)
	}

	// unmarshal
	var result []EntityRecord
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}

	t.Log(len(result))
}

func TestGetLabeledData(t *testing.T) {
	url := "https://storage.labelbox.com/clfcef7c406tt08z6emsmg04q%2F8dacbd9b-c10f-7327-fe81-cfbe6490aead-BIGZcRAliu.txt?Expires=1681750270830&KeyName=labelbox-assets-key-3&Signature=lTud_45BCHlK5ERAsgFg7nkLVNY"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	// 解码 utf-8 数据
	t.Log(string(body))
}
