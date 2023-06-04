package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //sama sperti let dan sebagainya dia mendeklarasikan titik dua dan sama dengan itu ngisi sekaligus deklarasiin

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	//error adalah bagian data yang dijalankan apabila function tidak berjalan
	// diatas bagian npointnya
	// function main adalah fungsi yang pertama kali dijalankan di package main
	// import merupakan kumpulan pckage-package yang kita pakai
	// get sebagai method menjalankan dua parameter yaitu npoint dan function
	// function string menjalankan dua parameter juga yaitu status dari string ketika berjalan statusnya apa kemudian value yag dijalankan
	// line cod logger fatal yaitu apabila ada kendala ketika menjalankan server dia akan matikan kemudian dikirim melalui lognya

	e.Logger.Fatal(e.Start("localhost:5000"))
}
