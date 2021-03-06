// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tke

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-tctest/sdk/go/tctest"
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
	case "tctest:Tke/addonAttachment:AddonAttachment":
		r = &AddonAttachment{}
	case "tctest:Tke/asScalingGroup:AsScalingGroup":
		r = &AsScalingGroup{}
	case "tctest:Tke/authAttachment:AuthAttachment":
		r = &AuthAttachment{}
	case "tctest:Tke/cluster:Cluster":
		r = &Cluster{}
	case "tctest:Tke/clusterAttachment:ClusterAttachment":
		r = &ClusterAttachment{}
	case "tctest:Tke/nodePool:NodePool":
		r = &NodePool{}
	case "tctest:Tke/scaleWorker:ScaleWorker":
		r = &ScaleWorker{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := tctest.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/addonAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/asScalingGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/authAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/clusterAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/nodePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Tke/scaleWorker",
		&module{version},
	)
}
