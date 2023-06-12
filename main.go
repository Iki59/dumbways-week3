package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //sama sperti let dan sebagainya dia mendeklarasikan titik dua dan sama dengan itu ngisi sekaligus deklarasiin

	// e = echo package
	// GET/POST =  run the method form Http
	// "/" = endpoint/routing (ex, localhost:5000'/', dumbways.id'/lms' jadi lms ini adalah end pointnya)
	// helloWorld = function that will run if the routes are opened

	e.Static("/public", "public")

	// Routing
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/project", Project)
	e.GET("/contact", contact)
	e.GET("/testimonials", testimonials)
	e.GET("project-detail/:id", projectDetail)
	e.GET("/form-project", formAddProject)
	e.POST("/project", addProject)

	// method GET untuk mendapatkan data dari halaman yang kita tuju

	e.Logger.Fatal(e.Start("localhost:5000"))
	//error adalah bagian data yang dijalankan apabila function tidak berjalan
	// diatas bagian npointnya
	// function main adalah fungsi yang pertama kali dijalankan di package main
	// import merupakan kumpulan pckage-package yang kita pakai
	// get sebagai method menjalankan dua parameter yaitu npoint dan function
	// function string menjalankan dua parameter juga yaitu status dari string ketika berjalan statusnya apa, kemudian value yang dijalankan
	// line cod logger fatal yaitu apabila ada kendala ketika menjalankan server dia akan matikan kemudian dikirim melalui lognya
	// kemudian function yang dijalankan didalamnya yaitu e.start berarti dia menjalankan aplikasi ini pada link apa
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
	// string yang dikirimkan ada dua value karena parameternya harus ngirim dua value
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} // method ini buat ngambil value err, jadi ketika error akan dikembalika JSON dengan pesan error diatas, jadi kalau errornya gk kosong maka return itu dijalanin

	return tmpl.Execute(c.Response(), nil)

	// tmpl execute ini mengeksekusi dari tmpl diatas, dan ada dua parameter yang harus dikirimkan di dalamnya

	// tmpl dan err menampung dua value yang dihasilkan dari template.Parsefiles
	// parse digunakan untuk mengurai file, file yang diurai yaitu halaman html yang kita buka
	// var disini akan nampung dua variable, kalau di golang bisa nampung di dua identifier berbeda
	// yang dijalankan disini udah bukan return c.String seperti diatas
	// pengkondisian di golang itu tidak lagi if dengan bracket () jadi langsung dipanggil saja
}

func Project(c echo.Context) error {
	data := map[string]interface{}{
		"Login": true,
	}

	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// tanda garis bawah/under score menandakan variable yang kita pakai disitu itu tidak dipakai, yang sebelumnya biasanya pake error, tapi karena tidak dipakai jadi pakai _
	// strconv.Atoi pada dasarnya mengirim dua value yaitu id dan error/err
	// strconv atau string converter digunakan untuk mengconvert string menjadi tipe data lain atau sebaliknya
	// kita memakai Atoi karena string akan diubah menjadi integer, karena dia menerima value string dan mengubahnya menjadi integer
	// karena nantinya dalam rootingan ada id, jadi tugas param/parameter yaitu menampung dari query string yang kita dapatkan/kirimkan, contohnya 1

	data := map[string]interface{}{
		"Id":      id,
		"Title":   "Dumbways Mobile App",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis. Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis.",
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("descriptionProject")
	checkOne := c.FormValue("inputJava")
	checkTwo := c.FormValue("inputPython")
	checkThree := c.FormValue("inputJavascript")
	checkFour := c.FormValue("inputPhp")

	println("Title:" + projectName + ", Description:" + description + ", Start Date:" + startDate + ", End Date:" + endDate + ", Java:" + checkOne + ", Python:" + checkTwo + ", Javascript:" + checkThree + ", PHP:" + checkFour)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}
