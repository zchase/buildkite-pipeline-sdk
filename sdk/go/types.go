// This file is auto-generated please do not edit

package buildkite

// A generic string map.
type StringRecord map[string]string

// An array of plugins for this step.
type PluginMap map[string]StringRecord

// Allow specified non-zero exit statuses not to fail the build.
type SoftFailOption map[string]int

// The conditions for retrying this step.
type RetryOptions string
const (
    AUTOMATIC RetryOptions = "automatic"
    MANUAL RetryOptions = "manual"
)