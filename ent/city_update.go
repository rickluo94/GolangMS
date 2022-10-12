// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/city"
	"awesomeProject/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CityUpdate is the builder for updating City entities.
type CityUpdate struct {
	config
	hooks    []Hook
	mutation *CityMutation
}

// Where appends a list predicates to the CityUpdate builder.
func (cu *CityUpdate) Where(ps ...predicate.City) *CityUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "Name" field.
func (cu *CityUpdate) SetName(s string) *CityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cu *CityUpdate) SetNillableName(s *string) *CityUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "Name" field.
func (cu *CityUpdate) ClearName() *CityUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetCountryCode sets the "CountryCode" field.
func (cu *CityUpdate) SetCountryCode(s string) *CityUpdate {
	cu.mutation.SetCountryCode(s)
	return cu
}

// SetNillableCountryCode sets the "CountryCode" field if the given value is not nil.
func (cu *CityUpdate) SetNillableCountryCode(s *string) *CityUpdate {
	if s != nil {
		cu.SetCountryCode(*s)
	}
	return cu
}

// ClearCountryCode clears the value of the "CountryCode" field.
func (cu *CityUpdate) ClearCountryCode() *CityUpdate {
	cu.mutation.ClearCountryCode()
	return cu
}

// SetDistrict sets the "District" field.
func (cu *CityUpdate) SetDistrict(s string) *CityUpdate {
	cu.mutation.SetDistrict(s)
	return cu
}

// SetNillableDistrict sets the "District" field if the given value is not nil.
func (cu *CityUpdate) SetNillableDistrict(s *string) *CityUpdate {
	if s != nil {
		cu.SetDistrict(*s)
	}
	return cu
}

// ClearDistrict clears the value of the "District" field.
func (cu *CityUpdate) ClearDistrict() *CityUpdate {
	cu.mutation.ClearDistrict()
	return cu
}

// SetPopulation sets the "Population" field.
func (cu *CityUpdate) SetPopulation(i int) *CityUpdate {
	cu.mutation.ResetPopulation()
	cu.mutation.SetPopulation(i)
	return cu
}

// SetNillablePopulation sets the "Population" field if the given value is not nil.
func (cu *CityUpdate) SetNillablePopulation(i *int) *CityUpdate {
	if i != nil {
		cu.SetPopulation(*i)
	}
	return cu
}

// AddPopulation adds i to the "Population" field.
func (cu *CityUpdate) AddPopulation(i int) *CityUpdate {
	cu.mutation.AddPopulation(i)
	return cu
}

// ClearPopulation clears the value of the "Population" field.
func (cu *CityUpdate) ClearPopulation() *CityUpdate {
	cu.mutation.ClearPopulation()
	return cu
}

// Mutation returns the CityMutation object of the builder.
func (cu *CityUpdate) Mutation() *CityMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CityUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := city.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "City.Name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.CountryCode(); ok {
		if err := city.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "CountryCode", err: fmt.Errorf(`ent: validator failed for field "City.CountryCode": %w`, err)}
		}
	}
	if v, ok := cu.mutation.District(); ok {
		if err := city.DistrictValidator(v); err != nil {
			return &ValidationError{Name: "District", err: fmt.Errorf(`ent: validator failed for field "City.District": %w`, err)}
		}
	}
	return nil
}

func (cu *CityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   city.Table,
			Columns: city.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
	}
	if cu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldName,
		})
	}
	if value, ok := cu.mutation.CountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldCountryCode,
		})
	}
	if cu.mutation.CountryCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldCountryCode,
		})
	}
	if value, ok := cu.mutation.District(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldDistrict,
		})
	}
	if cu.mutation.DistrictCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldDistrict,
		})
	}
	if value, ok := cu.mutation.Population(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldPopulation,
		})
	}
	if value, ok := cu.mutation.AddedPopulation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldPopulation,
		})
	}
	if cu.mutation.PopulationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: city.FieldPopulation,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CityUpdateOne is the builder for updating a single City entity.
type CityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CityMutation
}

