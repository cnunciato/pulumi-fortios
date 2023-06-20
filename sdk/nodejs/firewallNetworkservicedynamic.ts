// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure Dynamic Network Services. Applies to FortiOS Version `>= 7.2.1`.
 *
 * ## Import
 *
 * Firewall NetworkServiceDynamic can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallNetworkservicedynamic:FirewallNetworkservicedynamic labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallNetworkservicedynamic:FirewallNetworkservicedynamic labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallNetworkservicedynamic extends pulumi.CustomResource {
    /**
     * Get an existing FirewallNetworkservicedynamic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallNetworkservicedynamicState, opts?: pulumi.CustomResourceOptions): FirewallNetworkservicedynamic {
        return new FirewallNetworkservicedynamic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallNetworkservicedynamic:FirewallNetworkservicedynamic';

    /**
     * Returns true if the given object is an instance of FirewallNetworkservicedynamic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallNetworkservicedynamic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallNetworkservicedynamic.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Match criteria filter.
     */
    public readonly filter!: pulumi.Output<string | undefined>;
    /**
     * Dynamic Network Service name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SDN connector name.
     */
    public readonly sdn!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallNetworkservicedynamic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallNetworkservicedynamicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallNetworkservicedynamicArgs | FirewallNetworkservicedynamicState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallNetworkservicedynamicState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sdn"] = state ? state.sdn : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallNetworkservicedynamicArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sdn"] = args ? args.sdn : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallNetworkservicedynamic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallNetworkservicedynamic resources.
 */
export interface FirewallNetworkservicedynamicState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Match criteria filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Dynamic Network Service name.
     */
    name?: pulumi.Input<string>;
    /**
     * SDN connector name.
     */
    sdn?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallNetworkservicedynamic resource.
 */
export interface FirewallNetworkservicedynamicArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Match criteria filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Dynamic Network Service name.
     */
    name?: pulumi.Input<string>;
    /**
     * SDN connector name.
     */
    sdn?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
