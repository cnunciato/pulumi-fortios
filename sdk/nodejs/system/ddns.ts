// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure DDNS.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.system.Ddns("trname", {
 *     boundIp: "0.0.0.0",
 *     clearText: "disable",
 *     ddnsAuth: "disable",
 *     ddnsDomain: "www.s.com",
 *     ddnsPassword: "ewewcd",
 *     ddnsServer: "tzo.com",
 *     ddnsServerIp: "0.0.0.0",
 *     ddnsTtl: 300,
 *     ddnsUsername: "sie2ae",
 *     ddnsid: 1,
 *     monitorInterfaces: [{
 *         interfaceName: "port2",
 *     }],
 *     sslCertificate: "Fortinet_Factory",
 *     updateInterval: 300,
 *     usePublicIp: "disable",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * System Ddns can be imported using any of these accepted formats:
 *
 * ```sh
 * $ pulumi import fortios:system/ddns:Ddns labelname {{ddnsid}}
 * ```
 *
 * If you do not want to import arguments of block:
 *
 * $ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 * $ pulumi import fortios:system/ddns:Ddns labelname {{ddnsid}}
 * ```
 *
 * $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Ddns extends pulumi.CustomResource {
    /**
     * Get an existing Ddns resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DdnsState, opts?: pulumi.CustomResourceOptions): Ddns {
        return new Ddns(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:system/ddns:Ddns';

    /**
     * Returns true if the given object is an instance of Ddns.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ddns {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ddns.__pulumiType;
    }

    /**
     * Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
     */
    public readonly addrType!: pulumi.Output<string>;
    /**
     * Bound IP address.
     */
    public readonly boundIp!: pulumi.Output<string>;
    /**
     * Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
     */
    public readonly clearText!: pulumi.Output<string>;
    /**
     * Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
     */
    public readonly ddnsAuth!: pulumi.Output<string>;
    /**
     * Your fully qualified domain name (for example, yourname.DDNS.com).
     */
    public readonly ddnsDomain!: pulumi.Output<string>;
    /**
     * DDNS update key (base 64 encoding).
     */
    public readonly ddnsKey!: pulumi.Output<string>;
    /**
     * DDNS update key name.
     */
    public readonly ddnsKeyname!: pulumi.Output<string>;
    /**
     * DDNS password.
     */
    public readonly ddnsPassword!: pulumi.Output<string | undefined>;
    /**
     * Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
     */
    public readonly ddnsServer!: pulumi.Output<string>;
    /**
     * Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
     */
    public readonly ddnsServerAddrs!: pulumi.Output<outputs.system.DdnsDdnsServerAddr[] | undefined>;
    /**
     * Generic DDNS server IP.
     */
    public readonly ddnsServerIp!: pulumi.Output<string>;
    /**
     * DDNS Serial Number.
     */
    public readonly ddnsSn!: pulumi.Output<string>;
    /**
     * Time-to-live for DDNS packets.
     */
    public readonly ddnsTtl!: pulumi.Output<number>;
    /**
     * DDNS user name.
     */
    public readonly ddnsUsername!: pulumi.Output<string>;
    /**
     * Zone of your domain name (for example, DDNS.com).
     */
    public readonly ddnsZone!: pulumi.Output<string>;
    /**
     * DDNS ID.
     */
    public readonly ddnsid!: pulumi.Output<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Monitored interface. The structure of `monitorInterface` block is documented below.
     */
    public readonly monitorInterfaces!: pulumi.Output<outputs.system.DdnsMonitorInterface[]>;
    /**
     * Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    public readonly sslCertificate!: pulumi.Output<string>;
    /**
     * DDNS update interval (60 - 2592000 sec, default = 300).
     */
    public readonly updateInterval!: pulumi.Output<number>;
    /**
     * Enable/disable use of public IP address. Valid values: `disable`, `enable`.
     */
    public readonly usePublicIp!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Ddns resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DdnsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DdnsArgs | DdnsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DdnsState | undefined;
            resourceInputs["addrType"] = state ? state.addrType : undefined;
            resourceInputs["boundIp"] = state ? state.boundIp : undefined;
            resourceInputs["clearText"] = state ? state.clearText : undefined;
            resourceInputs["ddnsAuth"] = state ? state.ddnsAuth : undefined;
            resourceInputs["ddnsDomain"] = state ? state.ddnsDomain : undefined;
            resourceInputs["ddnsKey"] = state ? state.ddnsKey : undefined;
            resourceInputs["ddnsKeyname"] = state ? state.ddnsKeyname : undefined;
            resourceInputs["ddnsPassword"] = state ? state.ddnsPassword : undefined;
            resourceInputs["ddnsServer"] = state ? state.ddnsServer : undefined;
            resourceInputs["ddnsServerAddrs"] = state ? state.ddnsServerAddrs : undefined;
            resourceInputs["ddnsServerIp"] = state ? state.ddnsServerIp : undefined;
            resourceInputs["ddnsSn"] = state ? state.ddnsSn : undefined;
            resourceInputs["ddnsTtl"] = state ? state.ddnsTtl : undefined;
            resourceInputs["ddnsUsername"] = state ? state.ddnsUsername : undefined;
            resourceInputs["ddnsZone"] = state ? state.ddnsZone : undefined;
            resourceInputs["ddnsid"] = state ? state.ddnsid : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["monitorInterfaces"] = state ? state.monitorInterfaces : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["sslCertificate"] = state ? state.sslCertificate : undefined;
            resourceInputs["updateInterval"] = state ? state.updateInterval : undefined;
            resourceInputs["usePublicIp"] = state ? state.usePublicIp : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DdnsArgs | undefined;
            if ((!args || args.ddnsServer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ddnsServer'");
            }
            if ((!args || args.monitorInterfaces === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitorInterfaces'");
            }
            resourceInputs["addrType"] = args ? args.addrType : undefined;
            resourceInputs["boundIp"] = args ? args.boundIp : undefined;
            resourceInputs["clearText"] = args ? args.clearText : undefined;
            resourceInputs["ddnsAuth"] = args ? args.ddnsAuth : undefined;
            resourceInputs["ddnsDomain"] = args ? args.ddnsDomain : undefined;
            resourceInputs["ddnsKey"] = args?.ddnsKey ? pulumi.secret(args.ddnsKey) : undefined;
            resourceInputs["ddnsKeyname"] = args ? args.ddnsKeyname : undefined;
            resourceInputs["ddnsPassword"] = args?.ddnsPassword ? pulumi.secret(args.ddnsPassword) : undefined;
            resourceInputs["ddnsServer"] = args ? args.ddnsServer : undefined;
            resourceInputs["ddnsServerAddrs"] = args ? args.ddnsServerAddrs : undefined;
            resourceInputs["ddnsServerIp"] = args ? args.ddnsServerIp : undefined;
            resourceInputs["ddnsSn"] = args ? args.ddnsSn : undefined;
            resourceInputs["ddnsTtl"] = args ? args.ddnsTtl : undefined;
            resourceInputs["ddnsUsername"] = args ? args.ddnsUsername : undefined;
            resourceInputs["ddnsZone"] = args ? args.ddnsZone : undefined;
            resourceInputs["ddnsid"] = args ? args.ddnsid : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["monitorInterfaces"] = args ? args.monitorInterfaces : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["sslCertificate"] = args ? args.sslCertificate : undefined;
            resourceInputs["updateInterval"] = args ? args.updateInterval : undefined;
            resourceInputs["usePublicIp"] = args ? args.usePublicIp : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ddnsKey", "ddnsPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Ddns.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ddns resources.
 */
export interface DdnsState {
    /**
     * Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * Bound IP address.
     */
    boundIp?: pulumi.Input<string>;
    /**
     * Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
     */
    clearText?: pulumi.Input<string>;
    /**
     * Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
     */
    ddnsAuth?: pulumi.Input<string>;
    /**
     * Your fully qualified domain name (for example, yourname.DDNS.com).
     */
    ddnsDomain?: pulumi.Input<string>;
    /**
     * DDNS update key (base 64 encoding).
     */
    ddnsKey?: pulumi.Input<string>;
    /**
     * DDNS update key name.
     */
    ddnsKeyname?: pulumi.Input<string>;
    /**
     * DDNS password.
     */
    ddnsPassword?: pulumi.Input<string>;
    /**
     * Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
     */
    ddnsServer?: pulumi.Input<string>;
    /**
     * Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
     */
    ddnsServerAddrs?: pulumi.Input<pulumi.Input<inputs.system.DdnsDdnsServerAddr>[]>;
    /**
     * Generic DDNS server IP.
     */
    ddnsServerIp?: pulumi.Input<string>;
    /**
     * DDNS Serial Number.
     */
    ddnsSn?: pulumi.Input<string>;
    /**
     * Time-to-live for DDNS packets.
     */
    ddnsTtl?: pulumi.Input<number>;
    /**
     * DDNS user name.
     */
    ddnsUsername?: pulumi.Input<string>;
    /**
     * Zone of your domain name (for example, DDNS.com).
     */
    ddnsZone?: pulumi.Input<string>;
    /**
     * DDNS ID.
     */
    ddnsid?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Monitored interface. The structure of `monitorInterface` block is documented below.
     */
    monitorInterfaces?: pulumi.Input<pulumi.Input<inputs.system.DdnsMonitorInterface>[]>;
    /**
     * Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * DDNS update interval (60 - 2592000 sec, default = 300).
     */
    updateInterval?: pulumi.Input<number>;
    /**
     * Enable/disable use of public IP address. Valid values: `disable`, `enable`.
     */
    usePublicIp?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ddns resource.
 */
export interface DdnsArgs {
    /**
     * Address type of interface address in DDNS update. Valid values: `ipv4`, `ipv6`.
     */
    addrType?: pulumi.Input<string>;
    /**
     * Bound IP address.
     */
    boundIp?: pulumi.Input<string>;
    /**
     * Enable/disable use of clear text connections. Valid values: `disable`, `enable`.
     */
    clearText?: pulumi.Input<string>;
    /**
     * Enable/disable TSIG authentication for your DDNS server. Valid values: `disable`, `tsig`.
     */
    ddnsAuth?: pulumi.Input<string>;
    /**
     * Your fully qualified domain name (for example, yourname.DDNS.com).
     */
    ddnsDomain?: pulumi.Input<string>;
    /**
     * DDNS update key (base 64 encoding).
     */
    ddnsKey?: pulumi.Input<string>;
    /**
     * DDNS update key name.
     */
    ddnsKeyname?: pulumi.Input<string>;
    /**
     * DDNS password.
     */
    ddnsPassword?: pulumi.Input<string>;
    /**
     * Select a DDNS service provider. Valid values: `dyndns.org`, `dyns.net`, `tzo.com`, `vavic.com`, `dipdns.net`, `now.net.cn`, `dhs.org`, `easydns.com`, `genericDDNS`, `FortiGuardDDNS`, `noip.com`.
     */
    ddnsServer: pulumi.Input<string>;
    /**
     * Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
     */
    ddnsServerAddrs?: pulumi.Input<pulumi.Input<inputs.system.DdnsDdnsServerAddr>[]>;
    /**
     * Generic DDNS server IP.
     */
    ddnsServerIp?: pulumi.Input<string>;
    /**
     * DDNS Serial Number.
     */
    ddnsSn?: pulumi.Input<string>;
    /**
     * Time-to-live for DDNS packets.
     */
    ddnsTtl?: pulumi.Input<number>;
    /**
     * DDNS user name.
     */
    ddnsUsername?: pulumi.Input<string>;
    /**
     * Zone of your domain name (for example, DDNS.com).
     */
    ddnsZone?: pulumi.Input<string>;
    /**
     * DDNS ID.
     */
    ddnsid?: pulumi.Input<number>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Monitored interface. The structure of `monitorInterface` block is documented below.
     */
    monitorInterfaces: pulumi.Input<pulumi.Input<inputs.system.DdnsMonitorInterface>[]>;
    /**
     * Address type of the DDNS server. Valid values: `ipv4`, `ipv6`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Name of local certificate for SSL connections.
     */
    sslCertificate?: pulumi.Input<string>;
    /**
     * DDNS update interval (60 - 2592000 sec, default = 300).
     */
    updateInterval?: pulumi.Input<number>;
    /**
     * Enable/disable use of public IP address. Valid values: `disable`, `enable`.
     */
    usePublicIp?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}