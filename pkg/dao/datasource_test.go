package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestDBInitAndInsert(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./object_storage.db")
	if err != nil {
		t.Error(err)
		return
	}
	err = engine.Sync2(new(TrainingRecord), new(DataSet), new(DataSetInfo), new(Annotation))
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCreateDataSet(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./object_storage.db")
	if err != nil {
		t.Error(err)
		return
	}
	ds := &DataSet{
		Id:   uuid.New().String(),
		Name: "coding-chatops-intent-training",
	}
	t.Log(engine.Insert(ds))
}

func TestDataFile(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./object_storage.db")
	if err != nil {
		t.Error(err)
		return
	}
	//err = engine.Sync2(new(DataFile))
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//datafile := &DataFile{
	//	Id:        uuid.New().String(),
	//	Name:      "gitRepoCmd1.txt",
	//	DataSetId: "2fe976bb-a34e-4cf0-8d7f-9c0799718a80",
	//	Type:      1,
	//}
	//t.Error(engine.Insert(datafile))
	//t.Error(engine.Sync2(new(DataFile), new(FileAccess)))

	//fid := uuid.New().String()
	//fa := &FileAccess{
	//	Id:         fid,
	//	Type:       1,
	//	AccessInfo: "/Users/wangrui/go/src/github.com/wrpromail/annotate-helper/files/gitRepoCmd1.txt",
	//}
	//_, err = engine.Insert(fa)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	var rst = &DataFile{}
	_, err = engine.Get(rst)
	if err != nil {
		t.Error(err)
		return
	} else {
		rst.AccessInfo = "dad197b7-6948-42ab-b396-abeb1cda8d9a"
		t.Log(engine.Update(rst))
	}

}

func TestUpdateAnnotation(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite3", "./object_storage.db")
	if err != nil {
		t.Error(err)
		return
	}
	t.Error(engine.Sync2(new(Annotation)))
}
