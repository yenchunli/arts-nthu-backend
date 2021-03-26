package store

type Store interface {
	ListExhibitions(arg ListExhibitionsParams) ([]Exhibition, error)
	GetExhibition(id int8) (Exhibition, error)
	CreateExhibition(arg CreateExhibitionParams) (Exhibition, error)
	EditExhibitions() (Exhibition, error)
	DeleteExhibition() error
}
