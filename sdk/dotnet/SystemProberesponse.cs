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
    /// Configure system probe response.
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
    ///     var trname = new Fortios.SystemProberesponse("trname", new()
    ///     {
    ///         HttpProbeValue = "OK",
    ///         Mode = "none",
    ///         Port = 8008,
    ///         SecurityMode = "none",
    ///         Timeout = 300,
    ///         TtlMode = "retain",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System ProbeResponse can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemProberesponse:SystemProberesponse labelname SystemProbeResponse
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/systemProberesponse:SystemProberesponse labelname SystemProbeResponse
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/systemProberesponse:SystemProberesponse")]
    public partial class SystemProberesponse : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Value to respond to the monitoring server.
        /// </summary>
        [Output("httpProbeValue")]
        public Output<string> HttpProbeValue { get; private set; } = null!;

        /// <summary>
        /// SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Twamp respondor password in authentication mode
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Port number to response.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Twamp respondor security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Output("securityMode")]
        public Output<string> SecurityMode { get; private set; } = null!;

        /// <summary>
        /// An inactivity timer for a twamp test session.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        /// </summary>
        [Output("ttlMode")]
        public Output<string> TtlMode { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SystemProberesponse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemProberesponse(string name, SystemProberesponseArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemProberesponse:SystemProberesponse", name, args ?? new SystemProberesponseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemProberesponse(string name, Input<string> id, SystemProberesponseState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/systemProberesponse:SystemProberesponse", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SystemProberesponse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemProberesponse Get(string name, Input<string> id, SystemProberesponseState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemProberesponse(name, id, state, options);
        }
    }

    public sealed class SystemProberesponseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value to respond to the monitoring server.
        /// </summary>
        [Input("httpProbeValue")]
        public Input<string>? HttpProbeValue { get; set; }

        /// <summary>
        /// SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Twamp respondor password in authentication mode
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number to response.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Twamp respondor security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        /// <summary>
        /// An inactivity timer for a twamp test session.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        /// </summary>
        [Input("ttlMode")]
        public Input<string>? TtlMode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemProberesponseArgs()
        {
        }
        public static new SystemProberesponseArgs Empty => new SystemProberesponseArgs();
    }

    public sealed class SystemProberesponseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value to respond to the monitoring server.
        /// </summary>
        [Input("httpProbeValue")]
        public Input<string>? HttpProbeValue { get; set; }

        /// <summary>
        /// SLA response mode. Valid values: `none`, `http-probe`, `twamp`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Twamp respondor password in authentication mode
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Port number to response.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Twamp respondor security mode. Valid values: `none`, `authentication`.
        /// </summary>
        [Input("securityMode")]
        public Input<string>? SecurityMode { get; set; }

        /// <summary>
        /// An inactivity timer for a twamp test session.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Mode for TWAMP packet TTL modification. Valid values: `reinit`, `decrease`, `retain`.
        /// </summary>
        [Input("ttlMode")]
        public Input<string>? TtlMode { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SystemProberesponseState()
        {
        }
        public static new SystemProberesponseState Empty => new SystemProberesponseState();
    }
}
