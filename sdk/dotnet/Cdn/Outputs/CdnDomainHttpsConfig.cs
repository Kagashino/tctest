// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class CdnDomainHttpsConfig
    {
        public readonly Outputs.CdnDomainHttpsConfigClientCertificateConfig? ClientCertificateConfig;
        public readonly Outputs.CdnDomainHttpsConfigForceRedirect? ForceRedirect;
        public readonly string? Http2Switch;
        public readonly string HttpsSwitch;
        public readonly string? OcspStaplingSwitch;
        public readonly Outputs.CdnDomainHttpsConfigServerCertificateConfig? ServerCertificateConfig;
        public readonly string? SpdySwitch;
        public readonly string? VerifyClient;

        [OutputConstructor]
        private CdnDomainHttpsConfig(
            Outputs.CdnDomainHttpsConfigClientCertificateConfig? clientCertificateConfig,

            Outputs.CdnDomainHttpsConfigForceRedirect? forceRedirect,

            string? http2Switch,

            string httpsSwitch,

            string? ocspStaplingSwitch,

            Outputs.CdnDomainHttpsConfigServerCertificateConfig? serverCertificateConfig,

            string? spdySwitch,

            string? verifyClient)
        {
            ClientCertificateConfig = clientCertificateConfig;
            ForceRedirect = forceRedirect;
            Http2Switch = http2Switch;
            HttpsSwitch = httpsSwitch;
            OcspStaplingSwitch = ocspStaplingSwitch;
            ServerCertificateConfig = serverCertificateConfig;
            SpdySwitch = spdySwitch;
            VerifyClient = verifyClient;
        }
    }
}