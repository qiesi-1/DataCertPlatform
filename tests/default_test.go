package test

import (
	_ "data/routers"
	"github.com/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())

	convey.Convey("Subject: Test Station Endpoint\n", t, func() {
		convey.Convey("Status Code Should Be 200", func() {
			convey.So(w.Code, ShouldEqual, 200)
		})
		convey.Convey("The Result Should Not Be Empty", func() {
			convey.So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func ShouldBeGreaterThan(actual interface{}, expected ...interface{}) string {

}

func ShouldEqual(actual interface{}, expected ...interface{}) string {

}
