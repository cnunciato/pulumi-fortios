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
    public sealed class FirewallPolicyUser
    {
        /// <summary>
        /// Names of individual users that can authenticate with this policy.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private FirewallPolicyUser(string? name)
        {
            Name = name;
        }
    }
}