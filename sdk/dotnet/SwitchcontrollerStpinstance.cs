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
    /// Configure FortiSwitch multiple spanning tree protocol (MSTP) instances. Applies to FortiOS Version `&gt;= 6.2.4`.
    /// 
    /// ## Import
    /// 
    /// SwitchController StpInstance can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance")]
    public partial class SwitchcontrollerStpinstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("fosid")]
        public Output<string> Fosid { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// Configure VLAN range for STP instance. The structure of `vlan_range` block is documented below.
        /// </summary>
        [Output("vlanRanges")]
        public Output<ImmutableArray<Outputs.SwitchcontrollerStpinstanceVlanRange>> VlanRanges { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerStpinstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerStpinstance(string name, SwitchcontrollerStpinstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance", name, args ?? new SwitchcontrollerStpinstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerStpinstance(string name, Input<string> id, SwitchcontrollerStpinstanceState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerStpinstance:SwitchcontrollerStpinstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchcontrollerStpinstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerStpinstance Get(string name, Input<string> id, SwitchcontrollerStpinstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerStpinstance(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerStpinstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vlanRanges")]
        private InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeArgs>? _vlanRanges;

        /// <summary>
        /// Configure VLAN range for STP instance. The structure of `vlan_range` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeArgs> VlanRanges
        {
            get => _vlanRanges ?? (_vlanRanges = new InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeArgs>());
            set => _vlanRanges = value;
        }

        public SwitchcontrollerStpinstanceArgs()
        {
        }
        public static new SwitchcontrollerStpinstanceArgs Empty => new SwitchcontrollerStpinstanceArgs();
    }

    public sealed class SwitchcontrollerStpinstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("fosid")]
        public Input<string>? Fosid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("vlanRanges")]
        private InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeGetArgs>? _vlanRanges;

        /// <summary>
        /// Configure VLAN range for STP instance. The structure of `vlan_range` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeGetArgs> VlanRanges
        {
            get => _vlanRanges ?? (_vlanRanges = new InputList<Inputs.SwitchcontrollerStpinstanceVlanRangeGetArgs>());
            set => _vlanRanges = value;
        }

        public SwitchcontrollerStpinstanceState()
        {
        }
        public static new SwitchcontrollerStpinstanceState Empty => new SwitchcontrollerStpinstanceState();
    }
}
