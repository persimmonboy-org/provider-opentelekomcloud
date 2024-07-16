/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/persimmonboy-org/provider-opentelekomcloud/config/blockstorage"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/cce"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/compute"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/dcs"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/deh"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/dis"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/dms"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/dns"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/fg"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/fw"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/identity"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/image"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/lb"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/nat"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/networking"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/obs"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/rds"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/sfs"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/smn"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/vpcep"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/vpnaas"
	"github.com/persimmonboy-org/provider-opentelekomcloud/config/wafd"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "opentelekomcloud"
	modulePath     = "github.com/persimmonboy-org/provider-opentelekomcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("opentelekomcloud.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		blockstorage.Configure,
		cce.Configure,
		compute.Configure,
		dcs.Configure,
		deh.Configure,
		dis.Configure,
		dms.Configure,
		dns.Configure,
		fg.Configure,
		fw.Configure,
		identity.Configure,
		image.Configure,
		lb.Configure,
		nat.Configure,
		networking.Configure,
		obs.Configure,
		rds.Configure,
		sfs.Configure,
		smn.Configure,
		vpcep.Configure,
		vpnaas.Configure,
		wafd.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
