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
    /// Configure WTP groups.
    /// 
    /// ## Import
    /// 
    /// WirelessController WtpGroup can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup")]
    public partial class WirelesscontrollerWtpgroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// WTP group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// FortiAP models to define the WTP group platform type.
        /// </summary>
        [Output("platformType")]
        public Output<string> PlatformType { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// WTP list. The structure of `wtps` block is documented below.
        /// </summary>
        [Output("wtps")]
        public Output<ImmutableArray<Outputs.WirelesscontrollerWtpgroupWtp>> Wtps { get; private set; } = null!;


        /// <summary>
        /// Create a WirelesscontrollerWtpgroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WirelesscontrollerWtpgroup(string name, WirelesscontrollerWtpgroupArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup", name, args ?? new WirelesscontrollerWtpgroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WirelesscontrollerWtpgroup(string name, Input<string> id, WirelesscontrollerWtpgroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/wirelesscontrollerWtpgroup:WirelesscontrollerWtpgroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WirelesscontrollerWtpgroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WirelesscontrollerWtpgroup Get(string name, Input<string> id, WirelesscontrollerWtpgroupState? state = null, CustomResourceOptions? options = null)
        {
            return new WirelesscontrollerWtpgroup(name, id, state, options);
        }
    }

    public sealed class WirelesscontrollerWtpgroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// WTP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FortiAP models to define the WTP group platform type.
        /// </summary>
        [Input("platformType")]
        public Input<string>? PlatformType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("wtps")]
        private InputList<Inputs.WirelesscontrollerWtpgroupWtpArgs>? _wtps;

        /// <summary>
        /// WTP list. The structure of `wtps` block is documented below.
        /// </summary>
        public InputList<Inputs.WirelesscontrollerWtpgroupWtpArgs> Wtps
        {
            get => _wtps ?? (_wtps = new InputList<Inputs.WirelesscontrollerWtpgroupWtpArgs>());
            set => _wtps = value;
        }

        public WirelesscontrollerWtpgroupArgs()
        {
        }
        public static new WirelesscontrollerWtpgroupArgs Empty => new WirelesscontrollerWtpgroupArgs();
    }

    public sealed class WirelesscontrollerWtpgroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// WTP group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// FortiAP models to define the WTP group platform type.
        /// </summary>
        [Input("platformType")]
        public Input<string>? PlatformType { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        [Input("wtps")]
        private InputList<Inputs.WirelesscontrollerWtpgroupWtpGetArgs>? _wtps;

        /// <summary>
        /// WTP list. The structure of `wtps` block is documented below.
        /// </summary>
        public InputList<Inputs.WirelesscontrollerWtpgroupWtpGetArgs> Wtps
        {
            get => _wtps ?? (_wtps = new InputList<Inputs.WirelesscontrollerWtpgroupWtpGetArgs>());
            set => _wtps = value;
        }

        public WirelesscontrollerWtpgroupState()
        {
        }
        public static new WirelesscontrollerWtpgroupState Empty => new WirelesscontrollerWtpgroupState();
    }
}
