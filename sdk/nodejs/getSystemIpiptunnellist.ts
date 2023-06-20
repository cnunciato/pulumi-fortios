// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a list of `fortios.SystemIpiptunnel`.
 */
export function getSystemIpiptunnellist(args?: GetSystemIpiptunnellistArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemIpiptunnellistResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getSystemIpiptunnellist:getSystemIpiptunnellist", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystemIpiptunnellist.
 */
export interface GetSystemIpiptunnellistArgs {
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
 * A collection of values returned by getSystemIpiptunnellist.
 */
export interface GetSystemIpiptunnellistResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.SystemIpiptunnel`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.SystemIpiptunnel`.
 */
export function getSystemIpiptunnellistOutput(args?: GetSystemIpiptunnellistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemIpiptunnellistResult> {
    return pulumi.output(args).apply((a: any) => getSystemIpiptunnellist(a, opts))
}

/**
 * A collection of arguments for invoking getSystemIpiptunnellist.
 */
export interface GetSystemIpiptunnellistOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
