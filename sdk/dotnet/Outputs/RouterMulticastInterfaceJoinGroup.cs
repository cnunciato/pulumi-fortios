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
    public sealed class RouterMulticastInterfaceJoinGroup
    {
        /// <summary>
        /// Multicast group IP address.
        /// </summary>
        public readonly string? Address;

        [OutputConstructor]
        private RouterMulticastInterfaceJoinGroup(string? address)
        {
            Address = address;
        }
    }
}
