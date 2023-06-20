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
    public static class GetFirewallCentralsnatmap
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall centralsnatmap
        /// </summary>
        public static Task<GetFirewallCentralsnatmapResult> InvokeAsync(GetFirewallCentralsnatmapArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallCentralsnatmapResult>("fortios:index/getFirewallCentralsnatmap:getFirewallCentralsnatmap", args ?? new GetFirewallCentralsnatmapArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall centralsnatmap
        /// </summary>
        public static Output<GetFirewallCentralsnatmapResult> Invoke(GetFirewallCentralsnatmapInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallCentralsnatmapResult>("fortios:index/getFirewallCentralsnatmap:getFirewallCentralsnatmap", args ?? new GetFirewallCentralsnatmapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallCentralsnatmapArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall centralsnatmap.
        /// </summary>
        [Input("policyid", required: true)]
        public int Policyid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallCentralsnatmapArgs()
        {
        }
        public static new GetFirewallCentralsnatmapArgs Empty => new GetFirewallCentralsnatmapArgs();
    }

    public sealed class GetFirewallCentralsnatmapInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall centralsnatmap.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<int> Policyid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallCentralsnatmapInvokeArgs()
        {
        }
        public static new GetFirewallCentralsnatmapInvokeArgs Empty => new GetFirewallCentralsnatmapInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallCentralsnatmapResult
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// IPv6 Destination address. The structure of `dst_addr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapDstAddr6Result> DstAddr6s;
        /// <summary>
        /// Destination address name from available addresses. The structure of `dst_addr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapDstAddrResult> DstAddrs;
        /// <summary>
        /// Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapDstintfResult> Dstintfs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable source NAT.
        /// </summary>
        public readonly string Nat;
        /// <summary>
        /// Enable/disable NAT46.
        /// </summary>
        public readonly string Nat46;
        /// <summary>
        /// Enable/disable NAT64.
        /// </summary>
        public readonly string Nat64;
        /// <summary>
        /// IPv6 pools to be used for source NAT. The structure of `nat_ippool6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapNatIppool6Result> NatIppool6s;
        /// <summary>
        /// Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `nat_ippool` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapNatIppoolResult> NatIppools;
        /// <summary>
        /// Translated port or port range (0 to 65535).
        /// </summary>
        public readonly string NatPort;
        /// <summary>
        /// IPv6 Original address. The structure of `orig_addr6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapOrigAddr6Result> OrigAddr6s;
        /// <summary>
        /// Original address. The structure of `orig_addr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapOrigAddrResult> OrigAddrs;
        /// <summary>
        /// Original TCP port (0 to 65535).
        /// </summary>
        public readonly string OrigPort;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int Policyid;
        /// <summary>
        /// Integer value for the protocol type (0 - 255).
        /// </summary>
        public readonly int Protocol;
        /// <summary>
        /// Source interface name from available interfaces. The structure of `srcintf` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallCentralsnatmapSrcintfResult> Srcintfs;
        /// <summary>
        /// Enable/disable the active status of this policy.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// IPv4/IPv6 source NAT.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallCentralsnatmapResult(
            string comments,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapDstAddr6Result> dstAddr6s,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapDstAddrResult> dstAddrs,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapDstintfResult> dstintfs,

            string id,

            string nat,

            string nat46,

            string nat64,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapNatIppool6Result> natIppool6s,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapNatIppoolResult> natIppools,

            string natPort,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapOrigAddr6Result> origAddr6s,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapOrigAddrResult> origAddrs,

            string origPort,

            int policyid,

            int protocol,

            ImmutableArray<Outputs.GetFirewallCentralsnatmapSrcintfResult> srcintfs,

            string status,

            string type,

            string uuid,

            string? vdomparam)
        {
            Comments = comments;
            DstAddr6s = dstAddr6s;
            DstAddrs = dstAddrs;
            Dstintfs = dstintfs;
            Id = id;
            Nat = nat;
            Nat46 = nat46;
            Nat64 = nat64;
            NatIppool6s = natIppool6s;
            NatIppools = natIppools;
            NatPort = natPort;
            OrigAddr6s = origAddr6s;
            OrigAddrs = origAddrs;
            OrigPort = origPort;
            Policyid = policyid;
            Protocol = protocol;
            Srcintfs = srcintfs;
            Status = status;
            Type = type;
            Uuid = uuid;
            Vdomparam = vdomparam;
        }
    }
}
