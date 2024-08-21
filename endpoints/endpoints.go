package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"sample_api/models"

	"github.com/labstack/echo/v4"
)

var users []models.RequestData

func SetEndpoints(server *echo.Echo) {
	server.GET("/get_users", getUsersHandler)
	server.POST("/create_user", createUserHandler)
	server.PUT("/update_user", updateUserHandler)
	server.DELETE("/delete_user/:name", deleteUserHandler)

}

func getUsersHandler(c echo.Context) error {
	var w http.ResponseWriter

	err := getUsers(&users, w)
	if err != nil {
		return c.JSON(http.StatusOK, users)
	}

	return c.String(http.StatusInternalServerError, "Users can be gotten")
}

func getUsers(users *[]models.RequestData, w http.ResponseWriter) error {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return err
	}

	return nil
}

func createUserHandler(c echo.Context) error {
	err := createUser(c)
	if err != nil {
		return c.JSON(http.StatusOK, users)
	}

	return c.String(http.StatusBadRequest, err.Error())
}

func createUser(ctx echo.Context) error {
	var w http.ResponseWriter
	user, err := parseBody(ctx)

	if err != nil {
		return err
	}

	users := append(users, *user)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return errors.New("User data is incorrect")
	}

	return nil
}

func parseBody(ctx echo.Context) (*models.RequestData, error) {
	dataUser := new(models.RequestData)

	defer ctx.Request().Body.Close()

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataUser)
	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

func updateUserHandler(ctx echo.Context) error {
	err := updateUser(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}

func updateUser(ctx echo.Context) error {
	dataUser := new(models.RequestData)

	defer ctx.Request().Body.Close()

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataUser)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("users is empty")
	} else {
		for index, user := range users {
			if user.Name == dataUser.Name {
				users[index] = *dataUser
				return nil
			}
		}
	}

	return errors.New("User is not found")
}

func deleteUserHandler(ctx echo.Context) error {
	err := deleteUser(ctx)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, users)
}

func deleteUser(ctx echo.Context) error {
	name := ctx.Param("name")

	for index, user := range users {
		if user.Name == name {
			users = append(users[:index], users[index+1:]...)
			return nil
		}
	}

	return errors.New("User not found")
}
