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

    public sealed class FirewallPolicySgtGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Security group tag.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FirewallPolicySgtGetArgs()
        {
        }
        public static new FirewallPolicySgtGetArgs Empty => new FirewallPolicySgtGetArgs();
    }
}