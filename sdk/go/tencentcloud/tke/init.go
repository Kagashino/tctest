// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tke

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tencentcloud:Tke/addonAttachment:AddonAttachment":
		r = &AddonAttachment{}
	case "tencentcloud:Tke/asScalingGroup:AsScalingGroup":
		r = &AsScalingGroup{}
	case "tencentcloud:Tke/authAttachment:AuthAttachment":
		r = &AuthAttachment{}
	case "tencentcloud:Tke/cluster:Cluster":
		r = &Cluster{}
	case "tencentcloud:Tke/clusterAttachment:ClusterAttachment":
		r = &ClusterAttachment{}
	case "tencentcloud:Tke/nodePool:NodePool":
		r = &NodePool{}
	case "tencentcloud:Tke/scaleWorker:ScaleWorker":
		r = &ScaleWorker{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := tencentcloud.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/addonAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/asScalingGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/authAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/clusterAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/nodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tencentcloud",
		"Tke/scaleWorker",
		&module{version},
	)
}