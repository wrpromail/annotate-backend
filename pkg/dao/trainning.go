package dao

import "time"

type TrainingRecord struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Content   string    `json:"content"`
	DataSetId int       `xorm:"index" json:"data_set_id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

func (t *TrainingRecord) TableName() string {
	return "training_records"
}

type DataSet struct {
	Id        int       `xorm:"pk autoincr" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

func (d *DataSet) TableName() string {
	return "datasets"
}

type DataSetInfo struct {
	Id        int    `xorm:"pk autoincr" json:"id"`
	DataSetId int    `xorm:"index" json:"data_set_id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
}

func (d *DataSetInfo) TableName() string {
	return "dataset_infos"
}

type Annotation struct {
	Id               int    `xorm:"pk autoincr" json:"id"`
	Content          string `json:"content"`
	DataSetId        int    `xorm:"index" json:"data_set_id"`
	TrainingRecordId int    `xorm:"index" json:"training_record_id"`
}

func (a *Annotation) TableName() string {
	return "annotations"
}
