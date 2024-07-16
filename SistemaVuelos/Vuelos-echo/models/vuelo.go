package models

type Vuelo struct {
    ID      int    `db:"id" json:"id"`
    Origen  string `db:"origen" json:"origen"`
    Destino string `db:"destino" json:"destino"`
    Fecha   string `db:"fecha" json:"fecha"`
}
