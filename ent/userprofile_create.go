// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwitt/ent/media"
	"github.com/0xfzz/tuwitt/ent/useraccount"
	"github.com/0xfzz/tuwitt/ent/userprofile"
)

// UserProfileCreate is the builder for creating a UserProfile entity.
type UserProfileCreate struct {
	config
	mutation *UserProfileMutation
	hooks    []Hook
}

// SetDisplayName sets the "display_name" field.
func (upc *UserProfileCreate) SetDisplayName(s string) *UserProfileCreate {
	upc.mutation.SetDisplayName(s)
	return upc
}

// SetBio sets the "bio" field.
func (upc *UserProfileCreate) SetBio(s string) *UserProfileCreate {
	upc.mutation.SetBio(s)
	return upc
}

// SetNillableBio sets the "bio" field if the given value is not nil.
func (upc *UserProfileCreate) SetNillableBio(s *string) *UserProfileCreate {
	if s != nil {
		upc.SetBio(*s)
	}
	return upc
}

// SetProfilePictureID sets the "profile_picture_id" field.
func (upc *UserProfileCreate) SetProfilePictureID(i int) *UserProfileCreate {
	upc.mutation.SetProfilePictureID(i)
	return upc
}

// SetNillableProfilePictureID sets the "profile_picture_id" field if the given value is not nil.
func (upc *UserProfileCreate) SetNillableProfilePictureID(i *int) *UserProfileCreate {
	if i != nil {
		upc.SetProfilePictureID(*i)
	}
	return upc
}

// SetBannerID sets the "banner_id" field.
func (upc *UserProfileCreate) SetBannerID(i int) *UserProfileCreate {
	upc.mutation.SetBannerID(i)
	return upc
}

// SetNillableBannerID sets the "banner_id" field if the given value is not nil.
func (upc *UserProfileCreate) SetNillableBannerID(i *int) *UserProfileCreate {
	if i != nil {
		upc.SetBannerID(*i)
	}
	return upc
}

// SetAccountID sets the "account" edge to the UserAccount entity by ID.
func (upc *UserProfileCreate) SetAccountID(id int) *UserProfileCreate {
	upc.mutation.SetAccountID(id)
	return upc
}

// SetAccount sets the "account" edge to the UserAccount entity.
func (upc *UserProfileCreate) SetAccount(u *UserAccount) *UserProfileCreate {
	return upc.SetAccountID(u.ID)
}

// SetProfilePicture sets the "profile_picture" edge to the Media entity.
func (upc *UserProfileCreate) SetProfilePicture(m *Media) *UserProfileCreate {
	return upc.SetProfilePictureID(m.ID)
}

// SetBanner sets the "banner" edge to the Media entity.
func (upc *UserProfileCreate) SetBanner(m *Media) *UserProfileCreate {
	return upc.SetBannerID(m.ID)
}

// Mutation returns the UserProfileMutation object of the builder.
func (upc *UserProfileCreate) Mutation() *UserProfileMutation {
	return upc.mutation
}

// Save creates the UserProfile in the database.
func (upc *UserProfileCreate) Save(ctx context.Context) (*UserProfile, error) {
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserProfileCreate) SaveX(ctx context.Context) *UserProfile {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserProfileCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserProfileCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserProfileCreate) check() error {
	if _, ok := upc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "UserProfile.display_name"`)}
	}
	if _, ok := upc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "UserProfile.account"`)}
	}
	return nil
}

func (upc *UserProfileCreate) sqlSave(ctx context.Context) (*UserProfile, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserProfileCreate) createSpec() (*UserProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &UserProfile{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userprofile.Table, sqlgraph.NewFieldSpec(userprofile.FieldID, field.TypeInt))
	)
	if value, ok := upc.mutation.DisplayName(); ok {
		_spec.SetField(userprofile.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := upc.mutation.Bio(); ok {
		_spec.SetField(userprofile.FieldBio, field.TypeString, value)
		_node.Bio = value
	}
	if nodes := upc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userprofile.AccountTable,
			Columns: []string{userprofile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_account_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := upc.mutation.ProfilePictureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.ProfilePictureTable,
			Columns: []string{userprofile.ProfilePictureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProfilePictureID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := upc.mutation.BannerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userprofile.BannerTable,
			Columns: []string{userprofile.BannerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BannerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserProfileCreateBulk is the builder for creating many UserProfile entities in bulk.
type UserProfileCreateBulk struct {
	config
	err      error
	builders []*UserProfileCreate
}

// Save creates the UserProfile entities in the database.
func (upcb *UserProfileCreateBulk) Save(ctx context.Context) ([]*UserProfile, error) {
	if upcb.err != nil {
		return nil, upcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserProfile, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserProfileCreateBulk) SaveX(ctx context.Context) []*UserProfile {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserProfileCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
