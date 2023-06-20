// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class FirewallSslsshprofilePop3s
    {
        /// <summary>
        /// Action based on certificate validation failure. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        public readonly string? CertValidationFailure;
        /// <summary>
        /// Action based on certificate validation timeout. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        public readonly string? CertValidationTimeout;
        /// <summary>
        /// Action based on client certificate request. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        public readonly string? ClientCertRequest;
        /// <summary>
        /// Action based on received client certificate. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        public readonly string? ClientCertificate;
        /// <summary>
        /// Action based on server certificate is expired. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        public readonly string? ExpiredServerCert;
        /// <summary>
        /// Allow or block the invalid SSL session server certificate. Valid values: `allow`, `block`.
        /// </summary>
        public readonly string? InvalidServerCert;
        /// <summary>
        /// Ports to use for scanning (1 - 65535, default = 443).
        /// </summary>
        public readonly string? Ports;
        /// <summary>
        /// Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `enable`, `disable`.
        /// </summary>
        public readonly string? ProxyAfterTcpHandshake;
        /// <summary>
        /// Action based on server certificate is revoked. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        public readonly string? RevokedServerCert;
        /// <summary>
        /// Check the SNI in the client hello message with the CN or SAN fields in the returned server certificate. Valid values: `enable`, `strict`, `disable`.
        /// 
        /// The `pop3s` block supports:
        /// </summary>
        public readonly string? SniServerCertCheck;
        /// <summary>
        /// Configure protocol inspection status. Valid values: `disable`, `certificate-inspection`, `deep-inspection`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Action based on the SSL encryption used being unsupported. Valid values: `bypass`, `inspect`, `block`.
        /// </summary>
        public readonly string? UnsupportedSsl;
        /// <summary>
        /// Action based on the SSL cipher used being unsupported. Valid values: `allow`, `block`.
        /// </summary>
        public readonly string? UnsupportedSslCipher;
        /// <summary>
        /// Action based on the SSL negotiation used being unsupported. Valid values: `allow`, `block`.
        /// </summary>
        public readonly string? UnsupportedSslNegotiation;
        /// <summary>
        /// Action based on the SSL version used being unsupported.
        /// </summary>
        public readonly string? UnsupportedSslVersion;
        /// <summary>
        /// Allow, ignore, or block the untrusted SSL session server certificate. Valid values: `allow`, `block`, `ignore`.
        /// </summary>
        public readonly string? UntrustedServerCert;

        [OutputConstructor]
        private FirewallSslsshprofilePop3s(
            string? certValidationFailure,

            string? certValidationTimeout,

            string? clientCertRequest,

            string? clientCertificate,

            string? expiredServerCert,

            string? invalidServerCert,

            string? ports,

            string? proxyAfterTcpHandshake,

            string? revokedServerCert,

            string? sniServerCertCheck,

            string? status,

            string? unsupportedSsl,

            string? unsupportedSslCipher,

            string? unsupportedSslNegotiation,

            string? unsupportedSslVersion,

            string? untrustedServerCert)
        {
            CertValidationFailure = certValidationFailure;
            CertValidationTimeout = certValidationTimeout;
            ClientCertRequest = clientCertRequest;
            ClientCertificate = clientCertificate;
            ExpiredServerCert = expiredServerCert;
            InvalidServerCert = invalidServerCert;
            Ports = ports;
            ProxyAfterTcpHandshake = proxyAfterTcpHandshake;
            RevokedServerCert = revokedServerCert;
            SniServerCertCheck = sniServerCertCheck;
            Status = status;
            UnsupportedSsl = unsupportedSsl;
            UnsupportedSslCipher = unsupportedSslCipher;
            UnsupportedSslNegotiation = unsupportedSslNegotiation;
            UnsupportedSslVersion = unsupportedSslVersion;
            UntrustedServerCert = untrustedServerCert;
        }
    }
}
