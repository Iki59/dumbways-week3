package main

import (
	"context"
	"dumbwaysgolang/connection"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// nama dari structnya adalah Project
// yang membangun object/properties
type Project struct {
	ID                int
	ProjectName       string
	Description       string
	StartDate         time.Time
	EndDate           time.Time
	Duration          string
	Java              bool
	Python            bool
	Javascript        bool
	PHP               bool
	Author            string
	PostDate          time.Time
	Image             string
	CheckJava         string
	CheckPython       string
	CheckJavascript   string
	CheckPHP          string
	FormatDate        string
	Duration_Format   string
	Format_Start_Date string
	Format_End_Date   string
}

// cara ngisi valuenya
// data-data yang ditampung, yang kemudian data yang diisi harus sesuai dengan tipe data yang telah dibangun pada struct
// di golang indeksnya juga sama, jadi data di object pertama itu indeksnya 0 juga
// diatas kita sudah membuat data dalam bentuk array, yang di dalamnya ada beberapa object yang sudah dibangun oleh struct diatas
// jadi data di dalamnya harus sesuai dengan struct diatas
var dataProject = []Project{
	// {
	// 	ProjectName: "Dumbways Mobile App",
	// 	Description: "Lorem Ipsum ajalah dulu",
	// 	StartDate:   "2023-03-03",
	// 	EndDate:     "2023-06-06",
	// 	Duration:    "3 Bulan",
	// 	Java:        true,
	// 	Python:      true,
	// 	Javascript:  true,
	// 	PHP:         true,
	// },
	// {
	// 	ProjectName: "Pangahaku Mobile App",
	// 	Description: "Lorem Ipsum ajalah dulu",
	// 	StartDate:   "2023-03-03",
	// 	EndDate:     "2023-06-06",
	// 	Duration:    "3 Bulan",
	// 	Java:        true,
	// 	Python:      true,
	// 	Javascript:  true,
	// 	PHP:         true,
	// },
	// {
	// 	ProjectName: "Legopedia Mobile App",
	// 	Description: "Lorem Ipsum ajalah dulu",
	// 	StartDate:   "2023-03-03",
	// 	EndDate:     "2023-06-06",
	// 	Duration:    "3 Bulan",
	// 	Java:        true,
	// 	Python:      true,
	// 	Javascript:  true,
	// 	PHP:         true,
	// },
	// ini dinamakan slice
}

func main() {
	// db connection
	connection.DatabaseConnect()
	e := echo.New() //sama sperti let dan sebagainya dia mendeklarasikan titik dua dan sama dengan itu ngisi sekaligus deklarasiin

	// e = echo package
	// GET/POST =  run the method form Http
	// "/" = endpoint/routing (ex, localhost:5000'/', dumbways.id'/lms' jadi lms ini adalah end pointnya)
	// helloWorld = function that will run if the routes are opened

	e.Static("/public", "public")

	// Routing
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/", project)
	e.GET("/contact", contact)
	e.GET("/testimonials", testimonials)
	e.GET("project-detail/:id", projectDetail)
	e.GET("/form-project", formAddProject)
	e.GET("update-project/:id", updatingProject)

	e.POST("/", addProject)
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", updateProject)

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

func project(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, description, java, python, javascript, php, image, start_date, end_date, duration FROM tbl_project")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.ID, &each.ProjectName, &each.Description, &each.Java, &each.Python, &each.Javascript, &each.PHP, &each.Image, &each.StartDate, &each.EndDate, &each.Duration)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("5 September 1999")
		each.Author = "Muhammad Rizki B"
		// each.Duration = each.EndDate.Sub(each.StartDate)
		// each.Duration_Format = Durasi(each.Duration)

		result = append(result, each)
	}

	projects := map[string]interface{}{
		// "Projects": dataProject,
		"Projects": result,
	}
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// refetch

	// disini kita mapping projects dengan data Project yang ada dibawah
	// jadi kita sudah tidak memanggil dataProject dengan namanya sendiri, melainkan dengan nama aliasnya karena sudah di mapping oleh projects
	// sehingga di html nya itu pemag\nggilan datanya {{.projects.Title}}
	return tmpl.Execute(c.Response(), projects)
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

	// data := map[string]interface{}{
	// 	"Id":      id,
	// 	"Title":   "Dumbways Mobile App",
	// 	"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis. Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum nam excepturi reprehenderit molestias vero laboriosam, accusamus aliquam? Voluptatem, libero consectetur quaerat aspernatur porro quidem error facere reprehenderit omnis earum nisi quos aperiam soluta vel tempora dignissimos possimus facilis quas, animi eaque nostrum suscipit perferendis optio ullam? Praesentium excepturi animi eius illum autem voluptates labore. Libero excepturi nisi ipsam veritatis est voluptatibus voluptates recusandae sapiente dolore distinctio! Cumque asperiores corporis necessitatibus, quisquam neque adipisci. Itaque, natus harum sint eum nesciunt ea ipsa perferendis porro soluta magni, corporis asperiores accusamus sed minus? Laudantium aperiam rem beatae voluptatum ipsum ipsam at dignissimos nobis.",
	// }
	// cara mengisi project detail sesuai dengan yang diinput

	var ProjectDetail = Project{} // ini masih dipakai untuk database

	// kenapa pake query row karena yang dicari hanya satu baris data pada table
	// pakai $1 karena dia mendapatkan value dari id, maksudnya value pertamanya nanti diisi sesuai dengan id, $1 akan diisi sesuai dengan id kemudian bisa diisi value kedua dengan $2 diikuti value setelahnya.
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, description, java, python, javascript, php, image, start_date, end_date, duration FROM tbl_project WHERE id=$1", id).Scan(
		&ProjectDetail.ID, &ProjectDetail.ProjectName, &ProjectDetail.Description, &ProjectDetail.Java, &ProjectDetail.Python, &ProjectDetail.Javascript, &ProjectDetail.PHP, &ProjectDetail.Image, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration)
	// guna dari tanda & yaitu sebagai pointer, agar si ProjectDetail tau tipe datanya apa sesuai dengan struct yang dipasangkan

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}

	ProjectDetail.Author = "Muhammad Rizki B"
	ProjectDetail.FormatDate = ProjectDetail.PostDate.Format("2 January 2006")
	// ProjectDetail.Duration_Format = Durasi(ProjectDetail.Duration)
	ProjectDetail.Format_Start_Date = ProjectDetail.StartDate.Format("2 January 2006")
	ProjectDetail.Format_End_Date = ProjectDetail.EndDate.Format("2 January 2006")

	// for melakukan perulangan
	// i = penampung index
	// data = penampung data dari range
	// range = jarak/banyaknya data
	// dataProject = sumber data yang ingin dilakukan perulangan
	// for i, data := range dataProject {
	// 	if id == i {
	// 		ProjectDetail = Project{
	// 			ProjectName: data.ProjectName,
	// 			StartDate:   data.StartDate,
	// 			EndDate:     data.EndDate,
	// 			Duration:    data.Duration,
	// 			Description: data.Description,
	// 			Java:        data.Java,
	// 			Python:      data.Python,
	// 			Javascript:  data.Javascript,
	// 			PHP:         data.PHP,
	// 		}
	// 	}
	// }
	// data yang ditampilkan itu cuma data yang ketemu indeksnya, makanya diatas dilakukan pengecekan if id === i yaitu sesuai dengan loopingannya

	data := map[string]interface{}{
		"Project_Detail": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updatingProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectUpdate = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, description, java, python, javascript, php, image, start_date, end_date, duration FROM tbl_project WHERE id=$1", id).Scan(
		&ProjectUpdate.ID, &ProjectUpdate.ProjectName, &ProjectUpdate.Description, &ProjectUpdate.Java, &ProjectUpdate.Python, &ProjectUpdate.Javascript, &ProjectUpdate.PHP, &ProjectUpdate.Image, &ProjectUpdate.StartDate, &ProjectUpdate.EndDate, &ProjectUpdate.Duration)

	// for i, data := range dataProject {
	// 	if id == i {
	// 		ProjectDetail = Project{
	// 			ProjectName: data.ProjectName,
	// 			StartDate:   data.StartDate,
	// 			EndDate:     data.EndDate,
	// 			Duration:    data.Duration,
	// 			Description: data.Description,
	// 			Java:        data.Java,
	// 			Python:      data.Python,
	// 			Javascript:  data.Javascript,
	// 			PHP:         data.PHP,
	// 		}
	// 	}
	// }

	data := map[string]interface{}{
		"Project": ProjectUpdate,
	}

	var tmpl, errTemplate = template.ParseFiles("views/update-project.html")

	if errTemplate != nil {
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
	projectname := c.FormValue("projectName")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	duration := Durasi(startdate, enddate)
	description := c.FormValue("descriptionProject")
	// checkone := c.FormValue("inputJava")
	// checktwo := c.FormValue("inputPython")
	// checkthree := c.FormValue("inputJavascript")
	// checkfour := c.FormValue("inputPhp")
	// author:= "Muhammad Rizki B"

	var checkone bool
	if c.FormValue("inputJava") == "inputJava" {
		checkone = true
	}

	var checktwo bool
	if c.FormValue("inputPython") == "inputPython" {
		checktwo = true
	}

	var checkthree bool
	if c.FormValue("inputJavascript") == "inputJavascript" {
		checkthree = true
	}

	var checkfour bool
	if c.FormValue("inputPhp") == "inputPhp" {
		checkfour = true
	}

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tbl_project (name, description, java, python, javascript, php, image, start_date, end_date, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", projectname, description, checkone, checktwo, checkthree, checkfour, "apps-img-jpg", startdate, enddate, duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	// value pertama tidak dipanggil, kita cuma butuh errornya saja
	// kenapa pake exec karena dia mengeksekusi sesuatu yaitu menambahkan data
	// println("Title:" + projectname + ", Description:" + description + ", Start Date:" + startdate + ", End Date:" + enddate + ", Java:" + checkone + ", Python:" + checktwo + ", Javascript:" + checkthree + ", PHP:" + checkfour)

	// var newProject = Project{
	// 	ProjectName: projectname,
	// 	Description: description,
	// StartDate:   startdate,
	// EndDate:     enddate,
	// Duration:   duration,
	// Java:       (checkone == "inputJava"),
	// Python:     (checktwo == "inputPython"),
	// Javascript: (checkthree == "inputJavascript"),
	// PHP:        (checkfour == "inputPhp"),
	// PostDate:    time.Now().String(),
	// }

	// cara agar data yang kita dapatkan di newProject dimasukkan ke penampung data atau slice diatas
	// appaend adalah fungsi yang kita jalankan untuk menambahakan data newProject ke slice dataProject
	// kurang lebihnya mirip dengan fungsi push pada javascript
	// param 1 =  dimana datanya ditampung
	// param 2 = data yang akan ditampung
	// dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID :", id)

	projectname := c.FormValue("projectName")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	duration := Durasi(startdate, enddate)
	description := c.FormValue("descriptionProject")
	var checkone bool
	if c.FormValue("inputJava") == "inputJava" {
		checkone = true
	}

	var checktwo bool
	if c.FormValue("inputPython") == "inputPython" {
		checktwo = true
	}

	var checkthree bool
	if c.FormValue("inputJavascript") == "inputJavascript" {
		checkthree = true
	}

	var checkfour bool
	if c.FormValue("inputPhp") == "inputPhp" {
		checkfour = true
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tbl_project SET name=$1, description=$2, java=$3, python=$4, javascript=$5, php=$6, image=$7, start_date=$8, end_date=$9, duration=$10", projectname, description, checkone, checktwo, checkthree, checkfour, "apps-img-jpg", startdate, enddate, duration)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// var updateProject = Project{
	// 	ProjectName: projectname,
	// 	Description: description,
	// 	// StartDate:   startdate,
	// 	// EndDate:     enddate,
	// 	// Duration:   duration,
	// 	Java:       (checkone == "inputJava"),
	// 	Python:     (checktwo == "inputPython"),
	// 	Javascript: (checkthree == "inputJavascript"),
	// 	PHP:        (checkfour == "inputPhp"),
	// }

	// cara agar data yang kita dapatkan di newProject dimasukkan ke penampung data atau slice diatas
	// appaend adalah fungsi yang kita jalankan untuk menambahakan data newProject ke slice dataProject
	// kurang lebihnya mirip dengan fungsi push pada javascript
	// param 1 =  dimana datanya ditampung
	// param 2 = data yang akan ditampung
	// dataProject[id] = updateProject

	// fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

// trigger delete post
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID:", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tbl_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// dataProject = append(dataProject[:id], dataProject[id+1:]...)
	// id+1 dimaksudkan agar indeks setelahnya mengisi indeks yang sudah dihapus tadi
	// ditambah 3 titik karena diatas another slice
	// sebenernya di append itu bisa menambahkan slicing yang lain

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func Durasi(startdate string, enddate string) string {
	start, _ := time.Parse("2006-01-02", startdate)
	end, _ := time.Parse("2006-01-02", enddate)

	durTime := int(end.Sub(start).Hours())
	durDays := durTime / 24
	durWeeks := durDays / 7
	durMonths := durWeeks / 4
	durYears := durMonths / 12

	var duration string

	if durYears > 1 {
		duration = strconv.Itoa(durYears) + " years"
	} else if durYears > 0 {
		duration = strconv.Itoa(durYears) + " year"
	} else {
		if durMonths > 1 {
			duration = strconv.Itoa(durMonths) + " months"
		} else if durMonths > 0 {
			duration = strconv.Itoa(durMonths) + " month"
		} else {
			if durWeeks > 1 {
				duration = strconv.Itoa(durWeeks) + " weeks"
			} else if durWeeks > 0 {
				duration = strconv.Itoa(durWeeks) + " week"
			} else {
				if durDays > 1 {
					duration = strconv.Itoa(durDays) + " days"
				} else {
					duration = strconv.Itoa(durDays) + " day"
				}
			}
		}
	}

	return duration
}
