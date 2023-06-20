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

    public sealed class SwitchcontrollerGlobalCustomCommandGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of FortiSwitch commands.
        /// </summary>
        [Input("commandEntry")]
        public Input<string>? CommandEntry { get; set; }

        /// <summary>
        /// Name of custom command to push to all FortiSwitches in VDOM.
        /// </summary>
        [Input("commandName")]
        public Input<string>? CommandName { get; set; }

        public SwitchcontrollerGlobalCustomCommandGetArgs()
        {
        }
        public static new SwitchcontrollerGlobalCustomCommandGetArgs Empty => new SwitchcontrollerGlobalCustomCommandGetArgs();
    }
}
