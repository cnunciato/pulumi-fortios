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

    public sealed class RouterPolicySrcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP and mask.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public RouterPolicySrcGetArgs()
        {
        }
        public static new RouterPolicySrcGetArgs Empty => new RouterPolicySrcGetArgs();
    }
}
