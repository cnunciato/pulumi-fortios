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

    public sealed class SystemSdnconnectorExternalAccountListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AWS external ID.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("regionLists")]
        private InputList<Inputs.SystemSdnconnectorExternalAccountListRegionListArgs>? _regionLists;

        /// <summary>
        /// AWS region name list. The structure of `region_list` block is documented below.
        /// </summary>
        public InputList<Inputs.SystemSdnconnectorExternalAccountListRegionListArgs> RegionLists
        {
            get => _regionLists ?? (_regionLists = new InputList<Inputs.SystemSdnconnectorExternalAccountListRegionListArgs>());
            set => _regionLists = value;
        }

        /// <summary>
        /// AWS role ARN to assume.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public SystemSdnconnectorExternalAccountListArgs()
        {
        }
        public static new SystemSdnconnectorExternalAccountListArgs Empty => new SystemSdnconnectorExternalAccountListArgs();
    }
}
