package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"regexp"
	"strings"

	"github.com/yejin-lim/eks-workshop-sample-api-service-go"
)

//go:embed web/templates
var templatesFS embed.FS

func main() {

	r := gin.Default()
	pattern := "web/templates/*html"
	LoadHTMLFromEmbedFS(r, templatesFS, pattern)

	r.GET("/", index)
	r.GET("/", login)

	r.Run(":8080")
}

func LoadHTMLFromEmbedFS(engine *gin.Engine, embedFS embed.FS, pattern string) {
	root := template.New("")
	tmpl := template.Must(root, LoadAndAddToRoot(engine.FuncMap, root, embedFS, pattern))
	engine.SetHTMLTemplate(tmpl)
}

// Method version
// func (engine *gin.Engine) LoadHTMLFromFS(embedFS embed.FS, pattern string) {
// 	root := template.New("")
// 	tmpl := template.Must(root, LoadAndAddToRoot(engine.FuncMap, root, embedFS, pattern))
// 	engine.SetHTMLTemplate(tmpl)
// }

func LoadAndAddToRoot(funcMap template.FuncMap, rootTemplate *template.Template, embedFS embed.FS, pattern string) error {
	pattern = strings.ReplaceAll(pattern, ".", "\\.")
	pattern = strings.ReplaceAll(pattern, "*", ".*")

	err := fs.WalkDir(embedFS, ".", func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if matched, _ := regexp.MatchString(pattern, path); !d.IsDir() && matched {
			data, readErr := embedFS.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			t := rootTemplate.New(path).Funcs(funcMap)
			if _, parseErr := t.Parse(string(data)); parseErr != nil {
				return parseErr
			}
		}
		return nil
	})
	return err
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/index.html", nil)
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "web/templates/login.html", nil)
}
