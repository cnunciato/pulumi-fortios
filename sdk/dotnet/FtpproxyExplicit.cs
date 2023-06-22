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
    /// Configure explicit FTP proxy settings.
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
    ///     var trname = new Fortios.FtpproxyExplicit("trname", new()
    ///     {
    ///         IncomingIp = "0.0.0.0",
    ///         SecDefaultAction = "deny",
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FtpProxy Explicit can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/ftpproxyExplicit:FtpproxyExplicit labelname FtpProxyExplicit
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/ftpproxyExplicit:FtpproxyExplicit labelname FtpProxyExplicit
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/ftpproxyExplicit:FtpproxyExplicit")]
    public partial class FtpproxyExplicit : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Accept incoming FTP requests from this IP address. An interface must have this IP address.
        /// </summary>
        [Output("incomingIp")]
        public Output<string> IncomingIp { get; private set; } = null!;

        /// <summary>
        /// Accept incoming FTP requests on one or more ports.
        /// </summary>
        [Output("incomingPort")]
        public Output<string> IncomingPort { get; private set; } = null!;

        /// <summary>
        /// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        /// </summary>
        [Output("outgoingIp")]
        public Output<string> OutgoingIp { get; private set; } = null!;

        /// <summary>
        /// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Output("secDefaultAction")]
        public Output<string> SecDefaultAction { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("ssl")]
        public Output<string> Ssl { get; private set; } = null!;

        /// <summary>
        /// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Output("sslAlgorithm")]
        public Output<string> SslAlgorithm { get; private set; } = null!;

        /// <summary>
        /// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        /// </summary>
        [Output("sslCert")]
        public Output<string> SslCert { get; private set; } = null!;

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Output("sslDhBits")]
        public Output<string> SslDhBits { get; private set; } = null!;

        /// <summary>
        /// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FtpproxyExplicit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FtpproxyExplicit(string name, FtpproxyExplicitArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/ftpproxyExplicit:FtpproxyExplicit", name, args ?? new FtpproxyExplicitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FtpproxyExplicit(string name, Input<string> id, FtpproxyExplicitState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/ftpproxyExplicit:FtpproxyExplicit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FtpproxyExplicit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FtpproxyExplicit Get(string name, Input<string> id, FtpproxyExplicitState? state = null, CustomResourceOptions? options = null)
        {
            return new FtpproxyExplicit(name, id, state, options);
        }
    }

    public sealed class FtpproxyExplicitArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept incoming FTP requests from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("incomingIp")]
        public Input<string>? IncomingIp { get; set; }

        /// <summary>
        /// Accept incoming FTP requests on one or more ports.
        /// </summary>
        [Input("incomingPort")]
        public Input<string>? IncomingPort { get; set; }

        /// <summary>
        /// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("secDefaultAction")]
        public Input<string>? SecDefaultAction { get; set; }

        /// <summary>
        /// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssl")]
        public Input<string>? Ssl { get; set; }

        /// <summary>
        /// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("sslAlgorithm")]
        public Input<string>? SslAlgorithm { get; set; }

        /// <summary>
        /// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        /// </summary>
        [Input("sslCert")]
        public Input<string>? SslCert { get; set; }

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Input("sslDhBits")]
        public Input<string>? SslDhBits { get; set; }

        /// <summary>
        /// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtpproxyExplicitArgs()
        {
        }
        public static new FtpproxyExplicitArgs Empty => new FtpproxyExplicitArgs();
    }

    public sealed class FtpproxyExplicitState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Accept incoming FTP requests from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("incomingIp")]
        public Input<string>? IncomingIp { get; set; }

        /// <summary>
        /// Accept incoming FTP requests on one or more ports.
        /// </summary>
        [Input("incomingPort")]
        public Input<string>? IncomingPort { get; set; }

        /// <summary>
        /// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
        /// </summary>
        [Input("outgoingIp")]
        public Input<string>? OutgoingIp { get; set; }

        /// <summary>
        /// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
        /// </summary>
        [Input("secDefaultAction")]
        public Input<string>? SecDefaultAction { get; set; }

        /// <summary>
        /// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ssl")]
        public Input<string>? Ssl { get; set; }

        /// <summary>
        /// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
        /// </summary>
        [Input("sslAlgorithm")]
        public Input<string>? SslAlgorithm { get; set; }

        /// <summary>
        /// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
        /// </summary>
        [Input("sslCert")]
        public Input<string>? SslCert { get; set; }

        /// <summary>
        /// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        /// </summary>
        [Input("sslDhBits")]
        public Input<string>? SslDhBits { get; set; }

        /// <summary>
        /// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FtpproxyExplicitState()
        {
        }
        public static new FtpproxyExplicitState Empty => new FtpproxyExplicitState();
    }
}