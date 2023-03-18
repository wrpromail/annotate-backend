package reader

import (
	"bufio"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"mymlops/annotate-helper/pkg/dao"
	"os"
	"strings"
	"testing"
)

func TestInitDB(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	err = engine.Sync2(new(dao.Annotation), new(dao.DataSet), new(dao.DataSetInfo), new(dao.TrainingRecord))
	if err != nil {
		t.Error(err)
	}
}

func TestGetRecord(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	records := make([]*dao.TrainingRecord, 0)
	err = engine.Find(&records)
	if err != nil {
		t.Error(err)
		return
	}
	for _, record := range records {
		t.Log(*record)
	}
}

func TestReadIntentTrainingData(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer engine.Close()

	file, err := os.Open("coding-git-repository-intent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	ds := &dao.DataSet{
		Name: "coding-git-repository-intent-1",
	}
	datasetId, err := engine.Insert(ds)
	if err != nil {
		t.Error(err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && !strings.HasPrefix(line, "//") {
			// 处理非空行且不以 "//" 开头的情况
			// 这里可以根据业务需求进行相应处理
			t.Log(line)
			record := &dao.TrainingRecord{
				Content:   line,
				DataSetId: int(datasetId),
			}
			_, err := engine.Insert(record)
			if err != nil {
				t.Log("inert failed " + err.Error())
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
