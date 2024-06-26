// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Filters for memory buffer.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.log.memory.Filter("trname", {
 *     anomaly: "enable",
 *     dns: "enable",
 *     filterType: "include",
 *     forwardTraffic: "enable",
 *     gtp: "enable",
 *     localTraffic: "enable",
 *     multicastTraffic: "enable",
 *     severity: "information",
 *     snifferTraffic: "enable",
 *     ssh: "enable",
 *     voip: "enable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * LogMemory Filter can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:log/memory/filter:Filter labelname LogMemoryFilter
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:log/memory/filter:Filter labelname LogMemoryFilter
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Filter extends pulumi.CustomResource {
    /**
     * Get an existing Filter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FilterState, opts?: pulumi.CustomResourceOptions): Filter {
        return new Filter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:log/memory/filter:Filter';

    /**
     * Returns true if the given object is an instance of Filter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Filter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Filter.__pulumiType;
    }

    /**
     * Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
     */
    public readonly admin!: pulumi.Output<string>;
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    public readonly anomaly!: pulumi.Output<string>;
    /**
     * Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
     */
    public readonly auth!: pulumi.Output<string>;
    /**
     * Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
     */
    public readonly cpuMemoryUsage!: pulumi.Output<string>;
    /**
     * Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
     */
    public readonly dhcp!: pulumi.Output<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    public readonly dns!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    public readonly event!: pulumi.Output<string>;
    /**
     * Memory log filter.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    public readonly filterType!: pulumi.Output<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    public readonly fortiSwitch!: pulumi.Output<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly forwardTraffic!: pulumi.Output<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    public readonly freeStyles!: pulumi.Output<outputs.log.memory.FilterFreeStyle[] | undefined>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    public readonly getAllTables!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    public readonly gtp!: pulumi.Output<string>;
    /**
     * Enable/disable HA logging. Valid values: `enable`, `disable`.
     */
    public readonly ha!: pulumi.Output<string>;
    /**
     * Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
     */
    public readonly ipsec!: pulumi.Output<string>;
    /**
     * Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
     */
    public readonly ldbMonitor!: pulumi.Output<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly localTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly multicastTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    public readonly netscanDiscovery!: pulumi.Output<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    public readonly netscanVulnerability!: pulumi.Output<string>;
    /**
     * Enable/disable pattern update logging. Valid values: `enable`, `disable`.
     */
    public readonly pattern!: pulumi.Output<string>;
    /**
     * Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
     */
    public readonly ppp!: pulumi.Output<string>;
    /**
     * Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
     */
    public readonly radius!: pulumi.Output<string>;
    /**
     * Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly snifferTraffic!: pulumi.Output<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    public readonly ssh!: pulumi.Output<string>;
    /**
     * Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
     */
    public readonly sslvpnLogAdm!: pulumi.Output<string>;
    /**
     * Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
     */
    public readonly sslvpnLogAuth!: pulumi.Output<string>;
    /**
     * Enable/disable SSL session logging. Valid values: `enable`, `disable`.
     */
    public readonly sslvpnLogSession!: pulumi.Output<string>;
    /**
     * Enable/disable system activity logging. Valid values: `enable`, `disable`.
     */
    public readonly system!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
     */
    public readonly vipSsl!: pulumi.Output<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    public readonly voip!: pulumi.Output<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    public readonly wanOpt!: pulumi.Output<string>;
    /**
     * Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
     */
    public readonly wirelessActivity!: pulumi.Output<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    public readonly ztnaTraffic!: pulumi.Output<string>;

    /**
     * Create a Filter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FilterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FilterArgs | FilterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FilterState | undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["anomaly"] = state ? state.anomaly : undefined;
            resourceInputs["auth"] = state ? state.auth : undefined;
            resourceInputs["cpuMemoryUsage"] = state ? state.cpuMemoryUsage : undefined;
            resourceInputs["dhcp"] = state ? state.dhcp : undefined;
            resourceInputs["dns"] = state ? state.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["event"] = state ? state.event : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["filterType"] = state ? state.filterType : undefined;
            resourceInputs["fortiSwitch"] = state ? state.fortiSwitch : undefined;
            resourceInputs["forwardTraffic"] = state ? state.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = state ? state.freeStyles : undefined;
            resourceInputs["getAllTables"] = state ? state.getAllTables : undefined;
            resourceInputs["gtp"] = state ? state.gtp : undefined;
            resourceInputs["ha"] = state ? state.ha : undefined;
            resourceInputs["ipsec"] = state ? state.ipsec : undefined;
            resourceInputs["ldbMonitor"] = state ? state.ldbMonitor : undefined;
            resourceInputs["localTraffic"] = state ? state.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = state ? state.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = state ? state.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = state ? state.netscanVulnerability : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["ppp"] = state ? state.ppp : undefined;
            resourceInputs["radius"] = state ? state.radius : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["snifferTraffic"] = state ? state.snifferTraffic : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["sslvpnLogAdm"] = state ? state.sslvpnLogAdm : undefined;
            resourceInputs["sslvpnLogAuth"] = state ? state.sslvpnLogAuth : undefined;
            resourceInputs["sslvpnLogSession"] = state ? state.sslvpnLogSession : undefined;
            resourceInputs["system"] = state ? state.system : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vipSsl"] = state ? state.vipSsl : undefined;
            resourceInputs["voip"] = state ? state.voip : undefined;
            resourceInputs["wanOpt"] = state ? state.wanOpt : undefined;
            resourceInputs["wirelessActivity"] = state ? state.wirelessActivity : undefined;
            resourceInputs["ztnaTraffic"] = state ? state.ztnaTraffic : undefined;
        } else {
            const args = argsOrState as FilterArgs | undefined;
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["anomaly"] = args ? args.anomaly : undefined;
            resourceInputs["auth"] = args ? args.auth : undefined;
            resourceInputs["cpuMemoryUsage"] = args ? args.cpuMemoryUsage : undefined;
            resourceInputs["dhcp"] = args ? args.dhcp : undefined;
            resourceInputs["dns"] = args ? args.dns : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["event"] = args ? args.event : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["filterType"] = args ? args.filterType : undefined;
            resourceInputs["fortiSwitch"] = args ? args.fortiSwitch : undefined;
            resourceInputs["forwardTraffic"] = args ? args.forwardTraffic : undefined;
            resourceInputs["freeStyles"] = args ? args.freeStyles : undefined;
            resourceInputs["getAllTables"] = args ? args.getAllTables : undefined;
            resourceInputs["gtp"] = args ? args.gtp : undefined;
            resourceInputs["ha"] = args ? args.ha : undefined;
            resourceInputs["ipsec"] = args ? args.ipsec : undefined;
            resourceInputs["ldbMonitor"] = args ? args.ldbMonitor : undefined;
            resourceInputs["localTraffic"] = args ? args.localTraffic : undefined;
            resourceInputs["multicastTraffic"] = args ? args.multicastTraffic : undefined;
            resourceInputs["netscanDiscovery"] = args ? args.netscanDiscovery : undefined;
            resourceInputs["netscanVulnerability"] = args ? args.netscanVulnerability : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["ppp"] = args ? args.ppp : undefined;
            resourceInputs["radius"] = args ? args.radius : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["snifferTraffic"] = args ? args.snifferTraffic : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["sslvpnLogAdm"] = args ? args.sslvpnLogAdm : undefined;
            resourceInputs["sslvpnLogAuth"] = args ? args.sslvpnLogAuth : undefined;
            resourceInputs["sslvpnLogSession"] = args ? args.sslvpnLogSession : undefined;
            resourceInputs["system"] = args ? args.system : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vipSsl"] = args ? args.vipSsl : undefined;
            resourceInputs["voip"] = args ? args.voip : undefined;
            resourceInputs["wanOpt"] = args ? args.wanOpt : undefined;
            resourceInputs["wirelessActivity"] = args ? args.wirelessActivity : undefined;
            resourceInputs["ztnaTraffic"] = args ? args.ztnaTraffic : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Filter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Filter resources.
 */
export interface FilterState {
    /**
     * Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
     */
    admin?: pulumi.Input<string>;
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
     */
    auth?: pulumi.Input<string>;
    /**
     * Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
     */
    cpuMemoryUsage?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
     */
    dhcp?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    event?: pulumi.Input<string>;
    /**
     * Memory log filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    fortiSwitch?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.log.memory.FilterFreeStyle>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable HA logging. Valid values: `enable`, `disable`.
     */
    ha?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
     */
    ipsec?: pulumi.Input<string>;
    /**
     * Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
     */
    ldbMonitor?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Enable/disable pattern update logging. Valid values: `enable`, `disable`.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
     */
    ppp?: pulumi.Input<string>;
    /**
     * Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
     */
    radius?: pulumi.Input<string>;
    /**
     * Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogAdm?: pulumi.Input<string>;
    /**
     * Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogAuth?: pulumi.Input<string>;
    /**
     * Enable/disable SSL session logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogSession?: pulumi.Input<string>;
    /**
     * Enable/disable system activity logging. Valid values: `enable`, `disable`.
     */
    system?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
     */
    vipSsl?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    wanOpt?: pulumi.Input<string>;
    /**
     * Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
     */
    wirelessActivity?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Filter resource.
 */
export interface FilterArgs {
    /**
     * Enable/disable admin login/logout logging. Valid values: `enable`, `disable`.
     */
    admin?: pulumi.Input<string>;
    /**
     * Enable/disable anomaly logging. Valid values: `enable`, `disable`.
     */
    anomaly?: pulumi.Input<string>;
    /**
     * Enable/disable firewall authentication logging. Valid values: `enable`, `disable`.
     */
    auth?: pulumi.Input<string>;
    /**
     * Enable/disable CPU & memory usage logging every 5 minutes. Valid values: `enable`, `disable`.
     */
    cpuMemoryUsage?: pulumi.Input<string>;
    /**
     * Enable/disable DHCP service messages logging. Valid values: `enable`, `disable`.
     */
    dhcp?: pulumi.Input<string>;
    /**
     * Enable/disable detailed DNS event logging. Valid values: `enable`, `disable`.
     */
    dns?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable event logging. Valid values: `enable`, `disable`.
     */
    event?: pulumi.Input<string>;
    /**
     * Memory log filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * Include/exclude logs that match the filter. Valid values: `include`, `exclude`.
     */
    filterType?: pulumi.Input<string>;
    /**
     * Enable/disable Forti-Switch logging. Valid values: `enable`, `disable`.
     */
    fortiSwitch?: pulumi.Input<string>;
    /**
     * Enable/disable forward traffic logging. Valid values: `enable`, `disable`.
     */
    forwardTraffic?: pulumi.Input<string>;
    /**
     * Free Style Filters The structure of `freeStyle` block is documented below.
     */
    freeStyles?: pulumi.Input<pulumi.Input<inputs.log.memory.FilterFreeStyle>[]>;
    /**
     * Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables.
     */
    getAllTables?: pulumi.Input<string>;
    /**
     * Enable/disable GTP messages logging. Valid values: `enable`, `disable`.
     */
    gtp?: pulumi.Input<string>;
    /**
     * Enable/disable HA logging. Valid values: `enable`, `disable`.
     */
    ha?: pulumi.Input<string>;
    /**
     * Enable/disable IPsec negotiation messages logging. Valid values: `enable`, `disable`.
     */
    ipsec?: pulumi.Input<string>;
    /**
     * Enable/disable VIP real server health monitoring logging. Valid values: `enable`, `disable`.
     */
    ldbMonitor?: pulumi.Input<string>;
    /**
     * Enable/disable local in or out traffic logging. Valid values: `enable`, `disable`.
     */
    localTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable multicast traffic logging. Valid values: `enable`, `disable`.
     */
    multicastTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable netscan discovery event logging.
     */
    netscanDiscovery?: pulumi.Input<string>;
    /**
     * Enable/disable netscan vulnerability event logging.
     */
    netscanVulnerability?: pulumi.Input<string>;
    /**
     * Enable/disable pattern update logging. Valid values: `enable`, `disable`.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Enable/disable L2TP/PPTP/PPPoE logging. Valid values: `enable`, `disable`.
     */
    ppp?: pulumi.Input<string>;
    /**
     * Enable/disable RADIUS messages logging. Valid values: `enable`, `disable`.
     */
    radius?: pulumi.Input<string>;
    /**
     * Log every message above and including this severity level. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
     */
    severity?: pulumi.Input<string>;
    /**
     * Enable/disable sniffer traffic logging. Valid values: `enable`, `disable`.
     */
    snifferTraffic?: pulumi.Input<string>;
    /**
     * Enable/disable SSH logging. Valid values: `enable`, `disable`.
     */
    ssh?: pulumi.Input<string>;
    /**
     * Enable/disable SSL administrator login logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogAdm?: pulumi.Input<string>;
    /**
     * Enable/disable SSL user authentication logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogAuth?: pulumi.Input<string>;
    /**
     * Enable/disable SSL session logging. Valid values: `enable`, `disable`.
     */
    sslvpnLogSession?: pulumi.Input<string>;
    /**
     * Enable/disable system activity logging. Valid values: `enable`, `disable`.
     */
    system?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable VIP SSL logging. Valid values: `enable`, `disable`.
     */
    vipSsl?: pulumi.Input<string>;
    /**
     * Enable/disable VoIP logging. Valid values: `enable`, `disable`.
     */
    voip?: pulumi.Input<string>;
    /**
     * Enable/disable WAN optimization event logging. Valid values: `enable`, `disable`.
     */
    wanOpt?: pulumi.Input<string>;
    /**
     * Enable/disable wireless activity event logging. Valid values: `enable`, `disable`.
     */
    wirelessActivity?: pulumi.Input<string>;
    /**
     * Enable/disable ztna traffic logging. Valid values: `enable`, `disable`.
     */
    ztnaTraffic?: pulumi.Input<string>;
}
