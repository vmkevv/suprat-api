// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "color_activities", Type: field.TypeInt, Nullable: true},
		{Name: "icon_activities", Type: field.TypeInt, Nullable: true},
		{Name: "user_activities", Type: field.TypeInt, Nullable: true},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "activities_colors_activities",
				Columns: []*schema.Column{ActivitiesColumns[3]},

				RefColumns: []*schema.Column{ColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_icons_activities",
				Columns: []*schema.Column{ActivitiesColumns[4]},

				RefColumns: []*schema.Column{IconsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "activities_users_activities",
				Columns: []*schema.Column{ActivitiesColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ColorsColumns holds the columns for the "colors" table.
	ColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "rgb", Type: field.TypeString},
	}
	// ColorsTable holds the schema information for the "colors" table.
	ColorsTable = &schema.Table{
		Name:        "colors",
		Columns:     ColorsColumns,
		PrimaryKey:  []*schema.Column{ColorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// IconsColumns holds the columns for the "icons" table.
	IconsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// IconsTable holds the schema information for the "icons" table.
	IconsTable = &schema.Table{
		Name:        "icons",
		Columns:     IconsColumns,
		PrimaryKey:  []*schema.Column{IconsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MeasuresColumns holds the columns for the "measures" table.
	MeasuresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "measurement_measures", Type: field.TypeInt, Nullable: true},
		{Name: "record_measures", Type: field.TypeInt, Nullable: true},
	}
	// MeasuresTable holds the schema information for the "measures" table.
	MeasuresTable = &schema.Table{
		Name:       "measures",
		Columns:    MeasuresColumns,
		PrimaryKey: []*schema.Column{MeasuresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "measures_measurements_measures",
				Columns: []*schema.Column{MeasuresColumns[2]},

				RefColumns: []*schema.Column{MeasurementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "measures_records_measures",
				Columns: []*schema.Column{MeasuresColumns[3]},

				RefColumns: []*schema.Column{RecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MeasurementsColumns holds the columns for the "measurements" table.
	MeasurementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "enabled", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "color_measurements", Type: field.TypeInt, Nullable: true},
		{Name: "user_measurements", Type: field.TypeInt, Nullable: true},
	}
	// MeasurementsTable holds the schema information for the "measurements" table.
	MeasurementsTable = &schema.Table{
		Name:       "measurements",
		Columns:    MeasurementsColumns,
		PrimaryKey: []*schema.Column{MeasurementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "measurements_colors_measurements",
				Columns: []*schema.Column{MeasurementsColumns[4]},

				RefColumns: []*schema.Column{ColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "measurements_users_measurements",
				Columns: []*schema.Column{MeasurementsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RecordsColumns holds the columns for the "records" table.
	RecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "start", Type: field.TypeTime},
		{Name: "end", Type: field.TypeTime},
		{Name: "activity_records", Type: field.TypeInt, Nullable: true},
	}
	// RecordsTable holds the schema information for the "records" table.
	RecordsTable = &schema.Table{
		Name:       "records",
		Columns:    RecordsColumns,
		PrimaryKey: []*schema.Column{RecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "records_activities_records",
				Columns: []*schema.Column{RecordsColumns[3]},

				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ActivityMeasurementsColumns holds the columns for the "activity_measurements" table.
	ActivityMeasurementsColumns = []*schema.Column{
		{Name: "activity_id", Type: field.TypeInt},
		{Name: "measurement_id", Type: field.TypeInt},
	}
	// ActivityMeasurementsTable holds the schema information for the "activity_measurements" table.
	ActivityMeasurementsTable = &schema.Table{
		Name:       "activity_measurements",
		Columns:    ActivityMeasurementsColumns,
		PrimaryKey: []*schema.Column{ActivityMeasurementsColumns[0], ActivityMeasurementsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "activity_measurements_activity_id",
				Columns: []*schema.Column{ActivityMeasurementsColumns[0]},

				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "activity_measurements_measurement_id",
				Columns: []*schema.Column{ActivityMeasurementsColumns[1]},

				RefColumns: []*schema.Column{MeasurementsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		ColorsTable,
		IconsTable,
		MeasuresTable,
		MeasurementsTable,
		RecordsTable,
		UsersTable,
		ActivityMeasurementsTable,
	}
)

func init() {
	ActivitiesTable.ForeignKeys[0].RefTable = ColorsTable
	ActivitiesTable.ForeignKeys[1].RefTable = IconsTable
	ActivitiesTable.ForeignKeys[2].RefTable = UsersTable
	MeasuresTable.ForeignKeys[0].RefTable = MeasurementsTable
	MeasuresTable.ForeignKeys[1].RefTable = RecordsTable
	MeasurementsTable.ForeignKeys[0].RefTable = ColorsTable
	MeasurementsTable.ForeignKeys[1].RefTable = UsersTable
	RecordsTable.ForeignKeys[0].RefTable = ActivitiesTable
	ActivityMeasurementsTable.ForeignKeys[0].RefTable = ActivitiesTable
	ActivityMeasurementsTable.ForeignKeys[1].RefTable = MeasurementsTable
}
