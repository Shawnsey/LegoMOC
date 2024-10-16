package model

type Model interface {
	getID() int32
}

func (c *Creations) getID() int32 {
	return c.ID
}

func (o *Order) getID() int32 {
	return o.OrderID
}

func (p *Pieces) getID() int32 {
	return p.ID
}

func (i *Images) getID() int32 {
	return i.ID
}

func (u *Users) getID() int32 {
	return u.ID
}