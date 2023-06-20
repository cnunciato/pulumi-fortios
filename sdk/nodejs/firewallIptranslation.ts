// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure firewall IP-translation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.FirewallIptranslation("trname", {
 *     endip: "2.2.2.2",
 *     mapStartip: "0.0.0.0",
 *     startip: "1.1.1.1",
 *     transid: 1,
 *     type: "SCTP",
 * });
 * ```
 *
 * ## Import
 *
 * Firewall IpTranslation can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallIptranslation:FirewallIptranslation labelname {{transid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallIptranslation:FirewallIptranslation labelname {{transid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallIptranslation extends pulumi.CustomResource {
    /**
     * Get an existing FirewallIptranslation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallIptranslationState, opts?: pulumi.CustomResourceOptions): FirewallIptranslation {
        return new FirewallIptranslation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallIptranslation:FirewallIptranslation';

    /**
     * Returns true if the given object is an instance of FirewallIptranslation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallIptranslation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallIptranslation.__pulumiType;
    }

    /**
     * Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    public readonly endip!: pulumi.Output<string>;
    /**
     * Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    public readonly mapStartip!: pulumi.Output<string>;
    /**
     * First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    public readonly startip!: pulumi.Output<string>;
    /**
     * IP translation ID.
     */
    public readonly transid!: pulumi.Output<number>;
    /**
     * IP translation type (option: SCTP). Valid values: `SCTP`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a FirewallIptranslation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallIptranslationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallIptranslationArgs | FirewallIptranslationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallIptranslationState | undefined;
            resourceInputs["endip"] = state ? state.endip : undefined;
            resourceInputs["mapStartip"] = state ? state.mapStartip : undefined;
            resourceInputs["startip"] = state ? state.startip : undefined;
            resourceInputs["transid"] = state ? state.transid : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as FirewallIptranslationArgs | undefined;
            if ((!args || args.endip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endip'");
            }
            if ((!args || args.mapStartip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mapStartip'");
            }
            if ((!args || args.startip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'startip'");
            }
            resourceInputs["endip"] = args ? args.endip : undefined;
            resourceInputs["mapStartip"] = args ? args.mapStartip : undefined;
            resourceInputs["startip"] = args ? args.startip : undefined;
            resourceInputs["transid"] = args ? args.transid : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallIptranslation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallIptranslation resources.
 */
export interface FirewallIptranslationState {
    /**
     * Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    endip?: pulumi.Input<string>;
    /**
     * Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    mapStartip?: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    startip?: pulumi.Input<string>;
    /**
     * IP translation ID.
     */
    transid?: pulumi.Input<number>;
    /**
     * IP translation type (option: SCTP). Valid values: `SCTP`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallIptranslation resource.
 */
export interface FirewallIptranslationArgs {
    /**
     * Final IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    endip: pulumi.Input<string>;
    /**
     * Address to be used as the starting point for translation in the range (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    mapStartip: pulumi.Input<string>;
    /**
     * First IPv4 address (inclusive) in the range of the addresses to be translated (format xxx.xxx.xxx.xxx, default: 0.0.0.0).
     */
    startip: pulumi.Input<string>;
    /**
     * IP translation ID.
     */
    transid?: pulumi.Input<number>;
    /**
     * IP translation type (option: SCTP). Valid values: `SCTP`.
     */
    type?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
