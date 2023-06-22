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

    public sealed class IpsSensorOverrideGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action of override rule. Valid values: `pass`, `block`, `reset`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        [Input("exemptIps")]
        private InputList<Inputs.IpsSensorOverrideExemptIpGetArgs>? _exemptIps;

        /// <summary>
        /// Exempted IP. The structure of `exempt_ip` block is documented below.
        /// </summary>
        public InputList<Inputs.IpsSensorOverrideExemptIpGetArgs> ExemptIps
        {
            get => _exemptIps ?? (_exemptIps = new InputList<Inputs.IpsSensorOverrideExemptIpGetArgs>());
            set => _exemptIps = value;
        }

        /// <summary>
        /// Enable/disable logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("log")]
        public Input<string>? Log { get; set; }

        /// <summary>
        /// Enable/disable packet logging. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("logPacket")]
        public Input<string>? LogPacket { get; set; }

        /// <summary>
        /// Quarantine IP or interface. Valid values: `none`, `attacker`.
        /// </summary>
        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        /// <summary>
        /// Duration of quarantine in minute.
        /// </summary>
        [Input("quarantineExpiry")]
        public Input<int>? QuarantineExpiry { get; set; }

        /// <summary>
        /// Enable/disable logging of selected quarantine. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("quarantineLog")]
        public Input<string>? QuarantineLog { get; set; }

        /// <summary>
        /// Override rule ID.
        /// </summary>
        [Input("ruleId")]
        public Input<int>? RuleId { get; set; }

        /// <summary>
        /// Enable/disable status of override rule. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IpsSensorOverrideGetArgs()
        {
        }
        public static new IpsSensorOverrideGetArgs Empty => new IpsSensorOverrideGetArgs();
    }
}