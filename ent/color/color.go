// Code generated by entc, DO NOT EDIT.

package color

const (
	// Label holds the string label denoting the color type in the database.
	Label = "color"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRgb holds the string denoting the rgb field in the database.
	FieldRgb = "rgb"

	// EdgeActivities holds the string denoting the activities edge name in mutations.
	EdgeActivities = "activities"
	// EdgeMeasurements holds the string denoting the measurements edge name in mutations.
	EdgeMeasurements = "measurements"

	// Table holds the table name of the color in the database.
	Table = "colors"
	// ActivitiesTable is the table the holds the activities relation/edge.
	ActivitiesTable = "activities"
	// ActivitiesInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivitiesInverseTable = "activities"
	// ActivitiesColumn is the table column denoting the activities relation/edge.
	ActivitiesColumn = "color_activities"
	// MeasurementsTable is the table the holds the measurements relation/edge.
	MeasurementsTable = "measurements"
	// MeasurementsInverseTable is the table name for the Measurement entity.
	// It exists in this package in order to avoid circular dependency with the "measurement" package.
	MeasurementsInverseTable = "measurements"
	// MeasurementsColumn is the table column denoting the measurements relation/edge.
	MeasurementsColumn = "color_measurements"
)

// Columns holds all SQL columns for color fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRgb,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
