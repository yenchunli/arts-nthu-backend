package server

import (
	"log"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	db "github.com/yenchunli/arts-nthu-backend/internal/postgres"
	"github.com/yenchunli/arts-nthu-backend/util"
)

func TestShouldListExhibitions(t *testing.T) {

	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Fail to create mockDB")
	}

	store := db.NewDB(fakeDB)

	config := util.LoadTestConfig()

	server, _ := NewServer(config, store)

	rows := sqlmock.NewRows([]string{"id", "title", "title_en", "subtitle", "subtitle_en", "type", "cover", "start_date", "end_date", "draft", "host", "host_en", "performer", "location", "location_en", "daily_start_time", "daily_end_time", "category", "description", "description_en", "content", "content_en", "create_at", "update_at"}).
		AddRow(1, "exhibition 1", "", "subtitle", "", "visual_art", "https://i.imgur.com/12345", "2021-01-01", "", "false", "NTHU ArtsCenter", "", "performer", "NTHU", "", "11:00", "", "", "description123", "", "content", "", "123456789", "123456789").
		AddRow(2, "exhibition 2", "", "subtitle", "", "visual_art", "https://i.imgur.com/12345", "2021-01-01", "", "false", "NTHU ArtsCenter", "", "performer", "NTHU", "", "11:00", "", "", "description123", "", "content", "", "123456789", "123456789")

	mock.ExpectQuery("^SELECT (.+) FROM exhibitions WHERE id=(.+)").WillReturnRows(rows)

	w := performRequest(server.router, http.MethodGet, "/api/v1/exhibitions/1")
	assert.Equal(t, http.StatusOK, w.Code)

}
