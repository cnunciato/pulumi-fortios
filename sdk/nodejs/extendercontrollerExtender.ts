// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Extender controller configuration.
 *
 * > The resource applies to FortiOS Version < 6.4.2. For FortiOS Version >= 6.4.2, see `fortios.ExtendercontrollerExtender1`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.ExtendercontrollerExtender("trname", {
 *     admin: "disable",
 *     billingStartDay: 1,
 *     connStatus: 0,
 *     dialMode: "always-connect",
 *     dialStatus: 0,
 *     extName: "332",
 *     fosid: "1",
 *     initiatedUpdate: "disable",
 *     mode: "standalone",
 *     modemType: "gsm/lte",
 *     multiMode: "auto",
 *     pppAuthProtocol: "auto",
 *     pppEchoRequest: "disable",
 *     quotaLimitMb: 0,
 *     redial: "none",
 *     roaming: "disable",
 *     role: "primary",
 *     vdom: 0,
 *     wimaxAuthProtocol: "tls",
 * });
 * ```
 *
 * ## Import
 *
 * ExtenderController Extender can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/extendercontrollerExtender:ExtendercontrollerExtender labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/extendercontrollerExtender:ExtendercontrollerExtender labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class ExtendercontrollerExtender extends pulumi.CustomResource {
    /**
     * Get an existing ExtendercontrollerExtender resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtendercontrollerExtenderState, opts?: pulumi.CustomResourceOptions): ExtendercontrollerExtender {
        return new ExtendercontrollerExtender(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/extendercontrollerExtender:ExtendercontrollerExtender';

    /**
     * Returns true if the given object is an instance of ExtendercontrollerExtender.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExtendercontrollerExtender {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExtendercontrollerExtender.__pulumiType;
    }

    /**
     * AAA shared secret.
     */
    public readonly aaaSharedSecret!: pulumi.Output<string | undefined>;
    /**
     * Access point name(APN).
     */
    public readonly accessPointName!: pulumi.Output<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `discovered`, `enable`.
     */
    public readonly admin!: pulumi.Output<string>;
    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * Initialization AT commands specific to the MODEM.
     */
    public readonly atDialScript!: pulumi.Output<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
     */
    public readonly authorized!: pulumi.Output<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    public readonly bandwidthLimit!: pulumi.Output<number>;
    /**
     * Billing start day.
     */
    public readonly billingStartDay!: pulumi.Output<number>;
    /**
     * CDMA AAA SPI.
     */
    public readonly cdmaAaaSpi!: pulumi.Output<string>;
    /**
     * CDMA HA SPI.
     */
    public readonly cdmaHaSpi!: pulumi.Output<string>;
    /**
     * NAI for CDMA MODEMS.
     */
    public readonly cdmaNai!: pulumi.Output<string>;
    /**
     * Connection status.
     */
    public readonly connStatus!: pulumi.Output<number>;
    /**
     * FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
     */
    public readonly controllerReport!: pulumi.Output<outputs.ExtendercontrollerExtenderControllerReport>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * device-id
     */
    public readonly deviceId!: pulumi.Output<number>;
    /**
     * Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand`, `always-connect`.
     */
    public readonly dialMode!: pulumi.Output<string>;
    /**
     * Dial status.
     */
    public readonly dialStatus!: pulumi.Output<number>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    public readonly enforceBandwidth!: pulumi.Output<string>;
    /**
     * FortiExtender name.
     */
    public readonly extName!: pulumi.Output<string>;
    /**
     * Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
     */
    public readonly extensionType!: pulumi.Output<string>;
    /**
     * FortiExtender serial number.
     */
    public readonly fosid!: pulumi.Output<string>;
    /**
     * HA shared secret.
     */
    public readonly haSharedSecret!: pulumi.Output<string | undefined>;
    /**
     * FortiExtender interface name.
     */
    public readonly ifname!: pulumi.Output<string>;
    /**
     * Allow/disallow network initiated updates to the MODEM. Valid values: `enable`, `disable`.
     */
    public readonly initiatedUpdate!: pulumi.Output<string>;
    /**
     * FortiExtender login password.
     */
    public readonly loginPassword!: pulumi.Output<string | undefined>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    public readonly loginPasswordChange!: pulumi.Output<string>;
    /**
     * FortiExtender mode. Valid values: `standalone`, `redundant`.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Configuration options for modem 1. The structure of `modem1` block is documented below.
     */
    public readonly modem1!: pulumi.Output<outputs.ExtendercontrollerExtenderModem1>;
    /**
     * Configuration options for modem 2. The structure of `modem2` block is documented below.
     */
    public readonly modem2!: pulumi.Output<outputs.ExtendercontrollerExtenderModem2>;
    /**
     * MODEM password.
     */
    public readonly modemPasswd!: pulumi.Output<string | undefined>;
    /**
     * MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma`, `gsm/lte`, `wimax`.
     */
    public readonly modemType!: pulumi.Output<string>;
    /**
     * MODEM mode of operation(3G,LTE,etc). Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.
     */
    public readonly multiMode!: pulumi.Output<string>;
    /**
     * FortiExtender entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
     */
    public readonly overrideAllowaccess!: pulumi.Output<string>;
    /**
     * Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideEnforceBandwidth!: pulumi.Output<string>;
    /**
     * Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    public readonly overrideLoginPasswordChange!: pulumi.Output<string>;
    /**
     * PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto`, `pap`, `chap`.
     */
    public readonly pppAuthProtocol!: pulumi.Output<string>;
    /**
     * Enable/disable PPP echo request. Valid values: `enable`, `disable`.
     */
    public readonly pppEchoRequest!: pulumi.Output<string>;
    /**
     * PPP password.
     */
    public readonly pppPassword!: pulumi.Output<string | undefined>;
    /**
     * PPP username.
     */
    public readonly pppUsername!: pulumi.Output<string>;
    /**
     * Primary HA.
     */
    public readonly primaryHa!: pulumi.Output<string>;
    /**
     * FortiExtender profile configuration.
     */
    public readonly profile!: pulumi.Output<string>;
    /**
     * Monthly quota limit (MB).
     */
    public readonly quotaLimitMb!: pulumi.Output<number>;
    /**
     * Number of redials allowed based on failed attempts. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
     */
    public readonly redial!: pulumi.Output<string>;
    /**
     * Redundant interface.
     */
    public readonly redundantIntf!: pulumi.Output<string>;
    /**
     * Enable/disable MODEM roaming. Valid values: `enable`, `disable`.
     */
    public readonly roaming!: pulumi.Output<string>;
    /**
     * FortiExtender work role(Primary, Secondary, None). Valid values: `none`, `primary`, `secondary`.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Secondary HA.
     */
    public readonly secondaryHa!: pulumi.Output<string>;
    /**
     * SIM PIN.
     */
    public readonly simPin!: pulumi.Output<string | undefined>;
    /**
     * VDOM
     */
    public readonly vdom!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
     */
    public readonly wanExtension!: pulumi.Output<outputs.ExtendercontrollerExtenderWanExtension>;
    /**
     * WiMax authentication protocol(TLS or TTLS). Valid values: `tls`, `ttls`.
     */
    public readonly wimaxAuthProtocol!: pulumi.Output<string>;
    /**
     * WiMax carrier.
     */
    public readonly wimaxCarrier!: pulumi.Output<string>;
    /**
     * WiMax realm.
     */
    public readonly wimaxRealm!: pulumi.Output<string>;

    /**
     * Create a ExtendercontrollerExtender resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtendercontrollerExtenderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtendercontrollerExtenderArgs | ExtendercontrollerExtenderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExtendercontrollerExtenderState | undefined;
            resourceInputs["aaaSharedSecret"] = state ? state.aaaSharedSecret : undefined;
            resourceInputs["accessPointName"] = state ? state.accessPointName : undefined;
            resourceInputs["admin"] = state ? state.admin : undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["atDialScript"] = state ? state.atDialScript : undefined;
            resourceInputs["authorized"] = state ? state.authorized : undefined;
            resourceInputs["bandwidthLimit"] = state ? state.bandwidthLimit : undefined;
            resourceInputs["billingStartDay"] = state ? state.billingStartDay : undefined;
            resourceInputs["cdmaAaaSpi"] = state ? state.cdmaAaaSpi : undefined;
            resourceInputs["cdmaHaSpi"] = state ? state.cdmaHaSpi : undefined;
            resourceInputs["cdmaNai"] = state ? state.cdmaNai : undefined;
            resourceInputs["connStatus"] = state ? state.connStatus : undefined;
            resourceInputs["controllerReport"] = state ? state.controllerReport : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["dialMode"] = state ? state.dialMode : undefined;
            resourceInputs["dialStatus"] = state ? state.dialStatus : undefined;
            resourceInputs["enforceBandwidth"] = state ? state.enforceBandwidth : undefined;
            resourceInputs["extName"] = state ? state.extName : undefined;
            resourceInputs["extensionType"] = state ? state.extensionType : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["haSharedSecret"] = state ? state.haSharedSecret : undefined;
            resourceInputs["ifname"] = state ? state.ifname : undefined;
            resourceInputs["initiatedUpdate"] = state ? state.initiatedUpdate : undefined;
            resourceInputs["loginPassword"] = state ? state.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = state ? state.loginPasswordChange : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["modem1"] = state ? state.modem1 : undefined;
            resourceInputs["modem2"] = state ? state.modem2 : undefined;
            resourceInputs["modemPasswd"] = state ? state.modemPasswd : undefined;
            resourceInputs["modemType"] = state ? state.modemType : undefined;
            resourceInputs["multiMode"] = state ? state.multiMode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideAllowaccess"] = state ? state.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = state ? state.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = state ? state.overrideLoginPasswordChange : undefined;
            resourceInputs["pppAuthProtocol"] = state ? state.pppAuthProtocol : undefined;
            resourceInputs["pppEchoRequest"] = state ? state.pppEchoRequest : undefined;
            resourceInputs["pppPassword"] = state ? state.pppPassword : undefined;
            resourceInputs["pppUsername"] = state ? state.pppUsername : undefined;
            resourceInputs["primaryHa"] = state ? state.primaryHa : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["quotaLimitMb"] = state ? state.quotaLimitMb : undefined;
            resourceInputs["redial"] = state ? state.redial : undefined;
            resourceInputs["redundantIntf"] = state ? state.redundantIntf : undefined;
            resourceInputs["roaming"] = state ? state.roaming : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["secondaryHa"] = state ? state.secondaryHa : undefined;
            resourceInputs["simPin"] = state ? state.simPin : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["wanExtension"] = state ? state.wanExtension : undefined;
            resourceInputs["wimaxAuthProtocol"] = state ? state.wimaxAuthProtocol : undefined;
            resourceInputs["wimaxCarrier"] = state ? state.wimaxCarrier : undefined;
            resourceInputs["wimaxRealm"] = state ? state.wimaxRealm : undefined;
        } else {
            const args = argsOrState as ExtendercontrollerExtenderArgs | undefined;
            if ((!args || args.admin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'admin'");
            }
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["aaaSharedSecret"] = args?.aaaSharedSecret ? pulumi.secret(args.aaaSharedSecret) : undefined;
            resourceInputs["accessPointName"] = args ? args.accessPointName : undefined;
            resourceInputs["admin"] = args ? args.admin : undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["atDialScript"] = args ? args.atDialScript : undefined;
            resourceInputs["authorized"] = args ? args.authorized : undefined;
            resourceInputs["bandwidthLimit"] = args ? args.bandwidthLimit : undefined;
            resourceInputs["billingStartDay"] = args ? args.billingStartDay : undefined;
            resourceInputs["cdmaAaaSpi"] = args ? args.cdmaAaaSpi : undefined;
            resourceInputs["cdmaHaSpi"] = args ? args.cdmaHaSpi : undefined;
            resourceInputs["cdmaNai"] = args ? args.cdmaNai : undefined;
            resourceInputs["connStatus"] = args ? args.connStatus : undefined;
            resourceInputs["controllerReport"] = args ? args.controllerReport : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["dialMode"] = args ? args.dialMode : undefined;
            resourceInputs["dialStatus"] = args ? args.dialStatus : undefined;
            resourceInputs["enforceBandwidth"] = args ? args.enforceBandwidth : undefined;
            resourceInputs["extName"] = args ? args.extName : undefined;
            resourceInputs["extensionType"] = args ? args.extensionType : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["haSharedSecret"] = args?.haSharedSecret ? pulumi.secret(args.haSharedSecret) : undefined;
            resourceInputs["ifname"] = args ? args.ifname : undefined;
            resourceInputs["initiatedUpdate"] = args ? args.initiatedUpdate : undefined;
            resourceInputs["loginPassword"] = args ? args.loginPassword : undefined;
            resourceInputs["loginPasswordChange"] = args ? args.loginPasswordChange : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["modem1"] = args ? args.modem1 : undefined;
            resourceInputs["modem2"] = args ? args.modem2 : undefined;
            resourceInputs["modemPasswd"] = args?.modemPasswd ? pulumi.secret(args.modemPasswd) : undefined;
            resourceInputs["modemType"] = args ? args.modemType : undefined;
            resourceInputs["multiMode"] = args ? args.multiMode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideAllowaccess"] = args ? args.overrideAllowaccess : undefined;
            resourceInputs["overrideEnforceBandwidth"] = args ? args.overrideEnforceBandwidth : undefined;
            resourceInputs["overrideLoginPasswordChange"] = args ? args.overrideLoginPasswordChange : undefined;
            resourceInputs["pppAuthProtocol"] = args ? args.pppAuthProtocol : undefined;
            resourceInputs["pppEchoRequest"] = args ? args.pppEchoRequest : undefined;
            resourceInputs["pppPassword"] = args?.pppPassword ? pulumi.secret(args.pppPassword) : undefined;
            resourceInputs["pppUsername"] = args ? args.pppUsername : undefined;
            resourceInputs["primaryHa"] = args ? args.primaryHa : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["quotaLimitMb"] = args ? args.quotaLimitMb : undefined;
            resourceInputs["redial"] = args ? args.redial : undefined;
            resourceInputs["redundantIntf"] = args ? args.redundantIntf : undefined;
            resourceInputs["roaming"] = args ? args.roaming : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["secondaryHa"] = args ? args.secondaryHa : undefined;
            resourceInputs["simPin"] = args?.simPin ? pulumi.secret(args.simPin) : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["wanExtension"] = args ? args.wanExtension : undefined;
            resourceInputs["wimaxAuthProtocol"] = args ? args.wimaxAuthProtocol : undefined;
            resourceInputs["wimaxCarrier"] = args ? args.wimaxCarrier : undefined;
            resourceInputs["wimaxRealm"] = args ? args.wimaxRealm : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["aaaSharedSecret", "haSharedSecret", "modemPasswd", "pppPassword", "simPin"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ExtendercontrollerExtender.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExtendercontrollerExtender resources.
 */
