package main

import (
	"embed"
	"html/template"
	"net/http"
	"path"

	"github.com/yejin-lim/eks-workshop-sample-api-service-go"
)

var (
	//go:embed web/template/*.html
	templateFS embed.FS

	//go:embed web
	//staticFS embed.FS
)

func main() {
	r := gin.Default()

	LoadHTMLFromEmbedFS(r, templateFS, "template/*")

	r.GET("/", index)
	r.GET("template/*filepath", func(c *gin.Context) {
		c.FileFromFS(path.Join("/", c.Request.URL.Path), http.FS(templateFS))
	})
	r.Run()
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoadHTMLFromEmbedFS(engine *gin.Engine, em embed.FS, pattern string) {
	templ := template.Must(template.ParseFS(em, pattern))

	engine.SetHTMLTemplate(templ)
}
