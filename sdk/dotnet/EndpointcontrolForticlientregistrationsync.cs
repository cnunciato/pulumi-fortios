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
    /// Configure FortiClient registration synchronization settings. Applies to FortiOS Version `&lt;= 6.2.0`.
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
    ///     var trname = new Fortios.EndpointcontrolForticlientregistrationsync("trname", new()
    ///     {
    ///         PeerIp = "1.1.1.1",
    ///         PeerName = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync labelname {{peer_name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync")]
    public partial class EndpointcontrolForticlientregistrationsync : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address of the peer FortiGate for endpoint license synchronization.
        /// </summary>
        [Output("peerIp")]
        public Output<string> PeerIp { get; private set; } = null!;

        /// <summary>
        /// Peer name.
        /// </summary>
        [Output("peerName")]
        public Output<string> PeerName { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointcontrolForticlientregistrationsync resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointcontrolForticlientregistrationsync(string name, EndpointcontrolForticlientregistrationsyncArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync", name, args ?? new EndpointcontrolForticlientregistrationsyncArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointcontrolForticlientregistrationsync(string name, Input<string> id, EndpointcontrolForticlientregistrationsyncState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/endpointcontrolForticlientregistrationsync:EndpointcontrolForticlientregistrationsync", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointcontrolForticlientregistrationsync resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointcontrolForticlientregistrationsync Get(string name, Input<string> id, EndpointcontrolForticlientregistrationsyncState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointcontrolForticlientregistrationsync(name, id, state, options);
        }
    }

    public sealed class EndpointcontrolForticlientregistrationsyncArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of the peer FortiGate for endpoint license synchronization.
        /// </summary>
        [Input("peerIp", required: true)]
        public Input<string> PeerIp { get; set; } = null!;

        /// <summary>
        /// Peer name.
        /// </summary>
        [Input("peerName")]
        public Input<string>? PeerName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EndpointcontrolForticlientregistrationsyncArgs()
        {
        }
        public static new EndpointcontrolForticlientregistrationsyncArgs Empty => new EndpointcontrolForticlientregistrationsyncArgs();
    }

    public sealed class EndpointcontrolForticlientregistrationsyncState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of the peer FortiGate for endpoint license synchronization.
        /// </summary>
        [Input("peerIp")]
        public Input<string>? PeerIp { get; set; }

        /// <summary>
        /// Peer name.
        /// </summary>
        [Input("peerName")]
        public Input<string>? PeerName { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public EndpointcontrolForticlientregistrationsyncState()
        {
        }
        public static new EndpointcontrolForticlientregistrationsyncState Empty => new EndpointcontrolForticlientregistrationsyncState();
    }
}