// SetName sets the "Name" field.
func (cuo *CityUpdateOne) SetName(s string) *CityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (cuo *CityUpdateOne) SetNillableName(s *string) *CityUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "Name" field.
func (cuo *CityUpdateOne) ClearName() *CityUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetCountryCode sets the "CountryCode" field.
func (cuo *CityUpdateOne) SetCountryCode(s string) *CityUpdateOne {
	cuo.mutation.SetCountryCode(s)
	return cuo
}

// SetNillableCountryCode sets the "CountryCode" field if the given value is not nil.
func (cuo *CityUpdateOne) SetNillableCountryCode(s *string) *CityUpdateOne {
	if s != nil {
		cuo.SetCountryCode(*s)
	}
	return cuo
}

// ClearCountryCode clears the value of the "CountryCode" field.
func (cuo *CityUpdateOne) ClearCountryCode() *CityUpdateOne {
	cuo.mutation.ClearCountryCode()
	return cuo
}

// SetDistrict sets the "District" field.
func (cuo *CityUpdateOne) SetDistrict(s string) *CityUpdateOne {
	cuo.mutation.SetDistrict(s)
	return cuo
}

// SetNillableDistrict sets the "District" field if the given value is not nil.
func (cuo *CityUpdateOne) SetNillableDistrict(s *string) *CityUpdateOne {
	if s != nil {
		cuo.SetDistrict(*s)
	}
	return cuo
}

// ClearDistrict clears the value of the "District" field.
func (cuo *CityUpdateOne) ClearDistrict() *CityUpdateOne {
	cuo.mutation.ClearDistrict()
	return cuo
}

// SetPopulation sets the "Population" field.
func (cuo *CityUpdateOne) SetPopulation(i int) *CityUpdateOne {
	cuo.mutation.ResetPopulation()
	cuo.mutation.SetPopulation(i)
	return cuo
}

// SetNillablePopulation sets the "Population" field if the given value is not nil.
func (cuo *CityUpdateOne) SetNillablePopulation(i *int) *CityUpdateOne {
	if i != nil {
		cuo.SetPopulation(*i)
	}
	return cuo
}

// AddPopulation adds i to the "Population" field.
func (cuo *CityUpdateOne) AddPopulation(i int) *CityUpdateOne {
	cuo.mutation.AddPopulation(i)
	return cuo
}

// ClearPopulation clears the value of the "Population" field.
func (cuo *CityUpdateOne) ClearPopulation() *CityUpdateOne {
	cuo.mutation.ClearPopulation()
	return cuo
}

// Mutation returns the CityMutation object of the builder.
func (cuo *CityUpdateOne) Mutation() *CityMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CityUpdateOne) Select(field string, fields ...string) *CityUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated City entity.
func (cuo *CityUpdateOne) Save(ctx context.Context) (*City, error) {
	var (
		err  error
		node *City
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*City)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CityMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CityUpdateOne) SaveX(ctx context.Context) *City {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CityUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := city.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf(`ent: validator failed for field "City.Name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.CountryCode(); ok {
		if err := city.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "CountryCode", err: fmt.Errorf(`ent: validator failed for field "City.CountryCode": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.District(); ok {
		if err := city.DistrictValidator(v); err != nil {
			return &ValidationError{Name: "District", err: fmt.Errorf(`ent: validator failed for field "City.District": %w`, err)}
		}
	}
	return nil
}

func (cuo *CityUpdateOne) sqlSave(ctx context.Context) (_node *City, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   city.Table,
			Columns: city.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "City.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, city.FieldID)
		for _, f := range fields {
			if !city.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != city.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
	}
	if cuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldName,
		})
	}
	if value, ok := cuo.mutation.CountryCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldCountryCode,
		})
	}
	if cuo.mutation.CountryCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldCountryCode,
		})
	}
	if value, ok := cuo.mutation.District(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldDistrict,
		})
	}
	if cuo.mutation.DistrictCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: city.FieldDistrict,
		})
	}
	if value, ok := cuo.mutation.Population(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldPopulation,
		})
	}
	if value, ok := cuo.mutation.AddedPopulation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: city.FieldPopulation,
		})
	}
	if cuo.mutation.PopulationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: city.FieldPopulation,
		})
	}
	_node = &City{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
