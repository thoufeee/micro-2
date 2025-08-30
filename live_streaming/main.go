package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filepath := filename

		file, err := os.Open(filepath)
		if err != nil {
			c.JSON(http.StatusBadRequest, "File not found")
			return
		}

		defer file.Close()

		fi, _ := file.Stat()
		filesize := fi.Size()

		rangeHeader := c.GetHeader("Range")

		if rangeHeader == "" {
			c.Header("Content-Type", "auido/mp3")
			c.Header("Content-Length", fmt.Sprintf("%d", filesize))
			http.ServeFile(c.Writer, c.Request, filepath)
			return
		}

		var start, end int64
		fmt.Sscanf(strings.Replace(rangeHeader, "bytes=", "", 1), "%d-%d", &start, &end)
		if end == 0 || end >= filesize {
			end = filesize - 1
		}

		length := end - start + 1

		c.Status(http.StatusPartialContent)
		c.Header("Content-Type", "audio/mpeg")
		c.Header("Content-Length", fmt.Sprintf("%d", length))
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, filesize))

		file.Seek(start, 0)

		http.ServeContent(c.Writer, c.Request, filename, fi.ModTime(), file)
	})

	r.Run(":8080")

}
