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
    /// Configure custom services.
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
    ///     var trname = new Fortios.FirewallserviceCustom("trname", new()
    ///     {
    ///         AppServiceType = "disable",
    ///         Category = "General",
    ///         CheckResetRange = "default",
    ///         Color = 0,
    ///         Helper = "auto",
    ///         Iprange = "0.0.0.0",
    ///         Protocol = "TCP/UDP/SCTP",
    ///         ProtocolNumber = 6,
    ///         Proxy = "disable",
    ///         TcpHalfcloseTimer = 0,
    ///         TcpHalfopenTimer = 0,
    ///         TcpPortrange = "223-332",
    ///         TcpTimewaitTimer = 0,
    ///         UdpIdleTimer = 0,
    ///         Visibility = "enable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallService Custom can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallserviceCustom:FirewallserviceCustom labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallserviceCustom:FirewallserviceCustom labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallserviceCustom:FirewallserviceCustom")]
    public partial class FirewallserviceCustom : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application category ID. The structure of `app_category` block is documented below.
        /// </summary>
        [Output("appCategories")]
        public Output<ImmutableArray<Outputs.FirewallserviceCustomAppCategory>> AppCategories { get; private set; } = null!;

        /// <summary>
        /// Application service type. Valid values: `disable`, `app-id`, `app-category`.
        /// </summary>
        [Output("appServiceType")]
        public Output<string> AppServiceType { get; private set; } = null!;

        /// <summary>
        /// Application ID. The structure of `application` block is documented below.
        /// </summary>
        [Output("applications")]
        public Output<ImmutableArray<Outputs.FirewallserviceCustomApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// Service category.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
        /// </summary>
        [Output("checkResetRange")]
        public Output<string> CheckResetRange { get; private set; } = null!;

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Output("color")]
        public Output<int> Color { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("fabricObject")]
        public Output<string> FabricObject { get; private set; } = null!;

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Helper name.
        /// </summary>
        [Output("helper")]
        public Output<string> Helper { get; private set; } = null!;

        /// <summary>
        /// ICMP code.
        /// </summary>
        [Output("icmpcode")]
        public Output<int> Icmpcode { get; private set; } = null!;

        /// <summary>
        /// ICMP type.
        /// </summary>
        [Output("icmptype")]
        public Output<int> Icmptype { get; private set; } = null!;

        /// <summary>
        /// Start and end of the IP range associated with service.
        /// </summary>
        [Output("iprange")]
        public Output<string> Iprange { get; private set; } = null!;

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// IP protocol number.
        /// </summary>
        [Output("protocolNumber")]
        public Output<int> ProtocolNumber { get; private set; } = null!;

        /// <summary>
        /// Enable/disable web proxy service. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("proxy")]
        public Output<string> Proxy { get; private set; } = null!;

        /// <summary>
        /// Multiple SCTP port ranges.
        /// </summary>
        [Output("sctpPortrange")]
        public Output<string> SctpPortrange { get; private set; } = null!;

        /// <summary>
        /// Session TTL (300 - 604800, 0 = default).
        /// </summary>
        [Output("sessionTtl")]
        public Output<int> SessionTtl { get; private set; } = null!;

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Output("tcpHalfcloseTimer")]
        public Output<int> TcpHalfcloseTimer { get; private set; } = null!;

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Output("tcpHalfopenTimer")]
        public Output<int> TcpHalfopenTimer { get; private set; } = null!;

        /// <summary>
        /// Multiple TCP port ranges.
        /// </summary>
        [Output("tcpPortrange")]
        public Output<string> TcpPortrange { get; private set; } = null!;

        /// <summary>
        /// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
        /// </summary>
        [Output("tcpRstTimer")]
        public Output<int> TcpRstTimer { get; private set; } = null!;

        /// <summary>
        /// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
        /// </summary>
        [Output("tcpTimewaitTimer")]
        public Output<int> TcpTimewaitTimer { get; private set; } = null!;

        /// <summary>
        /// UDP half close timeout (0 - 86400 sec, 0 = default).
        /// </summary>
        [Output("udpIdleTimer")]
        public Output<int> UdpIdleTimer { get; private set; } = null!;

        /// <summary>
        /// Multiple UDP port ranges.
        /// </summary>
        [Output("udpPortrange")]
        public Output<string> UdpPortrange { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallserviceCustom resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallserviceCustom(string name, FirewallserviceCustomArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallserviceCustom:FirewallserviceCustom", name, args ?? new FirewallserviceCustomArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallserviceCustom(string name, Input<string> id, FirewallserviceCustomState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallserviceCustom:FirewallserviceCustom", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallserviceCustom resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallserviceCustom Get(string name, Input<string> id, FirewallserviceCustomState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallserviceCustom(name, id, state, options);
        }
    }

    public sealed class FirewallserviceCustomArgs : global::Pulumi.ResourceArgs
    {
        [Input("appCategories")]
        private InputList<Inputs.FirewallserviceCustomAppCategoryArgs>? _appCategories;

        /// <summary>
        /// Application category ID. The structure of `app_category` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallserviceCustomAppCategoryArgs> AppCategories
        {
            get => _appCategories ?? (_appCategories = new InputList<Inputs.FirewallserviceCustomAppCategoryArgs>());
            set => _appCategories = value;
        }

        /// <summary>
        /// Application service type. Valid values: `disable`, `app-id`, `app-category`.
        /// </summary>
        [Input("appServiceType")]
        public Input<string>? AppServiceType { get; set; }

        [Input("applications")]
        private InputList<Inputs.FirewallserviceCustomApplicationArgs>? _applications;

        /// <summary>
        /// Application ID. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallserviceCustomApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.FirewallserviceCustomApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Service category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
        /// </summary>
        [Input("checkResetRange")]
        public Input<string>? CheckResetRange { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Helper name.
        /// </summary>
        [Input("helper")]
        public Input<string>? Helper { get; set; }

        /// <summary>
        /// ICMP code.
        /// </summary>
        [Input("icmpcode")]
        public Input<int>? Icmpcode { get; set; }

        /// <summary>
        /// ICMP type.
        /// </summary>
        [Input("icmptype")]
        public Input<int>? Icmptype { get; set; }

        /// <summary>
        /// Start and end of the IP range associated with service.
        /// </summary>
        [Input("iprange")]
        public Input<string>? Iprange { get; set; }

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// IP protocol number.
        /// </summary>
        [Input("protocolNumber")]
        public Input<int>? ProtocolNumber { get; set; }

        /// <summary>
        /// Enable/disable web proxy service. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Multiple SCTP port ranges.
        /// </summary>
        [Input("sctpPortrange")]
        public Input<string>? SctpPortrange { get; set; }

        /// <summary>
        /// Session TTL (300 - 604800, 0 = default).
        /// </summary>
        [Input("sessionTtl")]
        public Input<int>? SessionTtl { get; set; }

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Input("tcpHalfcloseTimer")]
        public Input<int>? TcpHalfcloseTimer { get; set; }

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Input("tcpHalfopenTimer")]
        public Input<int>? TcpHalfopenTimer { get; set; }

        /// <summary>
        /// Multiple TCP port ranges.
        /// </summary>
        [Input("tcpPortrange")]
        public Input<string>? TcpPortrange { get; set; }

        /// <summary>
        /// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
        /// </summary>
        [Input("tcpRstTimer")]
        public Input<int>? TcpRstTimer { get; set; }

        /// <summary>
        /// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
        /// </summary>
        [Input("tcpTimewaitTimer")]
        public Input<int>? TcpTimewaitTimer { get; set; }

        /// <summary>
        /// UDP half close timeout (0 - 86400 sec, 0 = default).
        /// </summary>
        [Input("udpIdleTimer")]
        public Input<int>? UdpIdleTimer { get; set; }

        /// <summary>
        /// Multiple UDP port ranges.
        /// </summary>
        [Input("udpPortrange")]
        public Input<string>? UdpPortrange { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public FirewallserviceCustomArgs()
        {
        }
        public static new FirewallserviceCustomArgs Empty => new FirewallserviceCustomArgs();
    }

    public sealed class FirewallserviceCustomState : global::Pulumi.ResourceArgs
    {
        [Input("appCategories")]
        private InputList<Inputs.FirewallserviceCustomAppCategoryGetArgs>? _appCategories;

        /// <summary>
        /// Application category ID. The structure of `app_category` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallserviceCustomAppCategoryGetArgs> AppCategories
        {
            get => _appCategories ?? (_appCategories = new InputList<Inputs.FirewallserviceCustomAppCategoryGetArgs>());
            set => _appCategories = value;
        }

        /// <summary>
        /// Application service type. Valid values: `disable`, `app-id`, `app-category`.
        /// </summary>
        [Input("appServiceType")]
        public Input<string>? AppServiceType { get; set; }

        [Input("applications")]
        private InputList<Inputs.FirewallserviceCustomApplicationGetArgs>? _applications;

        /// <summary>
        /// Application ID. The structure of `application` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallserviceCustomApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.FirewallserviceCustomApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// Service category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
        /// </summary>
        [Input("checkResetRange")]
        public Input<string>? CheckResetRange { get; set; }

        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        [Input("color")]
        public Input<int>? Color { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Security Fabric global object setting. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fabricObject")]
        public Input<string>? FabricObject { get; set; }

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Helper name.
        /// </summary>
        [Input("helper")]
        public Input<string>? Helper { get; set; }

        /// <summary>
        /// ICMP code.
        /// </summary>
        [Input("icmpcode")]
        public Input<int>? Icmpcode { get; set; }

        /// <summary>
        /// ICMP type.
        /// </summary>
        [Input("icmptype")]
        public Input<int>? Icmptype { get; set; }

        /// <summary>
        /// Start and end of the IP range associated with service.
        /// </summary>
        [Input("iprange")]
        public Input<string>? Iprange { get; set; }

        /// <summary>
        /// Custom service name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// IP protocol number.
        /// </summary>
        [Input("protocolNumber")]
        public Input<int>? ProtocolNumber { get; set; }

        /// <summary>
        /// Enable/disable web proxy service. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        /// <summary>
        /// Multiple SCTP port ranges.
        /// </summary>
        [Input("sctpPortrange")]
        public Input<string>? SctpPortrange { get; set; }

        /// <summary>
        /// Session TTL (300 - 604800, 0 = default).
        /// </summary>
        [Input("sessionTtl")]
        public Input<int>? SessionTtl { get; set; }

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Input("tcpHalfcloseTimer")]
        public Input<int>? TcpHalfcloseTimer { get; set; }

        /// <summary>
        /// Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
        /// </summary>
        [Input("tcpHalfopenTimer")]
        public Input<int>? TcpHalfopenTimer { get; set; }

        /// <summary>
        /// Multiple TCP port ranges.
        /// </summary>
        [Input("tcpPortrange")]
        public Input<string>? TcpPortrange { get; set; }

        /// <summary>
        /// Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
        /// </summary>
        [Input("tcpRstTimer")]
        public Input<int>? TcpRstTimer { get; set; }

        /// <summary>
        /// Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
        /// </summary>
        [Input("tcpTimewaitTimer")]
        public Input<int>? TcpTimewaitTimer { get; set; }

        /// <summary>
        /// UDP half close timeout (0 - 86400 sec, 0 = default).
        /// </summary>
        [Input("udpIdleTimer")]
        public Input<int>? UdpIdleTimer { get; set; }

        /// <summary>
        /// Multiple UDP port ranges.
        /// </summary>
        [Input("udpPortrange")]
        public Input<string>? UdpPortrange { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public FirewallserviceCustomState()
        {
        }
        public static new FirewallserviceCustomState Empty => new FirewallserviceCustomState();
    }
}
