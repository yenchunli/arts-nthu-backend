package main

type MockDB struct{}

func NewMockStore() Store {
	return &MockDB{}
}

func (db *MockDB) ListExhibitions() ([]Exhibition, error) {
	var exhibitions []Exhibition
	var err error
	return exhibitions, err
}

func (db *MockDB) GetExhibition(id int8) (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) CreateExhibition() (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) EditExhibitions() (Exhibition, error) {
	var exhibition Exhibition
	var err error
	return exhibition, err
}

func (db *MockDB) DeleteExhibition() error {
	var err error
	return err
}
