// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.SystemApiuser`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test2 = new fortios.SystemApiuserSetting("test2", {
 *     accprofile: "restAPIprofile",
 *     trusthosts: [
 *         {
 *             ipv4Trusthost: "61.149.0.0 255.255.0.0",
 *             type: "ipv4-trusthost",
 *         },
 *         {
 *             ipv4Trusthost: "22.22.0.0 255.255.0.0",
 *             type: "ipv4-trusthost",
 *         },
 *     ],
 *     vdoms: ["root"],
 * });
 * ```
 */
export class SystemApiuserSetting extends pulumi.CustomResource {
    /**
     * Get an existing SystemApiuserSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemApiuserSettingState, opts?: pulumi.CustomResourceOptions): SystemApiuserSetting {
        return new SystemApiuserSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemApiuserSetting:SystemApiuserSetting';

    /**
     * Returns true if the given object is an instance of SystemApiuserSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemApiuserSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemApiuserSetting.__pulumiType;
    }

    /**
     * Admin user access profile.
     */
    public readonly accprofile!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * User name.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly trusthosts!: pulumi.Output<outputs.SystemApiuserSettingTrusthost[]>;
    /**
     * Virtual domains.
     * * `trusthost-Type` - (Required) Trusthost type.
     * * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
     */
    public readonly vdoms!: pulumi.Output<string[]>;

    /**
     * Create a SystemApiuserSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemApiuserSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemApiuserSettingArgs | SystemApiuserSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemApiuserSettingState | undefined;
            resourceInputs["accprofile"] = state ? state.accprofile : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["trusthosts"] = state ? state.trusthosts : undefined;
            resourceInputs["vdoms"] = state ? state.vdoms : undefined;
        } else {
            const args = argsOrState as SystemApiuserSettingArgs | undefined;
            if ((!args || args.accprofile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accprofile'");
            }
            if ((!args || args.trusthosts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trusthosts'");
            }
            if ((!args || args.vdoms === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vdoms'");
            }
            resourceInputs["accprofile"] = args ? args.accprofile : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["trusthosts"] = args ? args.trusthosts : undefined;
            resourceInputs["vdoms"] = args ? args.vdoms : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemApiuserSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemApiuserSetting resources.
 */
export interface SystemApiuserSettingState {
    /**
     * Admin user access profile.
     */
    accprofile?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    trusthosts?: pulumi.Input<pulumi.Input<inputs.SystemApiuserSettingTrusthost>[]>;
    /**
     * Virtual domains.
     * * `trusthost-Type` - (Required) Trusthost type.
     * * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
     */
    vdoms?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SystemApiuserSetting resource.
 */
export interface SystemApiuserSettingArgs {
    /**
     * Admin user access profile.
     */
    accprofile: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * User name.
     */
    name?: pulumi.Input<string>;
    trusthosts: pulumi.Input<pulumi.Input<inputs.SystemApiuserSettingTrusthost>[]>;
    /**
     * Virtual domains.
     * * `trusthost-Type` - (Required) Trusthost type.
     * * `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
     */
    vdoms: pulumi.Input<pulumi.Input<string>[]>;
}
