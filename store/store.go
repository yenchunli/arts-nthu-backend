package store

type Store interface {
	ListExhibitions(arg ListExhibitionsParams) ([]Exhibition, error)
	GetExhibition(id int) (Exhibition, error)
	CreateExhibition(arg CreateExhibitionParams) (Exhibition, error)
	EditExhibition(arg EditExhibitionParams) (Exhibition, error)
	DeleteExhibition(id int) error
	GetExhibitionsMaxSize() (int, error)

	CreateUser(arg CreateUserParams) (User, error)
	GetUser(username string) (User, error)

	ListNews(arg ListNewsParams) ([]News, error)
	GetNews(id int) (News, error)
	CreateNews(arg CreateNewsParams) (News, error)
	EditNews(arg EditNewsParams) (News, error)
	DeleteNews(id int) error
}
