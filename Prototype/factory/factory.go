package factory

import "factory-pattern/Prototype"

type Factory struct {
}

func (f Factory) CopyRobot(r *Prototype.Robot) *Prototype.Robot {
	return r.Clone().(*Prototype.Robot)
}
