package store

import (
	"encoding/json"
	"errors"
	"database/sql/driver"
)
type Performer struct {
	Persons []PerformerPerson `json:"persons"`
}

type PerformerPerson struct {
	Name     string `json:name`
	NameEn   string `json:name_en`
	Title    string `json:title`
	Title_en string `json:title_en`
}

func (p Performer) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Performer) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &p)
}

type Exhibition struct {
	ID             int32     `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      Performer `json:"performer"`
	Location       string    `json:"location"`
	LocationEn     string    `json:"location_en"`
	DailyStartTime string    `json:"daily_start_time"`
	DailyEndTime   string    `json:"daily_end_time"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	DescriptionEn  string    `json:"description_en"`
	Content        string    `json:"content"`
	ContentEn      string    `json:"content_en"`
	CreateAt       int64     `json:"create_at"`
	UpdateAt       int64     `json:"update_at"`
}

type CreateExhibitionParams struct {
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      Performer `json:"performer"`
	Location       string    `json:"location"`
	LocationEn     string    `json:"location_en"`
	DailyStartTime string    `json:"daily_start_time"`
	DailyEndTime   string    `json:"daily_end_time"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	DescriptionEn  string    `json:"description_en"`
	Content        string    `json:"content"`
	ContentEn      string    `json:"content_en"`
}

type ListExhibitionsParams struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type EditExhibitionParams struct {
	ID			   int32	 `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      Performer `json:"performer"`
	Location       string    `json:"location"`
	LocationEn     string    `json:"location_en"`
	DailyStartTime string    `json:"daily_start_time"`
	DailyEndTime   string    `json:"daily_end_time"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	DescriptionEn  string    `json:"description_en"`
	Content        string    `json:"content"`
	ContentEn      string    `json:"content_en"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt int64 	`json:"password_changed_at"`
	CreatedAt         int64 	`json:"created_at"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}
