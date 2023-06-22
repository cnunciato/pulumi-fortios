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
    public static class GetSystemHamonitor
    {
        /// <summary>
        /// Use this data source to get information on fortios system hamonitor
        /// </summary>
        public static Task<GetSystemHamonitorResult> InvokeAsync(GetSystemHamonitorArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemHamonitorResult>("fortios:index/getSystemHamonitor:getSystemHamonitor", args ?? new GetSystemHamonitorArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system hamonitor
        /// </summary>
        public static Output<GetSystemHamonitorResult> Invoke(GetSystemHamonitorInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemHamonitorResult>("fortios:index/getSystemHamonitor:getSystemHamonitor", args ?? new GetSystemHamonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemHamonitorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemHamonitorArgs()
        {
        }
        public static new GetSystemHamonitorArgs Empty => new GetSystemHamonitorArgs();
    }

    public sealed class GetSystemHamonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemHamonitorInvokeArgs()
        {
        }
        public static new GetSystemHamonitorInvokeArgs Empty => new GetSystemHamonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemHamonitorResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable monitor VLAN interfaces.
        /// </summary>
        public readonly string MonitorVlan;
        public readonly string? Vdomparam;
        /// <summary>
        /// Configure heartbeat interval (seconds).
        /// </summary>
        public readonly int VlanHbInterval;
        /// <summary>
        /// VLAN lost heartbeat threshold (1 - 60).
        /// </summary>
        public readonly int VlanHbLostThreshold;

        [OutputConstructor]
        private GetSystemHamonitorResult(
            string id,

            string monitorVlan,

            string? vdomparam,

            int vlanHbInterval,

            int vlanHbLostThreshold)
        {
            Id = id;
            MonitorVlan = monitorVlan;
            Vdomparam = vdomparam;
            VlanHbInterval = vlanHbInterval;
            VlanHbLostThreshold = vlanHbLostThreshold;
        }
    }
}