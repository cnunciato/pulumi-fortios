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

    public sealed class SystemSdnconnectorServerListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public SystemSdnconnectorServerListArgs()
        {
        }
        public static new SystemSdnconnectorServerListArgs Empty => new SystemSdnconnectorServerListArgs();
    }
}
