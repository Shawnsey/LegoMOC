//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Orders struct {
	OrderID     uuid.UUID `sql:"primary_key"`
	CustomerID  *uuid.UUID
	LineItems   *string
	CreatedAt   *time.Time
	ShippedAt   *time.Time
	CompletedAt *time.Time
}