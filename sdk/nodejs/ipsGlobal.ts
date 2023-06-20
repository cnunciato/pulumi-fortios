// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure IPS global parameter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.IpsGlobal("trname", {
 *     anomalyMode: "continuous",
 *     database: "regular",
 *     deepAppInspDbLimit: 0,
 *     deepAppInspTimeout: 0,
 *     engineCount: 0,
 *     excludeSignatures: "industrial",
 *     failOpen: "disable",
 *     intelligentMode: "enable",
 *     sessionLimitMode: "heuristic",
 *     socketSize: 0,
 *     syncSessionTtl: "enable",
 *     trafficSubmit: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * Ips Global can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/ipsGlobal:IpsGlobal labelname IpsGlobal
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/ipsGlobal:IpsGlobal labelname IpsGlobal
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class IpsGlobal extends pulumi.CustomResource {
    /**
     * Get an existing IpsGlobal resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsGlobalState, opts?: pulumi.CustomResourceOptions): IpsGlobal {
        return new IpsGlobal(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/ipsGlobal:IpsGlobal';

    /**
     * Returns true if the given object is an instance of IpsGlobal.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsGlobal {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsGlobal.__pulumiType;
    }

    /**
     * Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
     */
    public readonly anomalyMode!: pulumi.Output<string>;
    /**
     * IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
     */
    public readonly cpAccelMode!: pulumi.Output<string>;
    /**
     * Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
     */
    public readonly deepAppInspDbLimit!: pulumi.Output<number>;
    /**
     * Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
     */
    public readonly deepAppInspTimeout!: pulumi.Output<number>;
    /**
     * Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
     */
    public readonly engineCount!: pulumi.Output<number>;
    /**
     * Excluded signatures. Valid values: `none`, `industrial`.
     */
    public readonly excludeSignatures!: pulumi.Output<string>;
    /**
     * Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
     */
    public readonly failOpen!: pulumi.Output<string>;
    /**
     * Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
     */
    public readonly intelligentMode!: pulumi.Output<string>;
    /**
     * Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
     */
    public readonly ipsReserveCpu!: pulumi.Output<string>;
    /**
     * NGFW policy-mode app detection threshold.
     */
    public readonly ngfwMaxScanRange!: pulumi.Output<number>;
    /**
     * Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
     */
    public readonly npAccelMode!: pulumi.Output<string>;
    /**
     * Packet/pcap log queue depth per IPS engine.
     */
    public readonly packetLogQueueDepth!: pulumi.Output<number>;
    /**
     * Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
     */
    public readonly sessionLimitMode!: pulumi.Output<string>;
    /**
     * Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
     */
    public readonly skypeClientPublicIpaddr!: pulumi.Output<string | undefined>;
    /**
     * IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
     */
    public readonly socketSize!: pulumi.Output<number>;
    /**
     * Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
     */
    public readonly syncSessionTtl!: pulumi.Output<string>;
    /**
     * TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
     */
    public readonly tlsActiveProbe!: pulumi.Output<outputs.IpsGlobalTlsActiveProbe>;
    /**
     * Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
     */
    public readonly trafficSubmit!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a IpsGlobal resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpsGlobalArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsGlobalArgs | IpsGlobalState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsGlobalState | undefined;
            resourceInputs["anomalyMode"] = state ? state.anomalyMode : undefined;
            resourceInputs["cpAccelMode"] = state ? state.cpAccelMode : undefined;
            resourceInputs["database"] = state ? state.database : undefined;
            resourceInputs["deepAppInspDbLimit"] = state ? state.deepAppInspDbLimit : undefined;
            resourceInputs["deepAppInspTimeout"] = state ? state.deepAppInspTimeout : undefined;
            resourceInputs["engineCount"] = state ? state.engineCount : undefined;
            resourceInputs["excludeSignatures"] = state ? state.excludeSignatures : undefined;
            resourceInputs["failOpen"] = state ? state.failOpen : undefined;
            resourceInputs["intelligentMode"] = state ? state.intelligentMode : undefined;
            resourceInputs["ipsReserveCpu"] = state ? state.ipsReserveCpu : undefined;
            resourceInputs["ngfwMaxScanRange"] = state ? state.ngfwMaxScanRange : undefined;
            resourceInputs["npAccelMode"] = state ? state.npAccelMode : undefined;
            resourceInputs["packetLogQueueDepth"] = state ? state.packetLogQueueDepth : undefined;
            resourceInputs["sessionLimitMode"] = state ? state.sessionLimitMode : undefined;
            resourceInputs["skypeClientPublicIpaddr"] = state ? state.skypeClientPublicIpaddr : undefined;
            resourceInputs["socketSize"] = state ? state.socketSize : undefined;
            resourceInputs["syncSessionTtl"] = state ? state.syncSessionTtl : undefined;
            resourceInputs["tlsActiveProbe"] = state ? state.tlsActiveProbe : undefined;
            resourceInputs["trafficSubmit"] = state ? state.trafficSubmit : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as IpsGlobalArgs | undefined;
            resourceInputs["anomalyMode"] = args ? args.anomalyMode : undefined;
            resourceInputs["cpAccelMode"] = args ? args.cpAccelMode : undefined;
            resourceInputs["database"] = args ? args.database : undefined;
            resourceInputs["deepAppInspDbLimit"] = args ? args.deepAppInspDbLimit : undefined;
            resourceInputs["deepAppInspTimeout"] = args ? args.deepAppInspTimeout : undefined;
            resourceInputs["engineCount"] = args ? args.engineCount : undefined;
            resourceInputs["excludeSignatures"] = args ? args.excludeSignatures : undefined;
            resourceInputs["failOpen"] = args ? args.failOpen : undefined;
            resourceInputs["intelligentMode"] = args ? args.intelligentMode : undefined;
            resourceInputs["ipsReserveCpu"] = args ? args.ipsReserveCpu : undefined;
            resourceInputs["ngfwMaxScanRange"] = args ? args.ngfwMaxScanRange : undefined;
            resourceInputs["npAccelMode"] = args ? args.npAccelMode : undefined;
            resourceInputs["packetLogQueueDepth"] = args ? args.packetLogQueueDepth : undefined;
            resourceInputs["sessionLimitMode"] = args ? args.sessionLimitMode : undefined;
            resourceInputs["skypeClientPublicIpaddr"] = args ? args.skypeClientPublicIpaddr : undefined;
            resourceInputs["socketSize"] = args ? args.socketSize : undefined;
            resourceInputs["syncSessionTtl"] = args ? args.syncSessionTtl : undefined;
            resourceInputs["tlsActiveProbe"] = args ? args.tlsActiveProbe : undefined;
            resourceInputs["trafficSubmit"] = args ? args.trafficSubmit : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpsGlobal.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsGlobal resources.
 */
export interface IpsGlobalState {
    /**
     * Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
     */
    anomalyMode?: pulumi.Input<string>;
    /**
     * IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
     */
    cpAccelMode?: pulumi.Input<string>;
    /**
     * Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
     */
    database?: pulumi.Input<string>;
    /**
     * Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
     */
    deepAppInspDbLimit?: pulumi.Input<number>;
    /**
     * Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
     */
    deepAppInspTimeout?: pulumi.Input<number>;
    /**
     * Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
     */
    engineCount?: pulumi.Input<number>;
    /**
     * Excluded signatures. Valid values: `none`, `industrial`.
     */
    excludeSignatures?: pulumi.Input<string>;
    /**
     * Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
     */
    failOpen?: pulumi.Input<string>;
    /**
     * Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
     */
    intelligentMode?: pulumi.Input<string>;
    /**
     * Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
     */
    ipsReserveCpu?: pulumi.Input<string>;
    /**
     * NGFW policy-mode app detection threshold.
     */
    ngfwMaxScanRange?: pulumi.Input<number>;
    /**
     * Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
     */
    npAccelMode?: pulumi.Input<string>;
    /**
     * Packet/pcap log queue depth per IPS engine.
     */
    packetLogQueueDepth?: pulumi.Input<number>;
    /**
     * Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
     */
    sessionLimitMode?: pulumi.Input<string>;
    /**
     * Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
     */
    skypeClientPublicIpaddr?: pulumi.Input<string>;
    /**
     * IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
     */
    socketSize?: pulumi.Input<number>;
    /**
     * Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
     */
    syncSessionTtl?: pulumi.Input<string>;
    /**
     * TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
     */
    tlsActiveProbe?: pulumi.Input<inputs.IpsGlobalTlsActiveProbe>;
    /**
     * Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
     */
    trafficSubmit?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsGlobal resource.
 */
export interface IpsGlobalArgs {
    /**
     * Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
     */
    anomalyMode?: pulumi.Input<string>;
    /**
     * IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
     */
    cpAccelMode?: pulumi.Input<string>;
    /**
     * Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
     */
    database?: pulumi.Input<string>;
    /**
     * Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
     */
    deepAppInspDbLimit?: pulumi.Input<number>;
    /**
     * Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
     */
    deepAppInspTimeout?: pulumi.Input<number>;
    /**
     * Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
     */
    engineCount?: pulumi.Input<number>;
    /**
     * Excluded signatures. Valid values: `none`, `industrial`.
     */
    excludeSignatures?: pulumi.Input<string>;
    /**
     * Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
     */
    failOpen?: pulumi.Input<string>;
    /**
     * Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
     */
    intelligentMode?: pulumi.Input<string>;
    /**
     * Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
     */
    ipsReserveCpu?: pulumi.Input<string>;
    /**
     * NGFW policy-mode app detection threshold.
     */
    ngfwMaxScanRange?: pulumi.Input<number>;
    /**
     * Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
     */
    npAccelMode?: pulumi.Input<string>;
    /**
     * Packet/pcap log queue depth per IPS engine.
     */
    packetLogQueueDepth?: pulumi.Input<number>;
    /**
     * Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
     */
    sessionLimitMode?: pulumi.Input<string>;
    /**
     * Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
     */
    skypeClientPublicIpaddr?: pulumi.Input<string>;
    /**
     * IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
     */
    socketSize?: pulumi.Input<number>;
    /**
     * Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
     */
    syncSessionTtl?: pulumi.Input<string>;
    /**
     * TLS active probe configuration. The structure of `tlsActiveProbe` block is documented below.
     */
    tlsActiveProbe?: pulumi.Input<inputs.IpsGlobalTlsActiveProbe>;
    /**
     * Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
     */
    trafficSubmit?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
