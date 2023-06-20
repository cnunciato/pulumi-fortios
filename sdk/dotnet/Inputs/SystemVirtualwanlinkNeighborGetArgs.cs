// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemVirtualwanlinkNeighborGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SD-WAN health-check name.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// IP address of neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Member sequence number.
        /// </summary>
        [Input("member")]
        public Input<int>? Member { get; set; }

        /// <summary>
        /// Role of neighbor. Valid values: `standalone`, `primary`, `secondary`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// SLA ID.
        /// </summary>
        [Input("slaId")]
        public Input<int>? SlaId { get; set; }

        public SystemVirtualwanlinkNeighborGetArgs()
        {
        }
        public static new SystemVirtualwanlinkNeighborGetArgs Empty => new SystemVirtualwanlinkNeighborGetArgs();
    }
}
