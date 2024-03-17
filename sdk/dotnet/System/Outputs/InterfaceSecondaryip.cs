// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Outputs
{

    [OutputType]
    public sealed class InterfaceSecondaryip
    {
        /// <summary>
        /// Management access settings for the secondary IP address.
        /// </summary>
        public readonly string? Allowaccess;
        /// <summary>
        /// Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.
        /// </summary>
        public readonly string? Detectprotocol;
        /// <summary>
        /// Gateway's ping server for this IP.
        /// </summary>
        public readonly string? Detectserver;
        /// <summary>
        /// Enable/disable detect gateway alive for first. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Gwdetect;
        /// <summary>
        /// HA election priority for the PING server.
        /// </summary>
        public readonly int? HaPriority;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Secondary IP address of the interface.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// PING server status.
        /// </summary>
        public readonly int? PingServStatus;

        [OutputConstructor]
        private InterfaceSecondaryip(
            string? allowaccess,

            string? detectprotocol,

            string? detectserver,

            string? gwdetect,

            int? haPriority,

            int? id,

            string? ip,

            int? pingServStatus)
        {
            Allowaccess = allowaccess;
            Detectprotocol = detectprotocol;
            Detectserver = detectserver;
            Gwdetect = gwdetect;
            HaPriority = haPriority;
            Id = id;
            Ip = ip;
            PingServStatus = pingServStatus;
        }
    }
}