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
    public static class GetSystemAutomationtriggerlist
    {
        /// <summary>
        /// Provides a list of `fortios.SystemAutomationtrigger`.
        /// </summary>
        public static Task<GetSystemAutomationtriggerlistResult> InvokeAsync(GetSystemAutomationtriggerlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemAutomationtriggerlistResult>("fortios:index/getSystemAutomationtriggerlist:getSystemAutomationtriggerlist", args ?? new GetSystemAutomationtriggerlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.SystemAutomationtrigger`.
        /// </summary>
        public static Output<GetSystemAutomationtriggerlistResult> Invoke(GetSystemAutomationtriggerlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemAutomationtriggerlistResult>("fortios:index/getSystemAutomationtriggerlist:getSystemAutomationtriggerlist", args ?? new GetSystemAutomationtriggerlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemAutomationtriggerlistArgs : global::Pulumi.InvokeArgs
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

        public GetSystemAutomationtriggerlistArgs()
        {
        }
        public static new GetSystemAutomationtriggerlistArgs Empty => new GetSystemAutomationtriggerlistArgs();
    }

    public sealed class GetSystemAutomationtriggerlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetSystemAutomationtriggerlistInvokeArgs()
        {
        }
        public static new GetSystemAutomationtriggerlistInvokeArgs Empty => new GetSystemAutomationtriggerlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemAutomationtriggerlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.SystemAutomationtrigger`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemAutomationtriggerlistResult(
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
