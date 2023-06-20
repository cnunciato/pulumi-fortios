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
    public sealed class UserExchangeKdcIp
    {
        /// <summary>
        /// KDC IPv4 addresses for Kerberos authentication.
        /// </summary>
        public readonly string? Ipv4;

        [OutputConstructor]
        private UserExchangeKdcIp(string? ipv4)
        {
            Ipv4 = ipv4;
        }
    }
}
