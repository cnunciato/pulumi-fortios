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
    public static class GetSystemautoupdateTunneling
    {
        /// <summary>
        /// Use this data source to get information on fortios systemautoupdate tunneling
        /// </summary>
        public static Task<GetSystemautoupdateTunnelingResult> InvokeAsync(GetSystemautoupdateTunnelingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemautoupdateTunnelingResult>("fortios:index/getSystemautoupdateTunneling:getSystemautoupdateTunneling", args ?? new GetSystemautoupdateTunnelingArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios systemautoupdate tunneling
        /// </summary>
        public static Output<GetSystemautoupdateTunnelingResult> Invoke(GetSystemautoupdateTunnelingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemautoupdateTunnelingResult>("fortios:index/getSystemautoupdateTunneling:getSystemautoupdateTunneling", args ?? new GetSystemautoupdateTunnelingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemautoupdateTunnelingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemautoupdateTunnelingArgs()
        {
        }
        public static new GetSystemautoupdateTunnelingArgs Empty => new GetSystemautoupdateTunnelingArgs();
    }

    public sealed class GetSystemautoupdateTunnelingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemautoupdateTunnelingInvokeArgs()
        {
        }
        public static new GetSystemautoupdateTunnelingInvokeArgs Empty => new GetSystemautoupdateTunnelingInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemautoupdateTunnelingResult
    {
        /// <summary>
        /// Web proxy IP address or FQDN.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Web proxy password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Web proxy port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Enable/disable web proxy tunnelling.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Web proxy username.
        /// </summary>
        public readonly string Username;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemautoupdateTunnelingResult(
            string address,

            string id,

            string password,

            int port,

            string status,

            string username,

            string? vdomparam)
        {
            Address = address;
            Id = id;
            Password = password;
            Port = port;
            Status = status;
            Username = username;
            Vdomparam = vdomparam;
        }
    }
}
