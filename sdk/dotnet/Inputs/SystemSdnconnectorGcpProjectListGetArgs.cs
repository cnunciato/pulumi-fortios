// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemSdnconnectorGcpProjectListGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("gcpZoneLists")]
        private InputList<Inputs.SystemSdnconnectorGcpProjectListGcpZoneListGetArgs>? _gcpZoneLists;

        /// <summary>
        /// Configure GCP zone list. The structure of `gcp_zone_list` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdnconnectorGcpProjectListGcpZoneListGetArgs> GcpZoneLists
        {
            get => _gcpZoneLists ?? (_gcpZoneLists = new InputList<Inputs.SystemSdnconnectorGcpProjectListGcpZoneListGetArgs>());
            set => _gcpZoneLists = value;
        }

        /// <summary>
        /// GCP project ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public SystemSdnconnectorGcpProjectListGetArgs()
        {
        }
        public static new SystemSdnconnectorGcpProjectListGetArgs Empty => new SystemSdnconnectorGcpProjectListGetArgs();
    }
}
