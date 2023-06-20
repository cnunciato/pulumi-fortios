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
    public static class GetFirewallMulticastaddress6list
    {
        /// <summary>
        /// Provides a list of `fortios.FirewallMulticastaddress6`.
        /// </summary>
        public static Task<GetFirewallMulticastaddress6listResult> InvokeAsync(GetFirewallMulticastaddress6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallMulticastaddress6listResult>("fortios:index/getFirewallMulticastaddress6list:getFirewallMulticastaddress6list", args ?? new GetFirewallMulticastaddress6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.FirewallMulticastaddress6`.
        /// </summary>
        public static Output<GetFirewallMulticastaddress6listResult> Invoke(GetFirewallMulticastaddress6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallMulticastaddress6listResult>("fortios:index/getFirewallMulticastaddress6list:getFirewallMulticastaddress6list", args ?? new GetFirewallMulticastaddress6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallMulticastaddress6listArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallMulticastaddress6listArgs()
        {
        }
        public static new GetFirewallMulticastaddress6listArgs Empty => new GetFirewallMulticastaddress6listArgs();
    }

    public sealed class GetFirewallMulticastaddress6listInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetFirewallMulticastaddress6listInvokeArgs()
        {
        }
        public static new GetFirewallMulticastaddress6listInvokeArgs Empty => new GetFirewallMulticastaddress6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallMulticastaddress6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.FirewallMulticastaddress6`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallMulticastaddress6listResult(
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
