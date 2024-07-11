package models

type Pasajero struct {
    ID        int64  `json:"id"`
    Nombre    string `json:"nombre" validate:"required"`
    Apellido  string `json:"apellido" validate:"required"`
    Email     string `json:"email" validate:"required,email"`
}
