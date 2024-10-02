package svcuser_entuser_typ

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/teejays/gokutil/dalutil"
	filterlib "github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/naam"
	"github.com/teejays/gokutil/panics"
	"github.com/teejays/gokutil/scalars"

	app_typ "sampleapp/backend/.goku/generated/typ"
)

// User: <comments>
type User struct {
	ID          scalars.ID                 `json:"id"`
	Name        app_typ.PersonNameWithMeta `json:"name" validate:"required"`
	Email       scalars.Email              `json:"email" validate:"required"`
	Addresses   []app_typ.AddressWithMeta  `json:"addresses"`
	TeamID      scalars.ID                 `json:"teamID" validate:"required"`
	PastTeamIDs []scalars.ID               `json:"pastTeamIDs"`
	CreatedAt   scalars.Timestamp          `json:"createdAt"`
	UpdatedAt   scalars.Timestamp          `json:"updatedAt"`
	DeletedAt   *scalars.Timestamp         `json:"deletedAt"`
}

func (t User) GetID() scalars.ID {
	return t.ID
}
func (t User) GetUpdatedAt() scalars.Timestamp {
	return t.UpdatedAt
}
func (t User) SetUpdatedAt(tim scalars.Timestamp) {
	t.UpdatedAt = tim
}

// UserInput: <comments>
type UserInput struct {
	Name        app_typ.PersonNameInput `json:"name" validate:"required"`
	Email       scalars.Email           `json:"email" validate:"required"`
	Addresses   []app_typ.AddressInput  `json:"addresses"`
	TeamID      scalars.ID              `json:"teamID" validate:"required"`
	PastTeamIDs []scalars.ID            `json:"pastTeamIDs"`
}

func NewUserWithMetaArrayFromInputs(ctx context.Context, ins []UserInput) []User {
	var outs []User
	for _, in := range ins {
		outs = append(outs, NewUserWithMetaFromInput(ctx, in))
	}
	return outs
}

func NewUserWithMetaFromInput(ctx context.Context, in UserInput) User {
	return User{
		Name:        app_typ.NewPersonNameWithMetaFromInput(ctx, in.Name),
		Email:       in.Email,
		Addresses:   app_typ.NewAddressWithMetaArrayFromInputs(ctx, in.Addresses),
		TeamID:      in.TeamID,
		PastTeamIDs: in.PastTeamIDs,
	}
}

// UserField enum: <insert comment>
type UserField int

const (
	UserField_INVALID              UserField = 0
	UserField_ID                   UserField = 1
	UserField_Name                 UserField = 2
	UserField_Name_ParentID        UserField = 3
	UserField_Name_ID              UserField = 4
	UserField_Name_FirstName       UserField = 5
	UserField_Name_MiddleInitial   UserField = 6
	UserField_Name_LastName        UserField = 7
	UserField_Name_CreatedAt       UserField = 8
	UserField_Name_UpdatedAt       UserField = 9
	UserField_Name_DeletedAt       UserField = 10
	UserField_Email                UserField = 11
	UserField_Addresses            UserField = 12
	UserField_Addresses_ParentID   UserField = 13
	UserField_Addresses_ID         UserField = 14
	UserField_Addresses_Line1      UserField = 15
	UserField_Addresses_Line2      UserField = 16
	UserField_Addresses_City       UserField = 17
	UserField_Addresses_State      UserField = 18
	UserField_Addresses_PostalCode UserField = 19
	UserField_Addresses_Country    UserField = 20
	UserField_Addresses_CreatedAt  UserField = 21
	UserField_Addresses_UpdatedAt  UserField = 22
	UserField_Addresses_DeletedAt  UserField = 23
	UserField_TeamID               UserField = 24
	UserField_PastTeamIDs          UserField = 25
	UserField_CreatedAt            UserField = 26
	UserField_UpdatedAt            UserField = 27
	UserField_DeletedAt            UserField = 28
)

