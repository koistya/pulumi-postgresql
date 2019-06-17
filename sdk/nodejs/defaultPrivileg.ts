// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DefaultPrivileg extends pulumi.CustomResource {
    /**
     * Get an existing DefaultPrivileg resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultPrivilegState, opts?: pulumi.CustomResourceOptions): DefaultPrivileg {
        return new DefaultPrivileg(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/defaultPrivileg:DefaultPrivileg';

    /**
     * Returns true if the given object is an instance of DefaultPrivileg.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultPrivileg {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultPrivileg.__pulumiType;
    }

    /**
     * The database to grant default privileges for this role
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence)
     */
    public readonly objectType!: pulumi.Output<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by
     * yourself or by roles that you are a member of)
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * The list of privileges to apply as default privileges
     */
    public readonly privileges!: pulumi.Output<string[]>;
    /**
     * The name of the role to which grant default privileges on
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The database schema to set default privileges for this role
     */
    public readonly schema!: pulumi.Output<string>;

    /**
     * Create a DefaultPrivileg resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultPrivilegArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultPrivilegArgs | DefaultPrivilegState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DefaultPrivilegState | undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["objectType"] = state ? state.objectType : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["privileges"] = state ? state.privileges : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["schema"] = state ? state.schema : undefined;
        } else {
            const args = argsOrState as DefaultPrivilegArgs | undefined;
            if (!args || args.database === undefined) {
                throw new Error("Missing required property 'database'");
            }
            if (!args || args.objectType === undefined) {
                throw new Error("Missing required property 'objectType'");
            }
            if (!args || args.owner === undefined) {
                throw new Error("Missing required property 'owner'");
            }
            if (!args || args.privileges === undefined) {
                throw new Error("Missing required property 'privileges'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            if (!args || args.schema === undefined) {
                throw new Error("Missing required property 'schema'");
            }
            inputs["database"] = args ? args.database : undefined;
            inputs["objectType"] = args ? args.objectType : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["privileges"] = args ? args.privileges : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["schema"] = args ? args.schema : undefined;
        }
        super(DefaultPrivileg.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultPrivileg resources.
 */
export interface DefaultPrivilegState {
    /**
     * The database to grant default privileges for this role
     */
    readonly database?: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence)
     */
    readonly objectType?: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by
     * yourself or by roles that you are a member of)
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges
     */
    readonly privileges?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role
     */
    readonly schema?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultPrivileg resource.
 */
export interface DefaultPrivilegArgs {
    /**
     * The database to grant default privileges for this role
     */
    readonly database: pulumi.Input<string>;
    /**
     * The PostgreSQL object type to set the default privileges on (one of: table, sequence)
     */
    readonly objectType: pulumi.Input<string>;
    /**
     * Role for which apply default privileges (You can change default privileges only for objects that will be created by
     * yourself or by roles that you are a member of)
     */
    readonly owner: pulumi.Input<string>;
    /**
     * The list of privileges to apply as default privileges
     */
    readonly privileges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role to which grant default privileges on
     */
    readonly role: pulumi.Input<string>;
    /**
     * The database schema to set default privileges for this role
     */
    readonly schema: pulumi.Input<string>;
}
