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
    /// Configure RIP.
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
    ///     var trname = new Fortios.RouterRip("trname", new()
    ///     {
    ///         DefaultInformationOriginate = "disable",
    ///         DefaultMetric = 1,
    ///         GarbageTimer = 120,
    ///         MaxOutMetric = 0,
    ///         RecvBufferSize = 655360,
    ///         Redistributes = new[]
    ///         {
    ///             new Fortios.Inputs.RouterRipRedistributeArgs
    ///             {
    ///                 Metric = 10,
    ///                 Name = "connected",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Inputs.RouterRipRedistributeArgs
    ///             {
    ///                 Metric = 10,
    ///                 Name = "static",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Inputs.RouterRipRedistributeArgs
    ///             {
    ///                 Metric = 10,
    ///                 Name = "ospf",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Inputs.RouterRipRedistributeArgs
    ///             {
    ///                 Metric = 10,
    ///                 Name = "bgp",
    ///                 Status = "disable",
    ///             },
    ///             new Fortios.Inputs.RouterRipRedistributeArgs
    ///             {
    ///                 Metric = 10,
    ///                 Name = "isis",
    ///                 Status = "disable",
    ///             },
    ///         },
    ///         TimeoutTimer = 180,
    ///         UpdateTimer = 30,
    ///         Version = "2",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Router Rip can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerRip:RouterRip labelname RouterRip
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerRip:RouterRip labelname RouterRip
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/routerRip:RouterRip")]
    public partial class RouterRip : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("defaultInformationOriginate")]
        public Output<string> DefaultInformationOriginate { get; private set; } = null!;

        /// <summary>
        /// Default metric.
        /// </summary>
        [Output("defaultMetric")]
        public Output<int> DefaultMetric { get; private set; } = null!;

        /// <summary>
        /// distance The structure of `distance` block is documented below.
        /// </summary>
        [Output("distances")]
        public Output<ImmutableArray<Outputs.RouterRipDistance>> Distances { get; private set; } = null!;

        /// <summary>
        /// Distribute list. The structure of `distribute_list` block is documented below.
        /// </summary>
        [Output("distributeLists")]
        public Output<ImmutableArray<Outputs.RouterRipDistributeList>> DistributeLists { get; private set; } = null!;

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Output("dynamicSortSubtable")]
        public Output<string?> DynamicSortSubtable { get; private set; } = null!;

        /// <summary>
        /// Garbage timer in seconds.
        /// </summary>
        [Output("garbageTimer")]
        public Output<int> GarbageTimer { get; private set; } = null!;

        /// <summary>
        /// RIP interface configuration. The structure of `interface` block is documented below.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.RouterRipInterface>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// Maximum metric allowed to output(0 means 'not set').
        /// </summary>
        [Output("maxOutMetric")]
        public Output<int> MaxOutMetric { get; private set; } = null!;

        /// <summary>
        /// neighbor The structure of `neighbor` block is documented below.
        /// </summary>
        [Output("neighbors")]
        public Output<ImmutableArray<Outputs.RouterRipNeighbor>> Neighbors { get; private set; } = null!;

        /// <summary>
        /// network The structure of `network` block is documented below.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<Outputs.RouterRipNetwork>> Networks { get; private set; } = null!;

        /// <summary>
        /// Offset list. The structure of `offset_list` block is documented below.
        /// </summary>
        [Output("offsetLists")]
        public Output<ImmutableArray<Outputs.RouterRipOffsetList>> OffsetLists { get; private set; } = null!;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        [Output("passiveInterfaces")]
        public Output<ImmutableArray<Outputs.RouterRipPassiveInterface>> PassiveInterfaces { get; private set; } = null!;

        /// <summary>
        /// Receiving buffer size.
        /// </summary>
        [Output("recvBufferSize")]
        public Output<int> RecvBufferSize { get; private set; } = null!;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        [Output("redistributes")]
        public Output<ImmutableArray<Outputs.RouterRipRedistribute>> Redistributes { get; private set; } = null!;

        /// <summary>
        /// Timeout timer in seconds.
        /// </summary>
        [Output("timeoutTimer")]
        public Output<int> TimeoutTimer { get; private set; } = null!;

        /// <summary>
        /// Update timer in seconds.
        /// </summary>
        [Output("updateTimer")]
        public Output<int> UpdateTimer { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;

        /// <summary>
        /// RIP version. Valid values: `1`, `2`.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a RouterRip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterRip(string name, RouterRipArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerRip:RouterRip", name, args ?? new RouterRipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouterRip(string name, Input<string> id, RouterRipState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerRip:RouterRip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterRip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterRip Get(string name, Input<string> id, RouterRipState? state = null, CustomResourceOptions? options = null)
        {
            return new RouterRip(name, id, state, options);
        }
    }

    public sealed class RouterRipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default metric.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        [Input("distances")]
        private InputList<Inputs.RouterRipDistanceArgs>? _distances;

        /// <summary>
        /// distance The structure of `distance` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipDistanceArgs> Distances
        {
            get => _distances ?? (_distances = new InputList<Inputs.RouterRipDistanceArgs>());
            set => _distances = value;
        }

        [Input("distributeLists")]
        private InputList<Inputs.RouterRipDistributeListArgs>? _distributeLists;

        /// <summary>
        /// Distribute list. The structure of `distribute_list` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipDistributeListArgs> DistributeLists
        {
            get => _distributeLists ?? (_distributeLists = new InputList<Inputs.RouterRipDistributeListArgs>());
            set => _distributeLists = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Garbage timer in seconds.
        /// </summary>
        [Input("garbageTimer")]
        public Input<int>? GarbageTimer { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.RouterRipInterfaceArgs>? _interfaces;

        /// <summary>
        /// RIP interface configuration. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipInterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.RouterRipInterfaceArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Maximum metric allowed to output(0 means 'not set').
        /// </summary>
        [Input("maxOutMetric")]
        public Input<int>? MaxOutMetric { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.RouterRipNeighborArgs>? _neighbors;

        /// <summary>
        /// neighbor The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipNeighborArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.RouterRipNeighborArgs>());
            set => _neighbors = value;
        }

        [Input("networks")]
        private InputList<Inputs.RouterRipNetworkArgs>? _networks;

        /// <summary>
        /// network The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipNetworkArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.RouterRipNetworkArgs>());
            set => _networks = value;
        }

        [Input("offsetLists")]
        private InputList<Inputs.RouterRipOffsetListArgs>? _offsetLists;

        /// <summary>
        /// Offset list. The structure of `offset_list` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipOffsetListArgs> OffsetLists
        {
            get => _offsetLists ?? (_offsetLists = new InputList<Inputs.RouterRipOffsetListArgs>());
            set => _offsetLists = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.RouterRipPassiveInterfaceArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipPassiveInterfaceArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.RouterRipPassiveInterfaceArgs>());
            set => _passiveInterfaces = value;
        }

        /// <summary>
        /// Receiving buffer size.
        /// </summary>
        [Input("recvBufferSize")]
        public Input<int>? RecvBufferSize { get; set; }

        [Input("redistributes")]
        private InputList<Inputs.RouterRipRedistributeArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipRedistributeArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.RouterRipRedistributeArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// Timeout timer in seconds.
        /// </summary>
        [Input("timeoutTimer")]
        public Input<int>? TimeoutTimer { get; set; }

        /// <summary>
        /// Update timer in seconds.
        /// </summary>
        [Input("updateTimer")]
        public Input<int>? UpdateTimer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// RIP version. Valid values: `1`, `2`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public RouterRipArgs()
        {
        }
        public static new RouterRipArgs Empty => new RouterRipArgs();
    }

    public sealed class RouterRipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable generation of default route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("defaultInformationOriginate")]
        public Input<string>? DefaultInformationOriginate { get; set; }

        /// <summary>
        /// Default metric.
        /// </summary>
        [Input("defaultMetric")]
        public Input<int>? DefaultMetric { get; set; }

        [Input("distances")]
        private InputList<Inputs.RouterRipDistanceGetArgs>? _distances;

        /// <summary>
        /// distance The structure of `distance` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipDistanceGetArgs> Distances
        {
            get => _distances ?? (_distances = new InputList<Inputs.RouterRipDistanceGetArgs>());
            set => _distances = value;
        }

        [Input("distributeLists")]
        private InputList<Inputs.RouterRipDistributeListGetArgs>? _distributeLists;

        /// <summary>
        /// Distribute list. The structure of `distribute_list` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipDistributeListGetArgs> DistributeLists
        {
            get => _distributeLists ?? (_distributeLists = new InputList<Inputs.RouterRipDistributeListGetArgs>());
            set => _distributeLists = value;
        }

        /// <summary>
        /// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -&gt; [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -&gt; [ a10, a2 ].
        /// </summary>
        [Input("dynamicSortSubtable")]
        public Input<string>? DynamicSortSubtable { get; set; }

        /// <summary>
        /// Garbage timer in seconds.
        /// </summary>
        [Input("garbageTimer")]
        public Input<int>? GarbageTimer { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.RouterRipInterfaceGetArgs>? _interfaces;

        /// <summary>
        /// RIP interface configuration. The structure of `interface` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipInterfaceGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.RouterRipInterfaceGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Maximum metric allowed to output(0 means 'not set').
        /// </summary>
        [Input("maxOutMetric")]
        public Input<int>? MaxOutMetric { get; set; }

        [Input("neighbors")]
        private InputList<Inputs.RouterRipNeighborGetArgs>? _neighbors;

        /// <summary>
        /// neighbor The structure of `neighbor` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipNeighborGetArgs> Neighbors
        {
            get => _neighbors ?? (_neighbors = new InputList<Inputs.RouterRipNeighborGetArgs>());
            set => _neighbors = value;
        }

        [Input("networks")]
        private InputList<Inputs.RouterRipNetworkGetArgs>? _networks;

        /// <summary>
        /// network The structure of `network` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipNetworkGetArgs> Networks
        {
            get => _networks ?? (_networks = new InputList<Inputs.RouterRipNetworkGetArgs>());
            set => _networks = value;
        }

        [Input("offsetLists")]
        private InputList<Inputs.RouterRipOffsetListGetArgs>? _offsetLists;

        /// <summary>
        /// Offset list. The structure of `offset_list` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipOffsetListGetArgs> OffsetLists
        {
            get => _offsetLists ?? (_offsetLists = new InputList<Inputs.RouterRipOffsetListGetArgs>());
            set => _offsetLists = value;
        }

        [Input("passiveInterfaces")]
        private InputList<Inputs.RouterRipPassiveInterfaceGetArgs>? _passiveInterfaces;

        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipPassiveInterfaceGetArgs> PassiveInterfaces
        {
            get => _passiveInterfaces ?? (_passiveInterfaces = new InputList<Inputs.RouterRipPassiveInterfaceGetArgs>());
            set => _passiveInterfaces = value;
        }

        /// <summary>
        /// Receiving buffer size.
        /// </summary>
        [Input("recvBufferSize")]
        public Input<int>? RecvBufferSize { get; set; }

        [Input("redistributes")]
        private InputList<Inputs.RouterRipRedistributeGetArgs>? _redistributes;

        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public InputList<Inputs.RouterRipRedistributeGetArgs> Redistributes
        {
            get => _redistributes ?? (_redistributes = new InputList<Inputs.RouterRipRedistributeGetArgs>());
            set => _redistributes = value;
        }

        /// <summary>
        /// Timeout timer in seconds.
        /// </summary>
        [Input("timeoutTimer")]
        public Input<int>? TimeoutTimer { get; set; }

        /// <summary>
        /// Update timer in seconds.
        /// </summary>
        [Input("updateTimer")]
        public Input<int>? UpdateTimer { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        /// <summary>
        /// RIP version. Valid values: `1`, `2`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public RouterRipState()
        {
        }
        public static new RouterRipState Empty => new RouterRipState();
    }
}
