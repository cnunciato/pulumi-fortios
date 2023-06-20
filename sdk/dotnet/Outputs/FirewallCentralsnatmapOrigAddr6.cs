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
    public sealed class FirewallCentralsnatmapOrigAddr6
    {
        /// <summary>
        /// Address name.
        /// 
        /// The `orig_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `dst_addr6` block supports:
        /// 
        /// 
        /// 
        /// 
        /// The `nat_ippool6` block supports:
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private FirewallCentralsnatmapOrigAddr6(string? name)
        {
            Name = name;
        }
    }
}
