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
    public sealed class GetSystemSdnconnectorGcpProjectListResult
    {
        /// <summary>
        /// Configure GCP zone list. The structure of `gcp_zone_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorGcpProjectListGcpZoneListResult> GcpZoneLists;
        /// <summary>
        /// GCP project ID.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSystemSdnconnectorGcpProjectListResult(
            ImmutableArray<Outputs.GetSystemSdnconnectorGcpProjectListGcpZoneListResult> gcpZoneLists,

            string id)
        {
            GcpZoneLists = gcpZoneLists;
            Id = id;
        }
    }
}