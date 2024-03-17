// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure virtual hardware switch interfaces. Applies to FortiOS Version `7.0.4`.
 *
 * ## Import
 *
 * System VirtualSwitch can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/virtualswitch:Virtualswitch labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/virtualswitch:Virtualswitch labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Virtualswitch extends pulumi.CustomResource {
    /**
     * Get an existing Virtualswitch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualswitchState, opts?: pulumi.CustomResourceOptions): Virtualswitch {
        return new Virtualswitch(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/virtualswitch:Virtualswitch';

    /**
     * Returns true if the given object is an instance of Virtualswitch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Virtualswitch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Virtualswitch.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Name of the virtual switch.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Physical switch parent.
     */
    public readonly physicalSwitch!: pulumi.Output<string>;
    /**
     * Configure member ports. The structure of `port` block is documented below.
     */
    public readonly ports!: pulumi.Output<outputs.system.VirtualswitchPort[] | undefined>;
    /**
     * Enable/disable SPAN. Valid values: `disable`, `enable`.
     */
    public readonly span!: pulumi.Output<string>;
    /**
     * SPAN destination port.
     */
    public readonly spanDestPort!: pulumi.Output<string>;
    /**
     * SPAN direction. Valid values: `rx`, `tx`, `both`.
     */
    public readonly spanDirection!: pulumi.Output<string>;
    /**
     * SPAN source port.
     */
    public readonly spanSourcePort!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * VLAN.
     */
    public readonly vlan!: pulumi.Output<number>;

    /**
     * Create a Virtualswitch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VirtualswitchArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualswitchArgs | VirtualswitchState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualswitchState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["physicalSwitch"] = state ? state.physicalSwitch : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["span"] = state ? state.span : undefined;
            resourceInputs["spanDestPort"] = state ? state.spanDestPort : undefined;
            resourceInputs["spanDirection"] = state ? state.spanDirection : undefined;
            resourceInputs["spanSourcePort"] = state ? state.spanSourcePort : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as VirtualswitchArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["physicalSwitch"] = args ? args.physicalSwitch : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["span"] = args ? args.span : undefined;
            resourceInputs["spanDestPort"] = args ? args.spanDestPort : undefined;
            resourceInputs["spanDirection"] = args ? args.spanDirection : undefined;
            resourceInputs["spanSourcePort"] = args ? args.spanSourcePort : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Virtualswitch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Virtualswitch resources.
 */
export interface VirtualswitchState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of the virtual switch.
     */
    name?: pulumi.Input<string>;
    /**
     * Physical switch parent.
     */
    physicalSwitch?: pulumi.Input<string>;
    /**
     * Configure member ports. The structure of `port` block is documented below.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.system.VirtualswitchPort>[]>;
    /**
     * Enable/disable SPAN. Valid values: `disable`, `enable`.
     */
    span?: pulumi.Input<string>;
    /**
     * SPAN destination port.
     */
    spanDestPort?: pulumi.Input<string>;
    /**
     * SPAN direction. Valid values: `rx`, `tx`, `both`.
     */
    spanDirection?: pulumi.Input<string>;
    /**
     * SPAN source port.
     */
    spanSourcePort?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN.
     */
    vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Virtualswitch resource.
 */
export interface VirtualswitchArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Name of the virtual switch.
     */
    name?: pulumi.Input<string>;
    /**
     * Physical switch parent.
     */
    physicalSwitch?: pulumi.Input<string>;
    /**
     * Configure member ports. The structure of `port` block is documented below.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.system.VirtualswitchPort>[]>;
    /**
     * Enable/disable SPAN. Valid values: `disable`, `enable`.
     */
    span?: pulumi.Input<string>;
    /**
     * SPAN destination port.
     */
    spanDestPort?: pulumi.Input<string>;
    /**
     * SPAN direction. Valid values: `rx`, `tx`, `both`.
     */
    spanDirection?: pulumi.Input<string>;
    /**
     * SPAN source port.
     */
    spanSourcePort?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * VLAN.
     */
    vlan?: pulumi.Input<number>;
}