// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure redundant internet connections using SD-WAN (formerly virtual WAN link). Applies to FortiOS Version `<= 6.4.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SystemVirtualwanlink("trname", {
 *     failDetect: "disable",
 *     loadBalanceMode: "source-ip-based",
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * System VirtualWanLink can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemVirtualwanlink:SystemVirtualwanlink labelname SystemVirtualWanLink
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemVirtualwanlink:SystemVirtualwanlink labelname SystemVirtualWanLink
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemVirtualwanlink extends pulumi.CustomResource {
    /**
     * Get an existing SystemVirtualwanlink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemVirtualwanlinkState, opts?: pulumi.CustomResourceOptions): SystemVirtualwanlink {
        return new SystemVirtualwanlink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemVirtualwanlink:SystemVirtualwanlink';

    /**
     * Returns true if the given object is an instance of SystemVirtualwanlink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemVirtualwanlink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemVirtualwanlink.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
     */
    public readonly failAlertInterfaces!: pulumi.Output<outputs.SystemVirtualwanlinkFailAlertInterface[] | undefined>;
    /**
     * Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
     */
    public readonly failDetect!: pulumi.Output<string>;
    /**
     * SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
     */
    public readonly healthChecks!: pulumi.Output<outputs.SystemVirtualwanlinkHealthCheck[] | undefined>;
    /**
     * Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
     */
    public readonly loadBalanceMode!: pulumi.Output<string>;
    /**
     * Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.SystemVirtualwanlinkMember[] | undefined>;
    /**
     * Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
     */
    public readonly neighborHoldBootTime!: pulumi.Output<number>;
    /**
     * Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
     */
    public readonly neighborHoldDown!: pulumi.Output<string>;
    /**
     * Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
     */
    public readonly neighborHoldDownTime!: pulumi.Output<number>;
    /**
     * Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
     */
    public readonly neighbors!: pulumi.Output<outputs.SystemVirtualwanlinkNeighbor[] | undefined>;
    /**
     * Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
     */
    public readonly services!: pulumi.Output<outputs.SystemVirtualwanlinkService[] | undefined>;
    /**
     * Enable/disable SD-WAN. Valid values: `disable`, `enable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Configure SD-WAN zones. The structure of `zone` block is documented below.
     */
    public readonly zones!: pulumi.Output<outputs.SystemVirtualwanlinkZone[] | undefined>;

    /**
     * Create a SystemVirtualwanlink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemVirtualwanlinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemVirtualwanlinkArgs | SystemVirtualwanlinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemVirtualwanlinkState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["failAlertInterfaces"] = state ? state.failAlertInterfaces : undefined;
            resourceInputs["failDetect"] = state ? state.failDetect : undefined;
            resourceInputs["healthChecks"] = state ? state.healthChecks : undefined;
            resourceInputs["loadBalanceMode"] = state ? state.loadBalanceMode : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["neighborHoldBootTime"] = state ? state.neighborHoldBootTime : undefined;
            resourceInputs["neighborHoldDown"] = state ? state.neighborHoldDown : undefined;
            resourceInputs["neighborHoldDownTime"] = state ? state.neighborHoldDownTime : undefined;
            resourceInputs["neighbors"] = state ? state.neighbors : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as SystemVirtualwanlinkArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["failAlertInterfaces"] = args ? args.failAlertInterfaces : undefined;
            resourceInputs["failDetect"] = args ? args.failDetect : undefined;
            resourceInputs["healthChecks"] = args ? args.healthChecks : undefined;
            resourceInputs["loadBalanceMode"] = args ? args.loadBalanceMode : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["neighborHoldBootTime"] = args ? args.neighborHoldBootTime : undefined;
            resourceInputs["neighborHoldDown"] = args ? args.neighborHoldDown : undefined;
            resourceInputs["neighborHoldDownTime"] = args ? args.neighborHoldDownTime : undefined;
            resourceInputs["neighbors"] = args ? args.neighbors : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemVirtualwanlink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemVirtualwanlink resources.
 */
export interface SystemVirtualwanlinkState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
     */
    failAlertInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkFailAlertInterface>[]>;
    /**
     * Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
     */
    failDetect?: pulumi.Input<string>;
    /**
     * SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
     */
    healthChecks?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkHealthCheck>[]>;
    /**
     * Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
     */
    loadBalanceMode?: pulumi.Input<string>;
    /**
     * Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkMember>[]>;
    /**
     * Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
     */
    neighborHoldBootTime?: pulumi.Input<number>;
    /**
     * Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
     */
    neighborHoldDown?: pulumi.Input<string>;
    /**
     * Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
     */
    neighborHoldDownTime?: pulumi.Input<number>;
    /**
     * Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkNeighbor>[]>;
    /**
     * Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkService>[]>;
    /**
     * Enable/disable SD-WAN. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure SD-WAN zones. The structure of `zone` block is documented below.
     */
    zones?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkZone>[]>;
}

/**
 * The set of arguments for constructing a SystemVirtualwanlink resource.
 */
export interface SystemVirtualwanlinkArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Physical interfaces that will be alerted. The structure of `failAlertInterfaces` block is documented below.
     */
    failAlertInterfaces?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkFailAlertInterface>[]>;
    /**
     * Enable/disable SD-WAN Internet connection status checking (failure detection). Valid values: `enable`, `disable`.
     */
    failDetect?: pulumi.Input<string>;
    /**
     * SD-WAN status checking or health checking. Identify a server on the Internet and determine how SD-WAN verifies that the FortiGate can communicate with it. The structure of `healthCheck` block is documented below.
     */
    healthChecks?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkHealthCheck>[]>;
    /**
     * Algorithm or mode to use for load balancing Internet traffic to SD-WAN members. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`, `measured-volume-based`.
     */
    loadBalanceMode?: pulumi.Input<string>;
    /**
     * Physical FortiGate interfaces added to the virtual-wan-link. The structure of `members` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkMember>[]>;
    /**
     * Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
     */
    neighborHoldBootTime?: pulumi.Input<number>;
    /**
     * Enable/disable hold switching from the secondary neighbor to the primary neighbor. Valid values: `enable`, `disable`.
     */
    neighborHoldDown?: pulumi.Input<string>;
    /**
     * Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
     */
    neighborHoldDownTime?: pulumi.Input<number>;
    /**
     * Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
     */
    neighbors?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkNeighbor>[]>;
    /**
     * Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
     */
    services?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkService>[]>;
    /**
     * Enable/disable SD-WAN. Valid values: `disable`, `enable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Configure SD-WAN zones. The structure of `zone` block is documented below.
     */
    zones?: pulumi.Input<pulumi.Input<inputs.SystemVirtualwanlinkZone>[]>;
}
