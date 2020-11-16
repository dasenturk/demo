package postgres

import (
	"fmt"

	"github.com/dasenturk/demo"
	"github.com/jmoiron/sqlx"
)

type BreedStore struct {
	*sqlx.DB
}

func (s *BreedStore) Breed(name string) (petsy.Breed, error) {
	var b petsy.Breed
	if err := s.Get(&b, `SELECT * FROM breeds WHERE name = $1`, name); err != nil {
		return petsy.Breed{}, fmt.Errorf("error getting breed: %w", err)
	}
	return b, nil
}

func (s *BreedStore) BreedByType(btype string) ([]petsy.Breed, error){
	var bb []petsy.Breed
	if err := s.Get(&bb, `SELECT * FROM breeds WHERE btype=$1`, btype); err!= nil{
		return []petsy.Breed{}, fmt.Errorf("error getting breeds: %w", err)
	}
	return bb, nil
}

func (s *BreedStore) CreateBreed(u *petsy.Breed) error {
	if err := s.Get(u, `INSERT INTO breeds VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		u.Name,
		u.BType,
		u.AvgWeight,
		u.ActivityLevel,
		u.CoatType,
		u.Shedding); err != nil {
		return fmt.Errorf("error creating breed: %w", err)
	}
	return nil
}

func (s *BreedStore) UpdateBreed(u *petsy.Breed) error {
	if err := s.Get(u, `UPDATE breeds SET btype = $1, avgweight = $2, activity = $3, coat = $4, shedding = $5 WHERE name = $6 RETURNING *`,
		u.BType,
		u.AvgWeight,
		u.ActivityLevel,
		u.CoatType,
		u.Shedding,
		u.Name); err != nil {
		return fmt.Errorf("error updating breed: %w", err)
	}
	return nil
}

func (s *BreedStore) DeleteBreed(name string) error {
	if _, err := s.Exec(`DELETE FROM breeds WHERE name = $1`, name); err != nil {
		return fmt.Errorf("error deleting breed: %w", err)
	}
	return nil
}