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

    public sealed class SwitchcontrollerLldpprofileCustomTlvGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Organizationally defined information string (0 - 507 hexadecimal bytes).
        /// </summary>
        [Input("informationString")]
        public Input<string>? InformationString { get; set; }

        /// <summary>
        /// TLV name (not sent).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.
        /// </summary>
        [Input("oui")]
        public Input<string>? Oui { get; set; }

        /// <summary>
        /// Organizationally defined subtype (0 - 255).
        /// </summary>
        [Input("subtype")]
        public Input<int>? Subtype { get; set; }

        public SwitchcontrollerLldpprofileCustomTlvGetArgs()
        {
        }
        public static new SwitchcontrollerLldpprofileCustomTlvGetArgs Empty => new SwitchcontrollerLldpprofileCustomTlvGetArgs();
    }
}
