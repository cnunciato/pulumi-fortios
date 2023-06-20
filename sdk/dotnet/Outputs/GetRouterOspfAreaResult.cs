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
    public sealed class GetRouterOspfAreaResult
    {
        /// <summary>
        /// Authentication type.
        /// </summary>
        public readonly string Authentication;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Summary default cost of stub or NSSA area.
        /// </summary>
        public readonly int DefaultCost;
        /// <summary>
        /// OSPF area filter-list configuration. The structure of `filter_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspfAreaFilterListResult> FilterLists;
        /// <summary>
        /// Distribute list entry ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Redistribute, advertise, or do not originate Type-7 default route into NSSA area.
        /// </summary>
        public readonly string NssaDefaultInformationOriginate;
        /// <summary>
        /// OSPF default metric.
        /// </summary>
        public readonly int NssaDefaultInformationOriginateMetric;
        /// <summary>
        /// OSPF metric type for default routes.
        /// </summary>
        public readonly string NssaDefaultInformationOriginateMetricType;
        /// <summary>
        /// Enable/disable redistribute into NSSA area.
        /// </summary>
        public readonly string NssaRedistribution;
        /// <summary>
        /// NSSA translator role type.
        /// </summary>
        public readonly string NssaTranslatorRole;
        /// <summary>
        /// OSPF area range configuration. The structure of `range` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspfAreaRangeResult> Ranges;
        /// <summary>
        /// Enable/disable shortcut option.
        /// </summary>
        public readonly string Shortcut;
        /// <summary>
        /// Stub summary setting.
        /// </summary>
        public readonly string StubType;
        /// <summary>
        /// Area type setting.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// OSPF virtual link configuration. The structure of `virtual_link` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspfAreaVirtualLinkResult> VirtualLinks;

        [OutputConstructor]
        private GetRouterOspfAreaResult(
            string authentication,

            string comments,

            int defaultCost,

            ImmutableArray<Outputs.GetRouterOspfAreaFilterListResult> filterLists,

            string id,

            string nssaDefaultInformationOriginate,

            int nssaDefaultInformationOriginateMetric,

            string nssaDefaultInformationOriginateMetricType,

            string nssaRedistribution,

            string nssaTranslatorRole,

            ImmutableArray<Outputs.GetRouterOspfAreaRangeResult> ranges,

            string shortcut,

            string stubType,

            string type,

            ImmutableArray<Outputs.GetRouterOspfAreaVirtualLinkResult> virtualLinks)
        {
            Authentication = authentication;
            Comments = comments;
            DefaultCost = defaultCost;
            FilterLists = filterLists;
            Id = id;
            NssaDefaultInformationOriginate = nssaDefaultInformationOriginate;
            NssaDefaultInformationOriginateMetric = nssaDefaultInformationOriginateMetric;
            NssaDefaultInformationOriginateMetricType = nssaDefaultInformationOriginateMetricType;
            NssaRedistribution = nssaRedistribution;
            NssaTranslatorRole = nssaTranslatorRole;
            Ranges = ranges;
            Shortcut = shortcut;
            StubType = stubType;
            Type = type;
            VirtualLinks = virtualLinks;
        }
    }
}
