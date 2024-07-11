package controllers

import (
    "net/http"
    "strconv"
    "Vuelos-echo/models"
    "github.com/go-pg/pg/v10"
    "github.com/labstack/echo/v4"
    "log"
)

type PasajeroController struct {
    DB *pg.DB
}

func (controller *PasajeroController) GetAll(c echo.Context) error {
    var pasajeros []models.Pasajero
    err := controller.DB.Model(&pasajeros).Select()
    if err != nil {
        log.Printf("Error fetching pasajeros: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajeros)
}

func (controller *PasajeroController) Create(c echo.Context) error {
    var pasajero models.Pasajero
    if err := c.Bind(&pasajero); err != nil {
        log.Printf("Error binding pasajero: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(pasajero)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    log.Printf("Pasajero data: %+v", pasajero)
    _, err = controller.DB.Model(&pasajero).Insert()
    if err != nil {
        log.Printf("Error inserting pasajero: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, pasajero)
}

func (controller *PasajeroController) GetById(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    err := controller.DB.Model(pasajero).WherePK().Select()
    if err != nil {
        log.Printf("Error fetching pasajero by ID: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

func (controller *PasajeroController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    if err := c.Bind(pasajero); err != nil {
        log.Printf("Error binding pasajero: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    // Validar los datos
    err := models.Validate.Struct(pasajero)
    if err != nil {
        log.Printf("Validation error: %v", err)
        return c.JSON(http.StatusBadRequest, err)
    }

    _, err = controller.DB.Model(pasajero).WherePK().Update()
    if err != nil {
        log.Printf("Error updating pasajero: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

func (controller *PasajeroController) Delete(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero := &models.Pasajero{ID: int64(id)}
    _, err := controller.DB.Model(pasajero).WherePK().Delete()
    if err != nil {
        log.Printf("Error deleting pasajero: %v", err)
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
