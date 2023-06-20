// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure physical switches. Applies to FortiOS Version `7.0.4`.
 *
 * ## Import
 *
 * System PhysicalSwitch can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemPhysicalswitch:SystemPhysicalswitch labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemPhysicalswitch:SystemPhysicalswitch labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemPhysicalswitch extends pulumi.CustomResource {
    /**
     * Get an existing SystemPhysicalswitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemPhysicalswitchState, opts?: pulumi.CustomResourceOptions): SystemPhysicalswitch {
        return new SystemPhysicalswitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemPhysicalswitch:SystemPhysicalswitch';

    /**
     * Returns true if the given object is an instance of SystemPhysicalswitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemPhysicalswitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemPhysicalswitch.__pulumiType;
    }

    /**
     * Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
     */
    public readonly ageEnable!: pulumi.Output<string>;
    /**
     * Layer 2 table age timer value.
     */
    public readonly ageVal!: pulumi.Output<number>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemPhysicalswitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemPhysicalswitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemPhysicalswitchArgs | SystemPhysicalswitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemPhysicalswitchState | undefined;
            resourceInputs["ageEnable"] = state ? state.ageEnable : undefined;
            resourceInputs["ageVal"] = state ? state.ageVal : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemPhysicalswitchArgs | undefined;
            resourceInputs["ageEnable"] = args ? args.ageEnable : undefined;
            resourceInputs["ageVal"] = args ? args.ageVal : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemPhysicalswitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemPhysicalswitch resources.
 */
export interface SystemPhysicalswitchState {
    /**
     * Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
     */
    ageEnable?: pulumi.Input<string>;
    /**
     * Layer 2 table age timer value.
     */
    ageVal?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemPhysicalswitch resource.
 */
export interface SystemPhysicalswitchArgs {
    /**
     * Enable/disable layer 2 age timer. Valid values: `enable`, `disable`.
     */
    ageEnable?: pulumi.Input<string>;
    /**
     * Layer 2 table age timer value.
     */
    ageVal?: pulumi.Input<number>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
