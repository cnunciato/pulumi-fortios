// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Vpn.Ssl
{
    public static class GetSettings
    {
        /// <summary>
        /// Use this data source to get information on fortios vpnssl settings
        /// </summary>
        public static Task<GetSettingsResult> InvokeAsync(GetSettingsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSettingsResult>("fortios:vpn/ssl/getSettings:getSettings", args ?? new GetSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios vpnssl settings
        /// </summary>
        public static Output<GetSettingsResult> Invoke(GetSettingsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSettingsResult>("fortios:vpn/ssl/getSettings:getSettings", args ?? new GetSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSettingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSettingsArgs()
        {
        }
        public static new GetSettingsArgs Empty => new GetSettingsArgs();
    }

    public sealed class GetSettingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSettingsInvokeArgs()
        {
        }
        public static new GetSettingsInvokeArgs Empty => new GetSettingsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSettingsResult
    {
        /// <summary>
        /// Force the SSL-VPN security level. High allows only high. Medium allows medium and high. Low allows any.
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// Enable/disable checking of source IP for authentication session.
        /// </summary>
        public readonly string AuthSessionCheckSourceIp;
        /// <summary>
        /// SSL-VPN authentication timeout (1 - 259200 sec (3 days), 0 for no timeout).
        /// </summary>
        public readonly int AuthTimeout;
        /// <summary>
        /// Authentication rule for SSL VPN. The structure of `authentication_rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsAuthenticationRuleResult> AuthenticationRules;
        /// <summary>
        /// Enable to auto-create static routes for the SSL-VPN tunnel IP addresses.
        /// </summary>
        public readonly string AutoTunnelStaticRoute;
        /// <summary>
        /// Select one or more cipher technologies that cannot be used in SSL-VPN negotiations.
        /// </summary>
        public readonly string BannedCipher;
        /// <summary>
        /// Enable/disable overriding the configured system language based on the preferred language of the browser.
        /// </summary>
        public readonly string BrowserLanguageDetection;
        /// <summary>
        /// Enable/disable verification of referer field in HTTP request header.
        /// </summary>
        public readonly string CheckReferer;
        /// <summary>
        /// Select one or more TLS 1.3 ciphersuites to enable. Does not affect ciphers in TLS 1.2 and below. At least one must be enabled. To disable all, set ssl-max-proto-ver to tls1-2 or below.
        /// </summary>
        public readonly string Ciphersuite;
        /// <summary>
        /// Set signature algorithms related to client authentication. Affects TLS version &lt;= 1.2 only.
        /// </summary>
        public readonly string ClientSigalgs;
        /// <summary>
        /// Default SSL VPN portal.
        /// </summary>
        public readonly string DefaultPortal;
        /// <summary>
        /// Compression level (0~9).
        /// </summary>
        public readonly int DeflateCompressionLevel;
        /// <summary>
        /// Minimum amount of data that triggers compression (200 - 65535 bytes).
        /// </summary>
        public readonly int DeflateMinDataSize;
        /// <summary>
        /// DNS server 1.
        /// </summary>
        public readonly string DnsServer1;
        /// <summary>
        /// DNS server 2.
        /// </summary>
        public readonly string DnsServer2;
        /// <summary>
        /// DNS suffix used for SSL-VPN clients.
        /// </summary>
        public readonly string DnsSuffix;
        /// <summary>
        /// Number of missing heartbeats before the connection is considered dropped.
        /// </summary>
        public readonly int DtlsHeartbeatFailCount;
        /// <summary>
        /// Idle timeout before DTLS heartbeat is sent.
        /// </summary>
        public readonly int DtlsHeartbeatIdleTimeout;
        /// <summary>
        /// Interval between DTLS heartbeat.
        /// </summary>
        public readonly int DtlsHeartbeatInterval;
        /// <summary>
        /// SSLVPN maximum DTLS hello timeout (10 - 60 sec, default = 10).
        /// </summary>
        public readonly int DtlsHelloTimeout;
        /// <summary>
        /// DTLS maximum protocol version.
        /// </summary>
        public readonly string DtlsMaxProtoVer;
        /// <summary>
        /// DTLS minimum protocol version.
        /// </summary>
        public readonly string DtlsMinProtoVer;
        /// <summary>
        /// Enable DTLS to prevent eavesdropping, tampering, or message forgery.
        /// </summary>
        public readonly string DtlsTunnel;
        /// <summary>
        /// Tunnel mode: enable parallel IPv4 and IPv6 tunnel. Web mode: support IPv4 and IPv6 bookmarks in the portal.
        /// </summary>
        public readonly string DualStackMode;
        /// <summary>
        /// Encode \2F sequence to forward slash in URLs.
        /// </summary>
        public readonly string Encode2fSequence;
        /// <summary>
        /// Encrypt and store user passwords for SSL-VPN web sessions.
        /// </summary>
        public readonly string EncryptAndStorePassword;
        /// <summary>
        /// Enable to force two-factor authentication for all SSL-VPNs.
        /// </summary>
        public readonly string ForceTwoFactorAuth;
        /// <summary>
        /// Forward the same, add, or remove HTTP header.
        /// </summary>
        public readonly string HeaderXForwardedFor;
        /// <summary>
        /// Add HSTS includeSubDomains response header.
        /// </summary>
        public readonly string HstsIncludeSubdomains;
        /// <summary>
        /// Enable to allow HTTP compression over SSL-VPN tunnels.
        /// </summary>
        public readonly string HttpCompression;
        /// <summary>
        /// Enable/disable SSL-VPN support for HttpOnly cookies.
        /// </summary>
        public readonly string HttpOnlyCookie;
        /// <summary>
        /// SSL-VPN session is disconnected if an HTTP request body is not received within this time (1 - 60 sec, default = 20).
        /// </summary>
        public readonly int HttpRequestBodyTimeout;
        /// <summary>
        /// SSL-VPN session is disconnected if an HTTP request header is not received within this time (1 - 60 sec, default = 20).
        /// </summary>
        public readonly int HttpRequestHeaderTimeout;
        /// <summary>
        /// Enable/disable redirect of port 80 to SSL-VPN port.
        /// </summary>
        public readonly string HttpsRedirect;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// SSL VPN disconnects if idle for specified time in seconds.
        /// </summary>
        public readonly int IdleTimeout;
        /// <summary>
        /// IPv6 DNS server 1.
        /// </summary>
        public readonly string Ipv6DnsServer1;
        /// <summary>
        /// IPv6 DNS server 2.
        /// </summary>
        public readonly string Ipv6DnsServer2;
        /// <summary>
        /// IPv6 WINS server 1.
        /// </summary>
        public readonly string Ipv6WinsServer1;
        /// <summary>
        /// IPv6 WINS server 2.
        /// </summary>
        public readonly string Ipv6WinsServer2;
        /// <summary>
        /// SSL VPN maximum login attempt times before block (0 - 10, default = 2, 0 = no limit).
        /// </summary>
        public readonly int LoginAttemptLimit;
        /// <summary>
        /// Time for which a user is blocked from logging in after too many failed login attempts (0 - 86400 sec, default = 60).
        /// </summary>
        public readonly int LoginBlockTime;
        /// <summary>
        /// SSLVPN maximum login timeout (10 - 180 sec, default = 30).
        /// </summary>
        public readonly int LoginTimeout;
        /// <summary>
        /// SSL-VPN access port (1 - 65535).
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Enable means that if SSL-VPN connections are allowed on an interface admin GUI connections are blocked on that interface.
        /// </summary>
        public readonly string PortPrecedence;
        /// <summary>
        /// Enable to require client certificates for all SSL-VPN users.
        /// </summary>
        public readonly string Reqclientcert;
        /// <summary>
        /// Enable to allow SSL-VPN sessions to bypass routing and bind to the incoming interface.
        /// </summary>
        public readonly string RouteSourceInterface;
        /// <summary>
        /// SAML local redirect port in the machine running FCT (0 - 65535). 0 is to disable redirection on FGT side.
        /// </summary>
        public readonly int SamlRedirectPort;
        /// <summary>
        /// Server hostname for HTTPS. When set, will be used for SSL VPN web proxy host header for any redirection.
        /// </summary>
        public readonly string ServerHostname;
        /// <summary>
        /// Name of the server certificate to be used for SSL-VPNs.
        /// </summary>
        public readonly string Servercert;
        /// <summary>
        /// Enable/disable negated source IPv6 address match.
        /// </summary>
        public readonly string SourceAddress6Negate;
        /// <summary>
        /// IPv6 source address of incoming traffic. The structure of `source_address6` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsSourceAddress6Result> SourceAddress6s;
        /// <summary>
        /// Enable/disable negated source address match.
        /// </summary>
        public readonly string SourceAddressNegate;
        /// <summary>
        /// Source address of incoming traffic. The structure of `source_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsSourceAddressResult> SourceAddresses;
        /// <summary>
        /// SSL VPN source interface of incoming traffic. The structure of `source_interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsSourceInterfaceResult> SourceInterfaces;
        /// <summary>
        /// Enable to allow client renegotiation by the server if the tunnel goes down.
        /// </summary>
        public readonly string SslClientRenegotiation;
        /// <summary>
        /// Enable/disable insertion of empty fragment.
        /// </summary>
        public readonly string SslInsertEmptyFragment;
        /// <summary>
        /// SSL maximum protocol version.
        /// </summary>
        public readonly string SslMaxProtoVer;
        /// <summary>
        /// SSL minimum protocol version.
        /// </summary>
        public readonly string SslMinProtoVer;
        /// <summary>
        /// Enable/disable SSL-VPN.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Enable/disable TLSv1.0.
        /// </summary>
        public readonly string Tlsv10;
        /// <summary>
        /// Enable/disable TLSv1.1.
        /// </summary>
        public readonly string Tlsv11;
        /// <summary>
        /// Enable/disable TLSv1.2.
        /// </summary>
        public readonly string Tlsv12;
        /// <summary>
        /// Enable/disable TLSv1.3.
        /// </summary>
        public readonly string Tlsv13;
        /// <summary>
        /// Transform backward slashes to forward slashes in URLs.
        /// </summary>
        public readonly string TransformBackwardSlashes;
        /// <summary>
        /// Method used for assigning address for tunnel.
        /// </summary>
        public readonly string TunnelAddrAssignedMethod;
        /// <summary>
        /// Enable/disable tunnel connection without re-authorization if previous connection dropped.
        /// </summary>
        public readonly string TunnelConnectWithoutReauth;
        /// <summary>
        /// Names of the IPv4 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ip_pools` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsTunnelIpPoolResult> TunnelIpPools;
        /// <summary>
        /// Names of the IPv6 IP Pool firewall objects that define the IP addresses reserved for remote clients. The structure of `tunnel_ipv6_pools` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSettingsTunnelIpv6PoolResult> TunnelIpv6Pools;
        /// <summary>
        /// Time out value to clean up user session after tunnel connection is dropped (1 - 255 sec, default=30).
        /// </summary>
        public readonly int TunnelUserSessionTimeout;
        /// <summary>
        /// Enable/disable unsafe legacy re-negotiation.
        /// </summary>
        public readonly string UnsafeLegacyRenegotiation;
        /// <summary>
        /// Enable to obscure the host name of the URL of the web browser display.
        /// </summary>
        public readonly string UrlObscuration;
        /// <summary>
        /// Name of user peer.
        /// </summary>
        public readonly string UserPeer;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable use of IP pools defined in firewall policy while using web-mode.
        /// </summary>
        public readonly string WebModeSnat;
        /// <summary>
        /// WINS server 1.
        /// </summary>
        public readonly string WinsServer1;
        /// <summary>
        /// WINS server 2.
        /// </summary>
        public readonly string WinsServer2;
        /// <summary>
        /// Add HTTP X-Content-Type-Options header.
        /// </summary>
        public readonly string XContentTypeOptions;
        /// <summary>
        /// Enable/disable verification of device certificate for SSLVPN ZTNA session.
        /// </summary>
        public readonly string ZtnaTrustedClient;

        [OutputConstructor]
        private GetSettingsResult(
            string algorithm,

            string authSessionCheckSourceIp,

            int authTimeout,

            ImmutableArray<Outputs.GetSettingsAuthenticationRuleResult> authenticationRules,

            string autoTunnelStaticRoute,

            string bannedCipher,

            string browserLanguageDetection,

            string checkReferer,

            string ciphersuite,

            string clientSigalgs,

            string defaultPortal,

            int deflateCompressionLevel,

            int deflateMinDataSize,

            string dnsServer1,

            string dnsServer2,

            string dnsSuffix,

            int dtlsHeartbeatFailCount,

            int dtlsHeartbeatIdleTimeout,

            int dtlsHeartbeatInterval,

            int dtlsHelloTimeout,

            string dtlsMaxProtoVer,

            string dtlsMinProtoVer,

            string dtlsTunnel,

            string dualStackMode,

            string encode2fSequence,

            string encryptAndStorePassword,

            string forceTwoFactorAuth,

            string headerXForwardedFor,

            string hstsIncludeSubdomains,

            string httpCompression,

            string httpOnlyCookie,

            int httpRequestBodyTimeout,

            int httpRequestHeaderTimeout,

            string httpsRedirect,

            string id,

            int idleTimeout,

            string ipv6DnsServer1,

            string ipv6DnsServer2,

            string ipv6WinsServer1,

            string ipv6WinsServer2,

            int loginAttemptLimit,

            int loginBlockTime,

            int loginTimeout,

            int port,

            string portPrecedence,

            string reqclientcert,

            string routeSourceInterface,

            int samlRedirectPort,

            string serverHostname,

            string servercert,

            string sourceAddress6Negate,

            ImmutableArray<Outputs.GetSettingsSourceAddress6Result> sourceAddress6s,

            string sourceAddressNegate,

            ImmutableArray<Outputs.GetSettingsSourceAddressResult> sourceAddresses,

            ImmutableArray<Outputs.GetSettingsSourceInterfaceResult> sourceInterfaces,

            string sslClientRenegotiation,

            string sslInsertEmptyFragment,

            string sslMaxProtoVer,

            string sslMinProtoVer,

            string status,

            string tlsv10,

            string tlsv11,

            string tlsv12,

            string tlsv13,

            string transformBackwardSlashes,

            string tunnelAddrAssignedMethod,

            string tunnelConnectWithoutReauth,

            ImmutableArray<Outputs.GetSettingsTunnelIpPoolResult> tunnelIpPools,

            ImmutableArray<Outputs.GetSettingsTunnelIpv6PoolResult> tunnelIpv6Pools,

            int tunnelUserSessionTimeout,

            string unsafeLegacyRenegotiation,

            string urlObscuration,

            string userPeer,

            string? vdomparam,

            string webModeSnat,

            string winsServer1,

            string winsServer2,

            string xContentTypeOptions,

            string ztnaTrustedClient)
        {
            Algorithm = algorithm;
            AuthSessionCheckSourceIp = authSessionCheckSourceIp;
            AuthTimeout = authTimeout;
            AuthenticationRules = authenticationRules;
            AutoTunnelStaticRoute = autoTunnelStaticRoute;
            BannedCipher = bannedCipher;
            BrowserLanguageDetection = browserLanguageDetection;
            CheckReferer = checkReferer;
            Ciphersuite = ciphersuite;
            ClientSigalgs = clientSigalgs;
            DefaultPortal = defaultPortal;
            DeflateCompressionLevel = deflateCompressionLevel;
            DeflateMinDataSize = deflateMinDataSize;
            DnsServer1 = dnsServer1;
            DnsServer2 = dnsServer2;
            DnsSuffix = dnsSuffix;
            DtlsHeartbeatFailCount = dtlsHeartbeatFailCount;
            DtlsHeartbeatIdleTimeout = dtlsHeartbeatIdleTimeout;
            DtlsHeartbeatInterval = dtlsHeartbeatInterval;
            DtlsHelloTimeout = dtlsHelloTimeout;
            DtlsMaxProtoVer = dtlsMaxProtoVer;
            DtlsMinProtoVer = dtlsMinProtoVer;
            DtlsTunnel = dtlsTunnel;
            DualStackMode = dualStackMode;
            Encode2fSequence = encode2fSequence;
            EncryptAndStorePassword = encryptAndStorePassword;
            ForceTwoFactorAuth = forceTwoFactorAuth;
            HeaderXForwardedFor = headerXForwardedFor;
            HstsIncludeSubdomains = hstsIncludeSubdomains;
            HttpCompression = httpCompression;
            HttpOnlyCookie = httpOnlyCookie;
            HttpRequestBodyTimeout = httpRequestBodyTimeout;
            HttpRequestHeaderTimeout = httpRequestHeaderTimeout;
            HttpsRedirect = httpsRedirect;
            Id = id;
            IdleTimeout = idleTimeout;
            Ipv6DnsServer1 = ipv6DnsServer1;
            Ipv6DnsServer2 = ipv6DnsServer2;
            Ipv6WinsServer1 = ipv6WinsServer1;
            Ipv6WinsServer2 = ipv6WinsServer2;
            LoginAttemptLimit = loginAttemptLimit;
            LoginBlockTime = loginBlockTime;
            LoginTimeout = loginTimeout;
            Port = port;
            PortPrecedence = portPrecedence;
            Reqclientcert = reqclientcert;
            RouteSourceInterface = routeSourceInterface;
            SamlRedirectPort = samlRedirectPort;
            ServerHostname = serverHostname;
            Servercert = servercert;
            SourceAddress6Negate = sourceAddress6Negate;
            SourceAddress6s = sourceAddress6s;
            SourceAddressNegate = sourceAddressNegate;
            SourceAddresses = sourceAddresses;
            SourceInterfaces = sourceInterfaces;
            SslClientRenegotiation = sslClientRenegotiation;
            SslInsertEmptyFragment = sslInsertEmptyFragment;
            SslMaxProtoVer = sslMaxProtoVer;
            SslMinProtoVer = sslMinProtoVer;
            Status = status;
            Tlsv10 = tlsv10;
            Tlsv11 = tlsv11;
            Tlsv12 = tlsv12;
            Tlsv13 = tlsv13;
            TransformBackwardSlashes = transformBackwardSlashes;
            TunnelAddrAssignedMethod = tunnelAddrAssignedMethod;
            TunnelConnectWithoutReauth = tunnelConnectWithoutReauth;
            TunnelIpPools = tunnelIpPools;
            TunnelIpv6Pools = tunnelIpv6Pools;
            TunnelUserSessionTimeout = tunnelUserSessionTimeout;
            UnsafeLegacyRenegotiation = unsafeLegacyRenegotiation;
            UrlObscuration = urlObscuration;
            UserPeer = userPeer;
            Vdomparam = vdomparam;
            WebModeSnat = webModeSnat;
            WinsServer1 = winsServer1;
            WinsServer2 = winsServer2;
            XContentTypeOptions = xContentTypeOptions;
            ZtnaTrustedClient = ztnaTrustedClient;
        }
    }
}