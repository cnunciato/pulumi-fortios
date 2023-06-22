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
    public sealed class RouterBgpNetwork6
    {
        /// <summary>
        /// Enable/disable route as backdoor. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? Backdoor;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Enable/disable ensure BGP network route exists in IGP. Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? NetworkImportCheck;
        /// <summary>
        /// Aggregate IPv6 prefix.
        /// </summary>
        public readonly string? Prefix6;
        /// <summary>
        /// Route map to modify generated route.
        /// 
        /// The `network6` block supports:
        /// 
        /// 
        /// 
        /// The `redistribute6` block supports:
        /// </summary>
        public readonly string? RouteMap;

        [OutputConstructor]
        private RouterBgpNetwork6(
            string? backdoor,

            int? id,

            string? networkImportCheck,

            string? prefix6,

            string? routeMap)
        {
            Backdoor = backdoor;
            Id = id;
            NetworkImportCheck = networkImportCheck;
            Prefix6 = prefix6;
            RouteMap = routeMap;
        }
    }
}