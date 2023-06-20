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

    public sealed class SystemlldpNetworkpolicySoftphoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Differentiated Services Code Point (DSCP) value to advertise.
        /// </summary>
        [Input("dscp")]
        public Input<int>? Dscp { get; set; }

        /// <summary>
        /// 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Enable/disable advertising this policy. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.
        /// </summary>
        [Input("tag")]
        public Input<string>? Tag { get; set; }

        /// <summary>
        /// 802.1Q VLAN ID to advertise (1 - 4094).
        /// </summary>
        [Input("vlan")]
        public Input<int>? Vlan { get; set; }

        public SystemlldpNetworkpolicySoftphoneGetArgs()
        {
        }
        public static new SystemlldpNetworkpolicySoftphoneGetArgs Empty => new SystemlldpNetworkpolicySoftphoneGetArgs();
    }
}
