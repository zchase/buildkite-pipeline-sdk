// This file is auto-generated please do not edit

package buildkite

// Allow specified non-zero exit statuses not to fail the build.
type SoftFailOption map[string]int

type RetryOptions struct {
    // Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.
    Automatic bool `json:"automatic,omitempty"`

    // Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.
    Manual bool `json:"manual,omitempty"`

}
// A generic string map.
type StringRecord map[string]string

// An array of plugins for this step.
type PluginMap map[string]StringRecord
