package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.POST("/", normalPaste)
	router.GET("/:pbid", contentByPbid)
	router.GET("/:pbid/:lan", contentByPbidHighLight)

    // NOTE: https
	// go func() {
	//     router0 := httprouter.New()
	//     router0.POST("/", normalPaste)
	//     myLog.WithFields(
	//     logrus.Fields{
	//         "method": "main.go: main",
	//     },
	// ).Fatal(http.ListenAndServeTLS(":443", "aimisaka.site_bundle.crt", "aimisaka.site.key", router0))
	// }()

	myLog.WithFields(
		logrus.Fields{
			"method": "main.go: main",
		},
	).Fatal(http.ListenAndServe(":80", router))
}
