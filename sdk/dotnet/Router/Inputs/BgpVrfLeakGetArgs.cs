// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Inputs
{

    public sealed class BgpVrfLeakGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("targets")]
        private InputList<Inputs.BgpVrfLeakTargetGetArgs>? _targets;

        /// <summary>
        /// Target VRF table. The structure of `target` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpVrfLeakTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.BgpVrfLeakTargetGetArgs>());
            set => _targets = value;
        }

        /// <summary>
        /// Origin VRF ID &lt;0 - 31&gt;.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpVrfLeakGetArgs()
        {
        }
        public static new BgpVrfLeakGetArgs Empty => new BgpVrfLeakGetArgs();
    }
}