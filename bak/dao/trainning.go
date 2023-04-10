package dao

import "time"

type StorageService struct {
	ID       string `xorm:"varchar(64) pk notnull unique 'id'"` // id
	Account  string `xorm:"varchar(128) notnull 'account'"`     // 账号
	Password string `xorm:"varchar(256) notnull 'password'"`    // 密码
	Token    string `xorm:"varchar(256) notnull 'token'"`       // token
	Address  string `xorm:"varchar(256) notnull 'address'"`     // 地址
	Bucket   string `xorm:"varchar(128) notnull 'bucket'"`      // bucket 名称
	Type     int    `xorm:"default 0 notnull 'type'"`           // 类型：0-未知，1-阿里云，2-腾讯云，3-华为云 4- CODING generic
}

func (s *StorageService) TableName() string {
	return "storage_service"
}

type TrainingRecord struct {
	Id        string    `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	Content   string    `json:"content"`
	DataSetId int       `xorm:"index" json:"data_set_id"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

func (t *TrainingRecord) TableName() string {
	return "training_records"
}

type DataSet struct {
	Id        string    `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
}

func (d *DataSet) TableName() string {
	return "datasets"
}

type FileAccess struct {
	Id         string `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	Type       int    `xorm:"default 0 notnull 'type'"` // 1 Local
	AccessInfo string `xorm:"varchar(256) notnull 'access_info'"`
}

func (f *FileAccess) TableName() string {
	return "file_access"
}

type DataFile struct {
	Id         string    `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	Name       string    `xorm:"varchar(64) notnull unique 'name'" json:"name"`
	Type       int       `xorm:"default 0 notnull 'type'"` // 1 文本  2 图片  3 音频
	DataSetId  string    `xorm:"index" json:"data_set_id"`
	AccessInfo string    `xorm:"varchar(256) notnull 'access_info'"`
	CreatedAt  time.Time `xorm:"created" json:"created_at"`
	UpdatedAt  time.Time `xorm:"updated" json:"updated_at"`
}

func (d *DataFile) TableName() string {
	return "data_files"
}

type DataSetInfo struct {
	Id        string `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	DataSetId int    `xorm:"index" json:"data_set_id"`
	Type      string `json:"type"`
	Content   string `json:"content"`
}

func (d *DataSetInfo) TableName() string {
	return "dataset_infos"
}

type Annotation struct {
	Id      string `xorm:"varchar(64) pk notnull unique 'id'" json:"id"`
	Content string `xorm:"varchar(256)" json:"content"`
	Source  string `xorm:"varchar(512)" json:"source"`

	DataSetId        string `xorm:"index" json:"data_set_id"`
	DataFileId       string `xorm:"index" json:"data_file_id"`
	TrainingRecordId int    `xorm:"index" json:"training_record_id"`
}

func (a *Annotation) TableName() string {
	return "annotations"
}
