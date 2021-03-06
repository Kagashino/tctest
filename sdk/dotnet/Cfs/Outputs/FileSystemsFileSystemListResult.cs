// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cfs.Outputs
{

    [OutputType]
    public sealed class FileSystemsFileSystemListResult
    {
        public readonly string AccessGroupId;
        public readonly string AvailabilityZone;
        public readonly string CreateTime;
        public readonly string FileSystemId;
        public readonly string Name;
        public readonly string Protocol;
        public readonly int SizeLimit;
        public readonly int SizeUsed;
        public readonly string Status;
        public readonly string StorageType;

        [OutputConstructor]
        private FileSystemsFileSystemListResult(
            string accessGroupId,

            string availabilityZone,

            string createTime,

            string fileSystemId,

            string name,

            string protocol,

            int sizeLimit,

            int sizeUsed,

            string status,

            string storageType)
        {
            AccessGroupId = accessGroupId;
            AvailabilityZone = availabilityZone;
            CreateTime = createTime;
            FileSystemId = fileSystemId;
            Name = name;
            Protocol = protocol;
            SizeLimit = sizeLimit;
            SizeUsed = sizeUsed;
            Status = status;
            StorageType = storageType;
        }
    }
}
