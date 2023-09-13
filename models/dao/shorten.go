package dao

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Shorten struct {
	ID                    int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Shortcode             string     `json:"shortcode" gorm:"type:varchar(50);unique;notNull"`
	URL                   string     `json:"url" gorm:"type:varchar(255);notNull"`
	RedirectCount         int64      `json:"redirectCount" gorm:"type:int8;default:0"`
	LastSeenDate          *time.Time `json:"-" gorm:"default:null"`
	LastSeenDateFormatted string     `json:"lastSeenDate" gorm:"-"`
	StartDate             *time.Time `json:"startDate" gorm:"default:current_timestamp"`
	StartDateFormatted    string     `json:"startDateString"  gorm:"-"`
}

func (Shorten) TableName() string {
	return "shorten"
}

func MigrateShorten(db *gorm.DB) {
	if !db.Migrator().HasTable(&Shorten{}) {
		log.Println("migrate table shorten")
		db.AutoMigrate(&Shorten{})
	}
}
