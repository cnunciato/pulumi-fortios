// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Read/Update/Delete admin user for FortiManager
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.FmgSystemAdminUser("test1", {
 *     description: "local user",
 *     password: "123",
 *     profileid: "Standard_User",
 *     rpcPermit: "read",
 *     trusthost1: "1.1.1.1 255.255.255.255",
 *     userType: "local",
 *     userid: "user1",
 * });
 * const test2 = new fortios.FmgSystemAdminUser("test2", {
 *     description: "api user",
 *     password: "098",
 *     profileid: "Standard_User",
 *     rpcPermit: "read-write",
 *     trusthost1: "2.2.2.2 255.255.255.255",
 *     userid: "user2",
 * });
 * ```
 */
export class FmgSystemAdminUser extends pulumi.CustomResource {
    /**
     * Get an existing FmgSystemAdminUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FmgSystemAdminUserState, opts?: pulumi.CustomResourceOptions): FmgSystemAdminUser {
        return new FmgSystemAdminUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fmgSystemAdminUser:FmgSystemAdminUser';

    /**
     * Returns true if the given object is an instance of FmgSystemAdminUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FmgSystemAdminUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FmgSystemAdminUser.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Password.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Profile id.
     */
    public readonly profileid!: pulumi.Output<string | undefined>;
    /**
     * RADIUS server name.
     */
    public readonly radiusServer!: pulumi.Output<string | undefined>;
    /**
     * Rpc permit, Enum: ["read-write", "none", "read"]
     */
    public readonly rpcPermit!: pulumi.Output<string | undefined>;
    /**
     * Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
     */
    public readonly trusthost1!: pulumi.Output<string | undefined>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    public readonly trusthost2!: pulumi.Output<string | undefined>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    public readonly trusthost3!: pulumi.Output<string | undefined>;
    /**
     * User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
     */
    public readonly userType!: pulumi.Output<string | undefined>;
    /**
     * User name.
     */
    public readonly userid!: pulumi.Output<string>;

    /**
     * Create a FmgSystemAdminUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FmgSystemAdminUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FmgSystemAdminUserArgs | FmgSystemAdminUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FmgSystemAdminUserState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["profileid"] = state ? state.profileid : undefined;
            resourceInputs["radiusServer"] = state ? state.radiusServer : undefined;
            resourceInputs["rpcPermit"] = state ? state.rpcPermit : undefined;
            resourceInputs["trusthost1"] = state ? state.trusthost1 : undefined;
            resourceInputs["trusthost2"] = state ? state.trusthost2 : undefined;
            resourceInputs["trusthost3"] = state ? state.trusthost3 : undefined;
            resourceInputs["userType"] = state ? state.userType : undefined;
            resourceInputs["userid"] = state ? state.userid : undefined;
        } else {
            const args = argsOrState as FmgSystemAdminUserArgs | undefined;
            if ((!args || args.userid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userid'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["profileid"] = args ? args.profileid : undefined;
            resourceInputs["radiusServer"] = args ? args.radiusServer : undefined;
            resourceInputs["rpcPermit"] = args ? args.rpcPermit : undefined;
            resourceInputs["trusthost1"] = args ? args.trusthost1 : undefined;
            resourceInputs["trusthost2"] = args ? args.trusthost2 : undefined;
            resourceInputs["trusthost3"] = args ? args.trusthost3 : undefined;
            resourceInputs["userType"] = args ? args.userType : undefined;
            resourceInputs["userid"] = args ? args.userid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FmgSystemAdminUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FmgSystemAdminUser resources.
 */
export interface FmgSystemAdminUserState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Password.
     */
    password?: pulumi.Input<string>;
    /**
     * Profile id.
     */
    profileid?: pulumi.Input<string>;
    /**
     * RADIUS server name.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * Rpc permit, Enum: ["read-write", "none", "read"]
     */
    rpcPermit?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
     */
    trusthost1?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    trusthost2?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    trusthost3?: pulumi.Input<string>;
    /**
     * User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
     */
    userType?: pulumi.Input<string>;
    /**
     * User name.
     */
    userid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FmgSystemAdminUser resource.
 */
export interface FmgSystemAdminUserArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Password.
     */
    password?: pulumi.Input<string>;
    /**
     * Profile id.
     */
    profileid?: pulumi.Input<string>;
    /**
     * RADIUS server name.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * Rpc permit, Enum: ["read-write", "none", "read"]
     */
    rpcPermit?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
     */
    trusthost1?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    trusthost2?: pulumi.Input<string>;
    /**
     * Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
     */
    trusthost3?: pulumi.Input<string>;
    /**
     * User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
     */
    userType?: pulumi.Input<string>;
    /**
     * User name.
     */
    userid: pulumi.Input<string>;
}