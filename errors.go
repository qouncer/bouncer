package bouncer

type EntityNotFound struct{}

type EntityAlreadyExists struct{}

func (enf EntityNotFound) Error() string {
	return "Entity not found"
}

func (ee EntityAlreadyExists) Error() string {
	return "Entity exists"
}
