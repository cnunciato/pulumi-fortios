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
    public sealed class FirewallInternetserviceadditionEntryPortRange
    {
        /// <summary>
        /// Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).
        /// </summary>
        public readonly int? EndPort;
        /// <summary>
        /// Custom entry port range ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
        /// </summary>
        public readonly int? StartPort;

        [OutputConstructor]
        private FirewallInternetserviceadditionEntryPortRange(
            int? endPort,

            int? id,

            int? startPort)
        {
            EndPort = endPort;
            Id = id;
            StartPort = startPort;
        }
    }
}
