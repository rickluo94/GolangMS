// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/city"
	"awesomeProject/ent/predicate"
	"awesomeProject/ent/user"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCity = "City"
	TypeUser = "User"
)

// CityMutation represents an operation that mutates the City nodes in the graph.
type CityMutation struct {
	config
	op             Op
	typ            string
	id             *int
	_Name          *string
	_CountryCode   *string
	_District      *string
	_Population    *int
	add_Population *int
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*City, error)
	predicates     []predicate.City
}

var _ ent.Mutation = (*CityMutation)(nil)

// cityOption allows management of the mutation configuration using functional options.
type cityOption func(*CityMutation)

// newCityMutation creates new mutation for the City entity.
func newCityMutation(c config, op Op, opts ...cityOption) *CityMutation {
	m := &CityMutation{
		config:        c,
		op:            op,
		typ:           TypeCity,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCityID sets the ID field of the mutation.
func withCityID(id int) cityOption {
	return func(m *CityMutation) {
		var (
			err   error
			once  sync.Once
			value *City
		)
		m.oldValue = func(ctx context.Context) (*City, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().City.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCity sets the old City of the mutation.
func withCity(node *City) cityOption {
	return func(m *CityMutation) {
		m.oldValue = func(context.Context) (*City, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CityMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CityMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CityMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CityMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().City.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "Name" field.
func (m *CityMutation) SetName(s string) {
	m._Name = &s
}

// Name returns the value of the "Name" field in the mutation.
func (m *CityMutation) Name() (r string, exists bool) {
	v := m._Name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "Name" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "Name" field.
func (m *CityMutation) ClearName() {
	m._Name = nil
	m.clearedFields[city.FieldName] = struct{}{}
}

// NameCleared returns if the "Name" field was cleared in this mutation.
func (m *CityMutation) NameCleared() bool {
	_, ok := m.clearedFields[city.FieldName]
	return ok
}

// ResetName resets all changes to the "Name" field.
func (m *CityMutation) ResetName() {
	m._Name = nil
	delete(m.clearedFields, city.FieldName)
}

// SetCountryCode sets the "CountryCode" field.
func (m *CityMutation) SetCountryCode(s string) {
	m._CountryCode = &s
}

// CountryCode returns the value of the "CountryCode" field in the mutation.
func (m *CityMutation) CountryCode() (r string, exists bool) {
	v := m._CountryCode
	if v == nil {
		return
	}
	return *v, true
}

// OldCountryCode returns the old "CountryCode" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldCountryCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCountryCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCountryCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCountryCode: %w", err)
	}
	return oldValue.CountryCode, nil
}

// ClearCountryCode clears the value of the "CountryCode" field.
func (m *CityMutation) ClearCountryCode() {
	m._CountryCode = nil
	m.clearedFields[city.FieldCountryCode] = struct{}{}
}

// CountryCodeCleared returns if the "CountryCode" field was cleared in this mutation.
func (m *CityMutation) CountryCodeCleared() bool {
	_, ok := m.clearedFields[city.FieldCountryCode]
	return ok
}

// ResetCountryCode resets all changes to the "CountryCode" field.
func (m *CityMutation) ResetCountryCode() {
	m._CountryCode = nil
	delete(m.clearedFields, city.FieldCountryCode)
}

// SetDistrict sets the "District" field.
func (m *CityMutation) SetDistrict(s string) {
	m._District = &s
}

// District returns the value of the "District" field in the mutation.
func (m *CityMutation) District() (r string, exists bool) {
	v := m._District
	if v == nil {
		return
	}
	return *v, true
}

// OldDistrict returns the old "District" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldDistrict(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDistrict is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDistrict requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDistrict: %w", err)
	}
	return oldValue.District, nil
}

// ClearDistrict clears the value of the "District" field.
func (m *CityMutation) ClearDistrict() {
	m._District = nil
	m.clearedFields[city.FieldDistrict] = struct{}{}
}

// DistrictCleared returns if the "District" field was cleared in this mutation.
func (m *CityMutation) DistrictCleared() bool {
	_, ok := m.clearedFields[city.FieldDistrict]
	return ok
}

// ResetDistrict resets all changes to the "District" field.
func (m *CityMutation) ResetDistrict() {
	m._District = nil
	delete(m.clearedFields, city.FieldDistrict)
}

// SetPopulation sets the "Population" field.
func (m *CityMutation) SetPopulation(i int) {
	m._Population = &i
	m.add_Population = nil
}

// Population returns the value of the "Population" field in the mutation.
func (m *CityMutation) Population() (r int, exists bool) {
	v := m._Population
	if v == nil {
		return
	}
	return *v, true
}

// OldPopulation returns the old "Population" field's value of the City entity.
// If the City object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CityMutation) OldPopulation(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPopulation is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPopulation requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPopulation: %w", err)
	}
	return oldValue.Population, nil
}

// AddPopulation adds i to the "Population" field.
func (m *CityMutation) AddPopulation(i int) {
	if m.add_Population != nil {
		*m.add_Population += i
	} else {
		m.add_Population = &i
	}
}

// AddedPopulation returns the value that was added to the "Population" field in this mutation.
func (m *CityMutation) AddedPopulation() (r int, exists bool) {
	v := m.add_Population
	if v == nil {
		return
	}
	return *v, true
}

// ClearPopulation clears the value of the "Population" field.
func (m *CityMutation) ClearPopulation() {
	m._Population = nil
	m.add_Population = nil
	m.clearedFields[city.FieldPopulation] = struct{}{}
}

// PopulationCleared returns if the "Population" field was cleared in this mutation.
func (m *CityMutation) PopulationCleared() bool {
	_, ok := m.clearedFields[city.FieldPopulation]
	return ok
}

// ResetPopulation resets all changes to the "Population" field.
func (m *CityMutation) ResetPopulation() {
	m._Population = nil
	m.add_Population = nil
	delete(m.clearedFields, city.FieldPopulation)
}

// Where appends a list predicates to the CityMutation builder.
func (m *CityMutation) Where(ps ...predicate.City) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CityMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (City).
func (m *CityMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CityMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m._Name != nil {
		fields = append(fields, city.FieldName)
	}
	if m._CountryCode != nil {
		fields = append(fields, city.FieldCountryCode)
	}
	if m._District != nil {
		fields = append(fields, city.FieldDistrict)
	}
	if m._Population != nil {
		fields = append(fields, city.FieldPopulation)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CityMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case city.FieldName:
		return m.Name()
	case city.FieldCountryCode:
		return m.CountryCode()
	case city.FieldDistrict:
		return m.District()
	case city.FieldPopulation:
		return m.Population()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CityMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case city.FieldName:
		return m.OldName(ctx)
	case city.FieldCountryCode:
		return m.OldCountryCode(ctx)
	case city.FieldDistrict:
		return m.OldDistrict(ctx)
	case city.FieldPopulation:
		return m.OldPopulation(ctx)
	}
	return nil, fmt.Errorf("unknown City field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) SetField(name string, value ent.Value) error {
	switch name {
	case city.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case city.FieldCountryCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCountryCode(v)
		return nil
	case city.FieldDistrict:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDistrict(v)
		return nil
	case city.FieldPopulation:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPopulation(v)
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CityMutation) AddedFields() []string {
	var fields []string
	if m.add_Population != nil {
		fields = append(fields, city.FieldPopulation)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CityMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case city.FieldPopulation:
		return m.AddedPopulation()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CityMutation) AddField(name string, value ent.Value) error {
	switch name {
	case city.FieldPopulation:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPopulation(v)
		return nil
	}
	return fmt.Errorf("unknown City numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CityMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(city.FieldName) {
		fields = append(fields, city.FieldName)
	}
	if m.FieldCleared(city.FieldCountryCode) {
		fields = append(fields, city.FieldCountryCode)
	}
	if m.FieldCleared(city.FieldDistrict) {
		fields = append(fields, city.FieldDistrict)
	}
	if m.FieldCleared(city.FieldPopulation) {
		fields = append(fields, city.FieldPopulation)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CityMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CityMutation) ClearField(name string) error {
	switch name {
	case city.FieldName:
		m.ClearName()
		return nil
	case city.FieldCountryCode:
		m.ClearCountryCode()
		return nil
	case city.FieldDistrict:
		m.ClearDistrict()
		return nil
	case city.FieldPopulation:
		m.ClearPopulation()
		return nil
	}
	return fmt.Errorf("unknown City nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CityMutation) ResetField(name string) error {
	switch name {
	case city.FieldName:
		m.ResetName()
		return nil
	case city.FieldCountryCode:
		m.ResetCountryCode()
		return nil
	case city.FieldDistrict:
		m.ResetDistrict()
		return nil
	case city.FieldPopulation:
		m.ResetPopulation()
		return nil
	}
	return fmt.Errorf("unknown City field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CityMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CityMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CityMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CityMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CityMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CityMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CityMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown City unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CityMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown City edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	uuid          *uuid.UUID
	username      *string
	nickname      *string
	password      *string
	active        *bool
	state         *user.State
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uint64) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *UserMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *UserMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *UserMutation) ResetUUID() {
	m.uuid = nil
}

// SetUsername sets the "username" field.
func (m *UserMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *UserMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *UserMutation) ResetUsername() {
	m.username = nil
}

// SetNickname sets the "nickname" field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ResetNickname resets all changes to the "nickname" field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetActive sets the "active" field.
func (m *UserMutation) SetActive(b bool) {
	m.active = &b
}

// Active returns the value of the "active" field in the mutation.
func (m *UserMutation) Active() (r bool, exists bool) {
	v := m.active
	if v == nil {
		return
	}
	return *v, true
}

// OldActive returns the old "active" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldActive(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActive is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActive requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActive: %w", err)
	}
	return oldValue.Active, nil
}

// ResetActive resets all changes to the "active" field.
func (m *UserMutation) ResetActive() {
	m.active = nil
}

// SetState sets the "state" field.
func (m *UserMutation) SetState(u user.State) {
	m.state = &u
}

// State returns the value of the "state" field in the mutation.
func (m *UserMutation) State() (r user.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldState(ctx context.Context) (v user.State, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ClearState clears the value of the "state" field.
func (m *UserMutation) ClearState() {
	m.state = nil
	m.clearedFields[user.FieldState] = struct{}{}
}

// StateCleared returns if the "state" field was cleared in this mutation.
func (m *UserMutation) StateCleared() bool {
	_, ok := m.clearedFields[user.FieldState]
	return ok
}

// ResetState resets all changes to the "state" field.
func (m *UserMutation) ResetState() {
	m.state = nil
	delete(m.clearedFields, user.FieldState)
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.uuid != nil {
		fields = append(fields, user.FieldUUID)
	}
	if m.username != nil {
		fields = append(fields, user.FieldUsername)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.active != nil {
		fields = append(fields, user.FieldActive)
	}
	if m.state != nil {
		fields = append(fields, user.FieldState)
	}
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldUUID:
		return m.UUID()
	case user.FieldUsername:
		return m.Username()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldPassword:
		return m.Password()
	case user.FieldActive:
		return m.Active()
	case user.FieldState:
		return m.State()
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldUUID:
		return m.OldUUID(ctx)
	case user.FieldUsername:
		return m.OldUsername(ctx)
	case user.FieldNickname:
		return m.OldNickname(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldActive:
		return m.OldActive(ctx)
	case user.FieldState:
		return m.OldState(ctx)
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case user.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldActive:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActive(v)
		return nil
	case user.FieldState:
		v, ok := value.(user.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldState) {
		fields = append(fields, user.FieldState)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldState:
		m.ClearState()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldUUID:
		m.ResetUUID()
		return nil
	case user.FieldUsername:
		m.ResetUsername()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldActive:
		m.ResetActive()
		return nil
	case user.FieldState:
		m.ResetState()
		return nil
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
