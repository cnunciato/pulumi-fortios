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

    public sealed class SwitchcontrollerManagedswitchIpSourceGuardGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("bindingEntries")]
        private InputList<Inputs.SwitchcontrollerManagedswitchIpSourceGuardBindingEntryGetArgs>? _bindingEntries;

        /// <summary>
        /// IP and MAC address configuration. The structure of `binding_entry` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerManagedswitchIpSourceGuardBindingEntryGetArgs> BindingEntries
        {
            get => _bindingEntries ?? (_bindingEntries = new InputList<Inputs.SwitchcontrollerManagedswitchIpSourceGuardBindingEntryGetArgs>());
            set => _bindingEntries = value;
        }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Ingress interface to which source guard is bound.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public SwitchcontrollerManagedswitchIpSourceGuardGetArgs()
        {
        }
        public static new SwitchcontrollerManagedswitchIpSourceGuardGetArgs Empty => new SwitchcontrollerManagedswitchIpSourceGuardGetArgs();
    }
}
