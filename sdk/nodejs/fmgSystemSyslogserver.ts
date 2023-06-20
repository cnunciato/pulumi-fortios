// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource supports Create/Delete system syslog server for FortiManager.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.FmgSystemSyslogserver("test1", {
 *     ip: "1.1.1.1",
 *     port: 99,
 * });
 * ```
 */
export class FmgSystemSyslogserver extends pulumi.CustomResource {
    /**
     * Get an existing FmgSystemSyslogserver resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FmgSystemSyslogserverState, opts?: pulumi.CustomResourceOptions): FmgSystemSyslogserver {
        return new FmgSystemSyslogserver(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/fmgSystemSyslogserver:FmgSystemSyslogserver';

    /**
     * Returns true if the given object is an instance of FmgSystemSyslogserver.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FmgSystemSyslogserver {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FmgSystemSyslogserver.__pulumiType;
    }

    /**
     * Ipaddress of the syslog server.
     */
    public readonly ip!: pulumi.Output<string | undefined>;
    /**
     * Syslog server name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Port of the syslog server.
     */
    public readonly port!: pulumi.Output<number | undefined>;

    /**
     * Create a FmgSystemSyslogserver resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FmgSystemSyslogserverArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FmgSystemSyslogserverArgs | FmgSystemSyslogserverState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FmgSystemSyslogserverState | undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
        } else {
            const args = argsOrState as FmgSystemSyslogserverArgs | undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FmgSystemSyslogserver.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FmgSystemSyslogserver resources.
 */
export interface FmgSystemSyslogserverState {
    /**
     * Ipaddress of the syslog server.
     */
    ip?: pulumi.Input<string>;
    /**
     * Syslog server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the syslog server.
     */
    port?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a FmgSystemSyslogserver resource.
 */
export interface FmgSystemSyslogserverArgs {
    /**
     * Ipaddress of the syslog server.
     */
    ip?: pulumi.Input<string>;
    /**
     * Syslog server name.
     */
    name?: pulumi.Input<string>;
    /**
     * Port of the syslog server.
     */
    port?: pulumi.Input<number>;
}
