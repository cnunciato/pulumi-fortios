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
    /// Configure NetFlow per VDOM.
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
    ///     var trname = new Fortios.SystemVdomnetflow("trname", new()
    ///     {
    ///         CollectorIp = "0.0.0.0",
    ///         CollectorPort = 2055,
    ///         SourceIp = "0.0.0.0",
    ///         VdomNetflow = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System VdomNetflow can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemVdomnetflow:SystemVdomnetflow labelname SystemVdomNetflow
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemVdomnetflow:SystemVdomnetflow labelname SystemVdomNetflow
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemVdomnetflow:SystemVdomnetflow")]
    public partial class SystemVdomnetflow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// NetFlow collector IP address.
        /// </summary>
        [Output("collectorIp")]
        public Output<string> CollectorIp { get; private set; } = null!;

        /// <summary>
        /// NetFlow collector port number.
        /// </summary>
        [Output("collectorPort")]
        public Output<int> CollectorPort { get; private set; } = null!;

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Output("interfaceSelectMethod")]
        public Output<string> InterfaceSelectMethod { get; private set; } = null!;

        /// <summary>
        /// Source IP address for communication with the NetFlow agent.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("vdomNetflow")]
        public Output<string> VdomNetflow { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemVdomnetflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemVdomnetflow(string name, SystemVdomnetflowArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVdomnetflow:SystemVdomnetflow", name, args ?? new SystemVdomnetflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemVdomnetflow(string name, Input<string> id, SystemVdomnetflowState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVdomnetflow:SystemVdomnetflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemVdomnetflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemVdomnetflow Get(string name, Input<string> id, SystemVdomnetflowState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemVdomnetflow(name, id, state, options);
        }
    }

    public sealed class SystemVdomnetflowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NetFlow collector IP address.
        /// </summary>
        [Input("collectorIp")]
        public Input<string>? CollectorIp { get; set; }

        /// <summary>
        /// NetFlow collector port number.
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address for communication with the NetFlow agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vdomNetflow")]
        public Input<string>? VdomNetflow { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemVdomnetflowArgs()
        {
        }
        public static new SystemVdomnetflowArgs Empty => new SystemVdomnetflowArgs();
    }

    public sealed class SystemVdomnetflowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NetFlow collector IP address.
        /// </summary>
        [Input("collectorIp")]
        public Input<string>? CollectorIp { get; set; }

        /// <summary>
        /// NetFlow collector port number.
        /// </summary>
        [Input("collectorPort")]
        public Input<int>? CollectorPort { get; set; }

        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
        /// </summary>
        [Input("interfaceSelectMethod")]
        public Input<string>? InterfaceSelectMethod { get; set; }

        /// <summary>
        /// Source IP address for communication with the NetFlow agent.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("vdomNetflow")]
        public Input<string>? VdomNetflow { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemVdomnetflowState()
        {
        }
        public static new SystemVdomnetflowState Empty => new SystemVdomnetflowState();
    }
}
