## ECHO

Echo merupakan salah satu framework GoLang yang high performance, extensible, minimalist Go web framework. Framework Echo ini lebih baik karena sintaks minimalis, fitur yang cukup lengkap dan pastinya lebih mudah dipahami oleh pemula. Echo juga memiliki Forum & Chat untuk para developer.
cara install echo :

```cli
$ mkdir myapp && cd myapp
$ go mod init myapp
$ go get github.com/labstack/echo/v4
```

---

## URL PARAM

Method .Param() digunakan untuk mengambil data path parameter sesuai skema rute.

Contoh :

```Go
type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

  if err != nil{
    return c.JSON(400, map[string]interface{}{
      "error" : "invalid id",
    })
  }
  var indexData, found int

	for i , val := range users{
		if val.Id == id{
			indexData = i
      found = 1
      break
		}
	}

  if found == 0{
    return c.JSON(404, map[string]interface{}{
      "error" : "id not found",
    })
  }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user",
		"user":    users[indexData],
	})
}

func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/user/:id", GetUserController)
  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}
```

---

## BINDING DATA

Parsing request data adalah bagian yang penting dalam aplikasi web. Pada framework echo bisa dilakukan dengan proses yang dinamakan binding

contohnya :

```Go
type User struct {
  ID string `query:"id"`
}

// in the handler for /users?id=<userID>
var user User
err := c.Bind(&user); if err != nil {
    return c.String(http.StatusBadRequest, "bad request")
}

```
