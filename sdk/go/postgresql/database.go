// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``postgresql_database`` resource creates and manages [database
// objects](https://www.postgresql.org/docs/current/static/managing-databases.html)
// within a PostgreSQL server instance.
type Database struct {
	s *pulumi.ResourceState
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOpt) (*Database, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowConnections"] = nil
		inputs["connectionLimit"] = nil
		inputs["encoding"] = nil
		inputs["isTemplate"] = nil
		inputs["lcCollate"] = nil
		inputs["lcCtype"] = nil
		inputs["name"] = nil
		inputs["owner"] = nil
		inputs["tablespaceName"] = nil
		inputs["template"] = nil
	} else {
		inputs["allowConnections"] = args.AllowConnections
		inputs["connectionLimit"] = args.ConnectionLimit
		inputs["encoding"] = args.Encoding
		inputs["isTemplate"] = args.IsTemplate
		inputs["lcCollate"] = args.LcCollate
		inputs["lcCtype"] = args.LcCtype
		inputs["name"] = args.Name
		inputs["owner"] = args.Owner
		inputs["tablespaceName"] = args.TablespaceName
		inputs["template"] = args.Template
	}
	s, err := ctx.RegisterResource("postgresql:index/database:Database", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Database{s: s}, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatabaseState, opts ...pulumi.ResourceOpt) (*Database, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowConnections"] = state.AllowConnections
		inputs["connectionLimit"] = state.ConnectionLimit
		inputs["encoding"] = state.Encoding
		inputs["isTemplate"] = state.IsTemplate
		inputs["lcCollate"] = state.LcCollate
		inputs["lcCtype"] = state.LcCtype
		inputs["name"] = state.Name
		inputs["owner"] = state.Owner
		inputs["tablespaceName"] = state.TablespaceName
		inputs["template"] = state.Template
	}
	s, err := ctx.ReadResource("postgresql:index/database:Database", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Database{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Database) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Database) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// If `false` then no one can connect to this
// database. The default is `true`, allowing connections (except as restricted by
// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
func (r *Database) AllowConnections() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowConnections"])
}

// How many concurrent connections can be
// established to this database. `-1` (the default) means no limit.
func (r *Database) ConnectionLimit() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["connectionLimit"])
}

// Character set encoding to use in the database.
// Specify a string constant (e.g. `UTF8` or `SQL_ASCII`), or an integer encoding
// number.  If unset or set to an empty string the default encoding is set to
// `UTF8`.  If set to `DEFAULT` Terraform will use the same encoding as the
// template database.  Changing this value will force the creation of a new
// resource as this value can only be changed when a database is created.
func (r *Database) Encoding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["encoding"])
}

// If `true`, then this database can be cloned by any
// user with `CREATEDB` privileges; if `false` (the default), then only
// superusers or the owner of the database can clone it.
func (r *Database) IsTemplate() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["isTemplate"])
}

// Collation order (`LC_COLLATE`) to use in the
// database.  This affects the sort order applied to strings, e.g. in queries
// with `ORDER BY`, as well as the order used in indexes on text columns. If
// unset or set to an empty string the default collation is set to `C`.  If set
// to `DEFAULT` Terraform will use the same collation order as the specified
// `template` database.  Changing this value will force the creation of a new
// resource as this value can only be changed when a database is created.
func (r *Database) LcCollate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lcCollate"])
}

// Character classification (`LC_CTYPE`) to use in the
// database. This affects the categorization of characters, e.g. lower, upper and
// digit. If unset or set to an empty string the default character classification
// is set to `C`.  If set to `DEFAULT` Terraform will use the character
// classification of the specified `template` database.  Changing this value will
// force the creation of a new resource as this value can only be changed when a
// database is created.
func (r *Database) LcCtype() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lcCtype"])
}

// The name of the database. Must be unique on the PostgreSQL
// server instance where it is configured.
func (r *Database) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The role name of the user who will own the database, or
// `DEFAULT` to use the default (namely, the user executing the command). To
// create a database owned by another role or to change the owner of an existing
// database, you must be a direct or indirect member of the specified role, or
// the username in the provider is a superuser.
func (r *Database) Owner() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["owner"])
}

// The name of the tablespace that will be
// associated with the database, or `DEFAULT` to use the template database's
// tablespace.  This tablespace will be the default tablespace used for objects
// created in this database.
func (r *Database) TablespaceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["tablespaceName"])
}

