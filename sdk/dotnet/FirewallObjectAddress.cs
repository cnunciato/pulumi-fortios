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
    /// <summary>
    /// Provides a resource to configure firewall addresses used in firewall policies of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.FirewallAddress`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// ### Iprange Address
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s1 = new Fortios.FirewallObjectAddress("s1", new()
    ///     {
    ///         Comment = "dd",
    ///         EndIp = "2.0.0.0",
    ///         StartIp = "1.0.0.0",
    ///         Type = "iprange",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Geography Address
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s2 = new Fortios.FirewallObjectAddress("s2", new()
    ///     {
    ///         Comment = "dd",
    ///         Country = "AO",
    ///         Type = "geography",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Fqdn Address
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s3 = new Fortios.FirewallObjectAddress("s3", new()
    ///     {
    ///         AssociatedInterface = "port4",
    ///         Comment = "dd",
    ///         Fqdn = "baid.com",
    ///         ShowInAddressList = "disable",
    ///         StaticRouteConfigure = "enable",
    ///         Type = "fqdn",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Ipmask Address
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var s4 = new Fortios.FirewallObjectAddress("s4", new()
    ///     {
    ///         Comment = "dd",
    ///         Subnet = "0.0.0.0 0.0.0.0",
    ///         Type = "ipmask",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/firewallObjectAddress:FirewallObjectAddress")]
    public partial class FirewallObjectAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Output("associatedInterface")]
        public Output<string> AssociatedInterface { get; private set; } = null!;

        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Output("country")]
        public Output<string> Country { get; private set; } = null!;

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("endIp")]
        public Output<string> EndIp { get; private set; } = null!;

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Output("fqdn")]
        public Output<string> Fqdn { get; private set; } = null!;

        /// <summary>
        /// Address name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable address visibility in the GUI. default is "enable".
        /// </summary>
        [Output("showInAddressList")]
        public Output<string> ShowInAddressList { get; private set; } = null!;

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Output("startIp")]
        public Output<string> StartIp { get; private set; } = null!;

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Output("staticRouteConfigure")]
        public Output<string> StaticRouteConfigure { get; private set; } = null!;

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// Type of address(Support ipmask, iprange, fqdn and geography).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallObjectAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallObjectAddress(string name, FirewallObjectAddressArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/firewallObjectAddress:FirewallObjectAddress", name, args ?? new FirewallObjectAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallObjectAddress(string name, Input<string> id, FirewallObjectAddressState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallObjectAddress:FirewallObjectAddress", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallObjectAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallObjectAddress Get(string name, Input<string> id, FirewallObjectAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallObjectAddress(name, id, state, options);
        }
    }

    public sealed class FirewallObjectAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable address visibility in the GUI. default is "enable".
        /// </summary>
        [Input("showInAddressList")]
        public Input<string>? ShowInAddressList { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Input("staticRouteConfigure")]
        public Input<string>? StaticRouteConfigure { get; set; }

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Type of address(Support ipmask, iprange, fqdn and geography).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FirewallObjectAddressArgs()
        {
        }
        public static new FirewallObjectAddressArgs Empty => new FirewallObjectAddressArgs();
    }

    public sealed class FirewallObjectAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network interface associated with address.
        /// </summary>
        [Input("associatedInterface")]
        public Input<string>? AssociatedInterface { get; set; }

        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// IP addresses associated to a specific country.
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Final IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("endIp")]
        public Input<string>? EndIp { get; set; }

        /// <summary>
        /// Fully Qualified Domain Name address.
        /// </summary>
        [Input("fqdn")]
        public Input<string>? Fqdn { get; set; }

        /// <summary>
        /// Address name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable address visibility in the GUI. default is "enable".
        /// </summary>
        [Input("showInAddressList")]
        public Input<string>? ShowInAddressList { get; set; }

        /// <summary>
        /// First IP address (inclusive) in the range for the address.
        /// </summary>
        [Input("startIp")]
        public Input<string>? StartIp { get; set; }

        /// <summary>
        /// Enable/disable use of this address in the static route configuration. default is "disable".
        /// </summary>
        [Input("staticRouteConfigure")]
        public Input<string>? StaticRouteConfigure { get; set; }

        /// <summary>
        /// IP address and subnet mask of address.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Type of address(Support ipmask, iprange, fqdn and geography).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public FirewallObjectAddressState()
        {
        }
        public static new FirewallObjectAddressState Empty => new FirewallObjectAddressState();
    }
}
