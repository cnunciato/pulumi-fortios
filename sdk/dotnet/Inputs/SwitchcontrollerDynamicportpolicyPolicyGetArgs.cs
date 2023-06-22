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

    public sealed class SwitchcontrollerDynamicportpolicyPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("bouncePortLink")]
        public Input<string>? BouncePortLink { get; set; }

        /// <summary>
        /// Category of Dynamic port policy. Valid values: `device`, `interface-tag`.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// Description for the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Policy matching family.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// Policy matching host.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Match policy based on hardware vendor.
        /// </summary>
        [Input("hwVendor")]
        public Input<string>? HwVendor { get; set; }

        [Input("interfaceTags")]
        private InputList<Inputs.SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs>? _interfaceTags;

        /// <summary>
        /// Policy matching the FortiSwitch interface object tags. The structure of `interface_tags` block is documented below.
        /// </summary>
        public InputList<Inputs.SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs> InterfaceTags
        {
            get => _interfaceTags ?? (_interfaceTags = new InputList<Inputs.SwitchcontrollerDynamicportpolicyPolicyInterfaceTagGetArgs>());
            set => _interfaceTags = value;
        }

        /// <summary>
        /// LLDP profile to be applied when using this policy.
        /// </summary>
        [Input("lldpProfile")]
        public Input<string>? LldpProfile { get; set; }

        /// <summary>
        /// Policy matching MAC address.
        /// </summary>
        [Input("mac")]
        public Input<string>? Mac { get; set; }

        /// <summary>
        /// 802.1x security policy to be applied when using this policy.
        /// </summary>
        [Input("n8021x")]
        public Input<string>? N8021x { get; set; }

        /// <summary>
        /// Policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// QoS policy to be applied when using this policy.
        /// </summary>
        [Input("qosPolicy")]
        public Input<string>? QosPolicy { get; set; }

        /// <summary>
        /// Enable/disable policy. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Policy matching type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// VLAN policy to be applied when using this policy.
        /// </summary>
        [Input("vlanPolicy")]
        public Input<string>? VlanPolicy { get; set; }

        public SwitchcontrollerDynamicportpolicyPolicyGetArgs()
        {
        }
        public static new SwitchcontrollerDynamicportpolicyPolicyGetArgs Empty => new SwitchcontrollerDynamicportpolicyPolicyGetArgs();
    }
}