// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cam
{
    public static class SAMLProviders
    {
        public static Task<SAMLProvidersResult> InvokeAsync(SAMLProvidersArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SAMLProvidersResult>("tctest:Cam/sAMLProviders:SAMLProviders", args ?? new SAMLProvidersArgs(), options.WithDefaults());

        public static Output<SAMLProvidersResult> Invoke(SAMLProvidersInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SAMLProvidersResult>("tctest:Cam/sAMLProviders:SAMLProviders", args ?? new SAMLProvidersInvokeArgs(), options.WithDefaults());
    }


    public sealed class SAMLProvidersArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public SAMLProvidersArgs()
        {
        }
    }

    public sealed class SAMLProvidersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public SAMLProvidersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SAMLProvidersResult
    {
        public readonly string? Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.SAMLProvidersProviderListResult> ProviderLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private SAMLProvidersResult(
            string? description,

            string id,

            string? name,

            ImmutableArray<Outputs.SAMLProvidersProviderListResult> providerLists,

            string? resultOutputFile)
        {
            Description = description;
            Id = id;
            Name = name;
            ProviderLists = providerLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
