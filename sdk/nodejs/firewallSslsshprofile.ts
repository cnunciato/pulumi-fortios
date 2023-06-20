// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure SSL/SSH protocol options.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const t1 = new fortios.FirewallSslsshprofile("t1", {
 *     ftps: {
 *         ports: "990",
 *     },
 *     https: {
 *         ports: "443 127 422 392",
 *     },
 *     imaps: {
 *         ports: "993 1123",
 *     },
 *     pop3s: {
 *         ports: "995",
 *     },
 *     smtps: {
 *         ports: "465",
 *     },
 *     ssl: {
 *         inspectAll: "disable",
 *     },
 * });
 * const t2 = new fortios.FirewallSslsshprofile("t2", {
 *     https: {
 *         ports: "443",
 *     },
 *     ssl: {
 *         inspectAll: "deep-inspection",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Firewall SslSshProfile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSslsshprofile:FirewallSslsshprofile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/firewallSslsshprofile:FirewallSslsshprofile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class FirewallSslsshprofile extends pulumi.CustomResource {
    /**
     * Get an existing FirewallSslsshprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallSslsshprofileState, opts?: pulumi.CustomResourceOptions): FirewallSslsshprofile {
        return new FirewallSslsshprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/firewallSslsshprofile:FirewallSslsshprofile';

    /**
     * Returns true if the given object is an instance of FirewallSslsshprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallSslsshprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallSslsshprofile.__pulumiType;
    }

    /**
     * Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
     */
    public readonly allowlist!: pulumi.Output<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
     */
    public readonly blockBlacklistedCertificates!: pulumi.Output<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
     */
    public readonly blockBlocklistedCertificates!: pulumi.Output<string>;
    /**
     * CA certificate used by SSL Inspection.
     */
    public readonly caname!: pulumi.Output<string>;
    /**
     * Optional comments.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Configure DNS over TLS options. The structure of `dot` block is documented below.
     */
    public readonly dot!: pulumi.Output<outputs.FirewallSslsshprofileDot>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configure FTPS options. The structure of `ftps` block is documented below.
     */
    public readonly ftps!: pulumi.Output<outputs.FirewallSslsshprofileFtps>;
    /**
     * Configure HTTPS options. The structure of `https` block is documented below.
     */
    public readonly https!: pulumi.Output<outputs.FirewallSslsshprofileHttps>;
    /**
     * Configure IMAPS options. The structure of `imaps` block is documented below.
     */
    public readonly imaps!: pulumi.Output<outputs.FirewallSslsshprofileImaps>;
    /**
     * Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly mapiOverHttps!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configure POP3S options. The structure of `pop3s` block is documented below.
     */
    public readonly pop3s!: pulumi.Output<outputs.FirewallSslsshprofilePop3s>;
    /**
     * Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
     */
    public readonly rpcOverHttps!: pulumi.Output<string>;
    /**
     * Certificate used by SSL Inspection to replace server certificate.
     */
    public readonly serverCert!: pulumi.Output<string>;
    /**
     * Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
     */
    public readonly serverCertMode!: pulumi.Output<string>;
    /**
     * Configure SMTPS options. The structure of `smtps` block is documented below.
     */
    public readonly smtps!: pulumi.Output<outputs.FirewallSslsshprofileSmtps>;
    /**
     * Configure SSH options. The structure of `ssh` block is documented below.
     */
    public readonly ssh!: pulumi.Output<outputs.FirewallSslsshprofileSsh>;
    /**
     * Configure SSL options. The structure of `ssl` block is documented below.
     */
    public readonly ssl!: pulumi.Output<outputs.FirewallSslsshprofileSsl>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    public readonly sslAnomaliesLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    public readonly sslAnomalyLog!: pulumi.Output<string>;
    /**
     * Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
     */
    public readonly sslExemptionIpRating!: pulumi.Output<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    public readonly sslExemptionLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    public readonly sslExemptionsLog!: pulumi.Output<string>;
    /**
     * Servers to exempt from SSL inspection. The structure of `sslExempt` block is documented below.
     */
    public readonly sslExempts!: pulumi.Output<outputs.FirewallSslsshprofileSslExempt[] | undefined>;
    /**
     * Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
     */
    public readonly sslHandshakeLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
     */
    public readonly sslNegotiationLog!: pulumi.Output<string>;
    /**
     * Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
     */
    public readonly sslServerCertLog!: pulumi.Output<string>;
    /**
     * SSL servers. The structure of `sslServer` block is documented below.
     */
    public readonly sslServers!: pulumi.Output<outputs.FirewallSslsshprofileSslServer[] | undefined>;
    /**
     * Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
     */
    public readonly supportedAlpn!: pulumi.Output<string>;
    /**
     * Untrusted CA certificate used by SSL Inspection.
     */
    public readonly untrustedCaname!: pulumi.Output<string>;
    /**
     * Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
     */
    public readonly useSslServer!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
     */
    public readonly whitelist!: pulumi.Output<string>;

    /**
     * Create a FirewallSslsshprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FirewallSslsshprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallSslsshprofileArgs | FirewallSslsshprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallSslsshprofileState | undefined;
            resourceInputs["allowlist"] = state ? state.allowlist : undefined;
            resourceInputs["blockBlacklistedCertificates"] = state ? state.blockBlacklistedCertificates : undefined;
            resourceInputs["blockBlocklistedCertificates"] = state ? state.blockBlocklistedCertificates : undefined;
            resourceInputs["caname"] = state ? state.caname : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dot"] = state ? state.dot : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["ftps"] = state ? state.ftps : undefined;
            resourceInputs["https"] = state ? state.https : undefined;
            resourceInputs["imaps"] = state ? state.imaps : undefined;
            resourceInputs["mapiOverHttps"] = state ? state.mapiOverHttps : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["pop3s"] = state ? state.pop3s : undefined;
            resourceInputs["rpcOverHttps"] = state ? state.rpcOverHttps : undefined;
            resourceInputs["serverCert"] = state ? state.serverCert : undefined;
            resourceInputs["serverCertMode"] = state ? state.serverCertMode : undefined;
            resourceInputs["smtps"] = state ? state.smtps : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["ssl"] = state ? state.ssl : undefined;
            resourceInputs["sslAnomaliesLog"] = state ? state.sslAnomaliesLog : undefined;
            resourceInputs["sslAnomalyLog"] = state ? state.sslAnomalyLog : undefined;
            resourceInputs["sslExemptionIpRating"] = state ? state.sslExemptionIpRating : undefined;
            resourceInputs["sslExemptionLog"] = state ? state.sslExemptionLog : undefined;
            resourceInputs["sslExemptionsLog"] = state ? state.sslExemptionsLog : undefined;
            resourceInputs["sslExempts"] = state ? state.sslExempts : undefined;
            resourceInputs["sslHandshakeLog"] = state ? state.sslHandshakeLog : undefined;
            resourceInputs["sslNegotiationLog"] = state ? state.sslNegotiationLog : undefined;
            resourceInputs["sslServerCertLog"] = state ? state.sslServerCertLog : undefined;
            resourceInputs["sslServers"] = state ? state.sslServers : undefined;
            resourceInputs["supportedAlpn"] = state ? state.supportedAlpn : undefined;
            resourceInputs["untrustedCaname"] = state ? state.untrustedCaname : undefined;
            resourceInputs["useSslServer"] = state ? state.useSslServer : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["whitelist"] = state ? state.whitelist : undefined;
        } else {
            const args = argsOrState as FirewallSslsshprofileArgs | undefined;
            resourceInputs["allowlist"] = args ? args.allowlist : undefined;
            resourceInputs["blockBlacklistedCertificates"] = args ? args.blockBlacklistedCertificates : undefined;
            resourceInputs["blockBlocklistedCertificates"] = args ? args.blockBlocklistedCertificates : undefined;
            resourceInputs["caname"] = args ? args.caname : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dot"] = args ? args.dot : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["ftps"] = args ? args.ftps : undefined;
            resourceInputs["https"] = args ? args.https : undefined;
            resourceInputs["imaps"] = args ? args.imaps : undefined;
            resourceInputs["mapiOverHttps"] = args ? args.mapiOverHttps : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pop3s"] = args ? args.pop3s : undefined;
            resourceInputs["rpcOverHttps"] = args ? args.rpcOverHttps : undefined;
            resourceInputs["serverCert"] = args ? args.serverCert : undefined;
            resourceInputs["serverCertMode"] = args ? args.serverCertMode : undefined;
            resourceInputs["smtps"] = args ? args.smtps : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["ssl"] = args ? args.ssl : undefined;
            resourceInputs["sslAnomaliesLog"] = args ? args.sslAnomaliesLog : undefined;
            resourceInputs["sslAnomalyLog"] = args ? args.sslAnomalyLog : undefined;
            resourceInputs["sslExemptionIpRating"] = args ? args.sslExemptionIpRating : undefined;
            resourceInputs["sslExemptionLog"] = args ? args.sslExemptionLog : undefined;
            resourceInputs["sslExemptionsLog"] = args ? args.sslExemptionsLog : undefined;
            resourceInputs["sslExempts"] = args ? args.sslExempts : undefined;
            resourceInputs["sslHandshakeLog"] = args ? args.sslHandshakeLog : undefined;
            resourceInputs["sslNegotiationLog"] = args ? args.sslNegotiationLog : undefined;
            resourceInputs["sslServerCertLog"] = args ? args.sslServerCertLog : undefined;
            resourceInputs["sslServers"] = args ? args.sslServers : undefined;
            resourceInputs["supportedAlpn"] = args ? args.supportedAlpn : undefined;
            resourceInputs["untrustedCaname"] = args ? args.untrustedCaname : undefined;
            resourceInputs["useSslServer"] = args ? args.useSslServer : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["whitelist"] = args ? args.whitelist : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallSslsshprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallSslsshprofile resources.
 */
export interface FirewallSslsshprofileState {
    /**
     * Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
     */
    allowlist?: pulumi.Input<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
     */
    blockBlacklistedCertificates?: pulumi.Input<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
     */
    blockBlocklistedCertificates?: pulumi.Input<string>;
    /**
     * CA certificate used by SSL Inspection.
     */
    caname?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Configure DNS over TLS options. The structure of `dot` block is documented below.
     */
    dot?: pulumi.Input<inputs.FirewallSslsshprofileDot>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure FTPS options. The structure of `ftps` block is documented below.
     */
    ftps?: pulumi.Input<inputs.FirewallSslsshprofileFtps>;
    /**
     * Configure HTTPS options. The structure of `https` block is documented below.
     */
    https?: pulumi.Input<inputs.FirewallSslsshprofileHttps>;
    /**
     * Configure IMAPS options. The structure of `imaps` block is documented below.
     */
    imaps?: pulumi.Input<inputs.FirewallSslsshprofileImaps>;
    /**
     * Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
     */
    mapiOverHttps?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure POP3S options. The structure of `pop3s` block is documented below.
     */
    pop3s?: pulumi.Input<inputs.FirewallSslsshprofilePop3s>;
    /**
     * Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
     */
    rpcOverHttps?: pulumi.Input<string>;
    /**
     * Certificate used by SSL Inspection to replace server certificate.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
     */
    serverCertMode?: pulumi.Input<string>;
    /**
     * Configure SMTPS options. The structure of `smtps` block is documented below.
     */
    smtps?: pulumi.Input<inputs.FirewallSslsshprofileSmtps>;
    /**
     * Configure SSH options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.FirewallSslsshprofileSsh>;
    /**
     * Configure SSL options. The structure of `ssl` block is documented below.
     */
    ssl?: pulumi.Input<inputs.FirewallSslsshprofileSsl>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    sslAnomaliesLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    sslAnomalyLog?: pulumi.Input<string>;
    /**
     * Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
     */
    sslExemptionIpRating?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    sslExemptionLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    sslExemptionsLog?: pulumi.Input<string>;
    /**
     * Servers to exempt from SSL inspection. The structure of `sslExempt` block is documented below.
     */
    sslExempts?: pulumi.Input<pulumi.Input<inputs.FirewallSslsshprofileSslExempt>[]>;
    /**
     * Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
     */
    sslHandshakeLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
     */
    sslNegotiationLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
     */
    sslServerCertLog?: pulumi.Input<string>;
    /**
     * SSL servers. The structure of `sslServer` block is documented below.
     */
    sslServers?: pulumi.Input<pulumi.Input<inputs.FirewallSslsshprofileSslServer>[]>;
    /**
     * Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
     */
    supportedAlpn?: pulumi.Input<string>;
    /**
     * Untrusted CA certificate used by SSL Inspection.
     */
    untrustedCaname?: pulumi.Input<string>;
    /**
     * Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
     */
    useSslServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
     */
    whitelist?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallSslsshprofile resource.
 */
export interface FirewallSslsshprofileArgs {
    /**
     * Enable/disable exempting servers by FortiGuard allowlist. Valid values: `enable`, `disable`.
     */
    allowlist?: pulumi.Input<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blacklist. Valid values: `disable`, `enable`.
     */
    blockBlacklistedCertificates?: pulumi.Input<string>;
    /**
     * Enable/disable blocking SSL-based botnet communication by FortiGuard certificate blocklist. Valid values: `disable`, `enable`.
     */
    blockBlocklistedCertificates?: pulumi.Input<string>;
    /**
     * CA certificate used by SSL Inspection.
     */
    caname?: pulumi.Input<string>;
    /**
     * Optional comments.
     */
    comment?: pulumi.Input<string>;
    /**
     * Configure DNS over TLS options. The structure of `dot` block is documented below.
     */
    dot?: pulumi.Input<inputs.FirewallSslsshprofileDot>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configure FTPS options. The structure of `ftps` block is documented below.
     */
    ftps?: pulumi.Input<inputs.FirewallSslsshprofileFtps>;
    /**
     * Configure HTTPS options. The structure of `https` block is documented below.
     */
    https?: pulumi.Input<inputs.FirewallSslsshprofileHttps>;
    /**
     * Configure IMAPS options. The structure of `imaps` block is documented below.
     */
    imaps?: pulumi.Input<inputs.FirewallSslsshprofileImaps>;
    /**
     * Enable/disable inspection of MAPI over HTTPS. Valid values: `enable`, `disable`.
     */
    mapiOverHttps?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure POP3S options. The structure of `pop3s` block is documented below.
     */
    pop3s?: pulumi.Input<inputs.FirewallSslsshprofilePop3s>;
    /**
     * Enable/disable inspection of RPC over HTTPS. Valid values: `enable`, `disable`.
     */
    rpcOverHttps?: pulumi.Input<string>;
    /**
     * Certificate used by SSL Inspection to replace server certificate.
     */
    serverCert?: pulumi.Input<string>;
    /**
     * Re-sign or replace the server's certificate. Valid values: `re-sign`, `replace`.
     */
    serverCertMode?: pulumi.Input<string>;
    /**
     * Configure SMTPS options. The structure of `smtps` block is documented below.
     */
    smtps?: pulumi.Input<inputs.FirewallSslsshprofileSmtps>;
    /**
     * Configure SSH options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.FirewallSslsshprofileSsh>;
    /**
     * Configure SSL options. The structure of `ssl` block is documented below.
     */
    ssl?: pulumi.Input<inputs.FirewallSslsshprofileSsl>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    sslAnomaliesLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL anomalies. Valid values: `disable`, `enable`.
     */
    sslAnomalyLog?: pulumi.Input<string>;
    /**
     * Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
     */
    sslExemptionIpRating?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    sslExemptionLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL exemptions. Valid values: `disable`, `enable`.
     */
    sslExemptionsLog?: pulumi.Input<string>;
    /**
     * Servers to exempt from SSL inspection. The structure of `sslExempt` block is documented below.
     */
    sslExempts?: pulumi.Input<pulumi.Input<inputs.FirewallSslsshprofileSslExempt>[]>;
    /**
     * Enable/disable logging of TLS handshakes. Valid values: `disable`, `enable`.
     */
    sslHandshakeLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging SSL negotiation. Valid values: `disable`, `enable`.
     */
    sslNegotiationLog?: pulumi.Input<string>;
    /**
     * Enable/disable logging of server certificate information. Valid values: `disable`, `enable`.
     */
    sslServerCertLog?: pulumi.Input<string>;
    /**
     * SSL servers. The structure of `sslServer` block is documented below.
     */
    sslServers?: pulumi.Input<pulumi.Input<inputs.FirewallSslsshprofileSslServer>[]>;
    /**
     * Configure ALPN option. Valid values: `http1-1`, `http2`, `all`, `none`.
     */
    supportedAlpn?: pulumi.Input<string>;
    /**
     * Untrusted CA certificate used by SSL Inspection.
     */
    untrustedCaname?: pulumi.Input<string>;
    /**
     * Enable/disable the use of SSL server table for SSL offloading. Valid values: `disable`, `enable`.
     */
    useSslServer?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Enable/disable exempting servers by FortiGuard whitelist. Valid values: `enable`, `disable`.
     */
    whitelist?: pulumi.Input<string>;
}
