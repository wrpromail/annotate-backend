package labelbox

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type IntentionLabeledPair struct {
	Text   string `json:"text"`
	Intent string `json:"intent"`
}

func (e *EntityRecord) GetClassificationValue() string {
	classification := e.Label.Classifications
	if len(classification) == 0 {
		return ""
	}
	return classification[0].Answer.Value
}

func (e *EntityRecord) ReadLabeledDataContent() (result string, err error) {
	resp, err := http.Get(e.LabeledData)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 读取响应数据

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	result = string(body)
	return
}

type BIO struct {
	Text    string         `json:"text"`
	Objects []ObjectSimple `json:"objects"`
}

type ObjectSimple struct {
	ClassName string `json:"className"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
}

func checkOverlap(objects []ObjectSimple) bool {
	for i, obj1 := range objects {
		for j, obj2 := range objects {
			if i != j && ((obj1.Start >= obj2.Start && obj1.Start < obj2.End) || (obj1.End > obj2.Start && obj1.End <= obj2.End)) {
				return true
			}
		}
	}
	return false
}

func bioToTensorFlowFormat(bios []BIO, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, bio := range bios {
		textRunes := []rune(bio.Text)
		labels := make([]string, len(textRunes))

		// Initialize labels with "O"
		for i := range labels {
			labels[i] = "O"
		}

		for _, obj := range bio.Objects {
			if obj.Start >= 0 && obj.End <= len(textRunes) {
				for i := obj.Start; i < obj.End; i++ {
					prefix := "I-"
					if i == obj.Start {
						prefix = "B-"
					}
					labels[i] = prefix + obj.ClassName
				}
			}
		}

		for i, r := range textRunes {
			_, _ = file.WriteString(fmt.Sprintf("%c %s\n", r, labels[i]))
		}
		_, _ = file.WriteString("\n") // Add an empty line between sentences
	}

	return nil
}

func (e *EntityRecord) GetObjectsDetail() (result []ObjectSimple) {
	objects := e.Label.Objects
	if len(objects) == 0 {
		return
	}

	for _, rawObject := range objects {
		rb, _ := json.Marshal(rawObject)
		var rst ObjectEntity
		err := json.Unmarshal(rb, &rst)
		if err != nil {
			log.Error(err)
			continue
		}

		r := ObjectSimple{
			ClassName: rst.Title,
			Start:     rst.Data.Location.Start,
			End:       rst.Data.Location.End,
		}
		result = append(result, r)

	}

	return
}

type EntityRecordArray []EntityRecord

func EntityRecordListPreProcessing(source []EntityRecord) (result []EntityRecord) {
	if len(source) == 0 {
		return
	}
	for _, entity := range source {
		if strings.HasPrefix(entity.LabeledData, "http") {
			var newEntity EntityRecord
			newEntity = entity
			resp, err := http.Get(entity.LabeledData)
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			newEntity.LabeledData = string(body)
			result = append(result, newEntity)
		} else {
			result = append(result, entity)
		}
	}
	return

}

type EntityRecord struct {
	ID          string `json:"ID"`
	DataRowID   string `json:"DataRow ID"`
	LabeledData string `json:"Labeled Data"`
	Label       struct {
		Objects         []interface{} `json:"objects"`
		Classifications []struct {
			FeatureID string `json:"featureId"`
			SchemaID  string `json:"schemaId"`
			Scope     string `json:"scope"`
			Title     string `json:"title"`
			Value     string `json:"value"`
			Answer    struct {
				FeatureID string `json:"featureId"`
				SchemaID  string `json:"schemaId"`
				Title     string `json:"title"`
				Value     string `json:"value"`
				Position  int    `json:"position"`
			} `json:"answer"`
		} `json:"classifications"`
		Relationships []interface{} `json:"relationships"`
	} `json:"Label"`
	CreatedBy           string        `json:"Created By"`
	ProjectName         string        `json:"Project Name"`
	CreatedAt           time.Time     `json:"Created At"`
	UpdatedAt           time.Time     `json:"Updated At"`
	SecondsToLabel      float64       `json:"Seconds to Label"`
	SecondsToReview     float64       `json:"Seconds to Review"`
	SecondsToCreate     float64       `json:"Seconds to Create"`
	ExternalID          string        `json:"External ID"`
	GlobalKey           interface{}   `json:"Global Key"`
	Agreement           int           `json:"Agreement"`
	IsBenchmark         int           `json:"Is Benchmark"`
	BenchmarkAgreement  int           `json:"Benchmark Agreement"`
	BenchmarkID         interface{}   `json:"Benchmark ID"`
	DatasetName         string        `json:"Dataset Name"`
	Reviews             []interface{} `json:"Reviews"`
	ViewLabel           string        `json:"View Label"`
	HasOpenIssues       int           `json:"Has Open Issues"`
	Skipped             bool          `json:"Skipped"`
	DataRowWorkflowInfo struct {
		TaskName        string `json:"taskName"`
		WorkflowHistory []struct {
			ActorID          string    `json:"actorId"`
			Action           string    `json:"action"`
			CreatedAt        time.Time `json:"createdAt"`
			PreviousTaskID   string    `json:"previousTaskId,omitempty"`
			PreviousTaskName string    `json:"previousTaskName,omitempty"`
			NextTaskID       string    `json:"nextTaskId,omitempty"`
			NextTaskName     string    `json:"nextTaskName,omitempty"`
		} `json:"Workflow History"`
	} `json:"DataRow Workflow Info"`
}
