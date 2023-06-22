// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure WAN optimization profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.WanoptProfile("trname", {
 *     cifs: {
 *         byteCaching: "enable",
 *         logTraffic: "enable",
 *         port: 445,
 *         preferChunking: "fix",
 *         secureTunnel: "disable",
 *         status: "disable",
 *         tunnelSharing: "private",
 *     },
 *     comments: "test",
 *     ftp: {
 *         byteCaching: "enable",
 *         logTraffic: "enable",
 *         port: 21,
 *         preferChunking: "fix",
 *         secureTunnel: "disable",
 *         status: "disable",
 *         tunnelSharing: "private",
 *     },
 *     http: {
 *         byteCaching: "enable",
 *         logTraffic: "enable",
 *         port: 80,
 *         preferChunking: "fix",
 *         secureTunnel: "disable",
 *         ssl: "disable",
 *         sslPort: 443,
 *         status: "disable",
 *         tunnelNonHttp: "disable",
 *         tunnelSharing: "private",
 *         unknownHttpVersion: "tunnel",
 *     },
 *     mapi: {
 *         byteCaching: "enable",
 *         logTraffic: "enable",
 *         port: 135,
 *         secureTunnel: "disable",
 *         status: "disable",
 *         tunnelSharing: "private",
 *     },
 *     tcp: {
 *         byteCaching: "disable",
 *         byteCachingOpt: "mem-only",
 *         logTraffic: "enable",
 *         port: "1-65535",
 *         secureTunnel: "disable",
 *         ssl: "disable",
 *         sslPort: 443,
 *         status: "disable",
 *         tunnelSharing: "private",
 *     },
 *     transparent: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * Wanopt Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wanoptProfile:WanoptProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class WanoptProfile extends pulumi.CustomResource {
    /**
     * Get an existing WanoptProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WanoptProfileState, opts?: pulumi.CustomResourceOptions): WanoptProfile {
        return new WanoptProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wanoptProfile:WanoptProfile';

    /**
     * Returns true if the given object is an instance of WanoptProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WanoptProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WanoptProfile.__pulumiType;
    }

    /**
     * Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
     */
    public readonly authGroup!: pulumi.Output<string>;
    /**
     * Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
     */
    public readonly cifs!: pulumi.Output<outputs.WanoptProfileCifs>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
     */
    public readonly ftp!: pulumi.Output<outputs.WanoptProfileFtp>;
    /**
     * Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
     */
    public readonly http!: pulumi.Output<outputs.WanoptProfileHttp>;
    /**
     * Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
     */
    public readonly mapi!: pulumi.Output<outputs.WanoptProfileMapi>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
     */
    public readonly tcp!: pulumi.Output<outputs.WanoptProfileTcp>;
    /**
     * Enable/disable transparent mode. Valid values: `enable`, `disable`.
     */
    public readonly transparent!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a WanoptProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WanoptProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WanoptProfileArgs | WanoptProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WanoptProfileState | undefined;
            resourceInputs["authGroup"] = state ? state.authGroup : undefined;
            resourceInputs["cifs"] = state ? state.cifs : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["ftp"] = state ? state.ftp : undefined;
            resourceInputs["http"] = state ? state.http : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tcp"] = state ? state.tcp : undefined;
            resourceInputs["transparent"] = state ? state.transparent : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as WanoptProfileArgs | undefined;
            resourceInputs["authGroup"] = args ? args.authGroup : undefined;
            resourceInputs["cifs"] = args ? args.cifs : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["ftp"] = args ? args.ftp : undefined;
            resourceInputs["http"] = args ? args.http : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tcp"] = args ? args.tcp : undefined;
            resourceInputs["transparent"] = args ? args.transparent : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WanoptProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WanoptProfile resources.
 */
export interface WanoptProfileState {
    /**
     * Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
     */
    authGroup?: pulumi.Input<string>;
    /**
     * Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.WanoptProfileCifs>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.WanoptProfileFtp>;
    /**
     * Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.WanoptProfileHttp>;
    /**
     * Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.WanoptProfileMapi>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
     */
    tcp?: pulumi.Input<inputs.WanoptProfileTcp>;
    /**
     * Enable/disable transparent mode. Valid values: `enable`, `disable`.
     */
    transparent?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WanoptProfile resource.
 */
export interface WanoptProfileArgs {
    /**
     * Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
     */
    authGroup?: pulumi.Input<string>;
    /**
     * Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.WanoptProfileCifs>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.WanoptProfileFtp>;
    /**
     * Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.WanoptProfileHttp>;
    /**
     * Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.WanoptProfileMapi>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
     */
    tcp?: pulumi.Input<inputs.WanoptProfileTcp>;
    /**
     * Enable/disable transparent mode. Valid values: `enable`, `disable`.
     */
    transparent?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}