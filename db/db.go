package db

import (
	"database/sql"
	store "github.com/yenchunli/go-nthu-artscenter-server/store"
	_ "github.com/lib/pq"
	"time"
	"fmt"
)

type DB struct {
	conn *sql.DB
}

func NewDB(conn *sql.DB) store.Store {
	return &DB{conn: conn}
}

func (db *DB) ListExhibitions(arg store.ListExhibitionsParams) ([]store.Exhibition, error) {
	
	var command string
	var err error
	var rows *sql.Rows

	if arg.Type == "" {
		command = `
		SELECT * FROM exhibitions 
		ORDER by start_date
		LIMIT $1
		OFFSET $2
		`
		rows, err = db.conn.Query(command, arg.Limit, arg.Offset)
	} else {
		command = `
		SELECT * FROM exhibitions
		WHERE type=$1
		ORDER by start_date
		LIMIT $2
		OFFSET $3
		`
		rows, err = db.conn.Query(command, arg.Type, arg.Limit, arg.Offset)
	}
	

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []store.Exhibition{}
	
	for rows.Next() {
		var exhibition store.Exhibition
		if err := rows.Scan(
			&exhibition.ID,
			&exhibition.Title,
			&exhibition.TitleEn,
			&exhibition.Subtitle,
			&exhibition.SubtitleEn,
			&exhibition.Type,
			&exhibition.Cover,
			&exhibition.StartDate,
			&exhibition.EndDate,
			&exhibition.Draft,
			&exhibition.Host,
			&exhibition.HostEn,
			&exhibition.Performer,
			&exhibition.Location,
			&exhibition.LocationEn,
			&exhibition.DailyStartTime,
			&exhibition.DailyEndTime,
			&exhibition.Category,
			&exhibition.Description,
			&exhibition.DescriptionEn,
			&exhibition.Content,
			&exhibition.ContentEn,
			&exhibition.CreateAt,
			&exhibition.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, exhibition)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, err
}

func (db *DB) GetExhibition(id int32) (exhibition store.Exhibition, err error) {
	const command = `
	SELECT * FROM exhibitions WHERE id=$1`

	err = db.conn.QueryRow(command, id).Scan(
		&exhibition.ID,
		&exhibition.Title,
		&exhibition.TitleEn,
		&exhibition.Subtitle,
		&exhibition.SubtitleEn,
		&exhibition.Type,
		&exhibition.Cover,
		&exhibition.StartDate,
		&exhibition.EndDate,
		&exhibition.Draft,
		&exhibition.Host,
		&exhibition.HostEn,
		&exhibition.Performer,
		&exhibition.Location,
		&exhibition.LocationEn,
		&exhibition.DailyStartTime,
		&exhibition.DailyEndTime,
		&exhibition.Category,
		&exhibition.Description,
		&exhibition.DescriptionEn,
		&exhibition.Content,
		&exhibition.ContentEn,
		&exhibition.CreateAt,
		&exhibition.UpdateAt,
	)
	return exhibition, err
}

type CreateExhibitionParams struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	Type		   string    `json:"type"`
	Cover		   string    `json:"cover"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      store.Performer `json:"performer"`
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

func (db *DB) CreateExhibition(arg store.CreateExhibitionParams) (store.Exhibition, error) {

	const command = `
	INSERT INTO exhibitions (
		title,
		title_en,
		subtitle,
		subtitle_en,
		type,
		cover,
		start_date,
		end_date,
		draft,
		host,
		host_en,
		performer,
		location,
		location_en,
		daily_start_time,
		daily_end_time,
		category,
		description,
		description_en,
		content,
		content_en,
		create_at,
		update_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
	) RETURNING id,title,title_en,subtitle,
	subtitle_en,
	type,
	cover,
	start_date,
	end_date,
	draft,
	host,
	host_en,
	performer,
	location,
	location_en,
	daily_start_time,
	daily_end_time,
	category,
	description,
	description_en,
	content,
	content_en,
	create_at,
	update_at
	`
	currentTime := time.Now().Unix()
	row := db.conn.QueryRow(command,
		arg.Title,
		arg.TitleEn,
		arg.Subtitle,
		arg.SubtitleEn,
		arg.Type,
		arg.Cover,
		arg.StartDate,
		arg.EndDate,
		arg.Draft,
		arg.Host,
		arg.HostEn,
		arg.Performer,
		arg.Location,
		arg.LocationEn,
		arg.DailyStartTime,
		arg.DailyEndTime,
		arg.Category,
		arg.Description,
		arg.DescriptionEn,
		arg.Content,
		arg.ContentEn,
		currentTime,
		currentTime,
	)
	fmt.Println(arg.Draft)
	var e store.Exhibition

	err := row.Scan(
		&e.ID,
		&e.Title,
		&e.TitleEn,
		&e.Subtitle,
		&e.SubtitleEn,
		&e.Type,
		&e.Cover,
		&e.StartDate,
		&e.EndDate,
		&e.Draft,
		&e.Host,
		&e.HostEn,
		&e.Performer,
		&e.Location,
		&e.LocationEn,
		&e.DailyStartTime,
		&e.DailyEndTime,
		&e.Category,
		&e.Description,
		&e.DescriptionEn,
		&e.Content,
		&e.ContentEn,
		&e.CreateAt,
		&e.UpdateAt,
	)
	return e, err
}

func (db *DB) EditExhibitions(arg store.EditExhibitionParams) (store.Exhibition, error) {
	
	const command = `
	UPDATE exhibitions
	SET title=$2,
	title_en=$3,
	subtitle=$4,
	subtitle_en=$5,
	type=$6
	cover=$7,
	start_date=$8,
	end_date=$9,
	draft=$10,
	host=$11,
	host_en=$12,
	performer=$13,
	location=$14,
	location_en=$15,
	daily_start_time=$16,
	daily_end_time=$17,
	category=$18,
	description=$19,
	description_en=$20,
	content=$21,
	content_en=$22,
	update_at=$23
	WHERE id = $1
	RETURNING id,title,title_en,subtitle,subtitle_en,type,cover,start_date,end_date,draft,host,host_en,performer,location,location_en,daily_start_time,daily_end_time,category,description,description_en,content,content_en,create_at,update_at
	`
	
	currentTime := time.Now().Unix()
	row := db.conn.QueryRow(command,
		arg.ID,
		arg.Title,
		arg.TitleEn,
		arg.Subtitle,
		arg.SubtitleEn,
		arg.Type,
		arg.Cover,
		arg.StartDate,
		arg.EndDate,
		arg.Draft,
		arg.Host,
		arg.HostEn,
		arg.Performer,
		arg.Location,
		arg.LocationEn,
		arg.DailyStartTime,
		arg.DailyEndTime,
		arg.Category,
		arg.Description,
		arg.DescriptionEn,
		arg.Content,
		arg.ContentEn,
		currentTime,
	)
	
	var e store.Exhibition

	err := row.Scan(
		&e.ID,
		&e.Title,
		&e.TitleEn,
		&e.Subtitle,
		&e.SubtitleEn,
		&e.Type,
		&e.Cover,
		&e.StartDate,
		&e.EndDate,
		&e.Draft,
		&e.Host,
		&e.HostEn,
		&e.Performer,
		&e.Location,
		&e.LocationEn,
		&e.DailyStartTime,
		&e.DailyEndTime,
		&e.Category,
		&e.Description,
		&e.DescriptionEn,
		&e.Content,
		&e.ContentEn,
		&e.CreateAt,
		&e.UpdateAt,
	)
	return e, err
}

func (db *DB) DeleteExhibition(id int32) error {

	const command = `
	DELETE FROM exhibitions
	WHERE id = $1
	`

	_, err := db.conn.Exec(command, id)
	
	return err
}

func (db *DB) CreateUser(arg store.CreateUserParams) (store.User, error) {
	const command = `
	INSERT INTO users (
		username,
		hashed_password,
		full_name,
		email,
		password_change_at,
		create_at
	  ) VALUES (
		$1, $2, $3, $4, $5, $6
	  ) RETURNING username, hashed_password, full_name, email, password_change_at, create_at
	`
	currentTime := time.Now().Unix()
	row := db.conn.QueryRow(command,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		currentTime,
		currentTime,
	)
	var user store.User
	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangeAt,
		&user.CreateAt,
	)
	return user, err
}

func (db *DB) GetUser(username string) (store.User, error) {
	const command = `
	SELECT username, hashed_password, full_name, email, password_change_at, create_at FROM users
	WHERE username = $1 LIMIT 1
	`
	row := db.conn.QueryRow(command, username)
	var user store.User
	err := row.Scan(
		&user.Username,
		&user.HashedPassword,
		&user.FullName,
		&user.Email,
		&user.PasswordChangeAt,
		&user.CreateAt,
	)
	return user, err
}