package main

import (
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
	ProjectName string
	Description string
	StartDate   string
	EndDate     string
	Duration    string
	Java        bool
	Python      bool
	Javascript  bool
	PHP         bool
	Author      string
	PostDate    string
}

// cara ngisi valuenya
// data-data yang ditampung, yang kemudian data yang diisi harus sesuai dengan tipe data yang telah dibangun pada struct
// di golang indeksnya juga sama, jadi data di object pertama itu indeksnya 0 juga
// diatas kita sudah membuat data dalam bentuk array, yang di dalamnya ada beberapa object yang sudah dibangun oleh struct diatas
// jadi data di dalamnya harus sesuai dengan struct diatas
var dataProject = []Project{
	{
		ProjectName: "Dumbways Mobile App",
		Description: "Lorem Ipsum ajalah dulu",
		StartDate:   "2023-03-03",
		EndDate:     "2023-06-06",
		Duration:    "3 Bulan",
		Java:        true,
		Python:      true,
		Javascript:  true,
		PHP:         true,
		PostDate:    "07/06/2023",
	},
	{
		ProjectName: "Pangahaku Mobile App",
		Description: "Lorem Ipsum ajalah dulu",
		StartDate:   "2023-03-03",
		EndDate:     "2023-06-06",
		Duration:    "3 Bulan",
		Java:        true,
		Python:      true,
		Javascript:  true,
		PHP:         true,
		PostDate:    "07/06/2023",
	},
	{
		ProjectName: "Legopedia Mobile App",
		Description: "Lorem Ipsum ajalah dulu",
		StartDate:   "2023-03-03",
		EndDate:     "2023-06-06",
		Duration:    "3 Bulan",
		Java:        true,
		Python:      true,
		Javascript:  true,
		PHP:         true,
		PostDate:    "07/06/2023",
	},
	// ini dinamakan slice
}

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
	e.GET("/project", project)
	e.GET("/contact", contact)
	e.GET("/testimonials", testimonials)
	e.GET("project-detail/:id", projectDetail)
	e.GET("/form-project", formAddProject)
	e.GET("update-project/:id", updatingProject)

	e.POST("/project", addProject)
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

	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// refetch

	projects := map[string]interface{}{
		"Projects": dataProject,
	}
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

	var ProjectDetail = Project{}

	// for melakukan perulangan
	// i = penampung index
	// data = penampung data dari range
	// range = jarak/banyaknya data
	// dataProject = sumber data yang ingin dilakukan perulangan
	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Java:        data.Java,
				Python:      data.Python,
				Javascript:  data.Javascript,
				PHP:         data.PHP,
			}
		}
	}

	// data yang ditampilkan itu cuma data yang ketemu indeksnya, makanya diatas dilakukan pengecekan if id === i yaitu sesuai dengan loopingannya

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updatingProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Java:        data.Java,
				Python:      data.Python,
				Javascript:  data.Javascript,
				PHP:         data.PHP,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/update-project.html")

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
	projectname := c.FormValue("projectName")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	duration := Durasi(startdate, enddate)
	description := c.FormValue("descriptionProject")
	checkone := c.FormValue("inputJava")
	checktwo := c.FormValue("inputPython")
	checkthree := c.FormValue("inputJavascript")
	checkfour := c.FormValue("inputPhp")

	println("Title:" + projectname + ", Description:" + description + ", Start Date:" + startdate + ", End Date:" + enddate + ", Java:" + checkone + ", Python:" + checktwo + ", Javascript:" + checkthree + ", PHP:" + checkfour)

	var newProject = Project{
		ProjectName: projectname,
		Description: description,
		StartDate:   startdate,
		EndDate:     enddate,
		Duration:    duration,
		Java:        (checkone == "inputJava"),
		Python:      (checktwo == "inputPython"),
		Javascript:  (checkthree == "inputJavascript"),
		PHP:         (checkfour == "inputPhp"),
		PostDate:    time.Now().String(),
	}

	// cara agar data yang kita dapatkan di newProject dimasukkan ke penampung data atau slice diatas
	// appaend adalah fungsi yang kita jalankan untuk menambahakan data newProject ke slice dataProject
	// kurang lebihnya mirip dengan fungsi push pada javascript
	// param 1 =  dimana datanya ditampung
	// param 2 = data yang akan ditampung
	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}

func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	projectname := c.FormValue("projectName")
	startdate := c.FormValue("startDate")
	enddate := c.FormValue("endDate")
	duration := Durasi(startdate, enddate)
	description := c.FormValue("descriptionProject")
	checkone := c.FormValue("inputJava")
	checktwo := c.FormValue("inputPython")
	checkthree := c.FormValue("inputJavascript")
	checkfour := c.FormValue("inputPhp")

	fmt.Println("Index:", id)

	var updateProject = Project{
		ProjectName: projectname,
		Description: description,
		StartDate:   startdate,
		EndDate:     enddate,
		Duration:    duration,
		Java:        (checkone == "inputJava"),
		Python:      (checktwo == "inputPython"),
		Javascript:  (checkthree == "inputJavascript"),
		PHP:         (checkfour == "inputPhp"),
	}

	// cara agar data yang kita dapatkan di newProject dimasukkan ke penampung data atau slice diatas
	// appaend adalah fungsi yang kita jalankan untuk menambahakan data newProject ke slice dataProject
	// kurang lebihnya mirip dengan fungsi push pada javascript
	// param 1 =  dimana datanya ditampung
	// param 2 = data yang akan ditampung
	dataProject[id] = updateProject

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}

// trigger delete post
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index:", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)
	// id+1 dimaksudkan agar indeks setelahnya mengisi indeks yang sudah dihapus tadi
	// ditambah 3 titik karena diatas another slice
	// sebenernya di append itu bisa menambahkan slicing yang lain

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func Durasi(startdate, enddate string) string {
	startTime, _ := time.Parse("2006-01-02", startdate)
	endTime, _ := time.Parse("2006-01-02", enddate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " Tahun"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + "Tahun"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " Bulan"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + "Minggu"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + "Minggu"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " Hari"
				} else {
					duration = strconv.Itoa(durationDays) + " Hari"
				}
			}
		}
	}

	return duration
}
