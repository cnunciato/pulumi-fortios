// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure NSX-T setting. Applies to FortiOS Version `>= 6.4.10`.
 *
 * ## Import
 *
 * Nsxt Setting can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class NsxtSetting extends pulumi.CustomResource {
    /**
     * Get an existing NsxtSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtSettingState, opts?: pulumi.CustomResourceOptions): NsxtSetting {
        return new NsxtSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/nsxtSetting:NsxtSetting';

    /**
     * Returns true if the given object is an instance of NsxtSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtSetting.__pulumiType;
    }

    /**
     * Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
     */
    public readonly liveness!: pulumi.Output<string>;
    /**
     * Service name.
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a NsxtSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NsxtSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtSettingArgs | NsxtSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtSettingState | undefined;
            resourceInputs["liveness"] = state ? state.liveness : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as NsxtSettingArgs | undefined;
            resourceInputs["liveness"] = args ? args.liveness : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtSetting resources.
 */
export interface NsxtSettingState {
    /**
     * Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
     */
    liveness?: pulumi.Input<string>;
    /**
     * Service name.
     */
    service?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtSetting resource.
 */
export interface NsxtSettingArgs {
    /**
     * Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
     */
    liveness?: pulumi.Input<string>;
    /**
     * Service name.
     */
    service?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
