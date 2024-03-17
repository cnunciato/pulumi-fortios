// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios firewall DoSpolicy6
 */
export function getDoSpolicy6(args: GetDoSpolicy6Args, opts?: pulumi.InvokeOptions): Promise<GetDoSpolicy6Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getDoSpolicy6:getDoSpolicy6", {
        "policyid": args.policyid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getDoSpolicy6.
 */
export interface GetDoSpolicy6Args {
    /**
     * Specify the policyid of the desired firewall DoSpolicy6.
     */
    policyid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getDoSpolicy6.
 */
export interface GetDoSpolicy6Result {
    /**
     * Anomaly name. The structure of `anomaly` block is documented below.
     */
    readonly anomalies: outputs.firewall.GetDoSpolicy6Anomaly[];
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * Destination address name from available addresses. The structure of `dstaddr` block is documented below.
     */
    readonly dstaddrs: outputs.firewall.GetDoSpolicy6Dstaddr[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Incoming interface name from available interfaces.
     */
    readonly interface: string;
    /**
     * Anomaly name.
     */
    readonly name: string;
    /**
     * Policy ID.
     */
    readonly policyid: number;
    /**
     * Service object from available options. The structure of `service` block is documented below.
     */
    readonly services: outputs.firewall.GetDoSpolicy6Service[];
    /**
     * Source address name from available addresses. The structure of `srcaddr` block is documented below.
     */
    readonly srcaddrs: outputs.firewall.GetDoSpolicy6Srcaddr[];
    /**
     * Enable/disable this anomaly.
     */
    readonly status: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios firewall DoSpolicy6
 */
export function getDoSpolicy6Output(args: GetDoSpolicy6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDoSpolicy6Result> {
    return pulumi.output(args).apply((a: any) => getDoSpolicy6(a, opts))
}

/**
 * A collection of arguments for invoking getDoSpolicy6.
 */
export interface GetDoSpolicy6OutputArgs {
    /**
     * Specify the policyid of the desired firewall DoSpolicy6.
     */
    policyid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}