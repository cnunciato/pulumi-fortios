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

    public sealed class CommunitylistRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permit or deny route-based operations, based on the route's COMMUNITY attribute. Valid values: `deny`, `permit`.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Community specifications for matching a reserved community.
        /// </summary>
        [Input("match")]
        public Input<string>? Match { get; set; }

        /// <summary>
        /// Ordered list of COMMUNITY attributes as a regular expression.
        /// </summary>
        [Input("regexp")]
        public Input<string>? Regexp { get; set; }

        public CommunitylistRuleGetArgs()
        {
        }
        public static new CommunitylistRuleGetArgs Empty => new CommunitylistRuleGetArgs();
    }
}