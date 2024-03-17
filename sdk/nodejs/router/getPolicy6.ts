// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios router policy6
 */
export function getPolicy6(args: GetPolicy6Args, opts?: pulumi.InvokeOptions): Promise<GetPolicy6Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getPolicy6:getPolicy6", {
        "seqNum": args.seqNum,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy6.
 */
export interface GetPolicy6Args {
    /**
     * Specify the seqNum of the desired router policy6.
     */
    seqNum: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getPolicy6.
 */
export interface GetPolicy6Result {
    /**
     * Action of the policy route.
     */
    readonly action: string;
    /**
     * Optional comments.
     */
    readonly comments: string;
    /**
     * Destination IPv6 prefix.
     */
    readonly dst: string;
    /**
     * Enable/disable negating destination address match.
     */
    readonly dstNegate: string;
    /**
     * Destination address name. The structure of `dstaddr` block is documented below.
     */
    readonly dstaddrs: outputs.router.GetPolicy6Dstaddr[];
    /**
     * End destination port number (1 - 65535).
     */
    readonly endPort: number;
    /**
     * IPv6 address of the gateway.
     */
    readonly gateway: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Incoming interface name.
     */
    readonly inputDevice: string;
    /**
     * Enable/disable negation of input device match.
     */
    readonly inputDeviceNegate: string;
    /**
     * Custom Destination Internet Service name. The structure of `internetServiceCustom` block is documented below.
     */
    readonly internetServiceCustoms: outputs.router.GetPolicy6InternetServiceCustom[];
    /**
     * Destination Internet Service ID. The structure of `internetServiceId` block is documented below.
     */
    readonly internetServiceIds: outputs.router.GetPolicy6InternetServiceId[];
    /**
     * Outgoing interface name.
     */
    readonly outputDevice: string;
    /**
     * Protocol number (0 - 255).
     */
    readonly protocol: number;
    /**
     * Sequence number.
     */
    readonly seqNum: number;
    /**
     * Source IPv6 prefix.
     */
    readonly src: string;
    /**
     * Enable/disable negating source address match.
     */
    readonly srcNegate: string;
    /**
     * Source address name. The structure of `srcaddr` block is documented below.
     */
    readonly srcaddrs: outputs.router.GetPolicy6Srcaddr[];
    /**
     * Start destination port number (1 - 65535).
     */
    readonly startPort: number;
    /**
     * Enable/disable this policy route.
     */
    readonly status: string;
    /**
     * Type of service bit pattern.
     */
    readonly tos: string;
    /**
     * Type of service evaluated bits.
     */
    readonly tosMask: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router policy6
 */
export function getPolicy6Output(args: GetPolicy6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicy6Result> {
    return pulumi.output(args).apply((a: any) => getPolicy6(a, opts))
}

/**
 * A collection of arguments for invoking getPolicy6.
 */
export interface GetPolicy6OutputArgs {
    /**
     * Specify the seqNum of the desired router policy6.
     */
    seqNum: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}