func NewUserFieldFromString(s string) UserField {
	switch s {
	case "INVALID":
		return UserField_INVALID
	case "ID":
		return UserField_ID
	case "Name":
		return UserField_Name
	case "Name_ParentID":
		return UserField_Name_ParentID
	case "Name_ID":
		return UserField_Name_ID
	case "Name_FirstName":
		return UserField_Name_FirstName
	case "Name_MiddleInitial":
		return UserField_Name_MiddleInitial
	case "Name_LastName":
		return UserField_Name_LastName
	case "Name_CreatedAt":
		return UserField_Name_CreatedAt
	case "Name_UpdatedAt":
		return UserField_Name_UpdatedAt
	case "Name_DeletedAt":
		return UserField_Name_DeletedAt
	case "Email":
		return UserField_Email
	case "Addresses":
		return UserField_Addresses
	case "Addresses_ParentID":
		return UserField_Addresses_ParentID
	case "Addresses_ID":
		return UserField_Addresses_ID
	case "Addresses_Line1":
		return UserField_Addresses_Line1
	case "Addresses_Line2":
		return UserField_Addresses_Line2
	case "Addresses_City":
		return UserField_Addresses_City
	case "Addresses_State":
		return UserField_Addresses_State
	case "Addresses_PostalCode":
		return UserField_Addresses_PostalCode
	case "Addresses_Country":
		return UserField_Addresses_Country
	case "Addresses_CreatedAt":
		return UserField_Addresses_CreatedAt
	case "Addresses_UpdatedAt":
		return UserField_Addresses_UpdatedAt
	case "Addresses_DeletedAt":
		return UserField_Addresses_DeletedAt
	case "TeamID":
		return UserField_TeamID
	case "PastTeamIDs":
		return UserField_PastTeamIDs
	case "CreatedAt":
		return UserField_CreatedAt
	case "UpdatedAt":
		return UserField_UpdatedAt
	case "DeletedAt":
		return UserField_DeletedAt

	default:
		panic(fmt.Sprintf("'%s' is not a valid value for type '%s'", s, "UserField"))
	}
}

// String implements the `fmt.Stringer` interface for UserField. It allows us to print the enum values as strings.
func (f UserField) String() string {
	switch f {
	case UserField_INVALID:
		return "INVALID"
	case UserField_ID:
		return "ID"
	case UserField_Name:
		return "Name"
	case UserField_Name_ParentID:
		return "Name_ParentID"
	case UserField_Name_ID:
		return "Name_ID"
	case UserField_Name_FirstName:
		return "Name_FirstName"
	case UserField_Name_MiddleInitial:
		return "Name_MiddleInitial"
	case UserField_Name_LastName:
		return "Name_LastName"
	case UserField_Name_CreatedAt:
		return "Name_CreatedAt"
	case UserField_Name_UpdatedAt:
		return "Name_UpdatedAt"
	case UserField_Name_DeletedAt:
		return "Name_DeletedAt"
	case UserField_Email:
		return "Email"
	case UserField_Addresses:
		return "Addresses"
	case UserField_Addresses_ParentID:
		return "Addresses_ParentID"
	case UserField_Addresses_ID:
		return "Addresses_ID"
	case UserField_Addresses_Line1:
		return "Addresses_Line1"
	case UserField_Addresses_Line2:
		return "Addresses_Line2"
	case UserField_Addresses_City:
		return "Addresses_City"
	case UserField_Addresses_State:
		return "Addresses_State"
	case UserField_Addresses_PostalCode:
		return "Addresses_PostalCode"
	case UserField_Addresses_Country:
		return "Addresses_Country"
	case UserField_Addresses_CreatedAt:
		return "Addresses_CreatedAt"
	case UserField_Addresses_UpdatedAt:
		return "Addresses_UpdatedAt"
	case UserField_Addresses_DeletedAt:
		return "Addresses_DeletedAt"
	case UserField_TeamID:
		return "TeamID"
	case UserField_PastTeamIDs:
		return "PastTeamIDs"
	case UserField_CreatedAt:
		return "CreatedAt"
	case UserField_UpdatedAt:
		return "UpdatedAt"
	case UserField_DeletedAt:
		return "DeletedAt"

	default:
		panic(fmt.Sprintf("'%d' is not a valid type '%s'", f, "UserField"))
	}
}

