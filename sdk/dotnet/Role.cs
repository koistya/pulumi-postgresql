// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    /// <summary>
    /// The ``postgresql..Role`` resource creates and manages a role on a PostgreSQL
    /// server.
    /// 
    /// When a ``postgresql..Role`` resource is removed, the PostgreSQL ROLE will
    /// automatically run a [`REASSIGN
    /// OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html)
    /// and [`DROP
    /// OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html) to
    /// the `CURRENT_USER` (normally the connected user for the provider).  If the
    /// specified PostgreSQL ROLE owns objects in multiple PostgreSQL databases in the
    /// same PostgreSQL Cluster, one PostgreSQL provider per database must be created
    /// and all but the final ``postgresql..Role`` must specify a `skip_drop_role`.
    /// 
    /// &gt; **Note:** All arguments including role name and password will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/role.html.markdown.
    /// </summary>
    public partial class Role : Pulumi.CustomResource
    {
        /// <summary>
        /// Defines whether a role bypasses every
        /// row-level security (RLS) policy.  Default value is `false`.
        /// </summary>
        [Output("bypassRowLevelSecurity")]
        public Output<bool?> BypassRowLevelSecurity { get; private set; } = null!;

        /// <summary>
        /// If this role can log in, this specifies how
        /// many concurrent connections the role can establish. `-1` (the default) means no
        /// limit.
        /// </summary>
        [Output("connectionLimit")]
        public Output<int?> ConnectionLimit { get; private set; } = null!;

        /// <summary>
        /// Defines a role's ability to execute `CREATE
        /// DATABASE`.  Default value is `false`.
        /// </summary>
        [Output("createDatabase")]
        public Output<bool?> CreateDatabase { get; private set; } = null!;

        /// <summary>
        /// Defines a role's ability to execute `CREATE ROLE`.
        /// A role with this privilege can also alter and drop other roles.  Default value
        /// is `false`.
        /// </summary>
        [Output("createRole")]
        public Output<bool?> CreateRole { get; private set; } = null!;

        [Output("encrypted")]
        public Output<string?> Encrypted { get; private set; } = null!;

        /// <summary>
        /// Defines whether the password is stored
        /// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
        /// is always set (to the conservative and safe value), but may interfere with the
        /// behavior of
        /// [PostgreSQL's `password_encryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
        /// </summary>
        [Output("encryptedPassword")]
        public Output<bool?> EncryptedPassword { get; private set; } = null!;

        /// <summary>
        /// Defines whether a role "inherits" the privileges of
        /// roles it is a member of.  Default value is `true`.
        /// </summary>
        [Output("inherit")]
        public Output<bool?> Inherit { get; private set; } = null!;

        /// <summary>
        /// Defines whether role is allowed to log in.  Roles without
        /// this attribute are useful for managing database privileges, but are not users
        /// in the usual sense of the word.  Default value is `false`.
        /// </summary>
        [Output("login")]
        public Output<bool?> Login { get; private set; } = null!;

        /// <summary>
        /// The name of the role. Must be unique on the PostgreSQL
        /// server instance where it is configured.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets the role's password. A password is only of use
        /// for roles having the `login` attribute set to true.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Defines whether a role is allowed to initiate
        /// streaming replication or put the system in and out of backup mode.  Default
        /// value is `false`
        /// </summary>
        [Output("replication")]
        public Output<bool?> Replication { get; private set; } = null!;

        /// <summary>
        /// Defines list of roles which will be granted to this new role.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// Alters the search path of this new role. Note that
        /// due to limitations in the implementation, values cannot contain the substring
        /// `", "`.
        /// </summary>
        [Output("searchPaths")]
        public Output<ImmutableArray<string>> SearchPaths { get; private set; } = null!;

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, the
        /// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
        /// in each of the respective databases must occur before the ROLE can be dropped
        /// from the catalog.  Set this option to true when there are multiple databases
        /// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
        /// This is the third and final step taken when removing a ROLE from a database.
        /// </summary>
        [Output("skipDropRole")]
        public Output<bool?> SkipDropRole { get; private set; } = null!;

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, a
        /// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
        /// must be executed on each of the respective databases before the `DROP ROLE`
        /// can be executed to dropped the ROLE from the catalog.  This is the first and
        /// second steps taken when removing a ROLE from a database (the second step being
        /// an implicit
        /// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
        /// </summary>
        [Output("skipReassignOwned")]
        public Output<bool?> SkipReassignOwned { get; private set; } = null!;

        /// <summary>
        /// Defines whether the role is a "superuser", and
        /// therefore can override all access restrictions within the database.  Default
        /// value is `false`.
        /// </summary>
        [Output("superuser")]
        public Output<bool?> Superuser { get; private set; } = null!;

        /// <summary>
        /// Defines the date and time after which the role's
        /// password is no longer valid.  Established connections past this `valid_time`
        /// will have to be manually terminated.  This value corresponds to a PostgreSQL
        /// datetime. If omitted or the magic value `NULL` is used, `valid_until` will be
        /// set to `infinity`.  Default is `NULL`, therefore `infinity`.
        /// </summary>
        [Output("validUntil")]
        public Output<string?> ValidUntil { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql:index/role:Role", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/role:Role", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether a role bypasses every
        /// row-level security (RLS) policy.  Default value is `false`.
        /// </summary>
        [Input("bypassRowLevelSecurity")]
        public Input<bool>? BypassRowLevelSecurity { get; set; }

        /// <summary>
        /// If this role can log in, this specifies how
        /// many concurrent connections the role can establish. `-1` (the default) means no
        /// limit.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// Defines a role's ability to execute `CREATE
        /// DATABASE`.  Default value is `false`.
        /// </summary>
        [Input("createDatabase")]
        public Input<bool>? CreateDatabase { get; set; }

        /// <summary>
        /// Defines a role's ability to execute `CREATE ROLE`.
        /// A role with this privilege can also alter and drop other roles.  Default value
        /// is `false`.
        /// </summary>
        [Input("createRole")]
        public Input<bool>? CreateRole { get; set; }

        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        /// <summary>
        /// Defines whether the password is stored
        /// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
        /// is always set (to the conservative and safe value), but may interfere with the
        /// behavior of
        /// [PostgreSQL's `password_encryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
        /// </summary>
        [Input("encryptedPassword")]
        public Input<bool>? EncryptedPassword { get; set; }

        /// <summary>
        /// Defines whether a role "inherits" the privileges of
        /// roles it is a member of.  Default value is `true`.
        /// </summary>
        [Input("inherit")]
        public Input<bool>? Inherit { get; set; }

        /// <summary>
        /// Defines whether role is allowed to log in.  Roles without
        /// this attribute are useful for managing database privileges, but are not users
        /// in the usual sense of the word.  Default value is `false`.
        /// </summary>
        [Input("login")]
        public Input<bool>? Login { get; set; }

        /// <summary>
        /// The name of the role. Must be unique on the PostgreSQL
        /// server instance where it is configured.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the role's password. A password is only of use
        /// for roles having the `login` attribute set to true.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Defines whether a role is allowed to initiate
        /// streaming replication or put the system in and out of backup mode.  Default
        /// value is `false`
        /// </summary>
        [Input("replication")]
        public Input<bool>? Replication { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Defines list of roles which will be granted to this new role.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("searchPaths")]
        private InputList<string>? _searchPaths;

        /// <summary>
        /// Alters the search path of this new role. Note that
        /// due to limitations in the implementation, values cannot contain the substring
        /// `", "`.
        /// </summary>
        public InputList<string> SearchPaths
        {
            get => _searchPaths ?? (_searchPaths = new InputList<string>());
            set => _searchPaths = value;
        }

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, the
        /// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
        /// in each of the respective databases must occur before the ROLE can be dropped
        /// from the catalog.  Set this option to true when there are multiple databases
        /// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
        /// This is the third and final step taken when removing a ROLE from a database.
        /// </summary>
        [Input("skipDropRole")]
        public Input<bool>? SkipDropRole { get; set; }

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, a
        /// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
        /// must be executed on each of the respective databases before the `DROP ROLE`
        /// can be executed to dropped the ROLE from the catalog.  This is the first and
        /// second steps taken when removing a ROLE from a database (the second step being
        /// an implicit
        /// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
        /// </summary>
        [Input("skipReassignOwned")]
        public Input<bool>? SkipReassignOwned { get; set; }

        /// <summary>
        /// Defines whether the role is a "superuser", and
        /// therefore can override all access restrictions within the database.  Default
        /// value is `false`.
        /// </summary>
        [Input("superuser")]
        public Input<bool>? Superuser { get; set; }

        /// <summary>
        /// Defines the date and time after which the role's
        /// password is no longer valid.  Established connections past this `valid_time`
        /// will have to be manually terminated.  This value corresponds to a PostgreSQL
        /// datetime. If omitted or the magic value `NULL` is used, `valid_until` will be
        /// set to `infinity`.  Default is `NULL`, therefore `infinity`.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether a role bypasses every
        /// row-level security (RLS) policy.  Default value is `false`.
        /// </summary>
        [Input("bypassRowLevelSecurity")]
        public Input<bool>? BypassRowLevelSecurity { get; set; }

        /// <summary>
        /// If this role can log in, this specifies how
        /// many concurrent connections the role can establish. `-1` (the default) means no
        /// limit.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// Defines a role's ability to execute `CREATE
        /// DATABASE`.  Default value is `false`.
        /// </summary>
        [Input("createDatabase")]
        public Input<bool>? CreateDatabase { get; set; }

        /// <summary>
        /// Defines a role's ability to execute `CREATE ROLE`.
        /// A role with this privilege can also alter and drop other roles.  Default value
        /// is `false`.
        /// </summary>
        [Input("createRole")]
        public Input<bool>? CreateRole { get; set; }

        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        /// <summary>
        /// Defines whether the password is stored
        /// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
        /// is always set (to the conservative and safe value), but may interfere with the
        /// behavior of
        /// [PostgreSQL's `password_encryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
        /// </summary>
        [Input("encryptedPassword")]
        public Input<bool>? EncryptedPassword { get; set; }

        /// <summary>
        /// Defines whether a role "inherits" the privileges of
        /// roles it is a member of.  Default value is `true`.
        /// </summary>
        [Input("inherit")]
        public Input<bool>? Inherit { get; set; }

        /// <summary>
        /// Defines whether role is allowed to log in.  Roles without
        /// this attribute are useful for managing database privileges, but are not users
        /// in the usual sense of the word.  Default value is `false`.
        /// </summary>
        [Input("login")]
        public Input<bool>? Login { get; set; }

        /// <summary>
        /// The name of the role. Must be unique on the PostgreSQL
        /// server instance where it is configured.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the role's password. A password is only of use
        /// for roles having the `login` attribute set to true.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Defines whether a role is allowed to initiate
        /// streaming replication or put the system in and out of backup mode.  Default
        /// value is `false`
        /// </summary>
        [Input("replication")]
        public Input<bool>? Replication { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Defines list of roles which will be granted to this new role.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("searchPaths")]
        private InputList<string>? _searchPaths;

        /// <summary>
        /// Alters the search path of this new role. Note that
        /// due to limitations in the implementation, values cannot contain the substring
        /// `", "`.
        /// </summary>
        public InputList<string> SearchPaths
        {
            get => _searchPaths ?? (_searchPaths = new InputList<string>());
            set => _searchPaths = value;
        }

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, the
        /// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
        /// in each of the respective databases must occur before the ROLE can be dropped
        /// from the catalog.  Set this option to true when there are multiple databases
        /// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
        /// This is the third and final step taken when removing a ROLE from a database.
        /// </summary>
        [Input("skipDropRole")]
        public Input<bool>? SkipDropRole { get; set; }

        /// <summary>
        /// When a PostgreSQL ROLE exists in multiple
        /// databases and the ROLE is dropped, a
        /// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
        /// must be executed on each of the respective databases before the `DROP ROLE`
        /// can be executed to dropped the ROLE from the catalog.  This is the first and
        /// second steps taken when removing a ROLE from a database (the second step being
        /// an implicit
        /// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
        /// </summary>
        [Input("skipReassignOwned")]
        public Input<bool>? SkipReassignOwned { get; set; }

        /// <summary>
        /// Defines whether the role is a "superuser", and
        /// therefore can override all access restrictions within the database.  Default
        /// value is `false`.
        /// </summary>
        [Input("superuser")]
        public Input<bool>? Superuser { get; set; }

        /// <summary>
        /// Defines the date and time after which the role's
        /// password is no longer valid.  Established connections past this `valid_time`
        /// will have to be manually terminated.  This value corresponds to a PostgreSQL
        /// datetime. If omitted or the magic value `NULL` is used, `valid_until` will be
        /// set to `infinity`.  Default is `NULL`, therefore `infinity`.
        /// </summary>
        [Input("validUntil")]
        public Input<string>? ValidUntil { get; set; }

        public RoleState()
        {
        }
    }
}
