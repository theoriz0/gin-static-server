package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed templates/*
	files embed.FS
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(files, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)
	router.GET("/", listFiles)
	router.Static("/static", "./")

	absPath, _ := filepath.Abs("./")
	fmt.Printf("Serving %s\n", absPath)

	port := 8080
	// Listen and serve on 0.0.0.0:8080
	addr := portToAddr(port)
	fmt.Printf("Port %v", port)
	err := router.Run(addr)
	for {
		if err != nil {
			fmt.Print(" failed\n")
			port += 1
			fmt.Printf("Port %v", port)
			addr := portToAddr(port)
			err = router.Run(addr)
		}
	}
}

func portToAddr(port int) string {
	return ":" + strconv.FormatInt(int64(port), 10)
}

func listFiles(ctx *gin.Context) {
	folder, ok := ctx.GetQuery("folder")
	if !ok {
		folder = "."
	}
	files, err := os.ReadDir(folder)
	absPath, err := filepath.Abs(folder)
	if err != nil {
		ctx.Writer.WriteString(err.Error())
		ctx.Writer.Flush()
		return
	}
	ctx.HTML(200, "filelist.tmpl", gin.H{
		"curFolder":    folder,
		"curFolderAbs": absPath,
		"files":        files,
	})
	ctx.Writer.Flush()
}
