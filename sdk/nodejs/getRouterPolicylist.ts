// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.RouterPolicy`.
 */
export function getRouterPolicylist(args?: GetRouterPolicylistArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterPolicylistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterPolicylist:getRouterPolicylist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterPolicylist.
 */
export interface GetRouterPolicylistArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getRouterPolicylist.
 */
export interface GetRouterPolicylistResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.RouterPolicy`.
     */
    readonly seqNumlists: number[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.RouterPolicy`.
 */
export function getRouterPolicylistOutput(args?: GetRouterPolicylistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterPolicylistResult> {
    return pulumi.output(args).apply((a: any) => getRouterPolicylist(a, opts))
}

/**
 * A collection of arguments for invoking getRouterPolicylist.
 */
export interface GetRouterPolicylistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
