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

    public sealed class SystemVxlanRemoteIp6GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv6 address.
        /// </summary>
        [Input("ip6")]
        public Input<string>? Ip6 { get; set; }

        public SystemVxlanRemoteIp6GetArgs()
        {
        }
        public static new SystemVxlanRemoteIp6GetArgs Empty => new SystemVxlanRemoteIp6GetArgs();
    }
}