// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System.Snmp
{
    public static class GetCommunitylist
    {
        /// <summary>
        /// Provides a list of `fortios.system/snmp.Community`.
        /// </summary>
        public static Task<GetCommunitylistResult> InvokeAsync(GetCommunitylistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCommunitylistResult>("fortios:system/snmp/getCommunitylist:getCommunitylist", args ?? new GetCommunitylistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.system/snmp.Community`.
        /// </summary>
        public static Output<GetCommunitylistResult> Invoke(GetCommunitylistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCommunitylistResult>("fortios:system/snmp/getCommunitylist:getCommunitylist", args ?? new GetCommunitylistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCommunitylistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetCommunitylistArgs()
        {
        }
        public static new GetCommunitylistArgs Empty => new GetCommunitylistArgs();
    }

    public sealed class GetCommunitylistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetCommunitylistInvokeArgs()
        {
        }
        public static new GetCommunitylistInvokeArgs Empty => new GetCommunitylistInvokeArgs();
    }


    [OutputType]
    public sealed class GetCommunitylistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.system/snmp.Community`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetCommunitylistResult(
            string? filter,

            ImmutableArray<int> fosidlists,

            string id,

            string? vdomparam)
        {
            Filter = filter;
            Fosidlists = fosidlists;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}