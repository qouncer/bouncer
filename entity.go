package bouncer

const (
	TypeRole       = "role"
	TypePermission = "permission"
	TypeOperation  = "operation"
)

type Entity interface {
	GetName() string
	SetName(name string)
	GetType() string
}

func CanAddChild(parent Entity, child Entity) bool {
	if parent.GetType() == TypeRole && child.GetType() == TypeRole {
		return true
	}

	if parent.GetType() == TypeRole && child.GetType() == TypePermission {
		return true
	}

	if parent.GetType() == TypePermission && child.GetType() == TypeOperation {
		return true
	}

	return false
}
