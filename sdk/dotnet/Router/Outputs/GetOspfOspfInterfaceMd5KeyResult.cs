// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router.Outputs
{

    [OutputType]
    public sealed class GetOspfOspfInterfaceMd5KeyResult
    {
        /// <summary>
        /// Distribute list entry ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Password for the key.
        /// </summary>
        public readonly string KeyString;

        [OutputConstructor]
        private GetOspfOspfInterfaceMd5KeyResult(
            int id,

            string keyString)
        {
            Id = id;
            KeyString = keyString;
        }
    }
}