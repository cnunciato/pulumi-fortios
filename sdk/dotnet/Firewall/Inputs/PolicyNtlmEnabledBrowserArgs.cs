// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class PolicyNtlmEnabledBrowserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User agent string.
        /// </summary>
        [Input("userAgentString")]
        public Input<string>? UserAgentString { get; set; }

        public PolicyNtlmEnabledBrowserArgs()
        {
        }
        public static new PolicyNtlmEnabledBrowserArgs Empty => new PolicyNtlmEnabledBrowserArgs();
    }
}