package _map

import (
	"bouncer"
)

type Store struct {
	Data map[string][]bouncer.Entity
}

func (s *Store) Add(e bouncer.Entity) (bouncer.Entity, error) {
	found, err := s.Find(e)
	if found != nil || err == nil {
		return nil, bouncer.EntityAlreadyExists{}
	}
	s.Data[e.GetType()] = append(s.Data[e.GetType()], e)
	return e, nil
}

func (s *Store) Find(e bouncer.Entity) (bouncer.Entity, error) {
	entities := s.Data[e.GetType()]
	for _, entity := range entities {
		if e.GetName() == entity.GetName() {
			return entity, nil
		}
	}
	return nil, bouncer.EntityNotFound{}
}

func (s *Store) Remove(e bouncer.Entity) (bool, error) {
	var updated []bouncer.Entity
	found := false
	entities := s.Data[e.GetType()]

	for k, v := range entities {
		if v.GetName() != e.GetName() {
			updated = append(updated, entities[k])
		} else {
			found = true
		}
	}

	if found {
		s.Data[e.GetType()] = updated
		return true, nil
	}

	return false, bouncer.EntityNotFound{}
}
