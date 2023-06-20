// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure decrypted traffic mirror. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * Firewall DecryptedTrafficMirror can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallDecryptedtrafficmirror:FirewallDecryptedtrafficmirror labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallDecryptedtrafficmirror:FirewallDecryptedtrafficmirror labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallDecryptedtrafficmirror extends pulumi.CustomResource {
    /**
     * Get an existing FirewallDecryptedtrafficmirror resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallDecryptedtrafficmirrorState, opts?: pulumi.CustomResourceOptions): FirewallDecryptedtrafficmirror {
        return new FirewallDecryptedtrafficmirror(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallDecryptedtrafficmirror:FirewallDecryptedtrafficmirror';

    /**
     * Returns true if the given object is an instance of FirewallDecryptedtrafficmirror.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallDecryptedtrafficmirror {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallDecryptedtrafficmirror.__pulumiType;
    }

    /**
     * Set destination MAC address for mirrored traffic.
     */
    public readonly dstmac!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Decrypted traffic mirror interface The structure of `interface` block is documented below.
     */
    public readonly interfaces!: pulumi.Output<outputs.FirewallDecryptedtrafficmirrorInterface[] | undefined>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
     */
    public readonly trafficSource!: pulumi.Output<string>;
    /**
     * Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
     */
    public readonly trafficType!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallDecryptedtrafficmirror resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallDecryptedtrafficmirrorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallDecryptedtrafficmirrorArgs | FirewallDecryptedtrafficmirrorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallDecryptedtrafficmirrorState | undefined;
            resourceInputs["dstmac"] = state ? state.dstmac : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = state ? state.interfaces : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["trafficSource"] = state ? state.trafficSource : undefined;
            resourceInputs["trafficType"] = state ? state.trafficType : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallDecryptedtrafficmirrorArgs | undefined;
            resourceInputs["dstmac"] = args ? args.dstmac : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["interfaces"] = args ? args.interfaces : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["trafficSource"] = args ? args.trafficSource : undefined;
            resourceInputs["trafficType"] = args ? args.trafficType : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallDecryptedtrafficmirror.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallDecryptedtrafficmirror resources.
 */
export interface FirewallDecryptedtrafficmirrorState {
    /**
     * Set destination MAC address for mirrored traffic.
     */
    dstmac?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror interface The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.FirewallDecryptedtrafficmirrorInterface>[]>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
     */
    trafficSource?: pulumi.Input<string>;
    /**
     * Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallDecryptedtrafficmirror resource.
 */
export interface FirewallDecryptedtrafficmirrorArgs {
    /**
     * Set destination MAC address for mirrored traffic.
     */
    dstmac?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Decrypted traffic mirror interface The structure of `interface` block is documented below.
     */
    interfaces?: pulumi.Input<pulumi.Input<inputs.FirewallDecryptedtrafficmirrorInterface>[]>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Source of decrypted traffic to be mirrored. Valid values: `client`, `server`, `both`.
     */
    trafficSource?: pulumi.Input<string>;
    /**
     * Types of decrypted traffic to be mirrored. Valid values: `ssl`, `ssh`.
     */
    trafficType?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
