package repository

import (
    "Vuelos/models"
    "github.com/jmoiron/sqlx"
)

type PasajeroRepository interface {
    GetAll() ([]models.Pasajero, error)
    GetByID(id int) (models.Pasajero, error)
    Create(pasajero models.Pasajero) error
    Update(pasajero models.Pasajero) error
    Delete(id int) error
}

type pasajeroRepository struct {
    db *sqlx.DB
}

func NewPasajeroRepository(db *sqlx.DB) PasajeroRepository {
    return &pasajeroRepository{db: db}
}

func (r *pasajeroRepository) GetAll() ([]models.Pasajero, error) {
    var pasajeros []models.Pasajero
    err := r.db.Select(&pasajeros, "SELECT * FROM pasajero")
    return pasajeros, err
}

func (r *pasajeroRepository) GetByID(id int) (models.Pasajero, error) {
    var pasajero models.Pasajero
    err := r.db.Get(&pasajero, "SELECT * FROM pasajero WHERE id=$1", id)
    return pasajero, err
}

func (r *pasajeroRepository) Create(pasajero models.Pasajero) error {
    _, err := r.db.NamedExec(`INSERT INTO pasajero (nombre, apellido, email) VALUES (:nombre, :apellido, :email)`, &pasajero)
    return err
}

func (r *pasajeroRepository) Update(pasajero models.Pasajero) error {
    _, err := r.db.NamedExec(`UPDATE pasajero SET nombre=:nombre, apellido=:apellido, email=:email WHERE id=:id`, &pasajero)
    return err
}

func (r *pasajeroRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM pasajero WHERE id=$1", id)
    return err
}
