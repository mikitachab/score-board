package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/mikitachab/score-board/db"
	"github.com/mikitachab/score-board/mocks"
	"github.com/mikitachab/score-board/templateloader"
)

type TemplateLoaderSpy struct {
	tl           *templateloader.TemplateLoader
	RenderCalled map[string]bool
}

func (tls *TemplateLoaderSpy) getRenderSpy(rf templateloader.RTF, name string) templateloader.RTF {
	return func(wr io.Writer, data interface{}) error {
		tls.RenderCalled[name] = true
		return rf(wr, data)
	}
}

func (tls *TemplateLoaderSpy) GetRenderTemplateFunc(templateName string) (templateloader.RTF, error) {
	renderFunc, err := tls.tl.GetRenderTemplateFunc(templateName)
	tls.RenderCalled[templateName] = false
	return tls.getRenderSpy(renderFunc, templateName), err
}

func (tls *TemplateLoaderSpy) assertRenderCalledWith(templateName string, t *testing.T) {
	if !tls.RenderCalled[templateName] {
		t.Errorf("render template %s not called", templateName)
	}
}

func TestMain(m *testing.M) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir("..") // templates loads from project root
	os.Exit(m.Run())
	os.Chdir(cwd)
}

type TestCtx struct {
	server    *Server
	mockCtrl  *gomock.Controller
	loaderSpy *TemplateLoaderSpy
	mockDB    *mocks.MockRepositoryInterface
}

func makeTestCtx(t *testing.T) *TestCtx {
	mockCtrl := gomock.NewController(t)
	mockDB := mocks.NewMockRepositoryInterface(mockCtrl)
	loaderSpy := TemplateLoaderSpy{
		templateloader.NewTemplateLoader(),
		make(map[string]bool),
	}
	server := makeServer(&loaderSpy, mockDB)
	return &TestCtx{server, mockCtrl, &loaderSpy, mockDB}

}

func assertStatusOk(rr *httptest.ResponseRecorder, t *testing.T) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleIndex(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	testCtx.mockDB.EXPECT().GetAllPlays()

	handler := testCtx.server.handleIndex()
	handler.ServeHTTP(rr, req)

	assertStatusOk(rr, t)
	testCtx.loaderSpy.assertRenderCalledWith("index.html", t)
}

func TestHandlePlayersList(t *testing.T) {
	testCtx := makeTestCtx(t)
	defer testCtx.mockCtrl.Finish()

	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	players := []db.Player{{Name: "test"}}
	testCtx.mockDB.EXPECT().GetAllPlayers().Return(players)

	handler := testCtx.server.handlePlayersList()
	handler.ServeHTTP(rr, req)

	assertStatusOk(rr, t)
	testCtx.loaderSpy.assertRenderCalledWith("players_list.html", t)
}
