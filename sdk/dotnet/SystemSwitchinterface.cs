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
    /// Configure software switch interfaces by grouping physical and WiFi interfaces.
    /// 
    /// ## Import
    /// 
    /// System SwitchInterface can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSwitchinterface:SystemSwitchinterface labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemSwitchinterface:SystemSwitchinterface labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemSwitchinterface:SystemSwitchinterface")]
    public partial class SystemSwitchinterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        /// </summary>
        [Output("intraSwitchPolicy")]
        public Output<string> IntraSwitchPolicy { get; private set; } = null!;

        /// <summary>
        /// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        /// </summary>
        [Output("macTtl")]
        public Output<int> MacTtl { get; private set; } = null!;

        /// <summary>
        /// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<Outputs.SystemSwitchinterfaceMember>> Members { get; private set; } = null!;

        /// <summary>
        /// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        /// </summary>
        [Output("span")]
        public Output<string> Span { get; private set; } = null!;

        /// <summary>
        /// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        /// </summary>
        [Output("spanDestPort")]
        public Output<string> SpanDestPort { get; private set; } = null!;

        /// <summary>
        /// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Output("spanDirection")]
        public Output<string> SpanDirection { get; private set; } = null!;

        /// <summary>
        /// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        /// </summary>
        [Output("spanSourcePorts")]
        public Output<ImmutableArray<Outputs.SystemSwitchinterfaceSpanSourcePort>> SpanSourcePorts { get; private set; } = null!;

        /// <summary>
        /// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VDOM that the software switch belongs to.
        /// </summary>
        [Output("vdom")]
        public Output<string> Vdom { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemSwitchinterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemSwitchinterface(string name, SystemSwitchinterfaceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSwitchinterface:SystemSwitchinterface", name, args ?? new SystemSwitchinterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemSwitchinterface(string name, Input<string> id, SystemSwitchinterfaceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemSwitchinterface:SystemSwitchinterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemSwitchinterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemSwitchinterface Get(string name, Input<string> id, SystemSwitchinterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemSwitchinterface(name, id, state, options);
        }
    }

    public sealed class SystemSwitchinterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        /// </summary>
        [Input("intraSwitchPolicy")]
        public Input<string>? IntraSwitchPolicy { get; set; }

        /// <summary>
        /// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        /// </summary>
        [Input("macTtl")]
        public Input<int>? MacTtl { get; set; }

        [Input("members")]
        private InputList<Inputs.SystemSwitchinterfaceMemberArgs>? _members;

        /// <summary>
        /// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSwitchinterfaceMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.SystemSwitchinterfaceMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("span")]
        public Input<string>? Span { get; set; }

        /// <summary>
        /// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        /// </summary>
        [Input("spanDestPort")]
        public Input<string>? SpanDestPort { get; set; }

        /// <summary>
        /// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Input("spanDirection")]
        public Input<string>? SpanDirection { get; set; }

        [Input("spanSourcePorts")]
        private InputList<Inputs.SystemSwitchinterfaceSpanSourcePortArgs>? _spanSourcePorts;

        /// <summary>
        /// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSwitchinterfaceSpanSourcePortArgs> SpanSourcePorts
        {
            get => _spanSourcePorts ?? (_spanSourcePorts = new InputList<Inputs.SystemSwitchinterfaceSpanSourcePortArgs>());
            set => _spanSourcePorts = value;
        }

        /// <summary>
        /// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// VDOM that the software switch belongs to.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSwitchinterfaceArgs()
        {
        }
        public static new SystemSwitchinterfaceArgs Empty => new SystemSwitchinterfaceArgs();
    }

    public sealed class SystemSwitchinterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
        /// </summary>
        [Input("intraSwitchPolicy")]
        public Input<string>? IntraSwitchPolicy { get; set; }

        /// <summary>
        /// Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
        /// </summary>
        [Input("macTtl")]
        public Input<int>? MacTtl { get; set; }

        [Input("members")]
        private InputList<Inputs.SystemSwitchinterfaceMemberGetArgs>? _members;

        /// <summary>
        /// Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSwitchinterfaceMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.SystemSwitchinterfaceMemberGetArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("span")]
        public Input<string>? Span { get; set; }

        /// <summary>
        /// SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
        /// </summary>
        [Input("spanDestPort")]
        public Input<string>? SpanDestPort { get; set; }

        /// <summary>
        /// The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
        /// </summary>
        [Input("spanDirection")]
        public Input<string>? SpanDirection { get; set; }

        [Input("spanSourcePorts")]
        private InputList<Inputs.SystemSwitchinterfaceSpanSourcePortGetArgs>? _spanSourcePorts;

        /// <summary>
        /// Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSwitchinterfaceSpanSourcePortGetArgs> SpanSourcePorts
        {
            get => _spanSourcePorts ?? (_spanSourcePorts = new InputList<Inputs.SystemSwitchinterfaceSpanSourcePortGetArgs>());
            set => _spanSourcePorts = value;
        }

        /// <summary>
        /// Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// VDOM that the software switch belongs to.
        /// </summary>
        [Input("vdom")]
        public Input<string>? Vdom { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemSwitchinterfaceState()
        {
        }
        public static new SystemSwitchinterfaceState Empty => new SystemSwitchinterfaceState();
    }
}
