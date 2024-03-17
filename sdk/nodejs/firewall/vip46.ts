// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure IPv4 to IPv6 virtual IPs. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewall.Vip46("trname", {
 *     arpReply: "enable",
 *     color: 0,
 *     extip: "10.202.1.200",
 *     extport: "0-65535",
 *     fosid: 0,
 *     ldbMethod: "static",
 *     mappedip: "2001:1:1:2::200",
 *     mappedport: "0-65535",
 *     portforward: "disable",
 *     protocol: "tcp",
 *     type: "static-nat",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Firewall Vip46 can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:firewall/vip46:Vip46 labelname {{name}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:firewall/vip46:Vip46 labelname {{name}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Vip46 extends pulumi.CustomResource {
    /**
     * Get an existing Vip46 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Vip46State, opts?: pulumi.CustomResourceOptions): Vip46 {
        return new Vip46(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewall/vip46:Vip46';

    /**
     * Returns true if the given object is an instance of Vip46.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vip46 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vip46.__pulumiType;
    }

    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    public readonly arpReply!: pulumi.Output<string>;
    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    public readonly extip!: pulumi.Output<string>;
    /**
     * External service port.
     */
    public readonly extport!: pulumi.Output<string>;
    /**
     * Custom defined id.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    public readonly ldbMethod!: pulumi.Output<string>;
    /**
     * Start-mapped-IP [-end mapped-IP].
     */
    public readonly mappedip!: pulumi.Output<string>;
    /**
     * Mapped service port.
     */
    public readonly mappedport!: pulumi.Output<string>;
    /**
     * Health monitors. The structure of `monitor` block is documented below.
     */
    public readonly monitors!: pulumi.Output<outputs.firewall.Vip46Monitor[] | undefined>;
    /**
     * VIP46 name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    public readonly portforward!: pulumi.Output<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    public readonly realservers!: pulumi.Output<outputs.firewall.Vip46Realserver[] | undefined>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
     */
    public readonly srcFilters!: pulumi.Output<outputs.firewall.Vip46SrcFilter[] | undefined>;
    /**
     * Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
     */
    public readonly srcintfFilters!: pulumi.Output<outputs.firewall.Vip46SrcintfFilter[] | undefined>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    public readonly uuid!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Vip46 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Vip46Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Vip46Args | Vip46State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Vip46State | undefined;
            resourceInputs["arpReply"] = state ? state.arpReply : undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["extip"] = state ? state.extip : undefined;
            resourceInputs["extport"] = state ? state.extport : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ldbMethod"] = state ? state.ldbMethod : undefined;
            resourceInputs["mappedip"] = state ? state.mappedip : undefined;
            resourceInputs["mappedport"] = state ? state.mappedport : undefined;
            resourceInputs["monitors"] = state ? state.monitors : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portforward"] = state ? state.portforward : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["realservers"] = state ? state.realservers : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["srcFilters"] = state ? state.srcFilters : undefined;
            resourceInputs["srcintfFilters"] = state ? state.srcintfFilters : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Vip46Args | undefined;
            if ((!args || args.extip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'extip'");
            }
            if ((!args || args.mappedip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mappedip'");
            }
            resourceInputs["arpReply"] = args ? args.arpReply : undefined;
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["extip"] = args ? args.extip : undefined;
            resourceInputs["extport"] = args ? args.extport : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ldbMethod"] = args ? args.ldbMethod : undefined;
            resourceInputs["mappedip"] = args ? args.mappedip : undefined;
            resourceInputs["mappedport"] = args ? args.mappedport : undefined;
            resourceInputs["monitors"] = args ? args.monitors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portforward"] = args ? args.portforward : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["realservers"] = args ? args.realservers : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["srcFilters"] = args ? args.srcFilters : undefined;
            resourceInputs["srcintfFilters"] = args ? args.srcintfFilters : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uuid"] = args ? args.uuid : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vip46.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vip46 resources.
 */
export interface Vip46State {
    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    extip?: pulumi.Input<string>;
    /**
     * External service port.
     */
    extport?: pulumi.Input<string>;
    /**
     * Custom defined id.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Start-mapped-IP [-end mapped-IP].
     */
    mappedip?: pulumi.Input<string>;
    /**
     * Mapped service port.
     */
    mappedport?: pulumi.Input<string>;
    /**
     * Health monitors. The structure of `monitor` block is documented below.
     */
    monitors?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46Monitor>[]>;
    /**
     * VIP46 name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    portforward?: pulumi.Input<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46Realserver>[]>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
     */
    srcFilters?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46SrcFilter>[]>;
    /**
     * Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
     */
    srcintfFilters?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46SrcintfFilter>[]>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    type?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vip46 resource.
 */
export interface Vip46Args {
    /**
     * Enable ARP reply. Valid values: `disable`, `enable`.
     */
    arpReply?: pulumi.Input<string>;
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Start-external-IP [-end-external-IP].
     */
    extip: pulumi.Input<string>;
    /**
     * External service port.
     */
    extport?: pulumi.Input<string>;
    /**
     * Custom defined id.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Load balance method. Valid values: `static`, `round-robin`, `weighted`, `least-session`, `least-rtt`, `first-alive`.
     */
    ldbMethod?: pulumi.Input<string>;
    /**
     * Start-mapped-IP [-end mapped-IP].
     */
    mappedip: pulumi.Input<string>;
    /**
     * Mapped service port.
     */
    mappedport?: pulumi.Input<string>;
    /**
     * Health monitors. The structure of `monitor` block is documented below.
     */
    monitors?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46Monitor>[]>;
    /**
     * VIP46 name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable port forwarding. Valid values: `disable`, `enable`.
     */
    portforward?: pulumi.Input<string>;
    /**
     * Mapped port protocol. Valid values: `tcp`, `udp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Real servers. The structure of `realservers` block is documented below.
     */
    realservers?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46Realserver>[]>;
    /**
     * Server type. Valid values: `http`, `tcp`, `udp`, `ip`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Source IP filter (x.x.x.x/x). The structure of `srcFilter` block is documented below.
     */
    srcFilters?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46SrcFilter>[]>;
    /**
     * Interfaces to which the VIP46 applies. Separate the names with spaces. The structure of `srcintfFilter` block is documented below.
     */
    srcintfFilters?: pulumi.Input<pulumi.Input<inputs.firewall.Vip46SrcintfFilter>[]>;
    /**
     * VIP type: static NAT or server load balance. Valid values: `static-nat`, `server-load-balance`.
     */
    type?: pulumi.Input<string>;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    uuid?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}