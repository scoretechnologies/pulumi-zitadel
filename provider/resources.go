// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zitadel

import (
	"context"
	_ "embed"
	"fmt"
	"path/filepath"

	pftfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/scoretechnologies/pulumi-zitadel/provider/pkg/version"
	"github.com/zitadel/terraform-provider-zitadel/v2/zitadel"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "zitadel"
	// modules:
	mainMod = "index" // the zitadel module
)

//go:embed cmd/pulumi-resource-zitadel/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	ctx := context.Background()

	// Upstream serves its resources from two implementations muxed together: most of
	// them come from the Terraform Plugin SDKv2 provider, while the message text and
	// login text resources come from the Plugin Framework provider. Both have to be
	// bridged for the Pulumi provider to expose the full surface.
	//
	// The Plugin Framework provider also declares ephemeral resources. Those have no
	// Pulumi equivalent and are dropped by the bridge.
	p := pftfbridge.MuxShimWithPF(ctx,
		shimv2.NewProvider(zitadel.Provider()),
		zitadel.NewProviderPV6(),
	)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:       p,
		Name:    "zitadel",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "scoretechnologies",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/scoretechnologies",
		Description:       "A Pulumi package for creating and managing zitadel cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "zitadel", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/scoretechnologies/pulumi-zitadel",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "zitadel",
		// Upstream is released as v3.x but its go.mod still declares the /v2 module
		// path, so the module suffix stays "v2".
		TFProviderModuleVersion: "v2",
		// Muxed providers need the dispatch table that tfgen writes into
		// bridge-metadata.json to know which side owns each token.
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@scoretechnologies/zitadel",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},

			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "scoretechnologies_zitadel",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/scoretechnologies/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "scoretechnologies",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	// Derive every Pulumi token from its Terraform name. This reproduces the tokens
	// that used to be listed here by hand, so nothing is renamed, and resources added
	// upstream are picked up without editing this file.
	prov.MustComputeTokens(tokens.SingleModule("zitadel_", mainMod, tokens.MakeStandard(mainPkg)))

	return prov
}
