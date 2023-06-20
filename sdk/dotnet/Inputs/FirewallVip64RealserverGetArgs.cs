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

    public sealed class FirewallVip64RealserverGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Restrict server to a client IP in this range.
        /// </summary>
        [Input("clientIp")]
        public Input<string>? ClientIp { get; set; }

        /// <summary>
        /// Per server health check. Valid values: `disable`, `enable`, `vip`.
        /// </summary>
        [Input("healthcheck")]
        public Input<string>? Healthcheck { get; set; }

        /// <summary>
        /// Hold down interval.
        /// </summary>
        [Input("holddownInterval")]
        public Input<int>? HolddownInterval { get; set; }

        /// <summary>
        /// Real server ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Mapped server IP.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Maximum number of connections allowed to server.
        /// </summary>
        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        /// <summary>
        /// Health monitors.
        /// </summary>
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        /// <summary>
        /// Mapped server port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Server administrative status. Valid values: `active`, `standby`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// weight
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public FirewallVip64RealserverGetArgs()
        {
        }
        public static new FirewallVip64RealserverGetArgs Empty => new FirewallVip64RealserverGetArgs();
    }
}
