package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	InitRouter(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8080")
}
