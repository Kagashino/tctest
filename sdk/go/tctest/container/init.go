// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

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
	case "tctest:Container/cluster:Cluster":
		r = &Cluster{}
	case "tctest:Container/clusterInstance:ClusterInstance":
		r = &ClusterInstance{}
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
		"Container/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tctest",
		"Container/clusterInstance",
		&module{version},
	)
}