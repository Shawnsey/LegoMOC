//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Pieces struct {
	ID          int32 `sql:"primary_key"`
	Description string
	ImageID     int32
}
