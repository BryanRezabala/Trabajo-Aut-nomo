package repository

import (
    "Vuelos/models"
    "github.com/jmoiron/sqlx"
)

type VueloRepository interface {
    GetAll() ([]models.Vuelo, error)
    GetByID(id int) (models.Vuelo, error)
    Create(vuelo models.Vuelo) error
    Update(vuelo models.Vuelo) error
    Delete(id int) error
}

type vueloRepository struct {
    db *sqlx.DB
}

func NewVueloRepository(db *sqlx.DB) VueloRepository {
    return &vueloRepository{db: db}
}

func (r *vueloRepository) GetAll() ([]models.Vuelo, error) {
    var vuelos []models.Vuelo
    err := r.db.Select(&vuelos, "SELECT * FROM vuelo")
    return vuelos, err
}

func (r *vueloRepository) GetByID(id int) (models.Vuelo, error) {
    var vuelo models.Vuelo
    err := r.db.Get(&vuelo, "SELECT * FROM vuelo WHERE id=$1", id)
    return vuelo, err
}

func (r *vueloRepository) Create(vuelo models.Vuelo) error {
    _, err := r.db.NamedExec(`INSERT INTO vuelo (origen, destino, fecha) VALUES (:origen, :destino, :fecha)`, &vuelo)
    return err
}

func (r *vueloRepository) Update(vuelo models.Vuelo) error {
    _, err := r.db.NamedExec(`UPDATE vuelo SET origen=:origen, destino=:destino, fecha=:fecha WHERE id=:id`, &vuelo)
    return err
}

func (r *vueloRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM vuelo WHERE id=$1", id)
    return err
}
