package main

import (
	"net/http"
	"github.com/yejin-lim/eks-workshop-sample-api-service-go"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {		
        // OK 이면 index.html파일에 JSON데이터를 넘겨서 보여줌 
		c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Home Page",
			},
		)
	})

	r.Run()

}
