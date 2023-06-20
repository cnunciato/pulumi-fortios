// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure device categories. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * User DeviceCategory can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/userDevicecategory:UserDevicecategory labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/userDevicecategory:UserDevicecategory labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class UserDevicecategory extends pulumi.CustomResource {
    /**
     * Get an existing UserDevicecategory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserDevicecategoryState, opts?: pulumi.CustomResourceOptions): UserDevicecategory {
        return new UserDevicecategory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/userDevicecategory:UserDevicecategory';

    /**
     * Returns true if the given object is an instance of UserDevicecategory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserDevicecategory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserDevicecategory.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Device category description.
     */
    public readonly desc!: pulumi.Output<string | undefined>;
    /**
     * Device category name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a UserDevicecategory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserDevicecategoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserDevicecategoryArgs | UserDevicecategoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserDevicecategoryState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["desc"] = state ? state.desc : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as UserDevicecategoryArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["desc"] = args ? args.desc : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserDevicecategory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserDevicecategory resources.
 */
export interface UserDevicecategoryState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Device category description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Device category name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserDevicecategory resource.
 */
export interface UserDevicecategoryArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Device category description.
     */
    desc?: pulumi.Input<string>;
    /**
     * Device category name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
