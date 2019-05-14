package key

import "time"

type Key struct {
	ID        string `gorm:"type:serial;"`
	CreatedAt time.Time
	Name      string `gorm:"type:varchar(100);unique_index"`
	Public    string `gorm:"type:text"`
	Private   string `gorm:"type:text"`
}
