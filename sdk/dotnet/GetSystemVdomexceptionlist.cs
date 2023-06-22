// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetSystemVdomexceptionlist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemVdomexception`.
        /// </summary>
        public static Task<GetSystemVdomexceptionlistResult> InvokeAsync(GetSystemVdomexceptionlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemVdomexceptionlistResult>("fortios:index/getSystemVdomexceptionlist:getSystemVdomexceptionlist", args ?? new GetSystemVdomexceptionlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemVdomexception`.
        /// </summary>
        public static Output<GetSystemVdomexceptionlistResult> Invoke(GetSystemVdomexceptionlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemVdomexceptionlistResult>("fortios:index/getSystemVdomexceptionlist:getSystemVdomexceptionlist", args ?? new GetSystemVdomexceptionlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemVdomexceptionlistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemVdomexceptionlistArgs()
        {
        }
        public static new GetSystemVdomexceptionlistArgs Empty => new GetSystemVdomexceptionlistArgs();
    }

    public sealed class GetSystemVdomexceptionlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemVdomexceptionlistInvokeArgs()
        {
        }
        public static new GetSystemVdomexceptionlistInvokeArgs Empty => new GetSystemVdomexceptionlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemVdomexceptionlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.SystemVdomexception`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemVdomexceptionlistResult(
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