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
    public sealed class GetSystemSdnconnectorServerListResult
    {
        /// <summary>
        /// Configure IP configuration. The structure of `ip` block is documented below.
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private GetSystemSdnconnectorServerListResult(string ip)
        {
            Ip = ip;
        }
    }
}
