package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	echorenderer "github.com/bdreece/echo-renderer"
	"github.com/labstack/echo/v4"
)


func home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.gotmpl", echo.Map{
		"Name": "World",
	})
}

func main() {
	port := flag.Int("p", 3000, "port")
	flag.Parse()

	opts := echorenderer.Options{
		FS:      os.DirFS("templates"),
		Include: []string{"shared/*.gotmpl", "*.gotmpl"},
	}

	renderer, err := echorenderer.New(&opts)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Renderer = renderer
	e.GET("/", home)
	e.Start(fmt.Sprintf(":%d", *port))
}
