package controllers

import (
    "net/http"
    "strconv"

    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
)

// ReservaController maneja las operaciones CRUD para Reserva
type ReservaController struct {
    DB *pg.DB
}

// GetAll devuelve todas las reservas
func (controller *ReservaController) GetAll(c echo.Context) error {
    var reservas []models.Reserva
    err := controller.DB.Model(&reservas).Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reservas)
}

// Create crea una nueva reserva
func (controller *ReservaController) Create(c echo.Context) error {
    var reserva models.Reserva
    if err := c.Bind(&reserva); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(&reserva).Insert()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, reserva)
}

// GetById devuelve una reserva por ID
func (controller *ReservaController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    err := controller.DB.Model(reserva).WherePK().Select()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

// Update actualiza una reserva por ID
func (controller *ReservaController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    if err := c.Bind(reserva); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    _, err := controller.DB.Model(reserva).WherePK().Update()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

// Delete elimina una reserva por ID
func (controller *ReservaController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    _, err := controller.DB.Model(reserva).WherePK().Delete()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
