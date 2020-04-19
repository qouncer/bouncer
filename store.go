package bouncer

type Store interface {
	// Adds an entity to the store
	Add(e Entity) (Entity, error)
	// Adds child to parent entity
	AddChild(parent Entity, child Entity) error
	// Removes an entity to the store
	Remove(e Entity) (bool, error)
	// Finds an entity to the store
	Find(e Entity) (Entity, error)
}
