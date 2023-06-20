// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure AntiVirus profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.AntivirusProfile("trname", {
 *     analyticsBlFiletype: 0,
 *     analyticsDb: "disable",
 *     analyticsMaxUpload: 10,
 *     analyticsWlFiletype: 0,
 *     avBlockLog: "enable",
 *     avVirusLog: "enable",
 *     extendedLog: "disable",
 *     ftgdAnalytics: "disable",
 *     inspectionMode: "flow-based",
 *     mobileMalwareDb: "enable",
 *     scanMode: "quick",
 * });
 * ```
 *
 * ## Import
 *
 * Antivirus Profile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/antivirusProfile:AntivirusProfile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/antivirusProfile:AntivirusProfile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class AntivirusProfile extends pulumi.CustomResource {
    /**
     * Get an existing AntivirusProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AntivirusProfileState, opts?: pulumi.CustomResourceOptions): AntivirusProfile {
        return new AntivirusProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/antivirusProfile:AntivirusProfile';

    /**
     * Returns true if the given object is an instance of AntivirusProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AntivirusProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AntivirusProfile.__pulumiType;
    }

    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    public readonly analyticsAcceptFiletype!: pulumi.Output<number>;
    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    public readonly analyticsBlFiletype!: pulumi.Output<number>;
    /**
     * Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
     */
    public readonly analyticsDb!: pulumi.Output<string>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    public readonly analyticsIgnoreFiletype!: pulumi.Output<number>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
     */
    public readonly analyticsMaxUpload!: pulumi.Output<number>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    public readonly analyticsWlFiletype!: pulumi.Output<number>;
    /**
     * Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
     */
    public readonly avBlockLog!: pulumi.Output<string>;
    /**
     * Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
     */
    public readonly avVirusLog!: pulumi.Output<string>;
    /**
     * Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
     */
    public readonly cifs!: pulumi.Output<outputs.AntivirusProfileCifs>;
    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * AV Content Disarm and Reconstruction settings. The structure of `contentDisarm` block is documented below.
     */
    public readonly contentDisarm!: pulumi.Output<outputs.AntivirusProfileContentDisarm>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
     */
    public readonly emsThreatFeed!: pulumi.Output<string>;
    /**
     * Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
     */
    public readonly extendedLog!: pulumi.Output<string>;
    /**
     * Enable/disable all external blocklists. Valid values: `disable`, `enable`.
     */
    public readonly externalBlocklistEnableAll!: pulumi.Output<string>;
    /**
     * One or more external malware block lists. The structure of `externalBlocklist` block is documented below.
     */
    public readonly externalBlocklists!: pulumi.Output<outputs.AntivirusProfileExternalBlocklist[] | undefined>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    public readonly featureSet!: pulumi.Output<string>;
    /**
     * Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortiaiErrorAction!: pulumi.Output<string>;
    /**
     * Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortiaiTimeoutAction!: pulumi.Output<string>;
    /**
     * Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortindrErrorAction!: pulumi.Output<string>;
    /**
     * Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortindrTimeoutAction!: pulumi.Output<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortisandboxErrorAction!: pulumi.Output<string>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox.
     */
    public readonly fortisandboxMaxUpload!: pulumi.Output<number>;
    /**
     * FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
     */
    public readonly fortisandboxMode!: pulumi.Output<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    public readonly fortisandboxTimeoutAction!: pulumi.Output<string>;
    /**
     * Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
     */
    public readonly ftgdAnalytics!: pulumi.Output<string>;
    /**
     * Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
     */
    public readonly ftp!: pulumi.Output<outputs.AntivirusProfileFtp>;
    /**
     * Configure HTTP AntiVirus options. The structure of `http` block is documented below.
     */
    public readonly http!: pulumi.Output<outputs.AntivirusProfileHttp>;
    /**
     * Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
     */
    public readonly imap!: pulumi.Output<outputs.AntivirusProfileImap>;
    /**
     * Inspection mode. Valid values: `proxy`, `flow-based`.
     */
    public readonly inspectionMode!: pulumi.Output<string>;
    /**
     * Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
     */
    public readonly mapi!: pulumi.Output<outputs.AntivirusProfileMapi>;
    /**
     * Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
     */
    public readonly mobileMalwareDb!: pulumi.Output<string>;
    /**
     * Configure AntiVirus quarantine settings. The structure of `nacQuar` block is documented below.
     */
    public readonly nacQuar!: pulumi.Output<outputs.AntivirusProfileNacQuar>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
     */
    public readonly nntp!: pulumi.Output<outputs.AntivirusProfileNntp>;
    /**
     * Configure Virus Outbreak Prevention settings. The structure of `outbreakPrevention` block is documented below.
     */
    public readonly outbreakPrevention!: pulumi.Output<outputs.AntivirusProfileOutbreakPrevention>;
    /**
     * Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
     */
    public readonly outbreakPreventionArchiveScan!: pulumi.Output<string>;
    /**
     * Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
     */
    public readonly pop3!: pulumi.Output<outputs.AntivirusProfilePop3>;
    /**
     * Replacement message group customized for this profile.
     */
    public readonly replacemsgGroup!: pulumi.Output<string>;
    /**
     * Choose between full scan mode and quick scan mode.
     */
    public readonly scanMode!: pulumi.Output<string>;
    /**
     * Configure SMB AntiVirus options. The structure of `smb` block is documented below.
     */
    public readonly smb!: pulumi.Output<outputs.AntivirusProfileSmb>;
    /**
     * Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
     */
    public readonly smtp!: pulumi.Output<outputs.AntivirusProfileSmtp>;
    /**
     * Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
     */
    public readonly ssh!: pulumi.Output<outputs.AntivirusProfileSsh>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a AntivirusProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AntivirusProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AntivirusProfileArgs | AntivirusProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AntivirusProfileState | undefined;
            resourceInputs["analyticsAcceptFiletype"] = state ? state.analyticsAcceptFiletype : undefined;
            resourceInputs["analyticsBlFiletype"] = state ? state.analyticsBlFiletype : undefined;
            resourceInputs["analyticsDb"] = state ? state.analyticsDb : undefined;
            resourceInputs["analyticsIgnoreFiletype"] = state ? state.analyticsIgnoreFiletype : undefined;
            resourceInputs["analyticsMaxUpload"] = state ? state.analyticsMaxUpload : undefined;
            resourceInputs["analyticsWlFiletype"] = state ? state.analyticsWlFiletype : undefined;
            resourceInputs["avBlockLog"] = state ? state.avBlockLog : undefined;
            resourceInputs["avVirusLog"] = state ? state.avVirusLog : undefined;
            resourceInputs["cifs"] = state ? state.cifs : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["contentDisarm"] = state ? state.contentDisarm : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["emsThreatFeed"] = state ? state.emsThreatFeed : undefined;
            resourceInputs["extendedLog"] = state ? state.extendedLog : undefined;
            resourceInputs["externalBlocklistEnableAll"] = state ? state.externalBlocklistEnableAll : undefined;
            resourceInputs["externalBlocklists"] = state ? state.externalBlocklists : undefined;
            resourceInputs["featureSet"] = state ? state.featureSet : undefined;
            resourceInputs["fortiaiErrorAction"] = state ? state.fortiaiErrorAction : undefined;
            resourceInputs["fortiaiTimeoutAction"] = state ? state.fortiaiTimeoutAction : undefined;
            resourceInputs["fortindrErrorAction"] = state ? state.fortindrErrorAction : undefined;
            resourceInputs["fortindrTimeoutAction"] = state ? state.fortindrTimeoutAction : undefined;
            resourceInputs["fortisandboxErrorAction"] = state ? state.fortisandboxErrorAction : undefined;
            resourceInputs["fortisandboxMaxUpload"] = state ? state.fortisandboxMaxUpload : undefined;
            resourceInputs["fortisandboxMode"] = state ? state.fortisandboxMode : undefined;
            resourceInputs["fortisandboxTimeoutAction"] = state ? state.fortisandboxTimeoutAction : undefined;
            resourceInputs["ftgdAnalytics"] = state ? state.ftgdAnalytics : undefined;
            resourceInputs["ftp"] = state ? state.ftp : undefined;
            resourceInputs["http"] = state ? state.http : undefined;
            resourceInputs["imap"] = state ? state.imap : undefined;
            resourceInputs["inspectionMode"] = state ? state.inspectionMode : undefined;
            resourceInputs["mapi"] = state ? state.mapi : undefined;
            resourceInputs["mobileMalwareDb"] = state ? state.mobileMalwareDb : undefined;
            resourceInputs["nacQuar"] = state ? state.nacQuar : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nntp"] = state ? state.nntp : undefined;
            resourceInputs["outbreakPrevention"] = state ? state.outbreakPrevention : undefined;
            resourceInputs["outbreakPreventionArchiveScan"] = state ? state.outbreakPreventionArchiveScan : undefined;
            resourceInputs["pop3"] = state ? state.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = state ? state.replacemsgGroup : undefined;
            resourceInputs["scanMode"] = state ? state.scanMode : undefined;
            resourceInputs["smb"] = state ? state.smb : undefined;
            resourceInputs["smtp"] = state ? state.smtp : undefined;
            resourceInputs["ssh"] = state ? state.ssh : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AntivirusProfileArgs | undefined;
            resourceInputs["analyticsAcceptFiletype"] = args ? args.analyticsAcceptFiletype : undefined;
            resourceInputs["analyticsBlFiletype"] = args ? args.analyticsBlFiletype : undefined;
            resourceInputs["analyticsDb"] = args ? args.analyticsDb : undefined;
            resourceInputs["analyticsIgnoreFiletype"] = args ? args.analyticsIgnoreFiletype : undefined;
            resourceInputs["analyticsMaxUpload"] = args ? args.analyticsMaxUpload : undefined;
            resourceInputs["analyticsWlFiletype"] = args ? args.analyticsWlFiletype : undefined;
            resourceInputs["avBlockLog"] = args ? args.avBlockLog : undefined;
            resourceInputs["avVirusLog"] = args ? args.avVirusLog : undefined;
            resourceInputs["cifs"] = args ? args.cifs : undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["contentDisarm"] = args ? args.contentDisarm : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["emsThreatFeed"] = args ? args.emsThreatFeed : undefined;
            resourceInputs["extendedLog"] = args ? args.extendedLog : undefined;
            resourceInputs["externalBlocklistEnableAll"] = args ? args.externalBlocklistEnableAll : undefined;
            resourceInputs["externalBlocklists"] = args ? args.externalBlocklists : undefined;
            resourceInputs["featureSet"] = args ? args.featureSet : undefined;
            resourceInputs["fortiaiErrorAction"] = args ? args.fortiaiErrorAction : undefined;
            resourceInputs["fortiaiTimeoutAction"] = args ? args.fortiaiTimeoutAction : undefined;
            resourceInputs["fortindrErrorAction"] = args ? args.fortindrErrorAction : undefined;
            resourceInputs["fortindrTimeoutAction"] = args ? args.fortindrTimeoutAction : undefined;
            resourceInputs["fortisandboxErrorAction"] = args ? args.fortisandboxErrorAction : undefined;
            resourceInputs["fortisandboxMaxUpload"] = args ? args.fortisandboxMaxUpload : undefined;
            resourceInputs["fortisandboxMode"] = args ? args.fortisandboxMode : undefined;
            resourceInputs["fortisandboxTimeoutAction"] = args ? args.fortisandboxTimeoutAction : undefined;
            resourceInputs["ftgdAnalytics"] = args ? args.ftgdAnalytics : undefined;
            resourceInputs["ftp"] = args ? args.ftp : undefined;
            resourceInputs["http"] = args ? args.http : undefined;
            resourceInputs["imap"] = args ? args.imap : undefined;
            resourceInputs["inspectionMode"] = args ? args.inspectionMode : undefined;
            resourceInputs["mapi"] = args ? args.mapi : undefined;
            resourceInputs["mobileMalwareDb"] = args ? args.mobileMalwareDb : undefined;
            resourceInputs["nacQuar"] = args ? args.nacQuar : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nntp"] = args ? args.nntp : undefined;
            resourceInputs["outbreakPrevention"] = args ? args.outbreakPrevention : undefined;
            resourceInputs["outbreakPreventionArchiveScan"] = args ? args.outbreakPreventionArchiveScan : undefined;
            resourceInputs["pop3"] = args ? args.pop3 : undefined;
            resourceInputs["replacemsgGroup"] = args ? args.replacemsgGroup : undefined;
            resourceInputs["scanMode"] = args ? args.scanMode : undefined;
            resourceInputs["smb"] = args ? args.smb : undefined;
            resourceInputs["smtp"] = args ? args.smtp : undefined;
            resourceInputs["ssh"] = args ? args.ssh : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AntivirusProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AntivirusProfile resources.
 */
export interface AntivirusProfileState {
    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsAcceptFiletype?: pulumi.Input<number>;
    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsBlFiletype?: pulumi.Input<number>;
    /**
     * Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
     */
    analyticsDb?: pulumi.Input<string>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsIgnoreFiletype?: pulumi.Input<number>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
     */
    analyticsMaxUpload?: pulumi.Input<number>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsWlFiletype?: pulumi.Input<number>;
    /**
     * Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
     */
    avBlockLog?: pulumi.Input<string>;
    /**
     * Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
     */
    avVirusLog?: pulumi.Input<string>;
    /**
     * Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.AntivirusProfileCifs>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * AV Content Disarm and Reconstruction settings. The structure of `contentDisarm` block is documented below.
     */
    contentDisarm?: pulumi.Input<inputs.AntivirusProfileContentDisarm>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
     */
    emsThreatFeed?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Enable/disable all external blocklists. Valid values: `disable`, `enable`.
     */
    externalBlocklistEnableAll?: pulumi.Input<string>;
    /**
     * One or more external malware block lists. The structure of `externalBlocklist` block is documented below.
     */
    externalBlocklists?: pulumi.Input<pulumi.Input<inputs.AntivirusProfileExternalBlocklist>[]>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortiaiErrorAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortiaiTimeoutAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortindrErrorAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortindrTimeoutAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortisandboxErrorAction?: pulumi.Input<string>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox.
     */
    fortisandboxMaxUpload?: pulumi.Input<number>;
    /**
     * FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
     */
    fortisandboxMode?: pulumi.Input<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortisandboxTimeoutAction?: pulumi.Input<string>;
    /**
     * Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
     */
    ftgdAnalytics?: pulumi.Input<string>;
    /**
     * Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.AntivirusProfileFtp>;
    /**
     * Configure HTTP AntiVirus options. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.AntivirusProfileHttp>;
    /**
     * Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.AntivirusProfileImap>;
    /**
     * Inspection mode. Valid values: `proxy`, `flow-based`.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.AntivirusProfileMapi>;
    /**
     * Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
     */
    mobileMalwareDb?: pulumi.Input<string>;
    /**
     * Configure AntiVirus quarantine settings. The structure of `nacQuar` block is documented below.
     */
    nacQuar?: pulumi.Input<inputs.AntivirusProfileNacQuar>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
     */
    nntp?: pulumi.Input<inputs.AntivirusProfileNntp>;
    /**
     * Configure Virus Outbreak Prevention settings. The structure of `outbreakPrevention` block is documented below.
     */
    outbreakPrevention?: pulumi.Input<inputs.AntivirusProfileOutbreakPrevention>;
    /**
     * Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
     */
    outbreakPreventionArchiveScan?: pulumi.Input<string>;
    /**
     * Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.AntivirusProfilePop3>;
    /**
     * Replacement message group customized for this profile.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Choose between full scan mode and quick scan mode.
     */
    scanMode?: pulumi.Input<string>;
    /**
     * Configure SMB AntiVirus options. The structure of `smb` block is documented below.
     */
    smb?: pulumi.Input<inputs.AntivirusProfileSmb>;
    /**
     * Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.AntivirusProfileSmtp>;
    /**
     * Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.AntivirusProfileSsh>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AntivirusProfile resource.
 */
export interface AntivirusProfileArgs {
    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsAcceptFiletype?: pulumi.Input<number>;
    /**
     * Only submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsBlFiletype?: pulumi.Input<number>;
    /**
     * Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
     */
    analyticsDb?: pulumi.Input<string>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsIgnoreFiletype?: pulumi.Input<number>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
     */
    analyticsMaxUpload?: pulumi.Input<number>;
    /**
     * Do not submit files matching this DLP file-pattern to FortiSandbox.
     */
    analyticsWlFiletype?: pulumi.Input<number>;
    /**
     * Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
     */
    avBlockLog?: pulumi.Input<string>;
    /**
     * Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
     */
    avVirusLog?: pulumi.Input<string>;
    /**
     * Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
     */
    cifs?: pulumi.Input<inputs.AntivirusProfileCifs>;
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * AV Content Disarm and Reconstruction settings. The structure of `contentDisarm` block is documented below.
     */
    contentDisarm?: pulumi.Input<inputs.AntivirusProfileContentDisarm>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
     */
    emsThreatFeed?: pulumi.Input<string>;
    /**
     * Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
     */
    extendedLog?: pulumi.Input<string>;
    /**
     * Enable/disable all external blocklists. Valid values: `disable`, `enable`.
     */
    externalBlocklistEnableAll?: pulumi.Input<string>;
    /**
     * One or more external malware block lists. The structure of `externalBlocklist` block is documented below.
     */
    externalBlocklists?: pulumi.Input<pulumi.Input<inputs.AntivirusProfileExternalBlocklist>[]>;
    /**
     * Flow/proxy feature set. Valid values: `flow`, `proxy`.
     */
    featureSet?: pulumi.Input<string>;
    /**
     * Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortiaiErrorAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortiaiTimeoutAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiNDR encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortindrErrorAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiNDR encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortindrTimeoutAction?: pulumi.Input<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters an error. Valid values: `log-only`, `block`, `ignore`.
     */
    fortisandboxErrorAction?: pulumi.Input<string>;
    /**
     * Maximum size of files that can be uploaded to FortiSandbox.
     */
    fortisandboxMaxUpload?: pulumi.Input<number>;
    /**
     * FortiSandbox scan modes. Valid values: `inline`, `analytics-suspicious`, `analytics-everything`.
     */
    fortisandboxMode?: pulumi.Input<string>;
    /**
     * Action to take if FortiSandbox inline scan encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
     */
    fortisandboxTimeoutAction?: pulumi.Input<string>;
    /**
     * Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
     */
    ftgdAnalytics?: pulumi.Input<string>;
    /**
     * Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
     */
    ftp?: pulumi.Input<inputs.AntivirusProfileFtp>;
    /**
     * Configure HTTP AntiVirus options. The structure of `http` block is documented below.
     */
    http?: pulumi.Input<inputs.AntivirusProfileHttp>;
    /**
     * Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
     */
    imap?: pulumi.Input<inputs.AntivirusProfileImap>;
    /**
     * Inspection mode. Valid values: `proxy`, `flow-based`.
     */
    inspectionMode?: pulumi.Input<string>;
    /**
     * Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
     */
    mapi?: pulumi.Input<inputs.AntivirusProfileMapi>;
    /**
     * Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
     */
    mobileMalwareDb?: pulumi.Input<string>;
    /**
     * Configure AntiVirus quarantine settings. The structure of `nacQuar` block is documented below.
     */
    nacQuar?: pulumi.Input<inputs.AntivirusProfileNacQuar>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
     */
    nntp?: pulumi.Input<inputs.AntivirusProfileNntp>;
    /**
     * Configure Virus Outbreak Prevention settings. The structure of `outbreakPrevention` block is documented below.
     */
    outbreakPrevention?: pulumi.Input<inputs.AntivirusProfileOutbreakPrevention>;
    /**
     * Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
     */
    outbreakPreventionArchiveScan?: pulumi.Input<string>;
    /**
     * Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
     */
    pop3?: pulumi.Input<inputs.AntivirusProfilePop3>;
    /**
     * Replacement message group customized for this profile.
     */
    replacemsgGroup?: pulumi.Input<string>;
    /**
     * Choose between full scan mode and quick scan mode.
     */
    scanMode?: pulumi.Input<string>;
    /**
     * Configure SMB AntiVirus options. The structure of `smb` block is documented below.
     */
    smb?: pulumi.Input<inputs.AntivirusProfileSmb>;
    /**
     * Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
     */
    smtp?: pulumi.Input<inputs.AntivirusProfileSmtp>;
    /**
     * Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
     */
    ssh?: pulumi.Input<inputs.AntivirusProfileSsh>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