// Name gives a naam representation of the enum value
func (f UserField) Name() naam.Name {
	switch f {
	case UserField_ID:
		return naam.New("id")
	case UserField_Name:
		return naam.New("name")
	case UserField_Name_ParentID:
		return naam.New("name___parent_id")
	case UserField_Name_ID:
		return naam.New("name___id")
	case UserField_Name_FirstName:
		return naam.New("name___first_name")
	case UserField_Name_MiddleInitial:
		return naam.New("name___middle_initial")
	case UserField_Name_LastName:
		return naam.New("name___last_name")
	case UserField_Name_CreatedAt:
		return naam.New("name___created_at")
	case UserField_Name_UpdatedAt:
		return naam.New("name___updated_at")
	case UserField_Name_DeletedAt:
		return naam.New("name___deleted_at")
	case UserField_Email:
		return naam.New("email")
	case UserField_Addresses:
		return naam.New("addresses")
	case UserField_Addresses_ParentID:
		return naam.New("addresses___parent_id")
	case UserField_Addresses_ID:
		return naam.New("addresses___id")
	case UserField_Addresses_Line1:
		return naam.New("addresses___line_1")
	case UserField_Addresses_Line2:
		return naam.New("addresses___line_2")
	case UserField_Addresses_City:
		return naam.New("addresses___city")
	case UserField_Addresses_State:
		return naam.New("addresses___state")
	case UserField_Addresses_PostalCode:
		return naam.New("addresses___postal_code")
	case UserField_Addresses_Country:
		return naam.New("addresses___country")
	case UserField_Addresses_CreatedAt:
		return naam.New("addresses___created_at")
	case UserField_Addresses_UpdatedAt:
		return naam.New("addresses___updated_at")
	case UserField_Addresses_DeletedAt:
		return naam.New("addresses___deleted_at")
	case UserField_TeamID:
		return naam.New("team_id")
	case UserField_PastTeamIDs:
		return naam.New("past_team_ids")
	case UserField_CreatedAt:
		return naam.New("created_at")
	case UserField_UpdatedAt:
		return naam.New("updated_at")
	case UserField_DeletedAt:
		return naam.New("deleted_at")
	default:
		panics.P("UserField.Name(): Unrecognized field (%d)", f)
	}
	return naam.Nil()
}

// Value implements them the `drive.Valuer` interface for this enum. It allows us to save these enum values to the DB as a string.
func (f UserField) Value() (driver.Value, error) {
	switch f {
	case UserField_INVALID:
		return nil, nil
	case UserField_ID:
		return "ID", nil
	case UserField_Name:
		return "Name", nil
	case UserField_Name_ParentID:
		return "Name_ParentID", nil
	case UserField_Name_ID:
		return "Name_ID", nil
	case UserField_Name_FirstName:
		return "Name_FirstName", nil
	case UserField_Name_MiddleInitial:
		return "Name_MiddleInitial", nil
	case UserField_Name_LastName:
		return "Name_LastName", nil
	case UserField_Name_CreatedAt:
		return "Name_CreatedAt", nil
	case UserField_Name_UpdatedAt:
		return "Name_UpdatedAt", nil
	case UserField_Name_DeletedAt:
		return "Name_DeletedAt", nil
	case UserField_Email:
		return "Email", nil
	case UserField_Addresses:
		return "Addresses", nil
	case UserField_Addresses_ParentID:
		return "Addresses_ParentID", nil
	case UserField_Addresses_ID:
		return "Addresses_ID", nil
	case UserField_Addresses_Line1:
		return "Addresses_Line1", nil
	case UserField_Addresses_Line2:
		return "Addresses_Line2", nil
	case UserField_Addresses_City:
		return "Addresses_City", nil
	case UserField_Addresses_State:
		return "Addresses_State", nil
	case UserField_Addresses_PostalCode:
		return "Addresses_PostalCode", nil
	case UserField_Addresses_Country:
		return "Addresses_Country", nil
	case UserField_Addresses_CreatedAt:
		return "Addresses_CreatedAt", nil
	case UserField_Addresses_UpdatedAt:
		return "Addresses_UpdatedAt", nil
	case UserField_Addresses_DeletedAt:
		return "Addresses_DeletedAt", nil
	case UserField_TeamID:
		return "TeamID", nil
	case UserField_PastTeamIDs:
		return "PastTeamIDs", nil
	case UserField_CreatedAt:
		return "CreatedAt", nil
	case UserField_UpdatedAt:
		return "UpdatedAt", nil
	case UserField_DeletedAt:
		return "DeletedAt", nil

	default:
		return nil, fmt.Errorf("Cannot save enum UserField to DB: '%d' is not a valid value for enum UserField", f)
	}
}

