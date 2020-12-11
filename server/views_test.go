package server

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mikitachab/score-board/db"
	"github.com/mikitachab/score-board/mocks"
)

type TestCtx struct {
	mockCtrl     *gomock.Controller
	mockDB       *mocks.MockRepositoryInterface
	mockTemplate *mocks.MockTemplateInterface
	handlerCtx   *HandlerCtx
}

func makeTestCtx(t *testing.T) *TestCtx {
	mockCtrl := gomock.NewController(t)
	mockDB := mocks.NewMockRepositoryInterface(mockCtrl)
	mockTemplate := mocks.NewMockTemplateInterface(mockCtrl)
	handlerCtx := HandlerCtx{DB: mockDB, Template: mockTemplate}
	return &TestCtx{mockCtrl, mockDB, mockTemplate, &handlerCtx}
}

func assertStatusOk(rr *httptest.ResponseRecorder, t *testing.T) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v expected %v",
			status, http.StatusOK)
	}
}

func TestHandlePlayersList(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handlePlayersList(testCtx.handlerCtx)

	players := []db.Player{{Name: "test"}}
	testCtx.mockDB.EXPECT().GetAllPlayers().Return(players)
	testCtx.mockTemplate.EXPECT().Render(rr, players)

	handler.ServeHTTP(rr, req)

	assertStatusOk(rr, t)
}

func TestHandleAddPlayerGet(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	req, err := http.NewRequest("GET", "/players/add", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handleAddPlayer(testCtx.handlerCtx)

	testCtx.mockTemplate.EXPECT().Render(rr, nil)

	handler.ServeHTTP(rr, req)

	assertStatusOk(rr, t)
}

func createPostRequest(method, URL string, data map[string]string) (*http.Request, error) {
	formData := url.Values{}
	for k, v := range data {
		formData.Set(k, v)
	}
	dataEcoded := bytes.NewBufferString(formData.Encode())
	req, err := http.NewRequest(method, URL, dataEcoded)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	return req, nil
}

func TestHandleAddPlayerPostOk(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	playerName := "test"
	data := map[string]string{"playerName": playerName}
	req, err := createPostRequest("POST", "/players/add", data)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handleAddPlayer(testCtx.handlerCtx)

	testCtx.mockDB.EXPECT().CreatePlayer(playerName)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("expected redirect but got %v", status)
	}
}

func TestHandleAddPlayerPostNotOk(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	playerName := "test"
	data := map[string]string{"playerName": playerName}
	req, err := createPostRequest("POST", "/players/add", data)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handleAddPlayer(testCtx.handlerCtx)

	testCtx.mockDB.EXPECT().CreatePlayer(playerName).Return(nil, errors.New("exists"))
	errorMsg := fmt.Sprintf("Player with name %s already exist", playerName)
	testCtx.mockTemplate.EXPECT().Render(rr, AddPlayerFormCtx{true, []string{errorMsg}})

	handler.ServeHTTP(rr, req)

	assertStatusOk(rr, t)
}

func TestHandleAddScoreSelectPlayersGet(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	req, err := http.NewRequest("GET", "/score/add/select", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := handleAddScoreSelectPlayers(testCtx.handlerCtx)

	players := []db.Player{{Name: "test"}}
	testCtx.mockDB.EXPECT().GetAllPlayers().Return(players)
	testCtx.mockTemplate.EXPECT().Render(rr, players)

	handler.ServeHTTP(rr, req)
	assertStatusOk(rr, t)
}
