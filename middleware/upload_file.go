package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("input-project-image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()
		// sama seperti asynch kalau sudah dibuka harus ditutup

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer tempFile.Close()
		// fungsi defer kalau diatas terjadi kendala berarti tidak dijalani lagi

		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// io.Copy membawa dua parameter dan berfungsi menandai bahwa src adalah filenya sedangkan tempfile destinasinya, mau ditaruh dimana

		data := tempFile.Name()
		filename := data[8:] // uploads/

		c.Set("dataFile", filename)
		return next(c)
	}
}
