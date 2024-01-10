package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	Name      string     `json:"name"`
	Password  string     `json:"-"`
	Status    bool       `json:"status" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}
