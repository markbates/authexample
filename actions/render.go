package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
	"github.com/gobuffalo/velvet"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "application.html",
		TemplateEngine: velvet.BuffaloRenderer,
		FileResolverFunc: func() resolvers.FileResolver {
			return &resolvers.RiceBox{
				Box: rice.MustFindBox("../templates"),
			}
		},
		Helpers: map[string]interface{}{},
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../public/assets")
	return box.HTTPBox()
}
