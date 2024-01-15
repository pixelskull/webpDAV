package webdav

import (
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func CreateHandler(path string) (handler *webdav.Handler) {
	handler = &webdav.Handler{
		FileSystem: webdav.Dir(path),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Printf("WEBDAV [%s]: %s, ERROR: %s \n", r.Method, r.URL, err)
			}
			//    else {
			// 	log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
			// }
		},
	}
	return
}
