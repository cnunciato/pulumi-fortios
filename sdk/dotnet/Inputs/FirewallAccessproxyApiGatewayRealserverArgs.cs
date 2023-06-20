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

    public sealed class FirewallAccessproxyApiGatewayRealserverArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of address. Valid values: `ip`, `fqdn`.
        /// </summary>
        [Input("addrType")]
        public Input<string>? AddrType { get; set; }

        /// <summary>
        /// Address or address group of the real server.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Wildcard domain name of the real server.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Enable to check the responsiveness of the real server before forwarding traffic. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// Protocol of the health check monitor to use when polling to determine server's connectivity status. Valid values: `ping`, `http`, `tcp-connect`.
        /// </summary>
        [Input("healthCheckProto")]
        public Input<string>? HealthCheckProto { get; set; }

        /// <summary>
        /// Enable/disable holddown timer. Server will be considered active and reachable once the holddown period has expired (30 seconds). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("holddownInterval")]
        public Input<string>? HolddownInterval { get; set; }

        /// <summary>
        /// HTTP server domain name in HTTP header.
        /// </summary>
        [Input("httpHost")]
        public Input<string>? HttpHost { get; set; }

        /// <summary>
        /// Real server ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// IPv6 address of the real server.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Port for communicating with the real server.
        /// </summary>
        [Input("mappedport")]
        public Input<string>? Mappedport { get; set; }

        /// <summary>
        /// Port for communicating with the real server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Set access-proxy SSH client certificate profile.
        /// </summary>
        [Input("sshClientCert")]
        public Input<string>? SshClientCert { get; set; }

        /// <summary>
        /// Enable/disable SSH real server host key validation. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("sshHostKeyValidation")]
        public Input<string>? SshHostKeyValidation { get; set; }

        [Input("sshHostKeys")]
        private InputList<Inputs.FirewallAccessproxyApiGatewayRealserverSshHostKeyArgs>? _sshHostKeys;

        /// <summary>
        /// One or more server host key. The structure of `ssh_host_key` block is documented below.
        /// </summary>
        public InputList<Inputs.FirewallAccessproxyApiGatewayRealserverSshHostKeyArgs> SshHostKeys
        {
            get => _sshHostKeys ?? (_sshHostKeys = new InputList<Inputs.FirewallAccessproxyApiGatewayRealserverSshHostKeyArgs>());
            set => _sshHostKeys = value;
        }

        /// <summary>
        /// Set the status of the real server to active so that it can accept traffic, or on standby or disabled so no traffic is sent. Valid values: `active`, `standby`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// TCP forwarding server type. Valid values: `tcp-forwarding`, `ssh`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Weight of the real server. If weighted load balancing is enabled, the server with the highest weight gets more connections.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public FirewallAccessproxyApiGatewayRealserverArgs()
        {
        }
        public static new FirewallAccessproxyApiGatewayRealserverArgs Empty => new FirewallAccessproxyApiGatewayRealserverArgs();
    }
}
