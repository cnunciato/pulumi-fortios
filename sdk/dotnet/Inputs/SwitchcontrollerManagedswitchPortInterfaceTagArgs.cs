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

    public sealed class SwitchcontrollerManagedswitchPortInterfaceTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// FortiSwitch port tag name when exported to a virtual port pool or matched to dynamic port policy.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        public SwitchcontrollerManagedswitchPortInterfaceTagArgs()
        {
        }
        public static new SwitchcontrollerManagedswitchPortInterfaceTagArgs Empty => new SwitchcontrollerManagedswitchPortInterfaceTagArgs();
    }
}
