//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Order struct {
	OrderID     int32 `sql:"primary_key"`
	CustomerID  int32
	LineItems   *string
	CreatedAt   *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
}
