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

func TestJoinFiles(t *testing.T) {
	list1 := mustGetLabelBoxExportEntity("test.json")
	t.Log(len(list1))
	list2 := mustGetLabelBoxExportEntity("test1.json")
	t.Log(len(list2))

	for _, x := range list1 {
		list2 = append(list2, x)
	}

	if err := writeTextAndClassificationToTsv(list2, "data.tsv"); err != nil {
		t.Error(err)
	}

	//var count int
	//for _, entity := range list2 {
	//	cf := entity.Label.Classifications
	//	if len(cf) > 0 {
	//		count++
	//		t.Log(entity.LabeledData)
	//		t.Log(cf[0].Answer.Title)
	//	}
	//}
	//t.Log(len(list2))
	//t.Log(count)
}

func TestReadGptGenerated(t *testing.T) {
	data, err := os.ReadFile("export-2023-04-12T12_27_06.859Z.json")
	if err != nil {
		t.Fatal(err)
	}
	var result []EntityRecord
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Error(err)
		return
	}
	rst := EntityRecordListPreProcessing(result)
	rb, _ := json.Marshal(rst)
	err = ioutil.WriteFile("test1.json", rb, 0644)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestReadOld(t *testing.T) {
	data, err := os.ReadFile("export-2023-04-03T16_51_10.660Z.json")
	if err != nil {
		t.Error(err)
		return
	}

	var result []EntityRecord
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result[0].LabeledData)

	rst := EntityRecordListPreProcessing(result)
	t.Log(rst[0].LabeledData)

	rb, _ := json.Marshal(rst)
	err = ioutil.WriteFile("test.json", rb, 0644)

	if err != nil {
		t.Error(err)
		return
	}
}

func TestReadNERREsult(t *testing.T) {
	data, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}

	var result []EntityRecord
	err = json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	//result := EntityRecordListPreProcessing(rst)
	//if len(result) > 0 {
	//	t.Log(result[0].LabeledData)
	//}
	//
	//outputData, err := json.Marshal(result)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ioutil.WriteFile("test.json", outputData, 0644)
	//if err != nil {
	//	panic(err)
	//}

	var validBIO []BIO

	for _, file := range result {
		objectString := file.LabeledData
		objectDL := file.GetObjectsDetail()
		if len(objectDL) > 0 {
			bio := BIO{
				Text:    objectString,
				Objects: objectDL,
			}
			if checkOverlap(bio.Objects) {
				// overlap 的数据暂时跳过
			} else {
				validBIO = append(validBIO, bio)
			}
			//runes := []rune(objectString)
			//// 使用切片注意越界问题
			//t.Log(string(runes[objectDL[0].Start : objectDL[0].End+1]))
		}
	}
	err = bioToTensorFlowFormat(validBIO, "output.txt")
	if err != nil {
		t.Error(err)
	}
}

// TestWriteIntentionResult 将 labelbox 标注数据转化为 LSTM 意图训练所需要的数据格式
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

// TestGetLabeledData 从返回的数据 labelData 读取文件内容
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
