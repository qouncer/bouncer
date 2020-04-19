package _map

import (
	"bouncer"
	"testing"
)

func TestStore_Find(t *testing.T) {
	roles := []bouncer.Entity{&Role{"tester"}}
	data := make(map[string][]bouncer.Entity)
	data[bouncer.TypeRole] = roles

	var store bouncer.Store = &Store{data}
	e, err := store.Find(&Role{"tester"})

	if e == nil || err != nil {
		t.Fail()
	}

	e, err = store.Find(&Role{"failed-tester"})
	if e != nil || err == nil {
		t.Fail()
	}
}

func TestStore_Add(t *testing.T) {
	roles := []bouncer.Entity{&Role{"tester"}}
	data := make(map[string][]bouncer.Entity)
	data[bouncer.TypeRole] = roles

	store := &Store{data}

	role, _ := store.Add(&Role{"another-tester"})

	if role == nil {
		t.Fail()
	}
	if len(roles) != 1 {
		t.Fail()
	}
}

func TestStore_AddAnotherEntityType(t *testing.T) {
	data := make(map[string][]bouncer.Entity)
	store := &Store{data}

	_, _ = store.Add(&Role{"tester"})
	_, _ = store.Add(&Permission{"tester"})

	if len(data[bouncer.TypeRole]) != 1 || len(data[bouncer.TypePermission]) != 1 {
		t.Fail()
	}
}
func TestStore_Remove(t *testing.T) {
	roles := []bouncer.Entity{
		&Permission{"tester"},
		&Role{"will-be-removed"},
		&Role{"another-tester"},
	}

	data := make(map[string][]bouncer.Entity)
	data[bouncer.TypeRole] = roles
	store := &Store{data}

	result, err := store.Remove(&Role{"will-be-removed"})

	if result == false || err != nil || len(store.Data[bouncer.TypeRole]) != 2 {
		t.Fail()
	}

	roles = []bouncer.Entity{
		&Role{"tester"},
		&Role{"will-be-removed"},
	}
	data = make(map[string][]bouncer.Entity)
	data[bouncer.TypeRole] = roles
	store = &Store{data}
	result, err = store.Remove(&Role{"will-be-removed"})

	if result == false || err != nil || len(store.Data[bouncer.TypeRole]) != 1 {
		t.Fail()
	}

	roles = []bouncer.Entity{
		&Role{"will-be-removed"},
		&Role{"tester"},
	}
	data = make(map[string][]bouncer.Entity)
	data[bouncer.TypeRole] = roles
	store = &Store{data}
	result, err = store.Remove(&Role{"will-be-removed"})

	if result == false || err != nil || len(store.Data[bouncer.TypeRole]) != 1 {
		t.Fail()
	}
}
