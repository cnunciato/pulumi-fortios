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
    public static class GetSystemVirtualwanlink
    {
        /// <summary>
        /// Use this data source to get information on fortios system virtualwanlink
        /// </summary>
        public static Task<GetSystemVirtualwanlinkResult> InvokeAsync(GetSystemVirtualwanlinkArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemVirtualwanlinkResult>("fortios:index/getSystemVirtualwanlink:getSystemVirtualwanlink", args ?? new GetSystemVirtualwanlinkArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system virtualwanlink
        /// </summary>
        public static Output<GetSystemVirtualwanlinkResult> Invoke(GetSystemVirtualwanlinkInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemVirtualwanlinkResult>("fortios:index/getSystemVirtualwanlink:getSystemVirtualwanlink", args ?? new GetSystemVirtualwanlinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemVirtualwanlinkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemVirtualwanlinkArgs()
        {
        }
        public static new GetSystemVirtualwanlinkArgs Empty => new GetSystemVirtualwanlinkArgs();
    }

    public sealed class GetSystemVirtualwanlinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemVirtualwanlinkInvokeArgs()
        {
        }
        public static new GetSystemVirtualwanlinkInvokeArgs Empty => new GetSystemVirtualwanlinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemVirtualwanlinkResult
    {
        /// <summary>
        /// Physical interfaces that will be alerted. The structure of `fail_alert_interfaces` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkFailAlertInterfaceResult> FailAlertInterfaces;
        /// <summary>
        /// Enable/disable SD-WAN Internet connection status checking (failure detection).
        /// </summary>
        public readonly string FailDetect;
        /// <summary>
        /// Virtual WAN Link health-check.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkHealthCheckResult> HealthChecks;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Algorithm or mode to use for load balancing Internet traffic to SD-WAN members.
        /// </summary>
        public readonly string LoadBalanceMode;
        /// <summary>
        /// Member sequence number list. The structure of `members` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkMemberResult> Members;
        /// <summary>
        /// Waiting period in seconds when switching from the primary neighbor to the secondary neighbor from the neighbor start. (0 - 10000000, default = 0).
        /// </summary>
        public readonly int NeighborHoldBootTime;
        /// <summary>
        /// Enable/disable hold switching from the secondary neighbor to the primary neighbor.
        /// </summary>
        public readonly string NeighborHoldDown;
        /// <summary>
        /// Waiting period in seconds when switching from the secondary neighbor to the primary neighbor when hold-down is disabled. (0 - 10000000, default = 0).
        /// </summary>
        public readonly int NeighborHoldDownTime;
        /// <summary>
        /// Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status. The structure of `neighbor` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkNeighborResult> Neighbors;
        /// <summary>
        /// Create SD-WAN rules or priority rules (also called services) to control how sessions are distributed to physical interfaces in the SD-WAN. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkServiceResult> Services;
        /// <summary>
        /// Enable/disable SD-WAN service.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;
        /// <summary>
        /// Configure SD-WAN zones. The structure of `zone` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemVirtualwanlinkZoneResult> Zones;

        [OutputConstructor]
        private GetSystemVirtualwanlinkResult(
            ImmutableArray<Outputs.GetSystemVirtualwanlinkFailAlertInterfaceResult> failAlertInterfaces,

            string failDetect,

            ImmutableArray<Outputs.GetSystemVirtualwanlinkHealthCheckResult> healthChecks,

            string id,

            string loadBalanceMode,

            ImmutableArray<Outputs.GetSystemVirtualwanlinkMemberResult> members,

            int neighborHoldBootTime,

            string neighborHoldDown,

            int neighborHoldDownTime,

            ImmutableArray<Outputs.GetSystemVirtualwanlinkNeighborResult> neighbors,

            ImmutableArray<Outputs.GetSystemVirtualwanlinkServiceResult> services,

            string status,

            string? vdomparam,

            ImmutableArray<Outputs.GetSystemVirtualwanlinkZoneResult> zones)
        {
            FailAlertInterfaces = failAlertInterfaces;
            FailDetect = failDetect;
            HealthChecks = healthChecks;
            Id = id;
            LoadBalanceMode = loadBalanceMode;
            Members = members;
            NeighborHoldBootTime = neighborHoldBootTime;
            NeighborHoldDown = neighborHoldDown;
            NeighborHoldDownTime = neighborHoldDownTime;
            Neighbors = neighbors;
            Services = services;
            Status = status;
            Vdomparam = vdomparam;
            Zones = zones;
        }
    }
}
