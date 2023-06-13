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

# Struct Interface

## Struct

- Point dari struct adalah membangun beberapa object/properties, kurang lebihnya sama dengan membuat class pada js

Contoh code dari struct

```go
type Name struct {
    FirstName string
    LastName string
}
```

note: jadi diatas kita membuat struct yang di dalamnya terdapat beberapa properties dengan tipe datanya masing-masing.
Struct disini membangun struct atau propertiesnya saja bukan valuenya

- Create Blog with Slice Append. Ketika kita sudah membangun sebuah slice baru, kita bisa meng-append data baru di dalamnya


## Interface

- Poin dari interface yaitu membangun beberapa method, kemudian dia akan melakukan formatting atau returning. maksudnya dia bisa mengubah value atau mengambil data dalam struct. seperti di js yaitu konsep setter getter encapsulation. Jadi disini kita bisa membuat beberapa method yang tugasnya ngambil atau ngubah data

Contoh code dari interface:

```go
type Person interface {
    getFirstName() string
}

func () getFirstName() string {
    return FirstName
}
```

note: dalam fungsi diatas kita membangun method getFirstName dan return valuenya yaitu First Name dengan tipe datanya string


# StrConv

## Apa itu strconv?

strconv merupakan sebuah package yang membantu kita untuk mengkonversikan tipe data

- Atoi: mengubah string menjadi integer

# Referensi

[Slice Append Cheat Sheet](https://ueokande.github.io/go-slice-tericks)