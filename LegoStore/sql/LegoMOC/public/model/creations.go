//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Creations struct {
	ID               int32 `sql:"primary_key"`
	Name             string
	CreatorID        *int32
	InstructionsLink string
	ImageID          *int32
	Price            float64
}
