package controller

import (
	"github.com/avtara/mvc-go/lib/database"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllClientController(c echo.Context) error {
	clients, e := database.GetClients()
	if e != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, clients)
}

func GetClientByIDController(c echo.Context) error {
	ID := c.Param("id")
	i, _ := strconv.Atoi(ID)
	clients, e := database.GetClientByID(i)
	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusOK, ModelToClientDetailResponse(clients))
}

func PostNewClient(c echo.Context) (err error) {
	u := new(PostClientRequest)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	client, e := database.InsertClient(PostClientsRequestToModelRequest(u))
	if err != nil {
		if e != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": e.Error(),
			})
		}
	}
	return c.JSON(http.StatusCreated, ModelToClientDetailResponse(client))
}

func PutClient(c echo.Context) (err error) {
	u := new(PutClientRequest)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}
	_, e := database.UpdateClient(PutClientsRequestToModelRequest(u))
	client, err := database.GetClientByID(int(u.ID))
	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": e.Error(),
		})
	}
	return c.JSON(http.StatusOK, ModelToClientDetailResponse(client))
}

func DeleteClient(c echo.Context) (err error) {
	ID := c.Param("id")
	i, _ := strconv.Atoi(ID)
	_, e := database.DeleteClient(i)
	if e != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": e.Error(),
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "deleted",
	})
}