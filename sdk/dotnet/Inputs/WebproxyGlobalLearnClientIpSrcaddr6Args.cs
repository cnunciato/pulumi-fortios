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

    public sealed class WebproxyGlobalLearnClientIpSrcaddr6Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address name.
        /// 
        /// The `learn_client_ip_srcaddr6` block supports:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WebproxyGlobalLearnClientIpSrcaddr6Args()
        {
        }
        public static new WebproxyGlobalLearnClientIpSrcaddr6Args Empty => new WebproxyGlobalLearnClientIpSrcaddr6Args();
    }
}
