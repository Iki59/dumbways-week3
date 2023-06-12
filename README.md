# How to declare using Go?

## VAR

example

```go
var  x = 3
```

note: Var can be put inside a function or not

## := 

example

```go
func main() {
    x := 3
}
```

note: := cannot be outside of a function


variable di golang sama saja seperti kita mendeklarasikan variable pada umumnya. Let tidak ada di golang.
Pada dasarnya variable di golang sama seperti let, dia tidak bisa dideklarasikan ulang kadang-kadang apalagi kalau di dalam fungsi.
Tidak bisa pakai := kalau di luar fungsi
Const juga ada di golang, dan sama saja perannya, tidak bisa dideklarasikan ulang ataupun dganti valuenya, jadi fungsinya di golang sama saja.



# Query String

## Query String

Query string adalah

## Project Detail

```go
e.GET("project-detail/:id", projectDetail)
```

note: karena di project detail ada query string yaitu /:id maka untuk bagian image dan css di html-nya harus ditambah /(slice) sebelum public/...


# ROUTING

## Routing adalah cara kita untuk mengorganisir dari state sebuah aplikasi. Pada dasarnya di sebuah aplikasi banyak state atau fungsi yang kita kirimkan, baik itu merender data, get maupun post, jadi untuk memanage hal itu maka digunakanlah routing.

