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

    public sealed class FirewallAddrgrp6MemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address6/addrgrp6 name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallAddrgrp6MemberArgs()
        {
        }
        public static new FirewallAddrgrp6MemberArgs Empty => new FirewallAddrgrp6MemberArgs();
    }
}
