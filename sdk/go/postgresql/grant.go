// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``postgresql_grant`` resource creates and manages privileges given to a user for a database schema.
// 
// > **Note:** This resource needs Postgresql version 9 or above.
type Grant struct {
	s *pulumi.ResourceState
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOpt) (*Grant, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.ObjectType == nil {
		return nil, errors.New("missing required argument 'ObjectType'")
	}
	if args == nil || args.Privileges == nil {
		return nil, errors.New("missing required argument 'Privileges'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Schema == nil {
		return nil, errors.New("missing required argument 'Schema'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["database"] = nil
		inputs["objectType"] = nil
		inputs["privileges"] = nil
		inputs["role"] = nil
		inputs["schema"] = nil
	} else {
		inputs["database"] = args.Database
		inputs["objectType"] = args.ObjectType
		inputs["privileges"] = args.Privileges
		inputs["role"] = args.Role
		inputs["schema"] = args.Schema
	}
	s, err := ctx.RegisterResource("postgresql:index/grant:Grant", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Grant{s: s}, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.ID, state *GrantState, opts ...pulumi.ResourceOpt) (*Grant, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["database"] = state.Database
		inputs["objectType"] = state.ObjectType
		inputs["privileges"] = state.Privileges
		inputs["role"] = state.Role
		inputs["schema"] = state.Schema
	}
	s, err := ctx.ReadResource("postgresql:index/grant:Grant", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Grant{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Grant) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Grant) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The database to grant privileges on for this role.
func (r *Grant) Database() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["database"])
}

// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
func (r *Grant) ObjectType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["objectType"])
}

// The list of privileges to grant.
func (r *Grant) Privileges() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["privileges"])
}

// The name of the role to grant privileges on.
func (r *Grant) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The database schema to grant privileges on for this role.
func (r *Grant) Schema() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["schema"])
}

// Input properties used for looking up and filtering Grant resources.
type GrantState struct {
	// The database to grant privileges on for this role.
	Database interface{}
	// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
	ObjectType interface{}
	// The list of privileges to grant.
	Privileges interface{}
	// The name of the role to grant privileges on.
	Role interface{}
	// The database schema to grant privileges on for this role.
	Schema interface{}
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	// The database to grant privileges on for this role.
	Database interface{}
	// The PostgreSQL object type to grant the privileges on (one of: table, sequence).
	ObjectType interface{}
	// The list of privileges to grant.
	Privileges interface{}
	// The name of the role to grant privileges on.
	Role interface{}
	// The database schema to grant privileges on for this role.
	Schema interface{}
}