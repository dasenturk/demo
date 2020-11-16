package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/dasenturk/demo"
	"github.com/jmoiron/sqlx"
)

type PetStore struct {
	*sqlx.DB
}

func (s *PetStore) Pet(id uuid.UUID) (petsy.Pet, error) {
	var u petsy.Pet
	if err := s.Get(&u, `SELECT * FROM pets WHERE id = $1`, id); err != nil {
		return petsy.Pet{}, fmt.Errorf("error getting pet: %w", err)
	}
	return u, nil
}

func (s *PetStore) PetByName(name string) ([]petsy.Pet, error){
	var uu []petsy.Pet
	if err := s.Get(&uu, `SELECT * FROM pets WHERE name=$1`, name); err!= nil{
		return []petsy.Pet{}, fmt.Errorf("error getting pets: %w", err)
	}
	return uu, nil
}

func (s *PetStore) PetByType(pettype string) ([]petsy.Pet, error){
	var uu []petsy.Pet
	if err := s.Get(&uu, `SELECT * FROM pets WHERE pettype=$1`, pettype); err!= nil{
		return []petsy.Pet{}, fmt.Errorf("error getting pets: %w", err)
	}
	return uu, nil
}

func (s *PetStore) PetByOwner(ownerID uuid.UUID) ([]petsy.Pet, error){
	var uu []petsy.Pet
	if err := s.Get(&uu, `SELECT * FROM pets WHERE owner_id=$1`, ownerID); err!= nil{
		return []petsy.Pet{}, fmt.Errorf("error getting pets: %w", err)
	}
	return uu, nil
}

func (s *PetStore) PetByBreed(breed string) ([]petsy.Pet, error){
	var uu []petsy.Pet
	if err := s.Get(&uu, `SELECT * FROM pets WHERE breed_type=$1`, breed); err!= nil{
		return []petsy.Pet{}, fmt.Errorf("error getting pets: %w", err)
	}
	return uu, nil
}

func (s *PetStore) PetByGender(gender string) ([]petsy.Pet, error){
	var uu []petsy.Pet
	if err := s.Get(&uu, `SELECT * FROM pets WHERE gender=$1`, gender); err!= nil{
		return []petsy.Pet{}, fmt.Errorf("error getting pets: %w", err)
	}
	return uu, nil
}

func (s *PetStore) CreatePet(u *petsy.Pet) error {
	if err := s.Get(u, `INSERT INTO pets VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`,
		u.ID,
		u.Name,
		u.PetType,
		u.Breed,
		u.OwnerID,
		u.Gender,
		u.Weight,
		u.Age); err != nil {
		return fmt.Errorf("error creating pet: %w", err)
	}
	return nil
}

func (s *PetStore) UpdatePet(u *petsy.Pet) error {
	if err := s.Get(u, `UPDATE pets SET name = $1, pettype = $2, breed = $3, owner_id = $4, gender = $5, weight = $6, age = $7 WHERE id = $8 RETURNING *`,
		u.Name,
		u.PetType,
		u.Breed,
		u.OwnerID,
		u.Gender,
		u.Weight,
		u.Age,
		u.ID); err != nil {
		return fmt.Errorf("error updating pet: %w", err)
	}
	return nil
}

func (s *PetStore) DeletePet(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM pets WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting pet: %w", err)
	}
	return nil
}