// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Configure roaming consortium.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.Wirelesscontrollerhotspot20Anqproamingconsortium("trname", {});
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 AnqpRoamingConsortium can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerhotspot20Anqproamingconsortium:Wirelesscontrollerhotspot20Anqproamingconsortium labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:index/wirelesscontrollerhotspot20Anqproamingconsortium:Wirelesscontrollerhotspot20Anqproamingconsortium labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Wirelesscontrollerhotspot20Anqproamingconsortium extends pulumi.CustomResource {
    /**
     * Get an existing Wirelesscontrollerhotspot20Anqproamingconsortium resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Wirelesscontrollerhotspot20AnqproamingconsortiumState, opts?: pulumi.CustomResourceOptions): Wirelesscontrollerhotspot20Anqproamingconsortium {
        return new Wirelesscontrollerhotspot20Anqproamingconsortium(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/wirelesscontrollerhotspot20Anqproamingconsortium:Wirelesscontrollerhotspot20Anqproamingconsortium';

    /**
     * Returns true if the given object is an instance of Wirelesscontrollerhotspot20Anqproamingconsortium.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Wirelesscontrollerhotspot20Anqproamingconsortium {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Wirelesscontrollerhotspot20Anqproamingconsortium.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Roaming consortium name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    public readonly oiLists!: pulumi.Output<outputs.Wirelesscontrollerhotspot20AnqproamingconsortiumOiList[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Wirelesscontrollerhotspot20Anqproamingconsortium resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Wirelesscontrollerhotspot20AnqproamingconsortiumArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Wirelesscontrollerhotspot20AnqproamingconsortiumArgs | Wirelesscontrollerhotspot20AnqproamingconsortiumState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Wirelesscontrollerhotspot20AnqproamingconsortiumState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oiLists"] = state ? state.oiLists : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Wirelesscontrollerhotspot20AnqproamingconsortiumArgs | undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oiLists"] = args ? args.oiLists : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Wirelesscontrollerhotspot20Anqproamingconsortium.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Wirelesscontrollerhotspot20Anqproamingconsortium resources.
 */
export interface Wirelesscontrollerhotspot20AnqproamingconsortiumState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Roaming consortium name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    oiLists?: pulumi.Input<pulumi.Input<inputs.Wirelesscontrollerhotspot20AnqproamingconsortiumOiList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Wirelesscontrollerhotspot20Anqproamingconsortium resource.
 */
export interface Wirelesscontrollerhotspot20AnqproamingconsortiumArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Roaming consortium name.
     */
    name?: pulumi.Input<string>;
    /**
     * Organization identifier list. The structure of `oiList` block is documented below.
     */
    oiLists?: pulumi.Input<pulumi.Input<inputs.Wirelesscontrollerhotspot20AnqproamingconsortiumOiList>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}