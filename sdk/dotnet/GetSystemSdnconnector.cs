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
    public static class GetSystemSdnconnector
    {
        /// <summary>
        /// Use this data source to get information on an fortios system sdnconnector
        /// </summary>
        public static Task<GetSystemSdnconnectorResult> InvokeAsync(GetSystemSdnconnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemSdnconnectorResult>("fortios:index/getSystemSdnconnector:getSystemSdnconnector", args ?? new GetSystemSdnconnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system sdnconnector
        /// </summary>
        public static Output<GetSystemSdnconnectorResult> Invoke(GetSystemSdnconnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemSdnconnectorResult>("fortios:index/getSystemSdnconnector:getSystemSdnconnector", args ?? new GetSystemSdnconnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemSdnconnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system sdnconnector.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemSdnconnectorArgs()
        {
        }
        public static new GetSystemSdnconnectorArgs Empty => new GetSystemSdnconnectorArgs();
    }

    public sealed class GetSystemSdnconnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system sdnconnector.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemSdnconnectorInvokeArgs()
        {
        }
        public static new GetSystemSdnconnectorInvokeArgs Empty => new GetSystemSdnconnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemSdnconnectorResult
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        public readonly string AccessKey;
        /// <summary>
        /// IBM cloud API key or service ID API key.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Azure server region.
        /// </summary>
        public readonly string AzureRegion;
        /// <summary>
        /// Azure client ID (application ID).
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Azure client secret (application key).
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// Compartment ID.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// Compute generation for IBM cloud infrastructure.
        /// </summary>
        public readonly int ComputeGeneration;
        /// <summary>
        /// Domain name.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Configure AWS external account list. The structure of `external_account_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorExternalAccountListResult> ExternalAccountLists;
        /// <summary>
        /// Configure GCP external IP. The structure of `external_ip` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorExternalIpResult> ExternalIps;
        /// <summary>
        /// Configure GCP forwarding rule. The structure of `forwarding_rule` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorForwardingRuleResult> ForwardingRules;
        /// <summary>
        /// GCP project name.
        /// </summary>
        public readonly string GcpProject;
        /// <summary>
        /// Configure GCP project list. The structure of `gcp_project_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorGcpProjectListResult> GcpProjectLists;
        /// <summary>
        /// Group name of computers.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// Enable/disable use for FortiGate HA service.
        /// </summary>
        public readonly string HaStatus;
        /// <summary>
        /// IBM cloud region name.
        /// </summary>
        public readonly string IbmRegion;
        /// <summary>
        /// IBM cloud compute generation 1 region name.
        /// </summary>
        public readonly string IbmRegionGen1;
        /// <summary>
        /// IBM cloud compute generation 2 region name.
        /// </summary>
        public readonly string IbmRegionGen2;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Private key password.
        /// </summary>
        public readonly string KeyPasswd;
        /// <summary>
        /// Azure Stack login endpoint.
        /// </summary>
        public readonly string LoginEndpoint;
        /// <summary>
        /// GCP zone name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configure Azure network interface. The structure of `nic` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorNicResult> Nics;
        /// <summary>
        /// OCI certificate.
        /// </summary>
        public readonly string OciCert;
        /// <summary>
        /// OCI pubkey fingerprint.
        /// </summary>
        public readonly string OciFingerprint;
        /// <summary>
        /// OCI server region.
        /// </summary>
        public readonly string OciRegion;
        /// <summary>
        /// OCI region type.
        /// </summary>
        public readonly string OciRegionType;
        /// <summary>
        /// Password of the remote SDN connector as login credentials.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Private key of GCP service account.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// AWS region name.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Resource group of Azure route table.
        /// </summary>
        public readonly string ResourceGroup;
        /// <summary>
        /// Azure Stack resource URL.
        /// </summary>
        public readonly string ResourceUrl;
        /// <summary>
        /// Configure Azure route table. The structure of `route_table` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorRouteTableResult> RouteTables;
        /// <summary>
        /// Configure Azure route. The structure of `route` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorRouteResult> Routes;
        /// <summary>
        /// AWS secret access key.
        /// </summary>
        public readonly string SecretKey;
        /// <summary>
        /// Secret token of Kubernetes service account.
        /// </summary>
        public readonly string SecretToken;
        /// <summary>
        /// Server address of the remote SDN connector.
        /// </summary>
        public readonly string Server;
        /// <summary>
        /// Server address list of the remote SDN connector. The structure of `server_list` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSystemSdnconnectorServerListResult> ServerLists;
        /// <summary>
        /// Port number of the remote SDN connector.
        /// </summary>
        public readonly int ServerPort;
        /// <summary>
        /// GCP service account email.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Enable/disable connection to the remote SDN connector.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Subscription ID of Azure route table.
        /// </summary>
        public readonly string SubscriptionId;
        /// <summary>
        /// Tenant ID (directory ID).
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Type of SDN connector.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).
        /// </summary>
        public readonly int UpdateInterval;
        /// <summary>
        /// Enable/disable using IAM role from metadata to call API.
        /// </summary>
        public readonly string UseMetadataIam;
        /// <summary>
        /// User ID.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// Username of the remote SDN connector as login credentials.
        /// </summary>
        public readonly string Username;
        /// <summary>
        /// vCenter server password for NSX quarantine.
        /// </summary>
        public readonly string VcenterPassword;
        /// <summary>
        /// vCenter server address for NSX quarantine.
        /// </summary>
        public readonly string VcenterServer;
        /// <summary>
        /// vCenter server username for NSX quarantine.
        /// </summary>
        public readonly string VcenterUsername;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable server certificate verification.
        /// </summary>
        public readonly string VerifyCertificate;
        /// <summary>
        /// AWS VPC ID.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetSystemSdnconnectorResult(
            string accessKey,

            string apiKey,

            string azureRegion,

            string clientId,

            string clientSecret,

            string compartmentId,

            int computeGeneration,

            string domain,

            ImmutableArray<Outputs.GetSystemSdnconnectorExternalAccountListResult> externalAccountLists,

            ImmutableArray<Outputs.GetSystemSdnconnectorExternalIpResult> externalIps,

            ImmutableArray<Outputs.GetSystemSdnconnectorForwardingRuleResult> forwardingRules,

            string gcpProject,

            ImmutableArray<Outputs.GetSystemSdnconnectorGcpProjectListResult> gcpProjectLists,

            string groupName,

            string haStatus,

            string ibmRegion,

            string ibmRegionGen1,

            string ibmRegionGen2,

            string id,

            string keyPasswd,

            string loginEndpoint,

            string name,

            ImmutableArray<Outputs.GetSystemSdnconnectorNicResult> nics,

            string ociCert,

            string ociFingerprint,

            string ociRegion,

            string ociRegionType,

            string password,

            string privateKey,

            string region,

            string resourceGroup,

            string resourceUrl,

            ImmutableArray<Outputs.GetSystemSdnconnectorRouteTableResult> routeTables,

            ImmutableArray<Outputs.GetSystemSdnconnectorRouteResult> routes,

            string secretKey,

            string secretToken,

            string server,

            ImmutableArray<Outputs.GetSystemSdnconnectorServerListResult> serverLists,

            int serverPort,

            string serviceAccount,

            string status,

            string subscriptionId,

            string tenantId,

            string type,

            int updateInterval,

            string useMetadataIam,

            string userId,

            string username,

            string vcenterPassword,

            string vcenterServer,

            string vcenterUsername,

            string? vdomparam,

            string verifyCertificate,

            string vpcId)
        {
            AccessKey = accessKey;
            ApiKey = apiKey;
            AzureRegion = azureRegion;
            ClientId = clientId;
            ClientSecret = clientSecret;
            CompartmentId = compartmentId;
            ComputeGeneration = computeGeneration;
            Domain = domain;
            ExternalAccountLists = externalAccountLists;
            ExternalIps = externalIps;
            ForwardingRules = forwardingRules;
            GcpProject = gcpProject;
            GcpProjectLists = gcpProjectLists;
            GroupName = groupName;
            HaStatus = haStatus;
            IbmRegion = ibmRegion;
            IbmRegionGen1 = ibmRegionGen1;
            IbmRegionGen2 = ibmRegionGen2;
            Id = id;
            KeyPasswd = keyPasswd;
            LoginEndpoint = loginEndpoint;
            Name = name;
            Nics = nics;
            OciCert = ociCert;
            OciFingerprint = ociFingerprint;
            OciRegion = ociRegion;
            OciRegionType = ociRegionType;
            Password = password;
            PrivateKey = privateKey;
            Region = region;
            ResourceGroup = resourceGroup;
            ResourceUrl = resourceUrl;
            RouteTables = routeTables;
            Routes = routes;
            SecretKey = secretKey;
            SecretToken = secretToken;
            Server = server;
            ServerLists = serverLists;
            ServerPort = serverPort;
            ServiceAccount = serviceAccount;
            Status = status;
            SubscriptionId = subscriptionId;
            TenantId = tenantId;
            Type = type;
            UpdateInterval = updateInterval;
            UseMetadataIam = useMetadataIam;
            UserId = userId;
            Username = username;
            VcenterPassword = vcenterPassword;
            VcenterServer = vcenterServer;
            VcenterUsername = vcenterUsername;
            Vdomparam = vdomparam;
            VerifyCertificate = verifyCertificate;
            VpcId = vpcId;
        }
    }
}
