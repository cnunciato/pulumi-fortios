// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Configure IP address type availability.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Wirelesscontrollerhotspot20Anqpipaddresstype("trname", new()
    ///     {
    ///         Ipv4AddressType = "public",
    ///         Ipv6AddressType = "not-available",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype")]
    public partial class Wirelesscontrollerhotspot20Anqpipaddresstype : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
        /// </summary>
        [Output("ipv4AddressType")]
        public Output<string> Ipv4AddressType { get; private set; } = null!;

        /// <summary>
        /// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
        /// </summary>
        [Output("ipv6AddressType")]
        public Output<string> Ipv6AddressType { get; private set; } = null!;

        /// <summary>
        /// IP type name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Wirelesscontrollerhotspot20Anqpipaddresstype resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Wirelesscontrollerhotspot20Anqpipaddresstype(string name, Wirelesscontrollerhotspot20AnqpipaddresstypeArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype", name, args ?? new Wirelesscontrollerhotspot20AnqpipaddresstypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Wirelesscontrollerhotspot20Anqpipaddresstype(string name, Input<string> id, Wirelesscontrollerhotspot20AnqpipaddresstypeState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerhotspot20Anqpipaddresstype:Wirelesscontrollerhotspot20Anqpipaddresstype", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Wirelesscontrollerhotspot20Anqpipaddresstype resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Wirelesscontrollerhotspot20Anqpipaddresstype Get(string name, Input<string> id, Wirelesscontrollerhotspot20AnqpipaddresstypeState? state = null, CustomResourceOptions? options = null)
        {
            return new Wirelesscontrollerhotspot20Anqpipaddresstype(name, id, state, options);
        }
    }

    public sealed class Wirelesscontrollerhotspot20AnqpipaddresstypeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
        /// </summary>
        [Input("ipv4AddressType")]
        public Input<string>? Ipv4AddressType { get; set; }

        /// <summary>
        /// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
        /// </summary>
        [Input("ipv6AddressType")]
        public Input<string>? Ipv6AddressType { get; set; }

        /// <summary>
        /// IP type name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Wirelesscontrollerhotspot20AnqpipaddresstypeArgs()
        {
        }
        public static new Wirelesscontrollerhotspot20AnqpipaddresstypeArgs Empty => new Wirelesscontrollerhotspot20AnqpipaddresstypeArgs();
    }

    public sealed class Wirelesscontrollerhotspot20AnqpipaddresstypeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
        /// </summary>
        [Input("ipv4AddressType")]
        public Input<string>? Ipv4AddressType { get; set; }

        /// <summary>
        /// IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
        /// </summary>
        [Input("ipv6AddressType")]
        public Input<string>? Ipv6AddressType { get; set; }

        /// <summary>
        /// IP type name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public Wirelesscontrollerhotspot20AnqpipaddresstypeState()
        {
        }
        public static new Wirelesscontrollerhotspot20AnqpipaddresstypeState Empty => new Wirelesscontrollerhotspot20AnqpipaddresstypeState();
    }
}