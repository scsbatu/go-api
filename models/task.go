package models

import (
	"github.com/google/uuid"
	"github.com/scsbatu/go-api/core/helpers"
	"time"
)

type Task struct {
	ID               *[]byte    `gorm:"Column:id" sql:"type:binary(16);not null"`
	Title            *string    `gorm:"Column:title" sql:"type:varchar(255);not null"`
	TaskKey          *string    `gorm:"Column:task_key" sql:"type:varchar(255);default:null"`
	Details          *string    `gorm:"Column:details" sql:"type:varchar(255);default:null"`
	ExpectedDateTime *time.Time `gorm:"Column:expected_date_time"sql:"type:timestamp;default:current_timestamp"`
	Status           *int       `gorm:"Column:status" sql:"type:tinyint(4);not null"`
	CreatorID        *[]byte    `gorm:"Column:creator_id" sql:"type:binary(16);not null"`
	CreatedDate      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:current_timestamp"`
	UpdatedDate      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:current_timestamp"`
	StartDateTime    *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:0000-00-00 00:00:00"`
	EndDateTime      *time.Time `gorm:"Column:created_date" sql:"type:timestamp;default:0000-00-00 00:00:00"`
	CategoryID       *[]byte    `gorm:"Column:category_id" sql:"type:binary(16);default:null"`
}

// NewJob inits Job struct
func NewJob() *Task {
	return &Task{}
}

//TableName - returns name of the table
//Implement mysql.GenericTable interface
func (*Task) TableName() string {
	return "task"
}

func CreateTask(
	title, taskKey, details, creatorID, categoryID *string,

) (j *Task, err error) {
	id, _ := uuid.New().MarshalBinary()
	t := time.Now().Local()
	var cID, catID []byte
	if creatorID != nil {
		cID, err = helpers.StringToUUIDByte(*creatorID)
		if err != nil {
			return nil, err
		}
	}

	j = &Task{
		ID:               &id,
		Title:            title,
		TaskKey:          taskKey,
		Details:          details,
		CreatorID:        &cID,
		CategoryID:       &catID,
		ExpectedDateTime: &t,
		CreatedDate:      &t,
		UpdatedDate:      &t,
		StartDateTime:    &t,
	}
	if insert := db.Create(j); insert.Error != nil {
		return nil, insert.Error
	}
	return
}

func GetTaskByID(id string) (*Task, error) {
	idInBinary, err := helpers.StringToUUIDByte(id)
	if err != nil {
		return nil, err
	}
	var j Task
	j.ID = &idInBinary
	if find := db.First(&j); find.Error != nil {
		return nil, find.Error
	}
	return &j, nil
}

func UpdateTask(j *Task) (err error) {
	if update := db.Save(j); update.Error != nil {
		return update.Error
	}
	return nil
}
