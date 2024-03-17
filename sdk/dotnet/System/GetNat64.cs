// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.System
{
    public static class GetNat64
    {
        /// <summary>
        /// Use this data source to get information on fortios system nat64
        /// </summary>
        public static Task<GetNat64Result> InvokeAsync(GetNat64Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNat64Result>("fortios:system/getNat64:getNat64", args ?? new GetNat64Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system nat64
        /// </summary>
        public static Output<GetNat64Result> Invoke(GetNat64InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNat64Result>("fortios:system/getNat64:getNat64", args ?? new GetNat64InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNat64Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetNat64Args()
        {
        }
        public static new GetNat64Args Empty => new GetNat64Args();
    }

    public sealed class GetNat64InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetNat64InvokeArgs()
        {
        }
        public static new GetNat64InvokeArgs Empty => new GetNat64InvokeArgs();
    }


    [OutputType]
    public sealed class GetNat64Result
    {
        /// <summary>
        /// Enable/disable AAAA record synthesis (default = enable).
        /// </summary>
        public readonly string AlwaysSynthesizeAaaaRecord;
        /// <summary>
        /// Enable/disable IPv6 fragment header generation.
        /// </summary>
        public readonly string GenerateIpv6FragmentHeader;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable mandatory IPv4 packet forwarding in nat46.
        /// </summary>
        public readonly string Nat46ForceIpv4PacketForwarding;
        /// <summary>
        /// NAT64 prefix.
        /// </summary>
        public readonly string Nat64Prefix;
        /// <summary>
        /// Enable/disable secondary NAT64 prefix.
        /// </summary>
        public readonly string SecondaryPrefixStatus;
        /// <summary>
        /// Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNat64SecondaryPrefixResult> SecondaryPrefixes;
        /// <summary>
        /// Enable/disable NAT64 (default = disable).
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetNat64Result(
            string alwaysSynthesizeAaaaRecord,

            string generateIpv6FragmentHeader,

            string id,

            string nat46ForceIpv4PacketForwarding,

            string nat64Prefix,

            string secondaryPrefixStatus,

            ImmutableArray<Outputs.GetNat64SecondaryPrefixResult> secondaryPrefixes,

            string status,

            string? vdomparam)
        {
            AlwaysSynthesizeAaaaRecord = alwaysSynthesizeAaaaRecord;
            GenerateIpv6FragmentHeader = generateIpv6FragmentHeader;
            Id = id;
            Nat46ForceIpv4PacketForwarding = nat46ForceIpv4PacketForwarding;
            Nat64Prefix = nat64Prefix;
            SecondaryPrefixStatus = secondaryPrefixStatus;
            SecondaryPrefixes = secondaryPrefixes;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}