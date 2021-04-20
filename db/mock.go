package db

/*
import (
	"database/sql"
	store "github.com/yenchunli/arts-nthu-backend/store"
	"github.com/DATA-DOG/go-sqlmock"
)

type MockDB struct{
	conn *sql.DB
	mock sqlmock.Sqlmock
}

func NewMockDB() (*MockDB, error) {
	db, mock, err := sqlmock.New()

	return &MockDB{
		conn: db,
		mock: mock,
	}, err
}

func (db *MockDB) ListExhibitions(arg store.ListExhibitionsParams) ([]store.Exhibition, error) {
	var exhibitions []store.Exhibition
	var err error

	return exhibitions, err
}

func (db *MockDB) GetExhibition(id int) (store.Exhibition, error) {
	var exhibition store.Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) CreateExhibition(arg store.CreateExhibitionParams) (store.Exhibition, error) {
	var exhibition store.Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) EditExhibition(arg store.EditExhibitionParams) (store.Exhibition, error) {
	var exhibition store.Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) DeleteExhibition(id int) error {
	var err error
	return err
}

func (db *MockDB) GetExhibitionsMaxSize() (int, error) {
	var err error
	return -1, err
}

func (db *MockDB) CreateUser(arg store.CreateUserParams) (store.User, error) {
	var user store.User
	var err error
	return user, err
}

func (db *MockDB) GetUser(username string) (store.User, error) {
	var user store.User
	var err error
	return user, err
}

func (db *MockDB) ListNews(arg store.ListNewsParams) ([]store.News, error) {
	var news []store.News
	var err error
	return news, err
}
func (db *MockDB) GetNews(id int) (store.News, error) {
	var news store.News
	var err error
	return news, err
}

func (db *MockDB) CreateNews(arg store.CreateNewsParams) (store.News, error) {
	var news store.News
	var err error
	return news, err
}

func (db *MockDB) EditNews(arg store.EditNewsParams) (store.News, error) {
	var news store.News
	var err error
	return news, err
}

func (db *MockDB) DeleteNews(id int) error {
	var err error
	return err
}
*/
