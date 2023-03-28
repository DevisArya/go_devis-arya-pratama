package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
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
// delete user by id
func DeleteUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id")) 

  if err != nil{
    return c.JSON(400, map[string]interface{}{
      "error" : "invalid id",
    })
  }

  var found int
	for i , val := range users{
		if val.Id == id{
			users = append(users[:i], users[i+1:]...)
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
		"messages": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

  updateUser := new(User)
  if err := c.Bind(updateUser); err != nil{
    return err
  }

  id, _ := strconv.Atoi(c.Param("id")) 

  for i, val := range users{
    if val.Id == id{
			users[i].Name = updateUser.Name
			users[i].Email = updateUser.Email
			users[i].Password = updateUser.Password
      break
		}
  }

  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages" : "success update user",
  })
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}
// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.POST("/users", CreateUserController)
  e.GET("/user/:id", GetUserController)
  e.DELETE("/user/:id", DeleteUserController)
  e.PUT("/user/:id", UpdateUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}