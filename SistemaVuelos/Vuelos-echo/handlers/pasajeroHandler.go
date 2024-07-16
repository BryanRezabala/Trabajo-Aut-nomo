package handlers

import (
    "net/http"
    "Vuelos/models"
    "Vuelos/repository"
    "strconv"

    "github.com/labstack/echo/v4"
)

type PasajeroHandler struct {
    Repo repository.PasajeroRepository
}

func (h *PasajeroHandler) GetAllPasajeros(c echo.Context) error {
    pasajeros, err := h.Repo.GetAll()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajeros)
}

func (h *PasajeroHandler) GetPasajero(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    pasajero, err := h.Repo.GetByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

func (h *PasajeroHandler) CreatePasajero(c echo.Context) error {
    var pasajero models.Pasajero
    if err := c.Bind(&pasajero); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    if err := h.Repo.Create(pasajero); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, pasajero)
}

func (h *PasajeroHandler) UpdatePasajero(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var pasajero models.Pasajero
    if err := c.Bind(&pasajero); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    pasajero.ID = id
    if err := h.Repo.Update(pasajero); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, pasajero)
}

func (h *PasajeroHandler) DeletePasajero(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Repo.Delete(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
