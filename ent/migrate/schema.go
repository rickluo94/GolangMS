// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CitiesColumns holds the columns for the "cities" table.
	CitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 35},
		{Name: "country_code", Type: field.TypeString, Nullable: true, Size: 3},
		{Name: "district", Type: field.TypeString, Nullable: true, Size: 20},
		{Name: "population", Type: field.TypeInt, Nullable: true},
	}
	// CitiesTable holds the schema information for the "cities" table.
	CitiesTable = &schema.Table{
		Name:       "cities",
		Columns:    CitiesColumns,
		PrimaryKey: []*schema.Column{CitiesColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 40},
		{Name: "nickname", Type: field.TypeString, Size: 35, Default: ""},
		{Name: "password", Type: field.TypeString, Size: 50, Default: ""},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"on", "off"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0], UserColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CitiesTable,
		UserTable,
	}
)

func init() {
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
