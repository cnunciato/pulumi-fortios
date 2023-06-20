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
    public sealed class SystemSdwanHealthCheck
    {
        /// <summary>
        /// Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
        /// </summary>
        public readonly string? AddrMode;
        /// <summary>
        /// The mode determining how to detect the server.
        /// </summary>
        public readonly string? DetectMode;
        /// <summary>
        /// Differentiated services code point (DSCP) in the IP header of the probe packet.
        /// </summary>
        public readonly string? Diffservcode;
        /// <summary>
        /// Response IP expected from DNS server if the protocol is DNS.
        /// </summary>
        public readonly string? DnsMatchIp;
        /// <summary>
        /// Fully qualified domain name to resolve for the DNS probe.
        /// </summary>
        public readonly string? DnsRequestDomain;
        /// <summary>
        /// Enable/disable embedding measured health information. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? EmbedMeasuredHealth;
        /// <summary>
        /// Number of failures before server is considered lost (1 - 3600, default = 5).
        /// </summary>
        public readonly int? Failtime;
        /// <summary>
        /// Full path and file name on the FTP server to download for FTP health-check to probe.
        /// </summary>
        public readonly string? FtpFile;
        /// <summary>
        /// FTP mode. Valid values: `passive`, `port`.
        /// </summary>
        public readonly string? FtpMode;
        /// <summary>
        /// HA election priority (1 - 50).
        /// </summary>
        public readonly int? HaPriority;
        /// <summary>
        /// String in the http-agent field in the HTTP header.
        /// </summary>
        public readonly string? HttpAgent;
        /// <summary>
        /// URL used to communicate with the server if the protocol if the protocol is HTTP.
        /// </summary>
        public readonly string? HttpGet;
        /// <summary>
        /// Response string expected from the server if the protocol is HTTP.
        /// </summary>
        public readonly string? HttpMatch;
        /// <summary>
        /// Status check interval in milliseconds, or the time between attempting to connect to the server (500 - 3600*1000 msec, default = 500).
        /// </summary>
        public readonly int? Interval;
        /// <summary>
        /// Member sequence number list. The structure of `members` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SystemSdwanHealthCheckMember> Members;
        /// <summary>
        /// Codec to use for MOS calculation (default = g711). Valid values: `g711`, `g722`, `g729`.
        /// </summary>
        public readonly string? MosCodec;
        /// <summary>
        /// Health check name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Packet size of a twamp test session,
        /// </summary>
        public readonly int? PacketSize;
        /// <summary>
        /// Twamp controller password in authentication mode
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port number used to communicate with the server over the selected protocol (0-65535, default = 0, auto select. http, twamp: 80, udp-echo, tcp-echo: 7, dns: 53, ftp: 21).
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Number of most recent probes that should be used to calculate latency and jitter (5 - 30, default = 30).
        /// </summary>
        public readonly int? ProbeCount;
        /// <summary>
        /// Enable/disable transmission of probe packets. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? ProbePackets;
        /// <summary>
        /// Time to wait before a probe packet is considered lost (500 - 3600*1000 msec, default = 500).
        /// </summary>
        public readonly int? ProbeTimeout;
        /// <summary>
        /// Protocol used to determine if the FortiGate can communicate with the server.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Method to measure the quality of tcp-connect. Valid values: `half-open`, `half-close`.
        /// </summary>
        public readonly string? QualityMeasuredMethod;
        /// <summary>
        /// Number of successful responses received before server is considered recovered (1 - 3600, default = 5).
        /// </summary>
        public readonly int? Recoverytime;
        /// <summary>
        /// Twamp controller security mode. Valid values: `none`, `authentication`.
        /// </summary>
        public readonly string? SecurityMode;
        /// <summary>
        /// IP address or FQDN name of the server.
        /// </summary>
        public readonly string? Server;
        /// <summary>
        /// Time interval in seconds that SLA fail log messages will be generated (0 - 3600, default = 0).
        /// </summary>
        public readonly int? SlaFailLogPeriod;
        /// <summary>
        /// Select the ID from the SLA sub-table. The selected SLA's priority value will be distributed into the routing table (0 - 32, default = 0).
        /// </summary>
        public readonly int? SlaIdRedistribute;
        /// <summary>
        /// Time interval in seconds that SLA pass log messages will be generated (0 - 3600, default = 0).
        /// </summary>
        public readonly int? SlaPassLogPeriod;
        /// <summary>
        /// Service level agreement (SLA). The structure of `sla` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.SystemSdwanHealthCheckSla> Slas;
        /// <summary>
        /// Source IP address used in the health-check packet to the server.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// Enable/disable system DNS as the probe server. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? SystemDns;
        /// <summary>
        /// Alert threshold for jitter (ms, default = 0).
        /// </summary>
        public readonly int? ThresholdAlertJitter;
        /// <summary>
        /// Alert threshold for latency (ms, default = 0).
        /// </summary>
        public readonly int? ThresholdAlertLatency;
        /// <summary>
        /// Alert threshold for packet loss (percentage, default = 0).
        /// </summary>
        public readonly int? ThresholdAlertPacketloss;
        /// <summary>
        /// Warning threshold for jitter (ms, default = 0).
        /// </summary>
        public readonly int? ThresholdWarningJitter;
        /// <summary>
        /// Warning threshold for latency (ms, default = 0).
        /// </summary>
        public readonly int? ThresholdWarningLatency;
        /// <summary>
        /// Warning threshold for packet loss (percentage, default = 0).
        /// </summary>
        public readonly int? ThresholdWarningPacketloss;
        /// <summary>
        /// Enable/disable update cascade interface. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? UpdateCascadeInterface;
        /// <summary>
        /// Enable/disable updating the static route. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? UpdateStaticRoute;
        /// <summary>
        /// The user name to access probe server.
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// Virtual Routing Forwarding ID.
        /// </summary>
        public readonly int? Vrf;

        [OutputConstructor]
        private SystemSdwanHealthCheck(
            string? addrMode,

            string? detectMode,

            string? diffservcode,

            string? dnsMatchIp,

            string? dnsRequestDomain,

            string? embedMeasuredHealth,

            int? failtime,

            string? ftpFile,

            string? ftpMode,

            int? haPriority,

            string? httpAgent,

            string? httpGet,

            string? httpMatch,

            int? interval,

            ImmutableArray<Outputs.SystemSdwanHealthCheckMember> members,

            string? mosCodec,

            string? name,

            int? packetSize,

            string? password,

            int? port,

            int? probeCount,

            string? probePackets,

            int? probeTimeout,

            string? protocol,

            string? qualityMeasuredMethod,

            int? recoverytime,

            string? securityMode,

            string? server,

            int? slaFailLogPeriod,

            int? slaIdRedistribute,

            int? slaPassLogPeriod,

            ImmutableArray<Outputs.SystemSdwanHealthCheckSla> slas,

            string? source,

            string? systemDns,

            int? thresholdAlertJitter,

            int? thresholdAlertLatency,

            int? thresholdAlertPacketloss,

            int? thresholdWarningJitter,

            int? thresholdWarningLatency,

            int? thresholdWarningPacketloss,

            string? updateCascadeInterface,

            string? updateStaticRoute,

            string? user,

            int? vrf)
        {
            AddrMode = addrMode;
            DetectMode = detectMode;
            Diffservcode = diffservcode;
            DnsMatchIp = dnsMatchIp;
            DnsRequestDomain = dnsRequestDomain;
            EmbedMeasuredHealth = embedMeasuredHealth;
            Failtime = failtime;
            FtpFile = ftpFile;
            FtpMode = ftpMode;
            HaPriority = haPriority;
            HttpAgent = httpAgent;
            HttpGet = httpGet;
            HttpMatch = httpMatch;
            Interval = interval;
            Members = members;
            MosCodec = mosCodec;
            Name = name;
            PacketSize = packetSize;
            Password = password;
            Port = port;
            ProbeCount = probeCount;
            ProbePackets = probePackets;
            ProbeTimeout = probeTimeout;
            Protocol = protocol;
            QualityMeasuredMethod = qualityMeasuredMethod;
            Recoverytime = recoverytime;
            SecurityMode = securityMode;
            Server = server;
            SlaFailLogPeriod = slaFailLogPeriod;
            SlaIdRedistribute = slaIdRedistribute;
            SlaPassLogPeriod = slaPassLogPeriod;
            Slas = slas;
            Source = source;
            SystemDns = systemDns;
            ThresholdAlertJitter = thresholdAlertJitter;
            ThresholdAlertLatency = thresholdAlertLatency;
            ThresholdAlertPacketloss = thresholdAlertPacketloss;
            ThresholdWarningJitter = thresholdWarningJitter;
            ThresholdWarningLatency = thresholdWarningLatency;
            ThresholdWarningPacketloss = thresholdWarningPacketloss;
            UpdateCascadeInterface = updateCascadeInterface;
            UpdateStaticRoute = updateStaticRoute;
            User = user;
            Vrf = vrf;
        }
    }
}
