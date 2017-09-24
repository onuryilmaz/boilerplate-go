package server

import (
	"fmt"
	"testing"

	"net/http"
	"time"

	"github.com/onuryilmaz/boilerplate-go/pkg/commons"
	"github.com/phayes/freeport"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRESTServer(t *testing.T) {

	options := commons.Options{}
	options.ServerPort = fmt.Sprintf("%v", freeport.GetPort())
	RESTServer := NewREST(options)
	RESTServer.Start()

	Convey("Start and check RESTServer", t, func() {
		So(RESTServer, ShouldNotBeNil)

		// Wait for server is up
		time.Sleep(time.Second)

		Convey("Ping server", func() {

			response, err := http.Get("http://localhost:" + options.ServerPort + "/ping")
			So(err, ShouldBeNil)
			So(response.StatusCode, ShouldEqual, 200)

		})
	})

}
