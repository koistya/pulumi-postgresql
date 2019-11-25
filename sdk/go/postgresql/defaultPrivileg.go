// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.DefaultPrivileg`` resource creates and manages default privileges given to a user for a database schema.
// 
// > **Note:** This resource needs Postgresql version 9 or above.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/default_privileges.html.markdown.
type DefaultPrivileg struct {
	s *pulumi.ResourceState
}

// NewDefaultPrivileg registers a new resource with the given unique name, arguments, and options.
func NewDefaultPrivileg(ctx *pulumi.Context,
	name string, args *DefaultPrivilegArgs, opts ...pulumi.ResourceOpt) (*DefaultPrivileg, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.ObjectType == nil {
		return nil, errors.New("missing required argument 'ObjectType'")
	}
	if args == nil || args.Owner == nil {
		return nil, errors.New("missing required argument 'Owner'")
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
		inputs["owner"] = nil
		inputs["privileges"] = nil
		inputs["role"] = nil
		inputs["schema"] = nil
	} else {
		inputs["database"] = args.Database
		inputs["objectType"] = args.ObjectType
		inputs["owner"] = args.Owner
		inputs["privileges"] = args.Privileges
		inputs["role"] = args.Role
		inputs["schema"] = args.Schema
	}
	s, err := ctx.RegisterResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultPrivileg{s: s}, nil
}

// GetDefaultPrivileg gets an existing DefaultPrivileg resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPrivileg(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultPrivilegState, opts ...pulumi.ResourceOpt) (*DefaultPrivileg, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["database"] = state.Database
		inputs["objectType"] = state.ObjectType
		inputs["owner"] = state.Owner
		inputs["privileges"] = state.Privileges
		inputs["role"] = state.Role
		inputs["schema"] = state.Schema
	}
	s, err := ctx.ReadResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultPrivileg{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultPrivileg) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultPrivileg) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The database to grant default privileges for this role.
func (r *DefaultPrivileg) Database() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["database"])
}

// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
func (r *DefaultPrivileg) ObjectType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["objectType"])
}

// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
func (r *DefaultPrivileg) Owner() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["owner"])
}

// The list of privileges to apply as default privileges.
func (r *DefaultPrivileg) Privileges() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["privileges"])
}

// The name of the role to which grant default privileges on.
func (r *DefaultPrivileg) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// The database schema to set default privileges for this role.
func (r *DefaultPrivileg) Schema() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["schema"])
}

// Input properties used for looking up and filtering DefaultPrivileg resources.
type DefaultPrivilegState struct {
	// The database to grant default privileges for this role.
	Database interface{}
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
	ObjectType interface{}
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner interface{}
	// The list of privileges to apply as default privileges.
	Privileges interface{}
	// The name of the role to which grant default privileges on.
	Role interface{}
	// The database schema to set default privileges for this role.
	Schema interface{}
}

// The set of arguments for constructing a DefaultPrivileg resource.
type DefaultPrivilegArgs struct {
	// The database to grant default privileges for this role.
	Database interface{}
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
	ObjectType interface{}
	// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
	Owner interface{}
	// The list of privileges to apply as default privileges.
	Privileges interface{}
	// The name of the role to which grant default privileges on.
	Role interface{}
	// The database schema to set default privileges for this role.
	Schema interface{}
}
