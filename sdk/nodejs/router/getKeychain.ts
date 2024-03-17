// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios router keychain
 */
export function getKeychain(args: GetKeychainArgs, opts?: pulumi.InvokeOptions): Promise<GetKeychainResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getKeychain:getKeychain", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getKeychain.
 */
export interface GetKeychainArgs {
    /**
     * Specify the name of the desired router keychain.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getKeychain.
 */
export interface GetKeychainResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Configuration method to edit key settings. The structure of `key` block is documented below.
     */
    readonly keys: outputs.router.GetKeychainKey[];
    /**
     * Key-chain name.
     */
    readonly name: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router keychain
 */
export function getKeychainOutput(args: GetKeychainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeychainResult> {
    return pulumi.output(args).apply((a: any) => getKeychain(a, opts))
}

/**
 * A collection of arguments for invoking getKeychain.
 */
export interface GetKeychainOutputArgs {
    /**
     * Specify the name of the desired router keychain.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}