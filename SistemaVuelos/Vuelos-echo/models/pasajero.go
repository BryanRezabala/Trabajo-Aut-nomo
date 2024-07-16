package models

type Pasajero struct {
    ID       int    `db:"id" json:"id"`
    Nombre   string `db:"nombre" json:"nombre"`
    Apellido string `db:"apellido" json:"apellido"`
    Email    string `db:"email" json:"email"`
}
