package _map

import "bouncer"

type Role struct {
	Name string
}

func (e *Role) GetName() string {
	return e.Name
}

func (e *Role) SetName(name string) {
	e.Name = name
}

func (e *Role) GetType() string {
	return bouncer.TypeRole
}

func (e *Role) String() string {
	return "<role: " + e.GetName() + ">"
}

type Permission struct {
	Name string
}

func (p *Permission) GetName() string {
	return p.Name
}

func (p *Permission) SetName(name string) {
	p.Name = name
}

func (p *Permission) GetType() string {
	return bouncer.TypePermission
}

func (p *Permission) String() string {
	return "<permission: " + p.GetName() + ">"
}

type Operation struct {
	Name string
}

func (o *Operation) GetName() string {
	return o.Name
}

func (o *Operation) SetName(name string) {
	o.Name = name
}

func (o *Operation) GetType() string {
	return bouncer.TypeOperation
}
