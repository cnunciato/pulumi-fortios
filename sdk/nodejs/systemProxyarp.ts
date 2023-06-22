// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure proxy-ARP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.SystemProxyarp("trname", {
 *     endIp: "1.1.1.3",
 *     fosid: 1,
 *     "interface": "port4",
 *     ip: "1.1.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * System ProxyArp can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/systemProxyarp:SystemProxyarp labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/systemProxyarp:SystemProxyarp labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class SystemProxyarp extends pulumi.CustomResource {
    /**
     * Get an existing SystemProxyarp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemProxyarpState, opts?: pulumi.CustomResourceOptions): SystemProxyarp {
        return new SystemProxyarp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/systemProxyarp:SystemProxyarp';

    /**
     * Returns true if the given object is an instance of SystemProxyarp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SystemProxyarp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SystemProxyarp.__pulumiType;
    }

    /**
     * End IP of IP range to be proxied.
     */
    public readonly endIp!: pulumi.Output<string>;
    /**
     * Unique integer ID of the entry.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Interface acting proxy-ARP.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * IP address or start IP to be proxied.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a SystemProxyarp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SystemProxyarpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemProxyarpArgs | SystemProxyarpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemProxyarpState | undefined;
            resourceInputs["endIp"] = state ? state.endIp : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemProxyarpArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.interface === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interface'");
            }
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            resourceInputs["endIp"] = args ? args.endIp : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SystemProxyarp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SystemProxyarp resources.
 */
export interface SystemProxyarpState {
    /**
     * End IP of IP range to be proxied.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Unique integer ID of the entry.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Interface acting proxy-ARP.
     */
    interface?: pulumi.Input<string>;
    /**
     * IP address or start IP to be proxied.
     */
    ip?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SystemProxyarp resource.
 */
export interface SystemProxyarpArgs {
    /**
     * End IP of IP range to be proxied.
     */
    endIp?: pulumi.Input<string>;
    /**
     * Unique integer ID of the entry.
     */
    fosid: pulumi.Input<number>;
    /**
     * Interface acting proxy-ARP.
     */
    interface: pulumi.Input<string>;
    /**
     * IP address or start IP to be proxied.
     */
    ip: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}