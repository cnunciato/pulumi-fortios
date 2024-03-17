// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Wanopt
{
    /// <summary>
    /// Configure WAN optimization peers.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Wanopt.Peer("trname", new()
    ///     {
    ///         Ip = "1.1.1.1",
    ///         PeerHostId = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Wanopt Peer can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/peer:Peer labelname {{peer_host_id}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:wanopt/peer:Peer labelname {{peer_host_id}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:wanopt/peer:Peer")]
    public partial class Peer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Peer IP address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Peer host ID.
        /// </summary>
        [Output("peerHostId")]
        public Output<string> PeerHostId { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Peer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peer(string name, PeerArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/peer:Peer", name, args ?? new PeerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peer(string name, Input<string> id, PeerState? state = null, CustomResourceOptions? options = null)
            : base("fortios:wanopt/peer:Peer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Peer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peer Get(string name, Input<string> id, PeerState? state = null, CustomResourceOptions? options = null)
        {
            return new Peer(name, id, state, options);
        }
    }

    public sealed class PeerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peer IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Peer host ID.
        /// </summary>
        [Input("peerHostId")]
        public Input<string>? PeerHostId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeerArgs()
        {
        }
        public static new PeerArgs Empty => new PeerArgs();
    }

    public sealed class PeerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Peer IP address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Peer host ID.
        /// </summary>
        [Input("peerHostId")]
        public Input<string>? PeerHostId { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PeerState()
        {
        }
        public static new PeerState Empty => new PeerState();
    }
}