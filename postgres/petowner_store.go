package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/dasenturk/demo"
	"github.com/jmoiron/sqlx"
)

type PetOwnerStore struct {
	*sqlx.DB
}

func (s *PetOwnerStore) PetOwner(id uuid.UUID) (petsy.PetOwner, error) {
	var u petsy.PetOwner
	if err := s.Get(&u, `SELECT * FROM petowners WHERE id = $1`, id); err != nil {
		return petsy.PetOwner{}, fmt.Errorf("error getting petowner: %w", err)
	}
	return u, nil
}

func (s *PetOwnerStore) PetOwnerByUsername(username string) (petsy.PetOwner, error) {
	var u petsy.PetOwner
	if err := s.Get(&u, `SELECT * FROM petowners WHERE username = $1`, username); err != nil {
		return petsy.PetOwner{}, fmt.Errorf("error getting petowner: %w", err)
	}
	return u, nil
}

func (s *PetOwnerStore) PetOwners() ([]petsy.PetOwner, error) {
	var uu []petsy.PetOwner
	if err := s.Select(&uu, `SELECT * FROM petowners`); err != nil {
		return []petsy.PetOwner{}, fmt.Errorf("error getting petowners: %w", err)
	}
	return uu, nil
}

func (s *PetOwnerStore) PetOwnerByName(name, surname string) (petsy.PetOwner, error){
	var u petsy.PetOwner
	if err := s.Get(&u, `SELECT * FROM petowners WHERE name=$1 AND surname=$2`, name,surname); err!= nil{
		return petsy.PetOwner{}, fmt.Errorf("error getting petowner: %w", err)
	}
	return u, nil
}

func (s *PetOwnerStore) PetOwnerByEmail(email string) (petsy.PetOwner, error){
	var u petsy.PetOwner
	if err := s.Get(&u, `SELECT * FROM petowners WHERE email=$1`, email); err!= nil{
		return petsy.PetOwner{}, fmt.Errorf("error getting petowner: %w", err)
	}
	return u, nil
}

func (s *PetOwnerStore) PetOwnerByCity(cityname string) ([]petsy.PetOwner, error){
	var uu []petsy.PetOwner
	if err := s.Get(&uu, `SELECT * FROM petowners WHERE cityname=$1`, cityname); err!= nil{
		return []petsy.PetOwner{}, fmt.Errorf("error getting petowners: %w", err)
	}
	return uu, nil
}

func (s *PetOwnerStore) CreatePetOwner(u *petsy.PetOwner) error {
	if err := s.Get(u, `INSERT INTO petowners VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		u.ID,
		u.Name,
		u.Surname,
		u.Username,
		u.Password,
		u.Email,
		u.CityName); err != nil {
		return fmt.Errorf("error creating petowner: %w", err)
	}
	return nil
}

func (s *PetOwnerStore) UpdatePetOwner(u *petsy.PetOwner) error {
	if err := s.Get(u, `UPDATE petowners SET username = $1, password = $2, name = $3, surname = $4, email = $5, cityname = $6 WHERE id = $7 RETURNING *`,
		u.Username,
		u.Password,
		u.Name,
		u.Surname,
		u.Email,
		u.CityName,
		u.ID); err != nil {
		return fmt.Errorf("error updating petowner: %w", err)
	}
	return nil
}

func (s *PetOwnerStore) DeletePetOwner(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM petowners WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting PetOwner: %w", err)
	}
	return nil
}