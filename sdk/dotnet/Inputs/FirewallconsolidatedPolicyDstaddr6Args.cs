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

    public sealed class FirewallconsolidatedPolicyDstaddr6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallconsolidatedPolicyDstaddr6Args()
        {
        }
        public static new FirewallconsolidatedPolicyDstaddr6Args Empty => new FirewallconsolidatedPolicyDstaddr6Args();
    }
}
