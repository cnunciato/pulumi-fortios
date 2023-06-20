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

    public sealed class RouterRipngAggregateAddressGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Aggregate address entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Aggregate address prefix.
        /// </summary>
        [Input("prefix6")]
        public Input<string>? Prefix6 { get; set; }

        public RouterRipngAggregateAddressGetArgs()
        {
        }
        public static new RouterRipngAggregateAddressGetArgs Empty => new RouterRipngAggregateAddressGetArgs();
    }
}
