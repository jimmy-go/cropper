package prev

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jimmy-go/cropper"
)

// Handler endpoint.
func Handler(c *cropper.Cropper) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		v := r.URL.Query()
		width, _ := strconv.Atoi(v.Get("width"))
		height, _ := strconv.Atoi(v.Get("height"))
		uri := v.Get("url")

		// Validate exists.
		if !c.Exists(uri, width, height) {
			log.Printf("Screenshot : generating preview : url [%s]", uri)
			if err := c.Screenshot(ctx, uri, width, height); err != nil {
				log.Printf("Screenshot : err [%s]", err)
				http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
				return
			}
		}

		// Write preview.
		if err := c.Copy(w, uri, width, height); err != nil {
			log.Printf("Screenshot : copy : err [%s]", err)
			http.Error(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			return
		}
	})
}
