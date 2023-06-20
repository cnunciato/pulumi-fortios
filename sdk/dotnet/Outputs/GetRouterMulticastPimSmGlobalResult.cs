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
    public sealed class GetRouterMulticastPimSmGlobalResult
    {
        /// <summary>
        /// Sources allowed to register packets with this Rendezvous Point (RP).
        /// </summary>
        public readonly string AcceptRegisterList;
        /// <summary>
        /// Sources allowed to send multicast traffic.
        /// </summary>
        public readonly string AcceptSourceList;
        /// <summary>
        /// Enable/disable accept BSR quick refresh packets from neighbors.
        /// </summary>
        public readonly string BsrAllowQuickRefresh;
        /// <summary>
        /// Enable/disable allowing this router to become a bootstrap router (BSR).
        /// </summary>
        public readonly string BsrCandidate;
        /// <summary>
        /// BSR hash length (0 - 32, default = 10).
        /// </summary>
        public readonly int BsrHash;
        /// <summary>
        /// Interface to advertise as candidate BSR.
        /// </summary>
        public readonly string BsrInterface;
        /// <summary>
        /// BSR priority (0 - 255, default = 0).
        /// </summary>
        public readonly int BsrPriority;
        /// <summary>
        /// Enable/disable making candidate RP compatible with old Cisco IOS.
        /// </summary>
        public readonly string CiscoCrpPrefix;
        /// <summary>
        /// Use only hash for RP selection (compatibility with old Cisco IOS).
        /// </summary>
        public readonly string CiscoIgnoreRpSetPriority;
        /// <summary>
        /// Checksum entire register packet(for old Cisco IOS compatibility).
        /// </summary>
        public readonly string CiscoRegisterChecksum;
        /// <summary>
        /// Cisco register checksum only these groups.
        /// </summary>
        public readonly string CiscoRegisterChecksumGroup;
        /// <summary>
        /// Join/prune holdtime (1 - 65535, default = 210).
        /// </summary>
        public readonly int JoinPruneHoldtime;
        /// <summary>
        /// Period of time between sending periodic PIM join/prune messages in seconds (1 - 65535, default = 60).
        /// </summary>
        public readonly int MessageInterval;
        /// <summary>
        /// Maximum retries of null register (1 - 20, default = 1).
        /// </summary>
        public readonly int NullRegisterRetries;
        /// <summary>
        /// Limit of packets/sec per source registered through this RP (0 - 65535, default = 0 which means unlimited).
        /// </summary>
        public readonly int RegisterRateLimit;
        /// <summary>
        /// Enable/disable check RP is reachable before registering packets.
        /// </summary>
        public readonly string RegisterRpReachability;
        /// <summary>
        /// Override source address in register packets.
        /// </summary>
        public readonly string RegisterSource;
        /// <summary>
        /// Override with primary interface address.
        /// </summary>
        public readonly string RegisterSourceInterface;
        /// <summary>
        /// Override with local IP address.
        /// </summary>
        public readonly string RegisterSourceIp;
        /// <summary>
        /// Period of time to honor register-stop message (1 - 65535 sec, default = 60).
        /// </summary>
        public readonly int RegisterSupression;
        /// <summary>
        /// Statically configure RP addresses. The structure of `rp_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterMulticastPimSmGlobalRpAddressResult> RpAddresses;
        /// <summary>
        /// Timeout for RP receiving data on (S,G) tree (1 - 65535 sec, default = 185).
        /// </summary>
        public readonly int RpRegisterKeepalive;
        /// <summary>
        /// Enable/disable switching to source specific trees.
        /// </summary>
        public readonly string SptThreshold;
        /// <summary>
        /// Groups allowed to switch to source tree.
        /// </summary>
        public readonly string SptThresholdGroup;
        /// <summary>
        /// Enable/disable source specific multicast.
        /// </summary>
        public readonly string Ssm;
        /// <summary>
        /// Groups allowed to source specific multicast.
        /// </summary>
        public readonly string SsmRange;

        [OutputConstructor]
        private GetRouterMulticastPimSmGlobalResult(
            string acceptRegisterList,

            string acceptSourceList,

            string bsrAllowQuickRefresh,

            string bsrCandidate,

            int bsrHash,

            string bsrInterface,

            int bsrPriority,

            string ciscoCrpPrefix,

            string ciscoIgnoreRpSetPriority,

            string ciscoRegisterChecksum,

            string ciscoRegisterChecksumGroup,

            int joinPruneHoldtime,

            int messageInterval,

            int nullRegisterRetries,

            int registerRateLimit,

            string registerRpReachability,

            string registerSource,

            string registerSourceInterface,

            string registerSourceIp,

            int registerSupression,

            ImmutableArray<Outputs.GetRouterMulticastPimSmGlobalRpAddressResult> rpAddresses,

            int rpRegisterKeepalive,

            string sptThreshold,

            string sptThresholdGroup,

            string ssm,

            string ssmRange)
        {
            AcceptRegisterList = acceptRegisterList;
            AcceptSourceList = acceptSourceList;
            BsrAllowQuickRefresh = bsrAllowQuickRefresh;
            BsrCandidate = bsrCandidate;
            BsrHash = bsrHash;
            BsrInterface = bsrInterface;
            BsrPriority = bsrPriority;
            CiscoCrpPrefix = ciscoCrpPrefix;
            CiscoIgnoreRpSetPriority = ciscoIgnoreRpSetPriority;
            CiscoRegisterChecksum = ciscoRegisterChecksum;
            CiscoRegisterChecksumGroup = ciscoRegisterChecksumGroup;
            JoinPruneHoldtime = joinPruneHoldtime;
            MessageInterval = messageInterval;
            NullRegisterRetries = nullRegisterRetries;
            RegisterRateLimit = registerRateLimit;
            RegisterRpReachability = registerRpReachability;
            RegisterSource = registerSource;
            RegisterSourceInterface = registerSourceInterface;
            RegisterSourceIp = registerSourceIp;
            RegisterSupression = registerSupression;
            RpAddresses = rpAddresses;
            RpRegisterKeepalive = rpRegisterKeepalive;
            SptThreshold = sptThreshold;
            SptThresholdGroup = sptThresholdGroup;
            Ssm = ssm;
            SsmRange = ssmRange;
        }
    }
}
