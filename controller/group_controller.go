package controller

import (
    "net/http"

    "github.com/andersonlira/group-api/gateway/txtdb"
    "github.com/andersonlira/group-api/domain"
	"github.com/labstack/echo/v4"

)


//GetGroupList return all objects 
func GetGroupList(c echo.Context) error {

    list := txtdb.GetGroupList()

	return c.JSON(http.StatusOK, list)
}

func GetGroupByID(c echo.Context) error {
    ID := c.Param("id")
    it, err := txtdb.GetGroupByID(ID)
    if err != nil {
        return c.JSON(http.StatusNotFound,it)
    }
    return c.JSON(http.StatusOK, it)
}

func SaveGroup(c echo.Context) error {
    it := domain.Group{}
    c.Bind(&it)
    it = txtdb.SaveGroup(it)
    return c.JSON(http.StatusCreated, it)
}

func UpdateGroup(c echo.Context) error {
    ID := c.Param("id")
    it := domain.Group{}
    c.Bind(&it)
    it = txtdb.UpdateGroup(ID,it)
    return c.JSON(http.StatusOK, it)
}

func DeleteGroup(c echo.Context) error {
    ID := c.Param("id")
    result := txtdb.DeleteGroup(ID)
    return c.JSON(http.StatusOK, result)
}