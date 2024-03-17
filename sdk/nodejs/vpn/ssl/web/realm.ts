// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../../utilities";

/**
 * Realm.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.vpn.ssl.web.Realm("trname", {
 *     loginPage: "1.htm",
 *     maxConcurrentUser: 33,
 *     urlPath: "1",
 *     virtualHost: "2.2.2.2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * VpnSslWeb Realm can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/realm:Realm labelname {{url_path}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:vpn/ssl/web/realm:Realm labelname {{url_path}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Realm extends pulumi.CustomResource {
    /**
     * Get an existing Realm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmState, opts?: pulumi.CustomResourceOptions): Realm {
        return new Realm(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:vpn/ssl/web/realm:Realm';

    /**
     * Returns true if the given object is an instance of Realm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Realm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Realm.__pulumiType;
    }

    /**
     * Replacement HTML for SSL-VPN login page.
     */
    public readonly loginPage!: pulumi.Output<string | undefined>;
    /**
     * Maximum concurrent users (0 - 65535, 0 means unlimited).
     */
    public readonly maxConcurrentUser!: pulumi.Output<number>;
    /**
     * IP address used as a NAS-IP to communicate with the RADIUS server.
     */
    public readonly nasIp!: pulumi.Output<string>;
    /**
     * RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
     */
    public readonly radiusPort!: pulumi.Output<number>;
    /**
     * RADIUS server associated with realm.
     */
    public readonly radiusServer!: pulumi.Output<string>;
    /**
     * URL path to access SSL-VPN login page.
     */
    public readonly urlPath!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Virtual host name for realm.
     */
    public readonly virtualHost!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
     */
    public readonly virtualHostOnly!: pulumi.Output<string>;
    /**
     * Name of the server certificate to used for this realm.
     */
    public readonly virtualHostServerCert!: pulumi.Output<string>;

    /**
     * Create a Realm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RealmArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmArgs | RealmState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmState | undefined;
            resourceInputs["loginPage"] = state ? state.loginPage : undefined;
            resourceInputs["maxConcurrentUser"] = state ? state.maxConcurrentUser : undefined;
            resourceInputs["nasIp"] = state ? state.nasIp : undefined;
            resourceInputs["radiusPort"] = state ? state.radiusPort : undefined;
            resourceInputs["radiusServer"] = state ? state.radiusServer : undefined;
            resourceInputs["urlPath"] = state ? state.urlPath : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["virtualHost"] = state ? state.virtualHost : undefined;
            resourceInputs["virtualHostOnly"] = state ? state.virtualHostOnly : undefined;
            resourceInputs["virtualHostServerCert"] = state ? state.virtualHostServerCert : undefined;
        } else {
            const args = argsOrState as RealmArgs | undefined;
            resourceInputs["loginPage"] = args ? args.loginPage : undefined;
            resourceInputs["maxConcurrentUser"] = args ? args.maxConcurrentUser : undefined;
            resourceInputs["nasIp"] = args ? args.nasIp : undefined;
            resourceInputs["radiusPort"] = args ? args.radiusPort : undefined;
            resourceInputs["radiusServer"] = args ? args.radiusServer : undefined;
            resourceInputs["urlPath"] = args ? args.urlPath : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["virtualHost"] = args ? args.virtualHost : undefined;
            resourceInputs["virtualHostOnly"] = args ? args.virtualHostOnly : undefined;
            resourceInputs["virtualHostServerCert"] = args ? args.virtualHostServerCert : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Realm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Realm resources.
 */
export interface RealmState {
    /**
     * Replacement HTML for SSL-VPN login page.
     */
    loginPage?: pulumi.Input<string>;
    /**
     * Maximum concurrent users (0 - 65535, 0 means unlimited).
     */
    maxConcurrentUser?: pulumi.Input<number>;
    /**
     * IP address used as a NAS-IP to communicate with the RADIUS server.
     */
    nasIp?: pulumi.Input<string>;
    /**
     * RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
     */
    radiusPort?: pulumi.Input<number>;
    /**
     * RADIUS server associated with realm.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * URL path to access SSL-VPN login page.
     */
    urlPath?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual host name for realm.
     */
    virtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
     */
    virtualHostOnly?: pulumi.Input<string>;
    /**
     * Name of the server certificate to used for this realm.
     */
    virtualHostServerCert?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Realm resource.
 */
export interface RealmArgs {
    /**
     * Replacement HTML for SSL-VPN login page.
     */
    loginPage?: pulumi.Input<string>;
    /**
     * Maximum concurrent users (0 - 65535, 0 means unlimited).
     */
    maxConcurrentUser?: pulumi.Input<number>;
    /**
     * IP address used as a NAS-IP to communicate with the RADIUS server.
     */
    nasIp?: pulumi.Input<string>;
    /**
     * RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
     */
    radiusPort?: pulumi.Input<number>;
    /**
     * RADIUS server associated with realm.
     */
    radiusServer?: pulumi.Input<string>;
    /**
     * URL path to access SSL-VPN login page.
     */
    urlPath?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Virtual host name for realm.
     */
    virtualHost?: pulumi.Input<string>;
    /**
     * Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
     */
    virtualHostOnly?: pulumi.Input<string>;
    /**
     * Name of the server certificate to used for this realm.
     */
    virtualHostServerCert?: pulumi.Input<string>;
}