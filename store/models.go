package store

type Exhibition struct {
	ID             int     `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	Type           string    `json:"type"`
	Cover          string    `json:"cover"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      string	 `json:"performer"`
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
	Type           string    `json:"type"`
	Cover          string    `json:"cover"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      string	 `json:"performer"`
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
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Type   string `json:"type"`
}

type EditExhibitionParams struct {
	ID             int     	 `json:"id"`
	Title          string    `json:"title"`
	TitleEn        string    `json:"title_en"`
	Subtitle       string    `json:"subtitle"`
	SubtitleEn     string    `json:"subtitle_en"`
	Type           string    `json:"type"`
	Cover          string    `json:"cover"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	Draft          bool      `json:"draft"`
	Host           string    `json:"host"`
	HostEn         string    `json:"host_en"`
	Performer      string	 `json:"performer"`
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

// User Model
type User struct {
	Username         string `json:"username"`
	HashedPassword   string `json:"hashed_password"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	PasswordChangeAt int64  `json:"password_change_at"`
	CreateAt         int64  `json:"create_at"`
}

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

// News Model
type News struct {
	ID        int  `json:"id"`
	Username  string `json:"username"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	TitleEn   string `json:"title_en"`
	StartDate string `json:"start_date"`
	Type 	  string `json:"type"`
	Draft     bool `json:"draft"`
	Content   string `json:"content"`
	ContentEn string `json:"content_en"`
	CreateAt  int64  `json:"create_at"`
	UpdateAt  int64  `json:"update_at"`
}

type ListNewsParams struct {
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	Type   string `json:"type"`
}

type CreateNewsParams struct {
	Username  string `json:"username"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	TitleEn   string `json:"title_en"`
	StartDate string `json:"start_date"`
	Type 	  string `json:"type"`
	Draft     bool `json:"draft"`
	Content   string `json:"content"`
	ContentEn string `json:"content_en"`
}

type EditNewsParams struct {
	ID        int  `json:"id"`
	Username  string `json:"username"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	TitleEn   string `json:"title_en"`
	StartDate string `json:"start_date"`
	Type 	  string `json:"type"`
	Draft     bool 	 `json:"draft"`
	Content   string `json:"content"`
	ContentEn string `json:"content_en"`
}
