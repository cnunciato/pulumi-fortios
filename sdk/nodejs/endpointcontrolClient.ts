// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Configure endpoint control client lists. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * EndpointControl Client can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/endpointcontrolClient:EndpointcontrolClient labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/endpointcontrolClient:EndpointcontrolClient labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class EndpointcontrolClient extends pulumi.CustomResource {
    /**
     * Get an existing EndpointcontrolClient resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointcontrolClientState, opts?: pulumi.CustomResourceOptions): EndpointcontrolClient {
        return new EndpointcontrolClient(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/endpointcontrolClient:EndpointcontrolClient';

    /**
     * Returns true if the given object is an instance of EndpointcontrolClient.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointcontrolClient {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointcontrolClient.__pulumiType;
    }

    /**
     * Endpoint client AD logon groups.
     */
    public readonly adGroups!: pulumi.Output<string | undefined>;
    /**
     * Endpoint client ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Endpoint FortiClient UID.
     */
    public readonly ftclUid!: pulumi.Output<string>;
    /**
     * Endpoint client information.
     */
    public readonly info!: pulumi.Output<string>;
    /**
     * Endpoint client IP address.
     */
    public readonly srcIp!: pulumi.Output<string>;
    /**
     * Endpoint client MAC address.
     */
    public readonly srcMac!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a EndpointcontrolClient resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointcontrolClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointcontrolClientArgs | EndpointcontrolClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointcontrolClientState | undefined;
            resourceInputs["adGroups"] = state ? state.adGroups : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["ftclUid"] = state ? state.ftclUid : undefined;
            resourceInputs["info"] = state ? state.info : undefined;
            resourceInputs["srcIp"] = state ? state.srcIp : undefined;
            resourceInputs["srcMac"] = state ? state.srcMac : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as EndpointcontrolClientArgs | undefined;
            resourceInputs["adGroups"] = args ? args.adGroups : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["ftclUid"] = args ? args.ftclUid : undefined;
            resourceInputs["info"] = args ? args.info : undefined;
            resourceInputs["srcIp"] = args ? args.srcIp : undefined;
            resourceInputs["srcMac"] = args ? args.srcMac : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointcontrolClient.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointcontrolClient resources.
 */
export interface EndpointcontrolClientState {
    /**
     * Endpoint client AD logon groups.
     */
    adGroups?: pulumi.Input<string>;
    /**
     * Endpoint client ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Endpoint FortiClient UID.
     */
    ftclUid?: pulumi.Input<string>;
    /**
     * Endpoint client information.
     */
    info?: pulumi.Input<string>;
    /**
     * Endpoint client IP address.
     */
    srcIp?: pulumi.Input<string>;
    /**
     * Endpoint client MAC address.
     */
    srcMac?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndpointcontrolClient resource.
 */
export interface EndpointcontrolClientArgs {
    /**
     * Endpoint client AD logon groups.
     */
    adGroups?: pulumi.Input<string>;
    /**
     * Endpoint client ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Endpoint FortiClient UID.
     */
    ftclUid?: pulumi.Input<string>;
    /**
     * Endpoint client information.
     */
    info?: pulumi.Input<string>;
    /**
     * Endpoint client IP address.
     */
    srcIp?: pulumi.Input<string>;
    /**
     * Endpoint client MAC address.
     */
    srcMac?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
