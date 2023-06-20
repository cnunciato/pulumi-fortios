// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class FirewallProfileprotocoloptionsHttpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable/disable IP based URL rating. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("addressIpRating")]
        public Input<string>? AddressIpRating { get; set; }

        /// <summary>
        /// Code number returned for blocked HTTP pages (non-FortiGuard only) (100 - 599, default = 403).
        /// </summary>
        [Input("blockPageStatusCode")]
        public Input<int>? BlockPageStatusCode { get; set; }

        /// <summary>
        /// Amount of data to send in a transmission for client comforting (1 - 10240 bytes, default = 1).
        /// </summary>
        [Input("comfortAmount")]
        public Input<int>? ComfortAmount { get; set; }

        /// <summary>
        /// Period of time between start, or last transmission, and the next client comfort transmission of data (1 - 900 sec, default = 10).
        /// </summary>
        [Input("comfortInterval")]
        public Input<int>? ComfortInterval { get; set; }

        /// <summary>
        /// Enable/disable Fortinet bar on HTML content. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("fortinetBar")]
        public Input<string>? FortinetBar { get; set; }

        /// <summary>
        /// Port for use by Fortinet Bar (1 - 65535, default = 8011).
        /// </summary>
        [Input("fortinetBarPort")]
        public Input<int>? FortinetBarPort { get; set; }

        /// <summary>
        /// Enable/disable h2c HTTP connection upgrade. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("h2c")]
        public Input<string>? H2c { get; set; }

        /// <summary>
        /// Enable/disable HTTP policy check. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("httpPolicy")]
        public Input<string>? HttpPolicy { get; set; }

        /// <summary>
        /// Enable/disable the inspection of all ports for the protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("inspectAll")]
        public Input<string>? InspectAll { get; set; }

        /// <summary>
        /// One or more options that can be applied to the session. Valid values: `clientcomfort`, `servercomfort`, `oversize`, `chunkedbypass`.
        /// </summary>
        [Input("options")]
        public Input<string>? Options { get; set; }

        /// <summary>
        /// Maximum in-memory file size that can be scanned (1 - 383 MB, default = 10).
        /// </summary>
        [Input("oversizeLimit")]
        public Input<int>? OversizeLimit { get; set; }

        /// <summary>
        /// Ports to scan for content (1 - 65535, default = 80).
        /// </summary>
        [Input("ports")]
        public Input<int>? Ports { get; set; }

        /// <summary>
        /// ID codes for character sets to be used to convert to UTF-8 for banned words and DLP on HTTP posts (maximum of 5 character sets). Valid values: `jisx0201`, `jisx0208`, `jisx0212`, `gb2312`, `ksc5601-ex`, `euc-jp`, `sjis`, `iso2022-jp`, `iso2022-jp-1`, `iso2022-jp-2`, `euc-cn`, `ces-gbk`, `hz`, `ces-big5`, `euc-kr`, `iso2022-jp-3`, `iso8859-1`, `tis620`, `cp874`, `cp1252`, `cp1251`.
        /// </summary>
        [Input("postLang")]
        public Input<string>? PostLang { get; set; }

        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("proxyAfterTcpHandshake")]
        public Input<string>? ProxyAfterTcpHandshake { get; set; }

        /// <summary>
        /// Enable/disable blocking of partial downloads. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("rangeBlock")]
        public Input<string>? RangeBlock { get; set; }

        /// <summary>
        /// Number of attempts to retry HTTP connection (0 - 100, default = 0).
        /// </summary>
        [Input("retryCount")]
        public Input<int>? RetryCount { get; set; }

        /// <summary>
        /// Enable/disable scanning of BZip2 compressed files. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("scanBzip2")]
        public Input<string>? ScanBzip2 { get; set; }

        /// <summary>
        /// SSL decryption and encryption performed by an external device. Valid values: `no`, `yes`.
        /// </summary>
        [Input("sslOffloaded")]
        public Input<string>? SslOffloaded { get; set; }

        /// <summary>
        /// Enable/disable the active status of scanning for this protocol. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Maximum stream-based uncompressed data size that will be scanned (MB, 0 = unlimited (default).  Stream-based uncompression used only under certain conditions.).
        /// </summary>
        [Input("streamBasedUncompressedLimit")]
        public Input<int>? StreamBasedUncompressedLimit { get; set; }

        /// <summary>
        /// Enable/disable bypassing of streaming content from buffering. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("streamingContentBypass")]
        public Input<string>? StreamingContentBypass { get; set; }

        /// <summary>
        /// Enable/disable stripping of HTTP X-Forwarded-For header. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("stripXForwardedFor")]
        public Input<string>? StripXForwardedFor { get; set; }

        /// <summary>
        /// Bypass from scanning, or block a connection that attempts to switch protocol. Valid values: `bypass`, `block`.
        /// </summary>
        [Input("switchingProtocols")]
        public Input<string>? SwitchingProtocols { get; set; }

        /// <summary>
        /// Maximum dynamic TCP window size (default = 8MB).
        /// </summary>
        [Input("tcpWindowMaximum")]
        public Input<int>? TcpWindowMaximum { get; set; }

        /// <summary>
        /// Minimum dynamic TCP window size (default = 128KB).
        /// </summary>
        [Input("tcpWindowMinimum")]
        public Input<int>? TcpWindowMinimum { get; set; }

        /// <summary>
        /// Set TCP static window size (default = 256KB).
        /// </summary>
        [Input("tcpWindowSize")]
        public Input<int>? TcpWindowSize { get; set; }

        /// <summary>
        /// Specify type of TCP window to use for this protocol.
        /// </summary>
        [Input("tcpWindowType")]
        public Input<string>? TcpWindowType { get; set; }

        /// <summary>
        /// Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("tunnelNonHttp")]
        public Input<string>? TunnelNonHttp { get; set; }

        /// <summary>
        /// Maximum nested levels of compression that can be uncompressed and scanned (2 - 100, default = 12).
        /// </summary>
        [Input("uncompressedNestLimit")]
        public Input<int>? UncompressedNestLimit { get; set; }

        /// <summary>
        /// Maximum in-memory uncompressed file size that can be scanned (0 - 383 MB, 0 = unlimited, default = 10).
        /// </summary>
        [Input("uncompressedOversizeLimit")]
        public Input<int>? UncompressedOversizeLimit { get; set; }

        /// <summary>
        /// How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject`, `tunnel`, `best-effort`.
        /// </summary>
        [Input("unknownHttpVersion")]
        public Input<string>? UnknownHttpVersion { get; set; }

        /// <summary>
        /// Enable/disable verification of DNS for policy matching. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("verifyDnsForPolicyMatching")]
        public Input<string>? VerifyDnsForPolicyMatching { get; set; }

        public FirewallProfileprotocoloptionsHttpArgs()
        {
        }
        public static new FirewallProfileprotocoloptionsHttpArgs Empty => new FirewallProfileprotocoloptionsHttpArgs();
    }
}
