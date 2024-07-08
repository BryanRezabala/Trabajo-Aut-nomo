package controllers

import (
    "net/http"
    "strconv"

    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
)

// VueloController maneja las operaciones CRUD para Vuelo
type VueloController struct {
    DB *pg.DB
}

// GetAll devuelve todos los vuelos
func (controller *VueloController) GetAll(c echo.Context) error {
    var vuelos []models.Vuelo
    err := controller.DB.Model(&vuelos).Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelos)
}

// Create crea un nuevo vuelo
func (controller *VueloController) Create(c echo.Context) error {
    var vuelo models.Vuelo
    if err := c.Bind(&vuelo); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(&vuelo).Insert()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, vuelo)
}

// GetById devuelve un vuelo por ID
func (controller *VueloController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    err := controller.DB.Model(vuelo).WherePK().Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

// Update actualiza un vuelo por ID
func (controller *VueloController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    if err := c.Bind(vuelo); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(vuelo).WherePK().Update()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

// Delete elimina un vuelo por ID
func (controller *VueloController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo := &models.Vuelo{ID: int64(id)}
    _, err := controller.DB.Model(vuelo).WherePK().Delete()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
