package controllers

import (
    "net/http"
    "strconv"
    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
    "log"
)

type VueloController struct {
    DB *pg.DB
}

func (controller *VueloController) GetAll(c echo.Context) error {
    var vuelos []models.Vuelo
    err := controller.DB.Model(&vuelos).Select()
    if err != nil {
        log.Printf("Error fetching vuelos: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelos)
}

func (controller *VueloController) Create(c echo.Context) error {
    var vuelo models.Vuelo
    if err := c.Bind(&vuelo); err != nil {
        log.Printf("Error binding vuelo: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(vuelo)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    log.Printf("Vuelo data: %+v", vuelo)
    _, err = controller.DB.Model(&vuelo).Insert()
    if err != nil {
        log.Printf("Error inserting vuelo: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, vuelo)
}

func (controller *VueloController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    err := controller.DB.Model(vuelo).WherePK().Select()
    if err != nil {
        log.Printf("Error fetching vuelo by ID: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

func (controller *VueloController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    if err := c.Bind(vuelo); err != nil {
        log.Printf("Error binding vuelo: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(vuelo)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    _, err = controller.DB.Model(vuelo).WherePK().Update()
    if err != nil {
        log.Printf("Error updating vuelo: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

func (controller *VueloController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    _, err := controller.DB.Model(vuelo).WherePK().Delete()
    if err != nil {
        log.Printf("Error deleting vuelo: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
