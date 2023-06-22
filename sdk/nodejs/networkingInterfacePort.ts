// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to configure interface settings of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.SystemInterface`, we recommend that you use the new resource.
 *
 * ## Example Usage
 * ### Loopback Interface
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const loopback1 = new fortios.NetworkingInterfacePort("loopback1", {
 *     alias: "cc1",
 *     allowaccess: "ping http",
 *     description: "description",
 *     ip: "23.123.33.10 255.255.255.0",
 *     mode: "static",
 *     role: "lan",
 *     status: "up",
 *     type: "loopback",
 *     vdom: "root",
 * });
 * ```
 * ### VLAN Interface
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const vlan1 = new fortios.NetworkingInterfacePort("vlan1", {
 *     allowaccess: "ping",
 *     defaultgw: "enable",
 *     distance: "33",
 *     "interface": "port2",
 *     ip: "3.123.33.10 255.255.255.0",
 *     mode: "static",
 *     role: "lan",
 *     type: "vlan",
 *     vdom: "root",
 *     vlanid: "3",
 * });
 * ```
 * ### Physical Interface
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.NetworkingInterfacePort("test1", {
 *     alias: "dkeeew",
 *     allowaccess: "ping https",
 *     defaultgw: "enable",
 *     description: "description",
 *     deviceIdentification: "enable",
 *     distance: "33",
 *     dnsServerOverride: "enable",
 *     ip: "93.133.133.110 255.255.255.0",
 *     mode: "static",
 *     mtu: "2933",
 *     mtuOverride: "enable",
 *     role: "lan",
 *     speed: "auto",
 *     status: "up",
 *     tcpMss: "3232",
 *     type: "physical",
 * });
 * ```
 */
export class NetworkingInterfacePort extends pulumi.CustomResource {
    /**
     * Get an existing NetworkingInterfacePort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkingInterfacePortState, opts?: pulumi.CustomResourceOptions): NetworkingInterfacePort {
        return new NetworkingInterfacePort(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:index/networkingInterfacePort:NetworkingInterfacePort';

    /**
     * Returns true if the given object is an instance of NetworkingInterfacePort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkingInterfacePort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkingInterfacePort.__pulumiType;
    }

    /**
     * Alias will be displayed with the interface name to make it easier to distinguish.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * Permitted types of management access to this interface.
     */
    public readonly allowaccess!: pulumi.Output<string>;
    /**
     * Enable to get the gateway IP from the DHCP or PPPoE server.
     */
    public readonly defaultgw!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
     */
    public readonly deviceIdentification!: pulumi.Output<string>;
    /**
     * Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
     */
    public readonly distance!: pulumi.Output<string>;
    /**
     * Enable/disable use DNS acquired by DHCP or PPPoE.
     */
    public readonly dnsServerOverride!: pulumi.Output<string>;
    /**
     * Interface name.
     */
    public readonly interface!: pulumi.Output<string>;
    /**
     * Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Addressing mode.
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * MTU value for this interface.
     */
    public readonly mtu!: pulumi.Output<string>;
    /**
     * Enable to set a custom MTU for this interface.
     */
    public readonly mtuOverride!: pulumi.Output<string>;
    /**
     * Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Interface role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Interface speed. The default setting and the options available depend on the interface hardware.
     */
    public readonly speed!: pulumi.Output<string>;
    /**
     * Bring the interface up or shut the interface down.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * TCP maximum segment size. 0 means do not change segment size.
     */
    public readonly tcpMss!: pulumi.Output<string>;
    /**
     * Interface type (support physical, vlan, loopback).
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Interface is in this virtual domain (VDOM).
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * VLAN ID.
     */
    public readonly vlanid!: pulumi.Output<string>;

    /**
     * Create a NetworkingInterfacePort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkingInterfacePortArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkingInterfacePortArgs | NetworkingInterfacePortState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkingInterfacePortState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["allowaccess"] = state ? state.allowaccess : undefined;
            resourceInputs["defaultgw"] = state ? state.defaultgw : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceIdentification"] = state ? state.deviceIdentification : undefined;
            resourceInputs["distance"] = state ? state.distance : undefined;
            resourceInputs["dnsServerOverride"] = state ? state.dnsServerOverride : undefined;
            resourceInputs["interface"] = state ? state.interface : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["mtuOverride"] = state ? state.mtuOverride : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["speed"] = state ? state.speed : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tcpMss"] = state ? state.tcpMss : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vlanid"] = state ? state.vlanid : undefined;
        } else {
            const args = argsOrState as NetworkingInterfacePortArgs | undefined;
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["allowaccess"] = args ? args.allowaccess : undefined;
            resourceInputs["defaultgw"] = args ? args.defaultgw : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceIdentification"] = args ? args.deviceIdentification : undefined;
            resourceInputs["distance"] = args ? args.distance : undefined;
            resourceInputs["dnsServerOverride"] = args ? args.dnsServerOverride : undefined;
            resourceInputs["interface"] = args ? args.interface : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["mtuOverride"] = args ? args.mtuOverride : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["speed"] = args ? args.speed : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tcpMss"] = args ? args.tcpMss : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vlanid"] = args ? args.vlanid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkingInterfacePort.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkingInterfacePort resources.
 */
export interface NetworkingInterfacePortState {
    /**
     * Alias will be displayed with the interface name to make it easier to distinguish.
     */
    alias?: pulumi.Input<string>;
    /**
     * Permitted types of management access to this interface.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Enable to get the gateway IP from the DHCP or PPPoE server.
     */
    defaultgw?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
     */
    deviceIdentification?: pulumi.Input<string>;
    /**
     * Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
     */
    distance?: pulumi.Input<string>;
    /**
     * Enable/disable use DNS acquired by DHCP or PPPoE.
     */
    dnsServerOverride?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
     */
    ip?: pulumi.Input<string>;
    /**
     * Addressing mode.
     */
    mode?: pulumi.Input<string>;
    /**
     * MTU value for this interface.
     */
    mtu?: pulumi.Input<string>;
    /**
     * Enable to set a custom MTU for this interface.
     */
    mtuOverride?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Interface role.
     */
    role?: pulumi.Input<string>;
    /**
     * Interface speed. The default setting and the options available depend on the interface hardware.
     */
    speed?: pulumi.Input<string>;
    /**
     * Bring the interface up or shut the interface down.
     */
    status?: pulumi.Input<string>;
    /**
     * TCP maximum segment size. 0 means do not change segment size.
     */
    tcpMss?: pulumi.Input<string>;
    /**
     * Interface type (support physical, vlan, loopback).
     */
    type?: pulumi.Input<string>;
    /**
     * Interface is in this virtual domain (VDOM).
     */
    vdom?: pulumi.Input<string>;
    /**
     * VLAN ID.
     */
    vlanid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkingInterfacePort resource.
 */
export interface NetworkingInterfacePortArgs {
    /**
     * Alias will be displayed with the interface name to make it easier to distinguish.
     */
    alias?: pulumi.Input<string>;
    /**
     * Permitted types of management access to this interface.
     */
    allowaccess?: pulumi.Input<string>;
    /**
     * Enable to get the gateway IP from the DHCP or PPPoE server.
     */
    defaultgw?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
     */
    deviceIdentification?: pulumi.Input<string>;
    /**
     * Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
     */
    distance?: pulumi.Input<string>;
    /**
     * Enable/disable use DNS acquired by DHCP or PPPoE.
     */
    dnsServerOverride?: pulumi.Input<string>;
    /**
     * Interface name.
     */
    interface?: pulumi.Input<string>;
    /**
     * Interface IPv4 address and subnet mask, syntax` - X.X.X.X X.X.X.X.
     */
    ip?: pulumi.Input<string>;
    /**
     * Addressing mode.
     */
    mode?: pulumi.Input<string>;
    /**
     * MTU value for this interface.
     */
    mtu?: pulumi.Input<string>;
    /**
     * Enable to set a custom MTU for this interface.
     */
    mtuOverride?: pulumi.Input<string>;
    /**
     * Name.
     */
    name?: pulumi.Input<string>;
    /**
     * Interface role.
     */
    role?: pulumi.Input<string>;
    /**
     * Interface speed. The default setting and the options available depend on the interface hardware.
     */
    speed?: pulumi.Input<string>;
    /**
     * Bring the interface up or shut the interface down.
     */
    status?: pulumi.Input<string>;
    /**
     * TCP maximum segment size. 0 means do not change segment size.
     */
    tcpMss?: pulumi.Input<string>;
    /**
     * Interface type (support physical, vlan, loopback).
     */
    type: pulumi.Input<string>;
    /**
     * Interface is in this virtual domain (VDOM).
     */
    vdom?: pulumi.Input<string>;
    /**
     * VLAN ID.
     */
    vlanid?: pulumi.Input<string>;
}