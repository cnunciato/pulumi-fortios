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
    public static class GetSystemFortisandbox
    {
        /// <summary>
        /// Use this data source to get information on fortios system fortisandbox
        /// </summary>
        public static Task<GetSystemFortisandboxResult> InvokeAsync(GetSystemFortisandboxArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemFortisandboxResult>("fortios:index/getSystemFortisandbox:getSystemFortisandbox", args ?? new GetSystemFortisandboxArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fortisandbox
        /// </summary>
        public static Output<GetSystemFortisandboxResult> Invoke(GetSystemFortisandboxInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemFortisandboxResult>("fortios:index/getSystemFortisandbox:getSystemFortisandbox", args ?? new GetSystemFortisandboxInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemFortisandboxArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemFortisandboxArgs()
        {
        }
        public static new GetSystemFortisandboxArgs Empty => new GetSystemFortisandboxArgs();
    }

    public sealed class GetSystemFortisandboxInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemFortisandboxInvokeArgs()
        {
        }
        public static new GetSystemFortisandboxInvokeArgs Empty => new GetSystemFortisandboxInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemFortisandboxResult
    {
        /// <summary>
        /// Notifier email address.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Configure the level of SSL protection for secure communication with FortiSandbox.
        /// </summary>
        public readonly string EncAlgorithm;
        /// <summary>
        /// Enable/disable FortiSandbox Cloud.
        /// </summary>
        public readonly string Forticloud;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable FortiSandbox inline scan.
        /// </summary>
        public readonly string InlineScan;
        /// <summary>
        /// Specify outgoing interface to reach server.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Specify how to select outgoing interface to reach server.
        /// </summary>
        public readonly string InterfaceSelectMethod;
        /// <summary>
        /// IPv4 or IPv6 address of the remote FortiSandbox.
        /// </summary>
        public readonly string Server;
        /// <summary>
        /// Source IP address for communications to FortiSandbox.
        /// </summary>
        public readonly string SourceIp;
        /// <summary>
        /// Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
        /// </summary>
        public readonly string SslMinProtoVersion;
        /// <summary>
        /// Enable/disable FortiSandbox.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemFortisandboxResult(
            string email,

            string encAlgorithm,

            string forticloud,

            string id,

            string inlineScan,

            string @interface,

            string interfaceSelectMethod,

            string server,

            string sourceIp,

            string sslMinProtoVersion,

            string status,

            string? vdomparam)
        {
            Email = email;
            EncAlgorithm = encAlgorithm;
            Forticloud = forticloud;
            Id = id;
            InlineScan = inlineScan;
            Interface = @interface;
            InterfaceSelectMethod = interfaceSelectMethod;
            Server = server;
            SourceIp = sourceIp;
            SslMinProtoVersion = sslMinProtoVersion;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
