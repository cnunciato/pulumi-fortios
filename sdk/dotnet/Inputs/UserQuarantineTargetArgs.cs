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

    public sealed class UserQuarantineTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the quarantine entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Quarantine entry name.
        /// </summary>
        [Input("entry")]
        public Input<string>? Entry { get; set; }

        [Input("macs")]
        private InputList<Inputs.UserQuarantineTargetMacArgs>? _macs;

        /// <summary>
        /// Quarantine MACs. The structure of `macs` block is documented below.
        /// </summary>
        public InputList<Inputs.UserQuarantineTargetMacArgs> Macs
        {
            get => _macs ?? (_macs = new InputList<Inputs.UserQuarantineTargetMacArgs>());
            set => _macs = value;
        }

        public UserQuarantineTargetArgs()
        {
        }
        public static new UserQuarantineTargetArgs Empty => new UserQuarantineTargetArgs();
    }
}
