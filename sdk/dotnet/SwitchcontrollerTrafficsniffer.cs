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
    /// Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchController TrafficSniffer can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerTrafficsniffer:SwitchcontrollerTrafficsniffer labelname SwitchControllerTrafficSniffer
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerTrafficsniffer:SwitchcontrollerTrafficsniffer labelname SwitchControllerTrafficSniffer
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerTrafficsniffer:SwitchcontrollerTrafficsniffer")]
    public partial class SwitchcontrollerTrafficsniffer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Configure ERSPAN collector IP address.
        /// </summary>
        [Output("erspanIp")]
        public Output<string> ErspanIp { get; private set; } = null!;

        /// <summary>
        /// Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Sniffer IPs to filter. The structure of `target_ip` block is documented below.
        /// </summary>
        [Output("targetIps")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerTrafficsnifferTargetIp>> TargetIps { get; private set; } = null!;

        /// <summary>
        /// Sniffer MACs to filter. The structure of `target_mac` block is documented below.
        /// </summary>
        [Output("targetMacs")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerTrafficsnifferTargetMac>> TargetMacs { get; private set; } = null!;

        /// <summary>
        /// Sniffer ports to filter. The structure of `target_port` block is documented below.
        /// </summary>
        [Output("targetPorts")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerTrafficsnifferTargetPort>> TargetPorts { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerTrafficsniffer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerTrafficsniffer(string name, SwitchcontrollerTrafficsnifferArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerTrafficsniffer:SwitchcontrollerTrafficsniffer", name, args ?? new SwitchcontrollerTrafficsnifferArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerTrafficsniffer(string name, Input<string> id, SwitchcontrollerTrafficsnifferState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerTrafficsniffer:SwitchcontrollerTrafficsniffer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerTrafficsniffer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerTrafficsniffer Get(string name, Input<string> id, SwitchcontrollerTrafficsnifferState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerTrafficsniffer(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerTrafficsnifferArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Configure ERSPAN collector IP address.
        /// </summary>
        [Input("erspanIp")]
        public Input<string>? ErspanIp { get; set; }

        /// <summary>
        /// Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("targetIps")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpArgs>? _targetIps;

        /// <summary>
        /// Sniffer IPs to filter. The structure of `target_ip` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpArgs> TargetIps
        {
            get => _targetIps ?? (_targetIps = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpArgs>());
            set => _targetIps = value;
        }

        [Input("targetMacs")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacArgs>? _targetMacs;

        /// <summary>
        /// Sniffer MACs to filter. The structure of `target_mac` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacArgs> TargetMacs
        {
            get => _targetMacs ?? (_targetMacs = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacArgs>());
            set => _targetMacs = value;
        }

        [Input("targetPorts")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortArgs>? _targetPorts;

        /// <summary>
        /// Sniffer ports to filter. The structure of `target_port` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortArgs> TargetPorts
        {
            get => _targetPorts ?? (_targetPorts = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortArgs>());
            set => _targetPorts = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerTrafficsnifferArgs()
        {
        }
        public static new SwitchcontrollerTrafficsnifferArgs Empty => new SwitchcontrollerTrafficsnifferArgs();
    }

    public sealed class SwitchcontrollerTrafficsnifferState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Configure ERSPAN collector IP address.
        /// </summary>
        [Input("erspanIp")]
        public Input<string>? ErspanIp { get; set; }

        /// <summary>
        /// Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("targetIps")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpGetArgs>? _targetIps;

        /// <summary>
        /// Sniffer IPs to filter. The structure of `target_ip` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpGetArgs> TargetIps
        {
            get => _targetIps ?? (_targetIps = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetIpGetArgs>());
            set => _targetIps = value;
        }

        [Input("targetMacs")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacGetArgs>? _targetMacs;

        /// <summary>
        /// Sniffer MACs to filter. The structure of `target_mac` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacGetArgs> TargetMacs
        {
            get => _targetMacs ?? (_targetMacs = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetMacGetArgs>());
            set => _targetMacs = value;
        }

        [Input("targetPorts")]
        private InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortGetArgs>? _targetPorts;

        /// <summary>
        /// Sniffer ports to filter. The structure of `target_port` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortGetArgs> TargetPorts
        {
            get => _targetPorts ?? (_targetPorts = new InputList<Inputs.SwitchcontrollerTrafficsnifferTargetPortGetArgs>());
            set => _targetPorts = value;
        }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerTrafficsnifferState()
        {
        }
        public static new SwitchcontrollerTrafficsnifferState Empty => new SwitchcontrollerTrafficsnifferState();
    }
}
