// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class AccountPrivilege extends pulumi.CustomResource {
    /**
     * Get an existing AccountPrivilege resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountPrivilegeState, opts?: pulumi.CustomResourceOptions): AccountPrivilege {
        return new AccountPrivilege(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/accountPrivilege:AccountPrivilege';

    /**
     * Returns true if the given object is an instance of AccountPrivilege.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountPrivilege {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountPrivilege.__pulumiType;
    }

    /**
     * Account host, default is `%`.
     */
    public readonly accountHost!: pulumi.Output<string | undefined>;
    /**
     * Account name.
     */
    public readonly accountName!: pulumi.Output<string>;
    /**
     * List of specified database name.
     */
    public readonly databaseNames!: pulumi.Output<string[]>;
    /**
     * Instance ID.
     */
    public readonly mysqlId!: pulumi.Output<string>;
    /**
     * Database permissions. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`,
     * `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER
     * ROUTINE`, `EVENT` and `TRIGGER``.
     */
    public readonly privileges!: pulumi.Output<string[] | undefined>;

    /**
     * Create a AccountPrivilege resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountPrivilegeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountPrivilegeArgs | AccountPrivilegeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountPrivilegeState | undefined;
            resourceInputs["accountHost"] = state ? state.accountHost : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["databaseNames"] = state ? state.databaseNames : undefined;
            resourceInputs["mysqlId"] = state ? state.mysqlId : undefined;
            resourceInputs["privileges"] = state ? state.privileges : undefined;
        } else {
            const args = argsOrState as AccountPrivilegeArgs | undefined;
            if ((!args || args.accountName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.databaseNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseNames'");
            }
            if ((!args || args.mysqlId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mysqlId'");
            }
            resourceInputs["accountHost"] = args ? args.accountHost : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["databaseNames"] = args ? args.databaseNames : undefined;
            resourceInputs["mysqlId"] = args ? args.mysqlId : undefined;
            resourceInputs["privileges"] = args ? args.privileges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountPrivilege.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountPrivilege resources.
 */
export interface AccountPrivilegeState {
    /**
     * Account host, default is `%`.
     */
    accountHost?: pulumi.Input<string>;
    /**
     * Account name.
     */
    accountName?: pulumi.Input<string>;
    /**
     * List of specified database name.
     */
    databaseNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance ID.
     */
    mysqlId?: pulumi.Input<string>;
    /**
     * Database permissions. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`,
     * `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER
     * ROUTINE`, `EVENT` and `TRIGGER``.
     */
    privileges?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AccountPrivilege resource.
 */
export interface AccountPrivilegeArgs {
    /**
     * Account host, default is `%`.
     */
    accountHost?: pulumi.Input<string>;
    /**
     * Account name.
     */
    accountName: pulumi.Input<string>;
    /**
     * List of specified database name.
     */
    databaseNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance ID.
     */
    mysqlId: pulumi.Input<string>;
    /**
     * Database permissions. Valid values: `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, `REFERENCES`, `INDEX`,
     * `ALTER`, `CREATE TEMPORARY TABLES`, `LOCK TABLES`, `EXECUTE`, `CREATE VIEW`, `SHOW VIEW`, `CREATE ROUTINE`, `ALTER
     * ROUTINE`, `EVENT` and `TRIGGER``.
     */
    privileges?: pulumi.Input<pulumi.Input<string>[]>;
}