// The name of the template database from which to create
// the database, or `DEFAULT` to use the default template (`template0`).  NOTE:
// the default in Terraform is `template0`, not `template1`.  Changing this value
// will force the creation of a new resource as this value can only be changed
// when a database is created.
func (r *Database) Template() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["template"])
}

// Input properties used for looking up and filtering Database resources.
type DatabaseState struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections interface{}
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit interface{}
	// Character set encoding to use in the database.
	// Specify a string constant (e.g. `UTF8` or `SQL_ASCII`), or an integer encoding
	// number.  If unset or set to an empty string the default encoding is set to
	// `UTF8`.  If set to `DEFAULT` Terraform will use the same encoding as the
	// template database.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	Encoding interface{}
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate interface{}
	// Collation order (`LC_COLLATE`) to use in the
	// database.  This affects the sort order applied to strings, e.g. in queries
	// with `ORDER BY`, as well as the order used in indexes on text columns. If
	// unset or set to an empty string the default collation is set to `C`.  If set
	// to `DEFAULT` Terraform will use the same collation order as the specified
	// `template` database.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	LcCollate interface{}
	// Character classification (`LC_CTYPE`) to use in the
	// database. This affects the categorization of characters, e.g. lower, upper and
	// digit. If unset or set to an empty string the default character classification
	// is set to `C`.  If set to `DEFAULT` Terraform will use the character
	// classification of the specified `template` database.  Changing this value will
	// force the creation of a new resource as this value can only be changed when a
	// database is created.
	LcCtype interface{}
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name interface{}
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner interface{}
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName interface{}
	// The name of the template database from which to create
	// the database, or `DEFAULT` to use the default template (`template0`).  NOTE:
	// the default in Terraform is `template0`, not `template1`.  Changing this value
	// will force the creation of a new resource as this value can only be changed
	// when a database is created.
	Template interface{}
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// If `false` then no one can connect to this
	// database. The default is `true`, allowing connections (except as restricted by
	// other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
	AllowConnections interface{}
	// How many concurrent connections can be
	// established to this database. `-1` (the default) means no limit.
	ConnectionLimit interface{}
	// Character set encoding to use in the database.
	// Specify a string constant (e.g. `UTF8` or `SQL_ASCII`), or an integer encoding
	// number.  If unset or set to an empty string the default encoding is set to
	// `UTF8`.  If set to `DEFAULT` Terraform will use the same encoding as the
	// template database.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	Encoding interface{}
	// If `true`, then this database can be cloned by any
	// user with `CREATEDB` privileges; if `false` (the default), then only
	// superusers or the owner of the database can clone it.
	IsTemplate interface{}
	// Collation order (`LC_COLLATE`) to use in the
	// database.  This affects the sort order applied to strings, e.g. in queries
	// with `ORDER BY`, as well as the order used in indexes on text columns. If
	// unset or set to an empty string the default collation is set to `C`.  If set
	// to `DEFAULT` Terraform will use the same collation order as the specified
	// `template` database.  Changing this value will force the creation of a new
	// resource as this value can only be changed when a database is created.
	LcCollate interface{}
	// Character classification (`LC_CTYPE`) to use in the
	// database. This affects the categorization of characters, e.g. lower, upper and
	// digit. If unset or set to an empty string the default character classification
	// is set to `C`.  If set to `DEFAULT` Terraform will use the character
	// classification of the specified `template` database.  Changing this value will
	// force the creation of a new resource as this value can only be changed when a
	// database is created.
	LcCtype interface{}
	// The name of the database. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name interface{}
	// The role name of the user who will own the database, or
	// `DEFAULT` to use the default (namely, the user executing the command). To
	// create a database owned by another role or to change the owner of an existing
	// database, you must be a direct or indirect member of the specified role, or
	// the username in the provider is a superuser.
	Owner interface{}
	// The name of the tablespace that will be
	// associated with the database, or `DEFAULT` to use the template database's
	// tablespace.  This tablespace will be the default tablespace used for objects
	// created in this database.
	TablespaceName interface{}
	// The name of the template database from which to create
	// the database, or `DEFAULT` to use the default template (`template0`).  NOTE:
	// the default in Terraform is `template0`, not `template1`.  Changing this value
	// will force the creation of a new resource as this value can only be changed
	// when a database is created.
	Template interface{}
}
