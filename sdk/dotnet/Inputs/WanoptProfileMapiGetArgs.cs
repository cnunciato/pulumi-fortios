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

    public sealed class WanoptProfileMapiGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("byteCaching")]
        public Input<string>? ByteCaching { get; set; }

        /// <summary>
        /// Enable/disable logging. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("logTraffic")]
        public Input<string>? LogTraffic { get; set; }

        /// <summary>
        /// Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("secureTunnel")]
        public Input<string>? SecureTunnel { get; set; }

        /// <summary>
        /// Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
        /// </summary>
        [Input("tunnelSharing")]
        public Input<string>? TunnelSharing { get; set; }

        public WanoptProfileMapiGetArgs()
        {
        }
        public static new WanoptProfileMapiGetArgs Empty => new WanoptProfileMapiGetArgs();
    }
}
