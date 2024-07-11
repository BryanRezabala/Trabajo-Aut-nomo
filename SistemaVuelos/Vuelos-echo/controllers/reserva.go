package controllers

import (
    "net/http"
    "strconv"
    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
    "log"
)

type ReservaController struct {
    DB *pg.DB
}

func (controller *ReservaController) GetAll(c echo.Context) error {
    var reservas []models.Reserva
    err := controller.DB.Model(&reservas).Select()
    if err != nil {
        log.Printf("Error fetching reservas: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reservas)
}

func (controller *ReservaController) Create(c echo.Context) error {
    var reserva models.Reserva
    if err := c.Bind(&reserva); err != nil {
        log.Printf("Error binding reserva: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(reserva)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    log.Printf("Reserva data: %+v", reserva)
    _, err = controller.DB.Model(&reserva).Insert()
    if err != nil {
        log.Printf("Error inserting reserva: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, reserva)
}

func (controller *ReservaController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    err := controller.DB.Model(reserva).WherePK().Select()
    if err != nil {
        log.Printf("Error fetching reserva by ID: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

func (controller *ReservaController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    if err := c.Bind(reserva); err != nil {
        log.Printf("Error binding reserva: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(reserva)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    _, err = controller.DB.Model(reserva).WherePK().Update()
    if err != nil {
        log.Printf("Error updating reserva: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, reserva)
}

func (controller *ReservaController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    reserva := &models.Reserva{ID: int64(id)}
    _, err := controller.DB.Model(reserva).WherePK().Delete()
    if err != nil {
        log.Printf("Error deleting reserva: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