export interface ExtendercontrollerExtenderState {
    /**
     * AAA shared secret.
     */
    aaaSharedSecret?: pulumi.Input<string>;
    /**
     * Access point name(APN).
     */
    accessPointName?: pulumi.Input<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `discovered`, `enable`.
     */
    admin?: pulumi.Input<string>;
    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Initialization AT commands specific to the MODEM.
     */
    atDialScript?: pulumi.Input<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
     */
    authorized?: pulumi.Input<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * Billing start day.
     */
    billingStartDay?: pulumi.Input<number>;
    /**
     * CDMA AAA SPI.
     */
    cdmaAaaSpi?: pulumi.Input<string>;
    /**
     * CDMA HA SPI.
     */
    cdmaHaSpi?: pulumi.Input<string>;
    /**
     * NAI for CDMA MODEMS.
     */
    cdmaNai?: pulumi.Input<string>;
    /**
     * Connection status.
     */
    connStatus?: pulumi.Input<number>;
    /**
     * FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
     */
    controllerReport?: pulumi.Input<inputs.ExtendercontrollerExtenderControllerReport>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * device-id
     */
    deviceId?: pulumi.Input<number>;
    /**
     * Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand`, `always-connect`.
     */
    dialMode?: pulumi.Input<string>;
    /**
     * Dial status.
     */
    dialStatus?: pulumi.Input<number>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    enforceBandwidth?: pulumi.Input<string>;
    /**
     * FortiExtender name.
     */
    extName?: pulumi.Input<string>;
    /**
     * Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
     */
    extensionType?: pulumi.Input<string>;
    /**
     * FortiExtender serial number.
     */
    fosid?: pulumi.Input<string>;
    /**
     * HA shared secret.
     */
    haSharedSecret?: pulumi.Input<string>;
    /**
     * FortiExtender interface name.
     */
    ifname?: pulumi.Input<string>;
    /**
     * Allow/disallow network initiated updates to the MODEM. Valid values: `enable`, `disable`.
     */
    initiatedUpdate?: pulumi.Input<string>;
    /**
     * FortiExtender login password.
     */
    loginPassword?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswordChange?: pulumi.Input<string>;
    /**
     * FortiExtender mode. Valid values: `standalone`, `redundant`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Configuration options for modem 1. The structure of `modem1` block is documented below.
     */
    modem1?: pulumi.Input<inputs.ExtendercontrollerExtenderModem1>;
    /**
     * Configuration options for modem 2. The structure of `modem2` block is documented below.
     */
    modem2?: pulumi.Input<inputs.ExtendercontrollerExtenderModem2>;
    /**
     * MODEM password.
     */
    modemPasswd?: pulumi.Input<string>;
    /**
     * MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma`, `gsm/lte`, `wimax`.
     */
    modemType?: pulumi.Input<string>;
    /**
     * MODEM mode of operation(3G,LTE,etc). Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.
     */
    multiMode?: pulumi.Input<string>;
    /**
     * FortiExtender entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
     */
    overrideAllowaccess?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
     */
    overrideEnforceBandwidth?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    overrideLoginPasswordChange?: pulumi.Input<string>;
    /**
     * PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto`, `pap`, `chap`.
     */
    pppAuthProtocol?: pulumi.Input<string>;
    /**
     * Enable/disable PPP echo request. Valid values: `enable`, `disable`.
     */
    pppEchoRequest?: pulumi.Input<string>;
    /**
     * PPP password.
     */
    pppPassword?: pulumi.Input<string>;
    /**
     * PPP username.
     */
    pppUsername?: pulumi.Input<string>;
    /**
     * Primary HA.
     */
    primaryHa?: pulumi.Input<string>;
    /**
     * FortiExtender profile configuration.
     */
    profile?: pulumi.Input<string>;
    /**
     * Monthly quota limit (MB).
     */
    quotaLimitMb?: pulumi.Input<number>;
    /**
     * Number of redials allowed based on failed attempts. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
     */
    redial?: pulumi.Input<string>;
    /**
     * Redundant interface.
     */
    redundantIntf?: pulumi.Input<string>;
    /**
     * Enable/disable MODEM roaming. Valid values: `enable`, `disable`.
     */
    roaming?: pulumi.Input<string>;
    /**
     * FortiExtender work role(Primary, Secondary, None). Valid values: `none`, `primary`, `secondary`.
     */
    role?: pulumi.Input<string>;
    /**
     * Secondary HA.
     */
    secondaryHa?: pulumi.Input<string>;
    /**
     * SIM PIN.
     */
    simPin?: pulumi.Input<string>;
    /**
     * VDOM
     */
    vdom?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
     */
    wanExtension?: pulumi.Input<inputs.ExtendercontrollerExtenderWanExtension>;
    /**
     * WiMax authentication protocol(TLS or TTLS). Valid values: `tls`, `ttls`.
     */
    wimaxAuthProtocol?: pulumi.Input<string>;
    /**
     * WiMax carrier.
     */
    wimaxCarrier?: pulumi.Input<string>;
    /**
     * WiMax realm.
     */
    wimaxRealm?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExtendercontrollerExtender resource.
 */
export interface ExtendercontrollerExtenderArgs {
    /**
     * AAA shared secret.
     */
    aaaSharedSecret?: pulumi.Input<string>;
    /**
     * Access point name(APN).
     */
    accessPointName?: pulumi.Input<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `discovered`, `enable`.
     */
    admin: pulumi.Input<string>;
    /**
     * Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Initialization AT commands specific to the MODEM.
     */
    atDialScript?: pulumi.Input<string>;
    /**
     * FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
     */
    authorized?: pulumi.Input<string>;
    /**
     * FortiExtender LAN extension bandwidth limit (Mbps).
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * Billing start day.
     */
    billingStartDay?: pulumi.Input<number>;
    /**
     * CDMA AAA SPI.
     */
    cdmaAaaSpi?: pulumi.Input<string>;
    /**
     * CDMA HA SPI.
     */
    cdmaHaSpi?: pulumi.Input<string>;
    /**
     * NAI for CDMA MODEMS.
     */
    cdmaNai?: pulumi.Input<string>;
    /**
     * Connection status.
     */
    connStatus?: pulumi.Input<number>;
    /**
     * FortiExtender controller report configuration. The structure of `controllerReport` block is documented below.
     */
    controllerReport?: pulumi.Input<inputs.ExtendercontrollerExtenderControllerReport>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * device-id
     */
    deviceId?: pulumi.Input<number>;
    /**
     * Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand`, `always-connect`.
     */
    dialMode?: pulumi.Input<string>;
    /**
     * Dial status.
     */
    dialStatus?: pulumi.Input<number>;
    /**
     * Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
     */
    enforceBandwidth?: pulumi.Input<string>;
    /**
     * FortiExtender name.
     */
    extName?: pulumi.Input<string>;
    /**
     * Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
     */
    extensionType?: pulumi.Input<string>;
    /**
     * FortiExtender serial number.
     */
    fosid: pulumi.Input<string>;
    /**
     * HA shared secret.
     */
    haSharedSecret?: pulumi.Input<string>;
    /**
     * FortiExtender interface name.
     */
    ifname?: pulumi.Input<string>;
    /**
     * Allow/disallow network initiated updates to the MODEM. Valid values: `enable`, `disable`.
     */
    initiatedUpdate?: pulumi.Input<string>;
    /**
     * FortiExtender login password.
     */
    loginPassword?: pulumi.Input<string>;
    /**
     * Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
     */
    loginPasswordChange?: pulumi.Input<string>;
    /**
     * FortiExtender mode. Valid values: `standalone`, `redundant`.
     */
    mode?: pulumi.Input<string>;
    /**
     * Configuration options for modem 1. The structure of `modem1` block is documented below.
     */
    modem1?: pulumi.Input<inputs.ExtendercontrollerExtenderModem1>;
    /**
     * Configuration options for modem 2. The structure of `modem2` block is documented below.
     */
    modem2?: pulumi.Input<inputs.ExtendercontrollerExtenderModem2>;
    /**
     * MODEM password.
     */
    modemPasswd?: pulumi.Input<string>;
    /**
     * MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma`, `gsm/lte`, `wimax`.
     */
    modemType?: pulumi.Input<string>;
    /**
     * MODEM mode of operation(3G,LTE,etc). Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.
     */
    multiMode?: pulumi.Input<string>;
    /**
     * FortiExtender entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
     */
    overrideAllowaccess?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
     */
    overrideEnforceBandwidth?: pulumi.Input<string>;
    /**
     * Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
     */
    overrideLoginPasswordChange?: pulumi.Input<string>;
    /**
     * PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto`, `pap`, `chap`.
     */
    pppAuthProtocol?: pulumi.Input<string>;
    /**
     * Enable/disable PPP echo request. Valid values: `enable`, `disable`.
     */
    pppEchoRequest?: pulumi.Input<string>;
    /**
     * PPP password.
     */
    pppPassword?: pulumi.Input<string>;
    /**
     * PPP username.
     */
    pppUsername?: pulumi.Input<string>;
    /**
     * Primary HA.
     */
    primaryHa?: pulumi.Input<string>;
    /**
     * FortiExtender profile configuration.
     */
    profile?: pulumi.Input<string>;
    /**
     * Monthly quota limit (MB).
     */
    quotaLimitMb?: pulumi.Input<number>;
    /**
     * Number of redials allowed based on failed attempts. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
     */
    redial?: pulumi.Input<string>;
    /**
     * Redundant interface.
     */
    redundantIntf?: pulumi.Input<string>;
    /**
     * Enable/disable MODEM roaming. Valid values: `enable`, `disable`.
     */
    roaming?: pulumi.Input<string>;
    /**
     * FortiExtender work role(Primary, Secondary, None). Valid values: `none`, `primary`, `secondary`.
     */
    role: pulumi.Input<string>;
    /**
     * Secondary HA.
     */
    secondaryHa?: pulumi.Input<string>;
    /**
     * SIM PIN.
     */
    simPin?: pulumi.Input<string>;
    /**
     * VDOM
     */
    vdom?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * FortiExtender wan extension configuration. The structure of `wanExtension` block is documented below.
     */
    wanExtension?: pulumi.Input<inputs.ExtendercontrollerExtenderWanExtension>;
    /**
     * WiMax authentication protocol(TLS or TTLS). Valid values: `tls`, `ttls`.
     */
    wimaxAuthProtocol?: pulumi.Input<string>;
    /**
     * WiMax carrier.
     */
    wimaxCarrier?: pulumi.Input<string>;
    /**
     * WiMax realm.
     */
    wimaxRealm?: pulumi.Input<string>;
}
