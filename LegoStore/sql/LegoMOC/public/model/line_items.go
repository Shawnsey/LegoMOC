//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type LineItems struct {
	ID         int32 `sql:"primary_key"`
	CreationID *int32
	Quantity   *int32
	Price      *int32
}