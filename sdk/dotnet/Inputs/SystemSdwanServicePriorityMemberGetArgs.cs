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

    public sealed class SystemSdwanServicePriorityMemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Member sequence number.
        /// </summary>
        [Input("seqNum")]
        public Input<int>? SeqNum { get; set; }

        public SystemSdwanServicePriorityMemberGetArgs()
        {
        }
        public static new SystemSdwanServicePriorityMemberGetArgs Empty => new SystemSdwanServicePriorityMemberGetArgs();
    }
}
