// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Ssm
{
    public static class SecretVersions
    {
        public static Task<SecretVersionsResult> InvokeAsync(SecretVersionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<SecretVersionsResult>("tctest:Ssm/secretVersions:SecretVersions", args ?? new SecretVersionsArgs(), options.WithDefaults());

        public static Output<SecretVersionsResult> Invoke(SecretVersionsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<SecretVersionsResult>("tctest:Ssm/secretVersions:SecretVersions", args ?? new SecretVersionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class SecretVersionsArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("secretName", required: true)]
        public string SecretName { get; set; } = null!;

        [Input("versionId")]
        public string? VersionId { get; set; }

        public SecretVersionsArgs()
        {
        }
    }

    public sealed class SecretVersionsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public SecretVersionsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class SecretVersionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string SecretName;
        public readonly ImmutableArray<Outputs.SecretVersionsSecretVersionListResult> SecretVersionLists;
        public readonly string? VersionId;

        [OutputConstructor]
        private SecretVersionsResult(
            string id,

            string? resultOutputFile,

            string secretName,

            ImmutableArray<Outputs.SecretVersionsSecretVersionListResult> secretVersionLists,

            string? versionId)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SecretName = secretName;
            SecretVersionLists = secretVersionLists;
            VersionId = versionId;
        }
    }
}
