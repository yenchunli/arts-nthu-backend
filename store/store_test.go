package store

type MockDB struct{}

func NewMockStore() Store {
	return &MockDB{}
}

func (db *MockDB) ListExhibitions(arg ListExhibitionsParams) ([]Exhibition, error) {
	var exhibitions []Exhibition
	var err error
	return exhibitions, err
}

func (db *MockDB) GetExhibition(id int) (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) CreateExhibition(arg CreateExhibitionParams) (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) EditExhibitions(arg EditExhibitionParams) (Exhibition, error) {
	var exhibition Exhibition
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

func (db *MockDB) CreateUser(arg CreateUserParams) (User, error) {
	var user User
	var err error
	return user, err
}

func (db *MockDB) GetUser(username string) (User, error) {
	var user User
	var err error
	return user, err
}

func (db *MockDB) ListNews(arg ListNewsParams) ([]News, error) {
	var news []News
	var err error
	return news, err
}
func (db *MockDB) GetNews(id int) (News, error) {
	var news News
	var err error
	return news, err
}

func (db *MockDB) CreateNews(arg CreateNewsParams) (News, error) {
	var news News
	var err error
	return news, err
}

func (db *MockDB) EditNews(arg EditNewsParams) (News, error) {
	var news News
	var err error
	return news, err
}

func (db *MockDB) DeleteNews(id int) error {
	var err error
	return err
}
