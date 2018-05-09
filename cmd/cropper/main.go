package main

import (
	"flag"
	"log"

	"github.com/jimmy-go/cropper"
	"github.com/jimmy-go/cropper/api/prev"
	"github.com/jimmy-go/srest"
)

var (
	bin = flag.String("exec", "", "Chrome executable path.")
	dir = flag.String("screenshot-dir", "", "Screenshots directory.")
	// Server related config.
	port = flag.Int("port", -1, "Listen port for server mode.")
	w    = flag.Int("width", 0, "Preview width in pixels.")
	h    = flag.Int("height", 0, "Preview height in pixels.")
	// CLI related config.
	site   = flag.String("url", "", "Page URL for preview.")
	target = flag.String("target", "", "Target path name.")
)

func main() {
	flag.Parse()
	log.Printf("bin: %s", *bin)
	log.Printf("temporal dir: %s", *dir)
	log.Printf("url: %s", *site)

	//	if err := cropper.Preview(*site, *w, *h, *target); err != nil {
	//		log.Fatal(err)
	//	}

	c, err := cropper.New(*bin, *dir)
	if err != nil {
		log.Fatal(err)
	}

	m := srest.New(nil)
	m.Get("/v1/preview", prev.Handler(c))
	log.Printf("listen port: %d", *port)
	<-m.Run(*port)
}
