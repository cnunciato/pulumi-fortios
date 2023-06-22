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
    public static class GetSystemVxlan
    {
        /// <summary>
        /// Use this data source to get information on an fortios system vxlan
        /// </summary>
        public static Task<GetSystemVxlanResult> InvokeAsync(GetSystemVxlanArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemVxlanResult>("fortios:index/getSystemVxlan:getSystemVxlan", args ?? new GetSystemVxlanArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system vxlan
        /// </summary>
        public static Output<GetSystemVxlanResult> Invoke(GetSystemVxlanInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemVxlanResult>("fortios:index/getSystemVxlan:getSystemVxlan", args ?? new GetSystemVxlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemVxlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system vxlan.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemVxlanArgs()
        {
        }
        public static new GetSystemVxlanArgs Empty => new GetSystemVxlanArgs();
    }

    public sealed class GetSystemVxlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system vxlan.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemVxlanInvokeArgs()
        {
        }
        public static new GetSystemVxlanInvokeArgs Empty => new GetSystemVxlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemVxlanResult
    {
        /// <summary>
        /// VXLAN destination port (1 - 65535, default = 4789).
        /// </summary>
        public readonly int Dstport;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Outgoing interface for VXLAN encapsulated traffic.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
        /// </summary>
        public readonly string IpVersion;
        /// <summary>
        /// VXLAN multicast TTL (1-255, default = 0).
        /// </summary>
        public readonly int MulticastTtl;
        /// <summary>
        /// VXLAN device or interface name. Must be a unique interface name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVxlanRemoteIp6Result> RemoteIp6s;
        /// <summary>
        /// IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVxlanRemoteIpResult> RemoteIps;
        public readonly string? Vdomparam;
        /// <summary>
        /// VXLAN network ID.
        /// </summary>
        public readonly int Vni;

        [OutputConstructor]
        private GetSystemVxlanResult(
            int dstport,

            string id,

            string @interface,

            string ipVersion,

            int multicastTtl,

            string name,

            ImmutableArray<Outputs.GetSystemVxlanRemoteIp6Result> remoteIp6s,

            ImmutableArray<Outputs.GetSystemVxlanRemoteIpResult> remoteIps,

            string? vdomparam,

            int vni)
        {
            Dstport = dstport;
            Id = id;
            Interface = @interface;
            IpVersion = ipVersion;
            MulticastTtl = multicastTtl;
            Name = name;
            RemoteIp6s = remoteIp6s;
            RemoteIps = remoteIps;
            Vdomparam = vdomparam;
            Vni = vni;
        }
    }
}