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
    /// BGP IPv6 network table.
    /// 
    /// &gt; The provider supports the definition of Network6 in Router Bgp `fortios.RouterBgp`, and also allows the definition of separate Network6 resources `fortios.RouterbgpNetwork6`, but do not use a `fortios.RouterBgp` with in-line Network6 in conjunction with any `fortios.RouterbgpNetwork6` resources, otherwise conflicts and overwrite will occur.
    /// 
    /// ## Import
    /// 
    /// Routerbgp Network6 can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerbgpNetwork6:RouterbgpNetwork6 labelname {{fosid}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/routerbgpNetwork6:RouterbgpNetwork6 labelname {{fosid}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/routerbgpNetwork6:RouterbgpNetwork6")]
    public partial class RouterbgpNetwork6 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("backdoor")]
        public Output<string> Backdoor { get; private set; } = null!;

        /// <summary>
        /// ID.
        /// </summary>
        [Output("fosid")]
        public Output<int> Fosid { get; private set; } = null!;

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Output("networkImportCheck")]
        public Output<string> NetworkImportCheck { get; private set; } = null!;

        /// <summary>
        /// Network IPv6 prefix.
        /// </summary>
        [Output("prefix6")]
        public Output<string> Prefix6 { get; private set; } = null!;

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Output("routeMap")]
        public Output<string> RouteMap { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a RouterbgpNetwork6 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouterbgpNetwork6(string name, RouterbgpNetwork6Args? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerbgpNetwork6:RouterbgpNetwork6", name, args ?? new RouterbgpNetwork6Args(), MakeResourceOptions(options, ""))
        {
        }

        private RouterbgpNetwork6(string name, Input<string> id, RouterbgpNetwork6State? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/routerbgpNetwork6:RouterbgpNetwork6", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouterbgpNetwork6 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouterbgpNetwork6 Get(string name, Input<string> id, RouterbgpNetwork6State? state = null, CustomResourceOptions? options = null)
        {
            return new RouterbgpNetwork6(name, id, state, options);
        }
    }

    public sealed class RouterbgpNetwork6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("backdoor")]
        public Input<string>? Backdoor { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Input("networkImportCheck")]
        public Input<string>? NetworkImportCheck { get; set; }

        /// <summary>
        /// Network IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RouterbgpNetwork6Args()
        {
        }
        public static new RouterbgpNetwork6Args Empty => new RouterbgpNetwork6Args();
    }

    public sealed class RouterbgpNetwork6State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("backdoor")]
        public Input<string>? Backdoor { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("fosid")]
        public Input<int>? Fosid { get; set; }

        /// <summary>
        /// Configure insurance of BGP network route existence in IGP. Valid values: `global`, `enable`, `disable`.
        /// </summary>
        [Input("networkImportCheck")]
        public Input<string>? NetworkImportCheck { get; set; }

        /// <summary>
        /// Network IPv6 prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        /// <summary>
        /// Route map to modify generated route.
        /// </summary>
        [Input("routeMap")]
        public Input<string>? RouteMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public RouterbgpNetwork6State()
        {
        }
        public static new RouterbgpNetwork6State Empty => new RouterbgpNetwork6State();
    }
}
