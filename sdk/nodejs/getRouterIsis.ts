// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on fortios router isis
 */
export function getRouterIsis(args?: GetRouterIsisArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterIsisResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterIsis:getRouterIsis", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterIsis.
 */
export interface GetRouterIsisArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getRouterIsis.
 */
export interface GetRouterIsisResult {
    /**
     * Enable/disable adjacency check.
     */
    readonly adjacencyCheck: string;
    /**
     * Enable/disable IPv6 adjacency check.
     */
    readonly adjacencyCheck6: string;
    /**
     * Enable/disable IS-IS advertisement of passive interfaces only.
     */
    readonly advPassiveOnly: string;
    /**
     * Enable/disable IPv6 IS-IS advertisement of passive interfaces only.
     */
    readonly advPassiveOnly6: string;
    /**
     * Authentication key-chain for level 1 PDUs.
     */
    readonly authKeychainL1: string;
    /**
     * Authentication key-chain for level 2 PDUs.
     */
    readonly authKeychainL2: string;
    /**
     * Level 1 authentication mode.
     */
    readonly authModeL1: string;
    /**
     * Level 2 authentication mode.
     */
    readonly authModeL2: string;
    /**
     * Authentication password for level 1 PDUs.
     */
    readonly authPasswordL1: string;
    /**
     * Authentication password for level 2 PDUs.
     */
    readonly authPasswordL2: string;
    /**
     * Enable/disable level 1 authentication send-only.
     */
    readonly authSendonlyL1: string;
    /**
     * Enable/disable level 2 authentication send-only.
     */
    readonly authSendonlyL2: string;
    /**
     * Enable/disable distribution of default route information.
     */
    readonly defaultOriginate: string;
    /**
     * Enable/disable distribution of default IPv6 route information.
     */
    readonly defaultOriginate6: string;
    /**
     * Enable/disable dynamic hostname.
     */
    readonly dynamicHostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Enable/disable ignoring of LSP errors with bad checksums.
     */
    readonly ignoreLspErrors: string;
    /**
     * IS type.
     */
    readonly isType: string;
    /**
     * IS-IS interface configuration. The structure of `isisInterface` block is documented below.
     */
    readonly isisInterfaces: outputs.GetRouterIsisIsisInterface[];
    /**
     * IS-IS net configuration. The structure of `isisNet` block is documented below.
     */
    readonly isisNets: outputs.GetRouterIsisIsisNet[];
    /**
     * Minimum interval for level 1 LSP regenerating.
     */
    readonly lspGenIntervalL1: number;
    /**
     * Minimum interval for level 2 LSP regenerating.
     */
    readonly lspGenIntervalL2: number;
    /**
     * LSP refresh time in seconds.
     */
    readonly lspRefreshInterval: number;
    /**
     * Maximum LSP lifetime in seconds.
     */
    readonly maxLspLifetime: number;
    /**
     * Use old-style (ISO 10589) or new-style packet formats
     */
    readonly metricStyle: string;
    /**
     * Enable/disable signal other routers not to use us in SPF.
     */
    readonly overloadBit: string;
    /**
     * Overload-bit only temporarily after reboot.
     */
    readonly overloadBitOnStartup: number;
    /**
     * Suppress overload-bit for the specific prefixes.
     */
    readonly overloadBitSuppress: string;
    /**
     * Enable/disable redistribution of level 1 IPv6 routes into level 2.
     */
    readonly redistribute6L1: string;
    /**
     * Access-list for IPv6 route redistribution from l1 to l2.
     */
    readonly redistribute6L1List: string;
    /**
     * Enable/disable redistribution of level 2 IPv6 routes into level 1.
     */
    readonly redistribute6L2: string;
    /**
     * Access-list for IPv6 route redistribution from l2 to l1.
     */
    readonly redistribute6L2List: string;
    /**
     * IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
     */
    readonly redistribute6s: outputs.GetRouterIsisRedistribute6[];
    /**
     * Enable/disable redistribution of level 1 routes into level 2.
     */
    readonly redistributeL1: string;
    /**
     * Access-list for route redistribution from l1 to l2.
     */
    readonly redistributeL1List: string;
    /**
     * Enable/disable redistribution of level 2 routes into level 1.
     */
    readonly redistributeL2: string;
    /**
     * Access-list for route redistribution from l2 to l1.
     */
    readonly redistributeL2List: string;
    /**
     * IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
     */
    readonly redistributes: outputs.GetRouterIsisRedistribute[];
    /**
     * Level 1 SPF calculation delay.
     */
    readonly spfIntervalExpL1: string;
    /**
     * Level 2 SPF calculation delay.
     */
    readonly spfIntervalExpL2: string;
    /**
     * IS-IS IPv6 summary address. The structure of `summaryAddress6` block is documented below.
     */
    readonly summaryAddress6s: outputs.GetRouterIsisSummaryAddress6[];
    /**
     * IS-IS summary addresses. The structure of `summaryAddress` block is documented below.
     */
    readonly summaryAddresses: outputs.GetRouterIsisSummaryAddress[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios router isis
 */
export function getRouterIsisOutput(args?: GetRouterIsisOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterIsisResult> {
    return pulumi.output(args).apply((a: any) => getRouterIsis(a, opts))
}

/**
 * A collection of arguments for invoking getRouterIsis.
 */
export interface GetRouterIsisOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
