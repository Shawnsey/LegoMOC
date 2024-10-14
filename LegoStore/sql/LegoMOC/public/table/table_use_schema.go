//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	Creations = Creations.FromSchema(schema)
	Images = Images.FromSchema(schema)
	LineItems = LineItems.FromSchema(schema)
	Order = Order.FromSchema(schema)
	PaymentInfo = PaymentInfo.FromSchema(schema)
	Pieces = Pieces.FromSchema(schema)
	Users = Users.FromSchema(schema)
}
