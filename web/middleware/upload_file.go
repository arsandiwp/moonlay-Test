package middleware

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		files := form.File["files"]
		allowedExtensions := map[string]bool{
			".txt": true,
			".pdf": true,
		}

		var dataFiles []string
		for _, file := range files {
			ext := filepath.Ext(file.Filename)
			if !allowedExtensions[ext] {
				return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are .txt and .pdf")
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			tempFile, err := os.Create(filepath.Join("/uploads", file.Filename))
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			_, err = io.Copy(tempFile, src)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			dataFiles = append(dataFiles, file.Filename)
		}
		c.Set("dataFiles", dataFiles)
		return next(c)
	}
}
