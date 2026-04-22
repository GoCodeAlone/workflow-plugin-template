package internal

import (
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// Version is set at build time via -ldflags
// "-X github.com/GoCodeAlone/workflow-plugin-TEMPLATE/internal.Version=X.Y.Z".
// Keeping it as a package-level var (rather than baking the string into the
// Manifest literal) lets goreleaser inject the real release tag so the
// workflow engine's requires.plugins[].version check passes.
var Version = "dev"

// Manifest returns the plugin metadata used by the workflow engine for
// discovery and capability negotiation.
var Manifest = sdk.PluginManifest{
	Name:        "workflow-plugin-TEMPLATE",
	Description: "TEMPLATE plugin for the workflow engine",
	Author:      "GoCodeAlone",
	License:     "MIT",
	ModuleTypes: []string{
		// "example.module_type",
	},
	StepTypes: []string{
		// "step.example_action",
	},
}

type plugin struct{}

// NewPlugin creates a new plugin instance.
func NewPlugin() sdk.PluginProvider {
	return &plugin{}
}

func (p *plugin) Manifest() sdk.PluginManifest {
	m := Manifest
	m.Version = Version
	return m
}

func (p *plugin) ModuleFactories() map[string]sdk.ModuleFactory {
	return map[string]sdk.ModuleFactory{
		// "example.module_type": NewExampleModuleFactory(),
	}
}

func (p *plugin) StepFactories() map[string]sdk.StepFactory {
	return map[string]sdk.StepFactory{
		// "step.example_action": NewExampleStepFactory(),
	}
}
