package store

type Store interface {
	ListExhibitions(arg ListExhibitionsParams) ([]Exhibition, error)
	GetExhibition(id int32) (Exhibition, error)
	CreateExhibition(arg CreateExhibitionParams) (Exhibition, error)
	EditExhibitions(arg EditExhibitionParams) (Exhibition, error)
	DeleteExhibition(id int32) error

	CreateUser(arg CreateUserParams) (User, error)
	GetUser(username string) (User, error)

	ListNews(arg ListNewsParams) ([]News, error)
	GetNews(id int32) (News, error)
	CreateNews(arg CreateNewsParams) (News, error)
	EditNews(arg EditNewsParams) (News, error)
	DeleteNews(id int32) error
}
