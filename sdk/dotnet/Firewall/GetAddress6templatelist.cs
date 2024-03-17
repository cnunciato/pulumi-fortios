// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    public static class GetAddress6templatelist
    {
        /// <summary>
        /// Provides a list of `fortios.firewall.Address6template`.
        /// </summary>
        public static Task<GetAddress6templatelistResult> InvokeAsync(GetAddress6templatelistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddress6templatelistResult>("fortios:firewall/getAddress6templatelist:getAddress6templatelist", args ?? new GetAddress6templatelistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewall.Address6template`.
        /// </summary>
        public static Output<GetAddress6templatelistResult> Invoke(GetAddress6templatelistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddress6templatelistResult>("fortios:firewall/getAddress6templatelist:getAddress6templatelist", args ?? new GetAddress6templatelistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddress6templatelistArgs : global::Pulumi.InvokeArgs
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

        public GetAddress6templatelistArgs()
        {
        }
        public static new GetAddress6templatelistArgs Empty => new GetAddress6templatelistArgs();
    }

    public sealed class GetAddress6templatelistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAddress6templatelistInvokeArgs()
        {
        }
        public static new GetAddress6templatelistInvokeArgs Empty => new GetAddress6templatelistInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddress6templatelistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewall.Address6template`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAddress6templatelistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}