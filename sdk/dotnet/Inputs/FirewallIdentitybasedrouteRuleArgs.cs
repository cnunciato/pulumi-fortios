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

    public sealed class FirewallIdentitybasedrouteRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Outgoing interface for the rule.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// IPv4 address of the gateway (Format: xxx.xxx.xxx.xxx , Default: 0.0.0.0).
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("groups")]
        private InputList<Inputs.FirewallIdentitybasedrouteRuleGroupArgs>? _groups;

        /// <summary>
        /// Select one or more group(s) from available groups that are allowed to use this route. Separate group names with a space. The structure of `groups` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallIdentitybasedrouteRuleGroupArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.FirewallIdentitybasedrouteRuleGroupArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Rule ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        public FirewallIdentitybasedrouteRuleArgs()
        {
        }
        public static new FirewallIdentitybasedrouteRuleArgs Empty => new FirewallIdentitybasedrouteRuleArgs();
    }
}
