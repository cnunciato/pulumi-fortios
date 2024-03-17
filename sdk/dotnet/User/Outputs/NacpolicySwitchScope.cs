// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.User.Outputs
{

    [OutputType]
    public sealed class NacpolicySwitchScope
    {
        /// <summary>
        /// Managed FortiSwitch name from available options.
        /// </summary>
        public readonly string? SwitchId;

        [OutputConstructor]
        private NacpolicySwitchScope(string? switchId)
        {
            SwitchId = switchId;
        }
    }
}