// Scan implements them the `sql.Scanner` interface for this enum. It allows us to read these enum values from the DB,
// which are stored a string.
func (f *UserField) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		i := NewUserFieldFromString(v)
		*f = i
	default:
		return fmt.Errorf("Attempted to read data of type %T into enum %s from SQL", v, "UserField")
	}
	return nil
}

// ImplementsGraphQLType maps this custom Go type to the graphql scalar type in the schema.
func (f UserField) ImplementsGraphQLType(name string) bool {
	return name == "User_User_UserField"
}

func (f *UserField) UnmarshalGraphQL(input interface{}) error {
	var err error
	switch input := input.(type) {
	case string:
		i := NewUserFieldFromString(input)
		*f = i
	default:
		err = fmt.Errorf("wrong type for UserField: %T", input)
	}
	return err
}

func (f *UserField) UnmarshalJSON(data []byte) error {
	var enumStr string
	err := json.Unmarshal(data, &enumStr)
	if err != nil {
		return fmt.Errorf("cannot Unmarshal enum UserField to a string: %w", err)
	}
	i := NewUserFieldFromString(enumStr)
	*f = i
	return nil
}

func (f UserField) MarshalJSON() ([]byte, error) {
	panics.IfNil(f, "attempted to marshal nil UserField pointer to JSON")
	enumStr := f.String()

	data, err := json.Marshal(enumStr)
	if err != nil {
		return nil, fmt.Errorf("cannot Marshal enum \"%s\" into JSON: %w", enumStr, err)
	}
	return data, nil
}

type UserFieldCondition struct {
	Op     filterlib.Operator
	Values []UserField
}

func (c UserFieldCondition) GetOperator() filterlib.Operator {
	return c.Op
}
func (c UserFieldCondition) Len() int {
	return len(c.Values)
}
func (c UserFieldCondition) GetValue(i int) interface{} {
	return c.Values[i]
}

// UserFilter: <comments>
type UserFilter struct {
	ID                *filterlib.IDCondition        `json:"id"`
	Name              *app_typ.PersonNameFilter     `json:"name"`
	Email             *filterlib.EmailCondition     `json:"email"`
	HavingAddresses   *app_typ.AddressFilter        `json:"havingAddresses"`
	TeamID            *filterlib.IDCondition        `json:"teamID"`
	HavingPastTeamIDs *filterlib.IDCondition        `json:"havingPastTeamIDs"`
	CreatedAt         *filterlib.TimestampCondition `json:"createdAt"`
	UpdatedAt         *filterlib.TimestampCondition `json:"updatedAt"`
	DeletedAt         *filterlib.TimestampCondition `json:"deletedAt"`
	And               []UserFilter                  `json:"and"`
	Or                []UserFilter                  `json:"or"`
}

// StandardEntityRequest: <comments>
type StandardEntityRequest struct {
	ID scalars.ID `json:"id"`
}

// StandardEntityResponse: <comments>
type StandardEntityResponse struct {
	Object User `json:"object"`
}

// StandardEntityRequest:
// StandardEntityResponse:
// User:
// UserFieldCondition:
// UserFilter:
// AddRequest:
type AddRequest = dalutil.EntityAddRequest[UserInput]

// UpdateRequest:
type UpdateRequest = dalutil.UpdateEntityRequest[User, UserField]

// UpdateResponse:
type UpdateResponse = dalutil.UpdateEntityResponse[User]

// GetRequest:
type GetRequest = dalutil.GetEntityRequest[User]

// ListRequest:
type ListRequest = dalutil.ListEntityRequest[UserFilter]

// ListResponse:
type ListResponse = dalutil.ListEntityResponse[User]

// QueryByTextRequest:
type QueryByTextRequest = dalutil.QueryByTextEntityRequest[User]
