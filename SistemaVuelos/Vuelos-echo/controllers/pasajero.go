package controllers

import (
    "net/http"
    "strconv"

    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
)

// PasajeroController maneja las operaciones CRUD para Pasajero
type PasajeroController struct {
    DB *pg.DB
}

// GetAll devuelve todos los pasajeros
func (controller *PasajeroController) GetAll(c echo.Context) error {
    var pasajeros []models.Pasajero
    err := controller.DB.Model(&pasajeros).Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajeros)
}

// Create crea un nuevo pasajero
func (controller *PasajeroController) Create(c echo.Context) error {
    var pasajero models.Pasajero
    if err := c.Bind(&pasajero); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(&pasajero).Insert()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, pasajero)
}

// GetById devuelve un pasajero por ID
func (controller *PasajeroController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    err := controller.DB.Model(pasajero).WherePK().Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

// Update actualiza un pasajero por ID
func (controller *PasajeroController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    if err := c.Bind(pasajero); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(pasajero).WherePK().Update()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

// Delete elimina un pasajero por ID
func (controller *PasajeroController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    _, err := controller.DB.Model(pasajero).WherePK().Delete()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
