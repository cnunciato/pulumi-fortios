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

    public sealed class SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiSwitch port tag name.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        public SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs()
        {
        }
        public static new SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs Empty => new SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs();
    }
}
