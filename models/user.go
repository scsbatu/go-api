package models

import (
	"github.com/google/uuid"
	"github.com/scsbatu/go-api/core/helpers"
)

type User struct {
	ID            *[]byte `gorm:"Column:id;type:uuid" sql:"type:binary(16);not null"`
	FirstName     *string `gorm:"Column:first_name" sql:"type:varchar(255);not null"`
	LastName      *string `gorm:"Column:last_name" sql:"type:varchar(255);default:null"`
	DocumentNotes *string `gorm:"Column:document_notes" sql:"type:varchar(255);default:null"`
}

// NewUser inits User struct
func NewUser() *User {
	return &User{}
}

//TableName - returns name of the table
//Implement mysql.GenericTable interface
func (*User) TableName() string {
	return "user"
}

func CreateUser(
	firstName, lastName, docNotes *string,
) (
	u *User,
	err error,
) {
	id, _ := uuid.New().MarshalBinary()
	u = &User{
		ID:            &id,
		FirstName:     firstName,
		LastName:      lastName,
		DocumentNotes: docNotes,
	}
	if insert := db.Create(u); insert.Error != nil {
		return nil, insert.Error
	}
	return
}

func GetUserByID(id string) (*User, error) {
	idInBinary, err := helpers.StringToUUIDByte(id)
	if err != nil {
		return nil, err
	}
	var u User
	u.ID = &idInBinary
	if find := db.First(&u); find.Error != nil {
		return nil, find.Error
	}
	return &u, nil
}

func UpdateUser(u *User) (err error) {
	if update := db.Save(u); update.Error != nil {
		return update.Error
	}
	return nil
}
