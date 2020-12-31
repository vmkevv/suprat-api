// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/vmkevv/suprat-api/ent/activity"
	"github.com/vmkevv/suprat-api/ent/color"
	"github.com/vmkevv/suprat-api/ent/icon"
	"github.com/vmkevv/suprat-api/ent/user"
)

// Activity is the model entity for the Activity schema.
type Activity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActivityQuery when eager-loading is set.
	Edges            ActivityEdges `json:"edges"`
	color_activities *int
	icon_activities  *int
	user_activities  *int
}

// ActivityEdges holds the relations/edges for other nodes in the graph.
type ActivityEdges struct {
	// User holds the value of the user edge.
	User *User
	// Color holds the value of the color edge.
	Color *Color
	// Icon holds the value of the icon edge.
	Icon *Icon
	// Records holds the value of the records edge.
	Records []*Record
	// Measurements holds the value of the measurements edge.
	Measurements []*Measurement
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ColorOrErr returns the Color value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) ColorOrErr() (*Color, error) {
	if e.loadedTypes[1] {
		if e.Color == nil {
			// The edge color was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: color.Label}
		}
		return e.Color, nil
	}
	return nil, &NotLoadedError{edge: "color"}
}

// IconOrErr returns the Icon value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActivityEdges) IconOrErr() (*Icon, error) {
	if e.loadedTypes[2] {
		if e.Icon == nil {
			// The edge icon was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: icon.Label}
		}
		return e.Icon, nil
	}
	return nil, &NotLoadedError{edge: "icon"}
}

// RecordsOrErr returns the Records value or an error if the edge
// was not loaded in eager-loading.
func (e ActivityEdges) RecordsOrErr() ([]*Record, error) {
	if e.loadedTypes[3] {
		return e.Records, nil
	}
	return nil, &NotLoadedError{edge: "records"}
}

// MeasurementsOrErr returns the Measurements value or an error if the edge
// was not loaded in eager-loading.
func (e ActivityEdges) MeasurementsOrErr() ([]*Measurement, error) {
	if e.loadedTypes[4] {
		return e.Measurements, nil
	}
	return nil, &NotLoadedError{edge: "measurements"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Activity) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullTime{},   // created_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Activity) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // color_activities
		&sql.NullInt64{}, // icon_activities
		&sql.NullInt64{}, // user_activities
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Activity fields.
func (a *Activity) assignValues(values ...interface{}) error {
	if m, n := len(values), len(activity.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		a.Name = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[1])
	} else if value.Valid {
		a.CreatedAt = value.Time
	}
	values = values[2:]
	if len(values) == len(activity.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field color_activities", value)
		} else if value.Valid {
			a.color_activities = new(int)
			*a.color_activities = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field icon_activities", value)
		} else if value.Valid {
			a.icon_activities = new(int)
			*a.icon_activities = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_activities", value)
		} else if value.Valid {
			a.user_activities = new(int)
			*a.user_activities = int(value.Int64)
		}
	}
	return nil
}

// QueryUser queries the user edge of the Activity.
func (a *Activity) QueryUser() *UserQuery {
	return (&ActivityClient{config: a.config}).QueryUser(a)
}

// QueryColor queries the color edge of the Activity.
func (a *Activity) QueryColor() *ColorQuery {
	return (&ActivityClient{config: a.config}).QueryColor(a)
}

// QueryIcon queries the icon edge of the Activity.
func (a *Activity) QueryIcon() *IconQuery {
	return (&ActivityClient{config: a.config}).QueryIcon(a)
}

// QueryRecords queries the records edge of the Activity.
func (a *Activity) QueryRecords() *RecordQuery {
	return (&ActivityClient{config: a.config}).QueryRecords(a)
}

// QueryMeasurements queries the measurements edge of the Activity.
func (a *Activity) QueryMeasurements() *MeasurementQuery {
	return (&ActivityClient{config: a.config}).QueryMeasurements(a)
}

// Update returns a builder for updating this Activity.
// Note that, you need to call Activity.Unwrap() before calling this method, if this Activity
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Activity) Update() *ActivityUpdateOne {
	return (&ActivityClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Activity) Unwrap() *Activity {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Activity is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Activity) String() string {
	var builder strings.Builder
	builder.WriteString("Activity(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Activities is a parsable slice of Activity.
type Activities []*Activity

func (a Activities) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
