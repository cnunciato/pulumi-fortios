// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class VpnsslwebHostchecksoftwareCheckItemList
    {
        /// <summary>
        /// Action. Valid values: `require`, `deny`.
        /// </summary>
        public readonly string? Action;
        /// <summary>
        /// Hex string of MD5 checksum.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// MD5 checksum. The structure of `md5s` block is documented below.
        /// 
        /// The `md5s` block supports:
        /// </summary>
        public readonly ImmutableArray<Outputs.VpnsslwebHostchecksoftwareCheckItemListMd5> Md5s;
        /// <summary>
        /// Target.
        /// </summary>
        public readonly string? Target;
        /// <summary>
        /// Type. Valid values: `file`, `registry`, `process`.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private VpnsslwebHostchecksoftwareCheckItemList(
            string? action,

            int? id,

            ImmutableArray<Outputs.VpnsslwebHostchecksoftwareCheckItemListMd5> md5s,

            string? target,

            string? type,

            string? version)
        {
            Action = action;
            Id = id;
            Md5s = md5s;
            Target = target;
            Type = type;
            Version = version;
        }
    }
}