// SPDX-License-Identifier: Apache-2.0

package constants

// Repo get pipeline types.
const (
	// PipelineTypeYAML defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as yaml.
	PipelineTypeYAML = "yaml"

	// PipelineTypeGo defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as Go templates.
	PipelineTypeGo = "go"

	// PipelineTypeStarlark defines the pipeline type for allowing users
	// in Vela to control their pipeline being compiled as Starlark templates.
	PipelineTypeStarlark = "starlark"
)
