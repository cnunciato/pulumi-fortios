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
    public static class GetFirewallPolicy46
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall policy46
        /// </summary>
        public static Task<GetFirewallPolicy46Result> InvokeAsync(GetFirewallPolicy46Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallPolicy46Result>("fortios:index/getFirewallPolicy46:getFirewallPolicy46", args ?? new GetFirewallPolicy46Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall policy46
        /// </summary>
        public static Output<GetFirewallPolicy46Result> Invoke(GetFirewallPolicy46InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallPolicy46Result>("fortios:index/getFirewallPolicy46:getFirewallPolicy46", args ?? new GetFirewallPolicy46InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallPolicy46Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall policy46.
        /// </summary>
        [Input("policyid", required: true)]
        public int Policyid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallPolicy46Args()
        {
        }
        public static new GetFirewallPolicy46Args Empty => new GetFirewallPolicy46Args();
    }

    public sealed class GetFirewallPolicy46InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall policy46.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<int> Policyid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallPolicy46InvokeArgs()
        {
        }
        public static new GetFirewallPolicy46InvokeArgs Empty => new GetFirewallPolicy46InvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallPolicy46Result
    {
        /// <summary>
        /// Accept or deny traffic matching the policy.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Destination address objects. The structure of `dstaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallPolicy46DstaddrResult> Dstaddrs;
        /// <summary>
        /// Destination interface name.
        /// </summary>
        public readonly string Dstintf;
        /// <summary>
        /// Enable/disable fixed port for this policy.
        /// </summary>
        public readonly string Fixedport;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable use of IP Pools for source NAT.
        /// </summary>
        public readonly string Ippool;
        /// <summary>
        /// Enable/disable traffic logging for this policy.
        /// </summary>
        public readonly string Logtraffic;
        /// <summary>
        /// Record logs when a session starts and ends.
        /// </summary>
        public readonly string LogtrafficStart;
        /// <summary>
        /// IP pool name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Per IP traffic shaper.
        /// </summary>
        public readonly string PerIpShaper;
        /// <summary>
        /// Enable/disable allowing any host.
        /// </summary>
        public readonly string PermitAnyHost;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int Policyid;
        /// <summary>
        /// IP Pool names. The structure of `poolname` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallPolicy46PoolnameResult> Poolnames;
        /// <summary>
        /// Schedule name.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Service name. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallPolicy46ServiceResult> Services;
        /// <summary>
        /// Source address objects. The structure of `srcaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallPolicy46SrcaddrResult> Srcaddrs;
        /// <summary>
        /// Source interface name.
        /// </summary>
        public readonly string Srcintf;
        /// <summary>
        /// Enable/disable this policy.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
        /// </summary>
        public readonly int TcpMssReceiver;
        /// <summary>
        /// TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
        /// </summary>
        public readonly int TcpMssSender;
        /// <summary>
        /// Traffic shaper.
        /// </summary>
        public readonly string TrafficShaper;
        /// <summary>
        /// Reverse traffic shaper.
        /// </summary>
        public readonly string TrafficShaperReverse;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallPolicy46Result(
            string action,

            string comments,

            ImmutableArray<Outputs.GetFirewallPolicy46DstaddrResult> dstaddrs,

            string dstintf,

            string fixedport,

            string id,

            string ippool,

            string logtraffic,

            string logtrafficStart,

            string name,

            string perIpShaper,

            string permitAnyHost,

            int policyid,

            ImmutableArray<Outputs.GetFirewallPolicy46PoolnameResult> poolnames,

            string schedule,

            ImmutableArray<Outputs.GetFirewallPolicy46ServiceResult> services,

            ImmutableArray<Outputs.GetFirewallPolicy46SrcaddrResult> srcaddrs,

            string srcintf,

            string status,

            int tcpMssReceiver,

            int tcpMssSender,

            string trafficShaper,

            string trafficShaperReverse,

            string uuid,

            string? vdomparam)
        {
            Action = action;
            Comments = comments;
            Dstaddrs = dstaddrs;
            Dstintf = dstintf;
            Fixedport = fixedport;
            Id = id;
            Ippool = ippool;
            Logtraffic = logtraffic;
            LogtrafficStart = logtrafficStart;
            Name = name;
            PerIpShaper = perIpShaper;
            PermitAnyHost = permitAnyHost;
            Policyid = policyid;
            Poolnames = poolnames;
            Schedule = schedule;
            Services = services;
            Srcaddrs = srcaddrs;
            Srcintf = srcintf;
            Status = status;
            TcpMssReceiver = tcpMssReceiver;
            TcpMssSender = tcpMssSender;
            TrafficShaper = trafficShaper;
            TrafficShaperReverse = trafficShaperReverse;
            Uuid = uuid;
            Vdomparam = vdomparam;
        }
    }
}
