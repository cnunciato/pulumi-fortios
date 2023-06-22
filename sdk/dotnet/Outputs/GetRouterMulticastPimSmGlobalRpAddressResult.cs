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
    public sealed class GetRouterMulticastPimSmGlobalRpAddressResult
    {
        /// <summary>
        /// Groups to use this RP.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// RP router address.
        /// </summary>
        public readonly string IpAddress;

        [OutputConstructor]
        private GetRouterMulticastPimSmGlobalRpAddressResult(
            string group,

            int id,

            string ipAddress)
        {
            Group = group;
            Id = id;
            IpAddress = ipAddress;
        }
    }
}