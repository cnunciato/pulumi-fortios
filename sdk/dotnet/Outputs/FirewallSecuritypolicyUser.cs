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
    public sealed class FirewallSecuritypolicyUser
    {
        /// <summary>
        /// User name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private FirewallSecuritypolicyUser(string? name)
        {
            Name = name;
        }
    }
}
