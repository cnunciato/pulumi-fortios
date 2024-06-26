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

    public sealed class BgpNeighborArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable address family IPv4 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activate")]
        public Input<string>? Activate { get; set; }

        /// <summary>
        /// Enable/disable address family IPv6 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activate6")]
        public Input<string>? Activate6 { get; set; }

        /// <summary>
        /// Enable/disable address family L2VPN EVPN for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activateEvpn")]
        public Input<string>? ActivateEvpn { get; set; }

        /// <summary>
        /// Enable/disable address family VPNv4 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activateVpnv4")]
        public Input<string>? ActivateVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable address family VPNv6 for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("activateVpnv6")]
        public Input<string>? ActivateVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable IPv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        [Input("additionalPath")]
        public Input<string>? AdditionalPath { get; set; }

        /// <summary>
        /// Enable/disable IPv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        [Input("additionalPath6")]
        public Input<string>? AdditionalPath6 { get; set; }

        /// <summary>
        /// Enable/disable VPNv4 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        [Input("additionalPathVpnv4")]
        public Input<string>? AdditionalPathVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable VPNv6 additional-path capability. Valid values: `send`, `receive`, `both`, `disable`.
        /// </summary>
        [Input("additionalPathVpnv6")]
        public Input<string>? AdditionalPathVpnv6 { get; set; }

        /// <summary>
        /// Number of IPv4 additional paths that can be advertised to this neighbor.
        /// </summary>
        [Input("advAdditionalPath")]
        public Input<int>? AdvAdditionalPath { get; set; }

        /// <summary>
        /// Number of IPv6 additional paths that can be advertised to this neighbor.
        /// </summary>
        [Input("advAdditionalPath6")]
        public Input<int>? AdvAdditionalPath6 { get; set; }

        /// <summary>
        /// Number of VPNv4 additional paths that can be advertised to this neighbor.
        /// </summary>
        [Input("advAdditionalPathVpnv4")]
        public Input<int>? AdvAdditionalPathVpnv4 { get; set; }

        /// <summary>
        /// Number of VPNv6 additional paths that can be advertised to this neighbor.
        /// </summary>
        [Input("advAdditionalPathVpnv6")]
        public Input<int>? AdvAdditionalPathVpnv6 { get; set; }

        /// <summary>
        /// Minimum interval (sec) between sending updates.
        /// </summary>
        [Input("advertisementInterval")]
        public Input<int>? AdvertisementInterval { get; set; }

        /// <summary>
        /// IPv4 The maximum number of occurrence of my AS number allowed.
        /// </summary>
        [Input("allowasIn")]
        public Input<int>? AllowasIn { get; set; }

        /// <summary>
        /// IPv6 The maximum number of occurrence of my AS number allowed.
        /// </summary>
        [Input("allowasIn6")]
        public Input<int>? AllowasIn6 { get; set; }

        /// <summary>
        /// Enable/disable IPv4 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowasInEnable")]
        public Input<string>? AllowasInEnable { get; set; }

        /// <summary>
        /// Enable/disable IPv6 Enable to allow my AS in AS path. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowasInEnable6")]
        public Input<string>? AllowasInEnable6 { get; set; }

        /// <summary>
        /// Enable/disable to allow my AS in AS path for L2VPN EVPN route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowasInEnableEvpn")]
        public Input<string>? AllowasInEnableEvpn { get; set; }

        /// <summary>
        /// Enable/disable to allow my AS in AS path for VPNv4 route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowasInEnableVpnv4")]
        public Input<string>? AllowasInEnableVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable use of my AS in AS path for VPNv6 route. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("allowasInEnableVpnv6")]
        public Input<string>? AllowasInEnableVpnv6 { get; set; }

        /// <summary>
        /// The maximum number of occurrence of my AS number allowed for L2VPN EVPN route.
        /// </summary>
        [Input("allowasInEvpn")]
        public Input<int>? AllowasInEvpn { get; set; }

        /// <summary>
        /// The maximum number of occurrence of my AS number allowed for VPNv4 route.
        /// </summary>
        [Input("allowasInVpnv4")]
        public Input<int>? AllowasInVpnv4 { get; set; }

        /// <summary>
        /// The maximum number of occurrence of my AS number allowed for VPNv6 route.
        /// </summary>
        [Input("allowasInVpnv6")]
        public Input<int>? AllowasInVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable replace peer AS with own AS for IPv4. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("asOverride")]
        public Input<string>? AsOverride { get; set; }

        /// <summary>
        /// Enable/disable replace peer AS with own AS for IPv6. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("asOverride6")]
        public Input<string>? AsOverride6 { get; set; }

        /// <summary>
        /// IPv4 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        [Input("attributeUnchanged")]
        public Input<string>? AttributeUnchanged { get; set; }

        /// <summary>
        /// IPv6 List of attributes that should be unchanged. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        [Input("attributeUnchanged6")]
        public Input<string>? AttributeUnchanged6 { get; set; }

        /// <summary>
        /// List of attributes that should be unchanged for VPNv4 route. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        [Input("attributeUnchangedVpnv4")]
        public Input<string>? AttributeUnchangedVpnv4 { get; set; }

        /// <summary>
        /// List of attributes that should not be changed for VPNv6 route. Valid values: `as-path`, `med`, `next-hop`.
        /// </summary>
        [Input("attributeUnchangedVpnv6")]
        public Input<string>? AttributeUnchangedVpnv6 { get; set; }

        /// <summary>
        /// Key-chain name for TCP authentication options.
        /// </summary>
        [Input("authOptions")]
        public Input<string>? AuthOptions { get; set; }

        /// <summary>
        /// Enable/disable BFD for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("bfd")]
        public Input<string>? Bfd { get; set; }

        /// <summary>
        /// Enable/disable advertise default IPv4 route to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityDefaultOriginate")]
        public Input<string>? CapabilityDefaultOriginate { get; set; }

        /// <summary>
        /// Enable/disable advertise default IPv6 route to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityDefaultOriginate6")]
        public Input<string>? CapabilityDefaultOriginate6 { get; set; }

        /// <summary>
        /// Enable/disable advertise dynamic capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityDynamic")]
        public Input<string>? CapabilityDynamic { get; set; }

        /// <summary>
        /// Enable/disable advertise IPv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityGracefulRestart")]
        public Input<string>? CapabilityGracefulRestart { get; set; }

        /// <summary>
        /// Enable/disable advertise IPv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityGracefulRestart6")]
        public Input<string>? CapabilityGracefulRestart6 { get; set; }

        /// <summary>
        /// Enable/disable advertisement of L2VPN EVPN graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityGracefulRestartEvpn")]
        public Input<string>? CapabilityGracefulRestartEvpn { get; set; }

        /// <summary>
        /// Enable/disable advertise VPNv4 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityGracefulRestartVpnv4")]
        public Input<string>? CapabilityGracefulRestartVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable advertisement of VPNv6 graceful restart capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityGracefulRestartVpnv6")]
        public Input<string>? CapabilityGracefulRestartVpnv6 { get; set; }

        /// <summary>
        /// Accept/Send IPv4 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
        /// </summary>
        [Input("capabilityOrf")]
        public Input<string>? CapabilityOrf { get; set; }

        /// <summary>
        /// Accept/Send IPv6 ORF lists to/from this neighbor. Valid values: `none`, `receive`, `send`, `both`.
        /// </summary>
        [Input("capabilityOrf6")]
        public Input<string>? CapabilityOrf6 { get; set; }

        /// <summary>
        /// Enable/disable advertise route refresh capability to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("capabilityRouteRefresh")]
        public Input<string>? CapabilityRouteRefresh { get; set; }

        [Input("conditionalAdvertise6s")]
        private InputList<Inputs.BgpNeighborConditionalAdvertise6Args>? _conditionalAdvertise6s;

        /// <summary>
        /// IPv6 conditional advertisement. The structure of `conditional_advertise6` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpNeighborConditionalAdvertise6Args> ConditionalAdvertise6s
        {
            get => _conditionalAdvertise6s ?? (_conditionalAdvertise6s = new InputList<Inputs.BgpNeighborConditionalAdvertise6Args>());
            set => _conditionalAdvertise6s = value;
        }

        [Input("conditionalAdvertises")]
        private InputList<Inputs.BgpNeighborConditionalAdvertiseArgs>? _conditionalAdvertises;

        /// <summary>
        /// Conditional advertisement. The structure of `conditional_advertise` block is documented below.
        /// </summary>
        public InputList<Inputs.BgpNeighborConditionalAdvertiseArgs> ConditionalAdvertises
        {
            get => _conditionalAdvertises ?? (_conditionalAdvertises = new InputList<Inputs.BgpNeighborConditionalAdvertiseArgs>());
            set => _conditionalAdvertises = value;
        }

        /// <summary>
        /// Interval (sec) for connect timer.
        /// </summary>
        [Input("connectTimer")]
        public Input<int>? ConnectTimer { get; set; }

        /// <summary>
        /// Route map to specify criteria to originate IPv4 default.
        /// </summary>
        [Input("defaultOriginateRoutemap")]
        public Input<string>? DefaultOriginateRoutemap { get; set; }

        /// <summary>
        /// Route map to specify criteria to originate IPv6 default.
        /// </summary>
        [Input("defaultOriginateRoutemap6")]
        public Input<string>? DefaultOriginateRoutemap6 { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Filter for IPv4 updates from this neighbor.
        /// </summary>
        [Input("distributeListIn")]
        public Input<string>? DistributeListIn { get; set; }

        /// <summary>
        /// Filter for IPv6 updates from this neighbor.
        /// </summary>
        [Input("distributeListIn6")]
        public Input<string>? DistributeListIn6 { get; set; }

        /// <summary>
        /// Filter for VPNv4 updates from this neighbor.
        /// </summary>
        [Input("distributeListInVpnv4")]
        public Input<string>? DistributeListInVpnv4 { get; set; }

        /// <summary>
        /// Filter for VPNv6 updates from this neighbor.
        /// </summary>
        [Input("distributeListInVpnv6")]
        public Input<string>? DistributeListInVpnv6 { get; set; }

        /// <summary>
        /// Filter for IPv4 updates to this neighbor.
        /// </summary>
        [Input("distributeListOut")]
        public Input<string>? DistributeListOut { get; set; }

        /// <summary>
        /// Filter for IPv6 updates to this neighbor.
        /// </summary>
        [Input("distributeListOut6")]
        public Input<string>? DistributeListOut6 { get; set; }

        /// <summary>
        /// Filter for VPNv4 updates to this neighbor.
        /// </summary>
        [Input("distributeListOutVpnv4")]
        public Input<string>? DistributeListOutVpnv4 { get; set; }

        /// <summary>
        /// Filter for VPNv6 updates to this neighbor.
        /// </summary>
        [Input("distributeListOutVpnv6")]
        public Input<string>? DistributeListOutVpnv6 { get; set; }

        /// <summary>
        /// Don't negotiate capabilities with this neighbor Valid values: `enable`, `disable`.
        /// </summary>
        [Input("dontCapabilityNegotiate")]
        public Input<string>? DontCapabilityNegotiate { get; set; }

        /// <summary>
        /// Enable/disable allow multi-hop EBGP neighbors. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("ebgpEnforceMultihop")]
        public Input<string>? EbgpEnforceMultihop { get; set; }

        /// <summary>
        /// EBGP multihop TTL for this peer.
        /// </summary>
        [Input("ebgpMultihopTtl")]
        public Input<int>? EbgpMultihopTtl { get; set; }

        /// <summary>
        /// BGP filter for IPv4 inbound routes.
        /// </summary>
        [Input("filterListIn")]
        public Input<string>? FilterListIn { get; set; }

        /// <summary>
        /// BGP filter for IPv6 inbound routes.
        /// </summary>
        [Input("filterListIn6")]
        public Input<string>? FilterListIn6 { get; set; }

        /// <summary>
        /// BGP filter for VPNv4 inbound routes.
        /// </summary>
        [Input("filterListInVpnv4")]
        public Input<string>? FilterListInVpnv4 { get; set; }

        /// <summary>
        /// BGP filter for VPNv6 inbound routes.
        /// </summary>
        [Input("filterListInVpnv6")]
        public Input<string>? FilterListInVpnv6 { get; set; }

        /// <summary>
        /// BGP filter for IPv4 outbound routes.
        /// </summary>
        [Input("filterListOut")]
        public Input<string>? FilterListOut { get; set; }

        /// <summary>
        /// BGP filter for IPv6 outbound routes.
        /// </summary>
        [Input("filterListOut6")]
        public Input<string>? FilterListOut6 { get; set; }

        /// <summary>
        /// BGP filter for VPNv4 outbound routes.
        /// </summary>
        [Input("filterListOutVpnv4")]
        public Input<string>? FilterListOutVpnv4 { get; set; }

        /// <summary>
        /// BGP filter for VPNv6 outbound routes.
        /// </summary>
        [Input("filterListOutVpnv6")]
        public Input<string>? FilterListOutVpnv6 { get; set; }

        /// <summary>
        /// Interval (sec) before peer considered dead.
        /// </summary>
        [Input("holdtimeTimer")]
        public Input<int>? HoldtimeTimer { get; set; }

        /// <summary>
        /// Interface
        /// </summary>
        [Input("interface")]
        public Input<string>? Interface { get; set; }

        /// <summary>
        /// IP/IPv6 address of neighbor.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Keep alive timer interval (sec).
        /// </summary>
        [Input("keepAliveTimer")]
        public Input<int>? KeepAliveTimer { get; set; }

        /// <summary>
        /// Enable/disable failover upon link down. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("linkDownFailover")]
        public Input<string>? LinkDownFailover { get; set; }

        /// <summary>
        /// Local AS number of neighbor.
        /// </summary>
        [Input("localAs")]
        public Input<int>? LocalAs { get; set; }

        /// <summary>
        /// Do not prepend local-as to incoming updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localAsNoPrepend")]
        public Input<string>? LocalAsNoPrepend { get; set; }

        /// <summary>
        /// Replace real AS with local-as in outgoing updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("localAsReplaceAs")]
        public Input<string>? LocalAsReplaceAs { get; set; }

        /// <summary>
        /// Maximum number of IPv4 prefixes to accept from this peer.
        /// </summary>
        [Input("maximumPrefix")]
        public Input<int>? MaximumPrefix { get; set; }

        /// <summary>
        /// Maximum number of IPv6 prefixes to accept from this peer.
        /// </summary>
        [Input("maximumPrefix6")]
        public Input<int>? MaximumPrefix6 { get; set; }

        /// <summary>
        /// Maximum number of L2VPN EVPN prefixes to accept from this peer.
        /// </summary>
        [Input("maximumPrefixEvpn")]
        public Input<int>? MaximumPrefixEvpn { get; set; }

        /// <summary>
        /// Maximum IPv4 prefix threshold value (1 - 100 percent).
        /// </summary>
        [Input("maximumPrefixThreshold")]
        public Input<int>? MaximumPrefixThreshold { get; set; }

        /// <summary>
        /// Maximum IPv6 prefix threshold value (1 - 100 percent).
        /// </summary>
        [Input("maximumPrefixThreshold6")]
        public Input<int>? MaximumPrefixThreshold6 { get; set; }

        /// <summary>
        /// Maximum L2VPN EVPN prefix threshold value (1 - 100 percent).
        /// </summary>
        [Input("maximumPrefixThresholdEvpn")]
        public Input<int>? MaximumPrefixThresholdEvpn { get; set; }

        /// <summary>
        /// Maximum VPNv4 prefix threshold value (1 - 100 percent).
        /// </summary>
        [Input("maximumPrefixThresholdVpnv4")]
        public Input<int>? MaximumPrefixThresholdVpnv4 { get; set; }

        /// <summary>
        /// Maximum VPNv6 prefix threshold value (1 - 100 percent).
        /// </summary>
        [Input("maximumPrefixThresholdVpnv6")]
        public Input<int>? MaximumPrefixThresholdVpnv6 { get; set; }

        /// <summary>
        /// Maximum number of VPNv4 prefixes to accept from this peer.
        /// </summary>
        [Input("maximumPrefixVpnv4")]
        public Input<int>? MaximumPrefixVpnv4 { get; set; }

        /// <summary>
        /// Maximum number of VPNv6 prefixes to accept from this peer.
        /// </summary>
        [Input("maximumPrefixVpnv6")]
        public Input<int>? MaximumPrefixVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable IPv4 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maximumPrefixWarningOnly")]
        public Input<string>? MaximumPrefixWarningOnly { get; set; }

        /// <summary>
        /// Enable/disable IPv6 Only give warning message when limit is exceeded. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maximumPrefixWarningOnly6")]
        public Input<string>? MaximumPrefixWarningOnly6 { get; set; }

        /// <summary>
        /// Enable/disable only sending warning message when exceeding limit of L2VPN EVPN routes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maximumPrefixWarningOnlyEvpn")]
        public Input<string>? MaximumPrefixWarningOnlyEvpn { get; set; }

        /// <summary>
        /// Enable/disable only giving warning message when limit is exceeded for VPNv4 routes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maximumPrefixWarningOnlyVpnv4")]
        public Input<string>? MaximumPrefixWarningOnlyVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable warning message when limit is exceeded for VPNv6 routes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("maximumPrefixWarningOnlyVpnv6")]
        public Input<string>? MaximumPrefixWarningOnlyVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable IPv4 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelf")]
        public Input<string>? NextHopSelf { get; set; }

        /// <summary>
        /// Enable/disable IPv6 next-hop calculation for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelf6")]
        public Input<string>? NextHopSelf6 { get; set; }

        /// <summary>
        /// Enable/disable setting nexthop's address to interface's IPv4 address for route-reflector routes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelfRr")]
        public Input<string>? NextHopSelfRr { get; set; }

        /// <summary>
        /// Enable/disable setting nexthop's address to interface's IPv6 address for route-reflector routes. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelfRr6")]
        public Input<string>? NextHopSelfRr6 { get; set; }

        /// <summary>
        /// Enable/disable setting VPNv4 next-hop to interface's IP address for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelfVpnv4")]
        public Input<string>? NextHopSelfVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable use of outgoing interface's IP address as VPNv6 next-hop for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("nextHopSelfVpnv6")]
        public Input<string>? NextHopSelfVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable override result of capability negotiation. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("overrideCapability")]
        public Input<string>? OverrideCapability { get; set; }

        /// <summary>
        /// Enable/disable sending of open messages to this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("passive")]
        public Input<string>? Passive { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used in MD5 authentication.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// IPv4 Inbound filter for updates from this neighbor.
        /// </summary>
        [Input("prefixListIn")]
        public Input<string>? PrefixListIn { get; set; }

        /// <summary>
        /// IPv6 Inbound filter for updates from this neighbor.
        /// </summary>
        [Input("prefixListIn6")]
        public Input<string>? PrefixListIn6 { get; set; }

        /// <summary>
        /// Inbound filter for VPNv4 updates from this neighbor.
        /// </summary>
        [Input("prefixListInVpnv4")]
        public Input<string>? PrefixListInVpnv4 { get; set; }

        /// <summary>
        /// Inbound filter for VPNv6 updates from this neighbor.
        /// </summary>
        [Input("prefixListInVpnv6")]
        public Input<string>? PrefixListInVpnv6 { get; set; }

        /// <summary>
        /// IPv4 Outbound filter for updates to this neighbor.
        /// </summary>
        [Input("prefixListOut")]
        public Input<string>? PrefixListOut { get; set; }

        /// <summary>
        /// IPv6 Outbound filter for updates to this neighbor.
        /// </summary>
        [Input("prefixListOut6")]
        public Input<string>? PrefixListOut6 { get; set; }

        /// <summary>
        /// Outbound filter for VPNv4 updates to this neighbor.
        /// </summary>
        [Input("prefixListOutVpnv4")]
        public Input<string>? PrefixListOutVpnv4 { get; set; }

        /// <summary>
        /// Outbound filter for VPNv6 updates to this neighbor.
        /// </summary>
        [Input("prefixListOutVpnv6")]
        public Input<string>? PrefixListOutVpnv6 { get; set; }

        /// <summary>
        /// AS number of neighbor.
        /// </summary>
        [Input("remoteAs")]
        public Input<int>? RemoteAs { get; set; }

        /// <summary>
        /// Enable/disable remove private AS number from IPv4 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("removePrivateAs")]
        public Input<string>? RemovePrivateAs { get; set; }

        /// <summary>
        /// Enable/disable remove private AS number from IPv6 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("removePrivateAs6")]
        public Input<string>? RemovePrivateAs6 { get; set; }

        /// <summary>
        /// Enable/disable removing private AS number from L2VPN EVPN outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("removePrivateAsEvpn")]
        public Input<string>? RemovePrivateAsEvpn { get; set; }

        /// <summary>
        /// Enable/disable remove private AS number from VPNv4 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("removePrivateAsVpnv4")]
        public Input<string>? RemovePrivateAsVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable to remove private AS number from VPNv6 outbound updates. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("removePrivateAsVpnv6")]
        public Input<string>? RemovePrivateAsVpnv6 { get; set; }

        /// <summary>
        /// Graceful restart delay time (sec, 0 = global default).
        /// </summary>
        [Input("restartTime")]
        public Input<int>? RestartTime { get; set; }

        /// <summary>
        /// Time to retain stale routes.
        /// </summary>
        [Input("retainStaleTime")]
        public Input<int>? RetainStaleTime { get; set; }

        /// <summary>
        /// IPv4 Inbound route map filter.
        /// </summary>
        [Input("routeMapIn")]
        public Input<string>? RouteMapIn { get; set; }

        /// <summary>
        /// IPv6 Inbound route map filter.
        /// </summary>
        [Input("routeMapIn6")]
        public Input<string>? RouteMapIn6 { get; set; }

        /// <summary>
        /// L2VPN EVPN inbound route map filter.
        /// </summary>
        [Input("routeMapInEvpn")]
        public Input<string>? RouteMapInEvpn { get; set; }

        /// <summary>
        /// VPNv4 inbound route map filter.
        /// </summary>
        [Input("routeMapInVpnv4")]
        public Input<string>? RouteMapInVpnv4 { get; set; }

        /// <summary>
        /// VPNv6 inbound route map filter.
        /// </summary>
        [Input("routeMapInVpnv6")]
        public Input<string>? RouteMapInVpnv6 { get; set; }

        /// <summary>
        /// IPv4 Outbound route map filter.
        /// </summary>
        [Input("routeMapOut")]
        public Input<string>? RouteMapOut { get; set; }

        /// <summary>
        /// IPv6 Outbound route map filter.
        /// </summary>
        [Input("routeMapOut6")]
        public Input<string>? RouteMapOut6 { get; set; }

        /// <summary>
        /// IPv6 outbound route map filter if the peer is preferred.
        /// </summary>
        [Input("routeMapOut6Preferable")]
        public Input<string>? RouteMapOut6Preferable { get; set; }

        /// <summary>
        /// L2VPN EVPN outbound route map filter.
        /// </summary>
        [Input("routeMapOutEvpn")]
        public Input<string>? RouteMapOutEvpn { get; set; }

        /// <summary>
        /// IPv4 outbound route map filter if the peer is preferred.
        /// </summary>
        [Input("routeMapOutPreferable")]
        public Input<string>? RouteMapOutPreferable { get; set; }

        /// <summary>
        /// VPNv4 outbound route map filter.
        /// </summary>
        [Input("routeMapOutVpnv4")]
        public Input<string>? RouteMapOutVpnv4 { get; set; }

        /// <summary>
        /// VPNv4 outbound route map filter if the peer is preferred.
        /// </summary>
        [Input("routeMapOutVpnv4Preferable")]
        public Input<string>? RouteMapOutVpnv4Preferable { get; set; }

        /// <summary>
        /// VPNv6 outbound route map filter.
        /// </summary>
        [Input("routeMapOutVpnv6")]
        public Input<string>? RouteMapOutVpnv6 { get; set; }

        /// <summary>
        /// VPNv6 outbound route map filter if this neighbor is preferred.
        /// </summary>
        [Input("routeMapOutVpnv6Preferable")]
        public Input<string>? RouteMapOutVpnv6Preferable { get; set; }

        /// <summary>
        /// Enable/disable IPv4 AS route reflector client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeReflectorClient")]
        public Input<string>? RouteReflectorClient { get; set; }

        /// <summary>
        /// Enable/disable IPv6 AS route reflector client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeReflectorClient6")]
        public Input<string>? RouteReflectorClient6 { get; set; }

        /// <summary>
        /// Enable/disable L2VPN EVPN AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeReflectorClientEvpn")]
        public Input<string>? RouteReflectorClientEvpn { get; set; }

        /// <summary>
        /// Enable/disable VPNv4 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeReflectorClientVpnv4")]
        public Input<string>? RouteReflectorClientVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable VPNv6 AS route reflector client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeReflectorClientVpnv6")]
        public Input<string>? RouteReflectorClientVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable IPv4 AS route server client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeServerClient")]
        public Input<string>? RouteServerClient { get; set; }

        /// <summary>
        /// Enable/disable IPv6 AS route server client. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeServerClient6")]
        public Input<string>? RouteServerClient6 { get; set; }

        /// <summary>
        /// Enable/disable L2VPN EVPN AS route server client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeServerClientEvpn")]
        public Input<string>? RouteServerClientEvpn { get; set; }

        /// <summary>
        /// Enable/disable VPNv4 AS route server client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeServerClientVpnv4")]
        public Input<string>? RouteServerClientVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable VPNv6 AS route server client for this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("routeServerClientVpnv6")]
        public Input<string>? RouteServerClientVpnv6 { get; set; }

        /// <summary>
        /// IPv4 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        [Input("sendCommunity")]
        public Input<string>? SendCommunity { get; set; }

        /// <summary>
        /// IPv6 Send community attribute to neighbor. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        [Input("sendCommunity6")]
        public Input<string>? SendCommunity6 { get; set; }

        /// <summary>
        /// Enable/disable sending community attribute to neighbor for L2VPN EVPN address family. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        [Input("sendCommunityEvpn")]
        public Input<string>? SendCommunityEvpn { get; set; }

        /// <summary>
        /// Send community attribute to neighbor for VPNv4 address family. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        [Input("sendCommunityVpnv4")]
        public Input<string>? SendCommunityVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable sending community attribute to this neighbor for VPNv6 address family. Valid values: `standard`, `extended`, `both`, `disable`.
        /// </summary>
        [Input("sendCommunityVpnv6")]
        public Input<string>? SendCommunityVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable shutdown this neighbor. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("shutdown")]
        public Input<string>? Shutdown { get; set; }

        /// <summary>
        /// Enable/disable allow IPv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("softReconfiguration")]
        public Input<string>? SoftReconfiguration { get; set; }

        /// <summary>
        /// Enable/disable allow IPv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("softReconfiguration6")]
        public Input<string>? SoftReconfiguration6 { get; set; }

        /// <summary>
        /// Enable/disable L2VPN EVPN inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("softReconfigurationEvpn")]
        public Input<string>? SoftReconfigurationEvpn { get; set; }

        /// <summary>
        /// Enable/disable allow VPNv4 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("softReconfigurationVpnv4")]
        public Input<string>? SoftReconfigurationVpnv4 { get; set; }

        /// <summary>
        /// Enable/disable VPNv6 inbound soft reconfiguration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("softReconfigurationVpnv6")]
        public Input<string>? SoftReconfigurationVpnv6 { get; set; }

        /// <summary>
        /// Enable/disable stale route after neighbor down. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("staleRoute")]
        public Input<string>? StaleRoute { get; set; }

        /// <summary>
        /// Enable/disable strict capability matching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("strictCapabilityMatch")]
        public Input<string>? StrictCapabilityMatch { get; set; }

        /// <summary>
        /// IPv4 Route map to selectively unsuppress suppressed routes.
        /// </summary>
        [Input("unsuppressMap")]
        public Input<string>? UnsuppressMap { get; set; }

        /// <summary>
        /// IPv6 Route map to selectively unsuppress suppressed routes.
        /// </summary>
        [Input("unsuppressMap6")]
        public Input<string>? UnsuppressMap6 { get; set; }

        /// <summary>
        /// Interface to use as source IP/IPv6 address of TCP connections.
        /// </summary>
        [Input("updateSource")]
        public Input<string>? UpdateSource { get; set; }

        /// <summary>
        /// Neighbor weight.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public BgpNeighborArgs()
        {
        }
        public static new BgpNeighborArgs Empty => new BgpNeighborArgs();
    }
}
