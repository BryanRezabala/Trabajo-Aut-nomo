package handlers

import (
    "net/http"
    "Vuelos/models"
    "Vuelos/repository"
    "strconv"

    "github.com/labstack/echo/v4"
)

type VueloHandler struct {
    Repo repository.VueloRepository
}

func (h *VueloHandler) GetAllVuelos(c echo.Context) error {
    vuelos, err := h.Repo.GetAll()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelos)
}

func (h *VueloHandler) GetVuelo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    vuelo, err := h.Repo.GetByID(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

func (h *VueloHandler) CreateVuelo(c echo.Context) error {
    var vuelo models.Vuelo
    if err := c.Bind(&vuelo); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    if err := h.Repo.Create(vuelo); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, vuelo)
}

func (h *VueloHandler) UpdateVuelo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    var vuelo models.Vuelo
    if err := c.Bind(&vuelo); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    vuelo.ID = id
    if err := h.Repo.Update(vuelo); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusOK, vuelo)
}

func (h *VueloHandler) DeleteVuelo(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Repo.Delete(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.NoContent(http.StatusNoContent)
}
