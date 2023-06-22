// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get information on an fortios router accesslist
 */
export function getRouterAccesslist(args: GetRouterAccesslistArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterAccesslistResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:index/getRouterAccesslist:getRouterAccesslist", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterAccesslist.
 */
export interface GetRouterAccesslistArgs {
    /**
     * Specify the name of the desired router accesslist.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getRouterAccesslist.
 */
export interface GetRouterAccesslistResult {
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name.
     */
    readonly name: string;
    /**
     * Rule. The structure of `rule` block is documented below.
     */
    readonly rules: outputs.GetRouterAccesslistRule[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router accesslist
 */
export function getRouterAccesslistOutput(args: GetRouterAccesslistOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterAccesslistResult> {
    return pulumi.output(args).apply((a: any) => getRouterAccesslist(a, opts))
}

/**
 * A collection of arguments for invoking getRouterAccesslist.
 */
export interface GetRouterAccesslistOutputArgs {
    /**
     * Specify the name of the desired router accesslist.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}