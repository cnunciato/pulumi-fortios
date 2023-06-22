// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class GetFirewallProfileprotocoloptionsPop3Result
    {
        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol.
        /// </summary>
        public readonly string InspectAll;
        /// <summary>
        /// One or more options that can be applied to the session.
        /// </summary>
        public readonly string Options;
        /// <summary>
        /// Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
        /// </summary>
        public readonly int OversizeLimit;
        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 445).
        /// </summary>
        public readonly int Ports;
        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before).
        /// </summary>
        public readonly string ProxyAfterTcpHandshake;
        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files.
        /// </summary>
        public readonly string ScanBzip2;
        /// <summary>
        /// SSL decryption and encryption performed by an external device.
        /// </summary>
        public readonly string SslOffloaded;
        /// <summary>
        /// Enable/disable adding an email signature to SMTP email messages as they pass through the FortiGate.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        public readonly int UncompressedNestLimit;
        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
        /// </summary>
        public readonly int UncompressedOversizeLimit;

        [OutputConstructor]
        private GetFirewallProfileprotocoloptionsPop3Result(
            string inspectAll,

            string options,

            int oversizeLimit,

            int ports,

            string proxyAfterTcpHandshake,

            string scanBzip2,

            string sslOffloaded,

            string status,

            int uncompressedNestLimit,

            int uncompressedOversizeLimit)
        {
            InspectAll = inspectAll;
            Options = options;
            OversizeLimit = oversizeLimit;
            Ports = ports;
            ProxyAfterTcpHandshake = proxyAfterTcpHandshake;
            ScanBzip2 = scanBzip2;
            SslOffloaded = sslOffloaded;
            Status = status;
            UncompressedNestLimit = uncompressedNestLimit;
            UncompressedOversizeLimit = uncompressedOversizeLimit;
        }
    }
}