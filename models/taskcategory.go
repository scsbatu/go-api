package models

import (
	"fmt"
	"github.com/google/uuid"
)

type JobCategory struct {
	ID          *[]byte `gorm:"Column:id" sql:"type:binary(16);not null"`
	CategoryKey *int    `gorm:"Column:category_key" sql:"type:tinyint(4);not null"`
	Name        *string `gorm:"Column:name" sql:"type:varchar(255);default:null"`
}

// NewJobCategory inits JobCategory struct
func NewJobCategory() *JobCategory {
	return &JobCategory{}
}

//TableName - returns name of the table
//Implement mysql.GenericTable interface
func (*JobCategory) TableName() string {
	return "job_category"
}

func CreateJobCategory(name, parentID *string, categoryKey *int) (j *JobCategory, err error) {
	id, _ := uuid.New().MarshalBinary()
	var p *[]byte
	if parentID != nil {
		*p = []byte(*parentID)
	}
	j = &JobCategory{
		ID:          &id,
		Name:        name,
		CategoryKey: categoryKey,
	}
	if insert := db.Create(j); insert.Error != nil {
		return nil, insert.Error
	}
	return
}

func GetJobCategoryByID(id string) (*JobCategory, error) {
	uuID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("unable to convert the string to UUID")
	}
	idInBinary, err := uuID.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("unable to convert the UUID")
	}
	var j JobCategory
	j.ID = &idInBinary
	if find := db.First(&j); find.Error != nil {
		return nil, find.Error
	}
	return &j, nil
}
