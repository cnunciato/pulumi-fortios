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
    /// Configure virtual network enabler tunnel. Applies to FortiOS Version `&gt;= 6.4.1`.
    /// 
    /// ## Import
    /// 
    /// System VneTunnel can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemVnetunnel:SystemVnetunnel labelname SystemVneTunnel
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemVnetunnel:SystemVnetunnel labelname SystemVneTunnel
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemVnetunnel:SystemVnetunnel")]
    public partial class SystemVnetunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("autoAsicOffload")]
        public Output<string> AutoAsicOffload { get; private set; } = null!;

        /// <summary>
        /// BMR hostname.
        /// </summary>
        [Output("bmrHostname")]
        public Output<string?> BmrHostname { get; private set; } = null!;

        /// <summary>
        /// Border relay IPv6 address.
        /// </summary>
        [Output("br")]
        public Output<string> Br { get; private set; } = null!;

        /// <summary>
        /// HTTP authentication password.
        /// </summary>
        [Output("httpPassword")]
        public Output<string?> HttpPassword { get; private set; } = null!;

        /// <summary>
        /// HTTP authentication user name.
        /// </summary>
        [Output("httpUsername")]
        public Output<string> HttpUsername { get; private set; } = null!;

        /// <summary>
        /// Interface name.
        /// </summary>
        [Output("interface")]
        public Output<string> Interface { get; private set; } = null!;

        /// <summary>
        /// Tunnel IPv4 address and netmask.
        /// </summary>
        [Output("ipv4Address")]
        public Output<string> Ipv4Address { get; private set; } = null!;

        /// <summary>
        /// VNE tunnel mode.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Output("sslCertificate")]
        public Output<string> SslCertificate { get; private set; } = null!;

        /// <summary>
        /// Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// URL of provisioning server.
        /// </summary>
        [Output("updateUrl")]
        public Output<string> UpdateUrl { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemVnetunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemVnetunnel(string name, SystemVnetunnelArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVnetunnel:SystemVnetunnel", name, args ?? new SystemVnetunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemVnetunnel(string name, Input<string> id, SystemVnetunnelState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemVnetunnel:SystemVnetunnel", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "bmrHostname",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SystemVnetunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemVnetunnel Get(string name, Input<string> id, SystemVnetunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemVnetunnel(name, id, state, options);
        }
    }

    public sealed class SystemVnetunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        [Input("bmrHostname")]
        private Input<string>? _bmrHostname;

        /// <summary>
        /// BMR hostname.
        /// </summary>
        public Input<string>? BmrHostname
        {
            get => _bmrHostname;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bmrHostname = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Border relay IPv6 address.
        /// </summary>
        [Input("br")]
        public Input<string>? Br { get; set; }

        /// <summary>
        /// HTTP authentication password.
        /// </summary>
        [Input("httpPassword")]
        public Input<string>? HttpPassword { get; set; }

        /// <summary>
        /// HTTP authentication user name.
        /// </summary>
        [Input("httpUsername")]
        public Input<string>? HttpUsername { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Tunnel IPv4 address and netmask.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        /// <summary>
        /// VNE tunnel mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL of provisioning server.
        /// </summary>
        [Input("updateUrl")]
        public Input<string>? UpdateUrl { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemVnetunnelArgs()
        {
        }
        public static new SystemVnetunnelArgs Empty => new SystemVnetunnelArgs();
    }

    public sealed class SystemVnetunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("autoAsicOffload")]
        public Input<string>? AutoAsicOffload { get; set; }

        [Input("bmrHostname")]
        private Input<string>? _bmrHostname;

        /// <summary>
        /// BMR hostname.
        /// </summary>
        public Input<string>? BmrHostname
        {
            get => _bmrHostname;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bmrHostname = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Border relay IPv6 address.
        /// </summary>
        [Input("br")]
        public Input<string>? Br { get; set; }

        /// <summary>
        /// HTTP authentication password.
        /// </summary>
        [Input("httpPassword")]
        public Input<string>? HttpPassword { get; set; }

        /// <summary>
        /// HTTP authentication user name.
        /// </summary>
        [Input("httpUsername")]
        public Input<string>? HttpUsername { get; set; }

        /// <summary>
        /// Interface name.
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// Tunnel IPv4 address and netmask.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        /// <summary>
        /// VNE tunnel mode.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of local certificate for SSL connections.
        /// </summary>
        [Input("sslCertificate")]
        public Input<string>? SslCertificate { get; set; }

        /// <summary>
        /// Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// URL of provisioning server.
        /// </summary>
        [Input("updateUrl")]
        public Input<string>? UpdateUrl { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemVnetunnelState()
        {
        }
        public static new SystemVnetunnelState Empty => new SystemVnetunnelState();
    }
}
