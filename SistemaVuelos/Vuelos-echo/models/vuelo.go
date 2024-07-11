package models

type Vuelo struct {
    ID      int64  `json:"id"`
    Origen  string `json:"origen" validate:"required"`
    Destino string `json:"destino" validate:"required"`
    Fecha   string `json:"fecha" validate:"required"`
}
