package server

import (
	"net/http"
	"net/http/httptest"
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
