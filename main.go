package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Static("/", "./")
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
