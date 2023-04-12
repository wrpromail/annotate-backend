package labelbox

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func mustGetLabelBoxExportEntity(filename string) (result []EntityRecord) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Error(err)
		return
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Error(err)
	}
	return
}

func writeTextAndClassificationToTsv(source []EntityRecord, target string) error {
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, data := range source {
		cf := data.Label.Classifications
		if len(cf) > 0 {
			line := fmt.Sprintf("%s\t%s\n", data.LabeledData, cf[0].Answer.Title)
			_, e := writer.WriteString(line)
			if e != nil {
				log.Warn(e)
			}
		} else {
			continue
		}
	}
	return writer.Flush()
}
