// This file is auto-generated please do not edit

package buildkite
import (
    "encoding/json"
    "fmt"
)
type stepBuilder struct {
	Steps []interface{} `yaml:"steps"`
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{}
}
type CommandStepArgs struct {
    // An array of plugins for this step.
    Plugins PluginMap `json:"plugins,omitempty"`

    // The branch pattern defining which branches will include this step in their builds.
    Branches string `json:"branches,omitempty"`

    // The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.
    Concurrency int `json:"concurrency,omitempty"`

    // A map of environment variables for this step.
    Env map[string]string `json:"env,omitempty"`

    // A unique string to identify the step. The value is available in the BUILDKITE_STEP_KEY environment variable.
    // Keys can not have the same pattern as a UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
    Key string `json:"key,omitempty"`

    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    Label string `json:"label,omitempty"`

    // A map of agent tag keys to values to target specific agents for this step.
    Agents map[string]string `json:"agents,omitempty"`

    // Whether to skip this step or not. Passing a string (with a 70-character limit) provides a reason for skipping this command. Passing an empty string is equivalent to false. Note: Skipped steps will be hidden in the pipeline view by default, but can be made visible by toggling the 'Skipped jobs' icon.
    Skip string `json:"skip,omitempty"`

    // The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.
    TimeoutInMinutes int `json:"timeout_in_minutes,omitempty"`

    // The shell command to run during this step.
    Commands []string `json:"commands,omitempty"`

    // The glob path or paths of artifacts to upload from this step.
    ArtifactPaths []string `json:"artifact_paths,omitempty"`

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // Allow all non-zero exit statuses not to fail the build.
    SoftFailAll bool `json:"soft_fail_all,omitempty"`

    // The number of parallel jobs that will be created based on this step.
    Parallelism int `json:"parallelism,omitempty"`

    // Adjust the priority for a specific job, as a positive or negative integer.
    Priority int `json:"priority,omitempty"`

    // The conditions for retrying this step.
    Retry RetryOptions `json:"retry,omitempty"`

    // Whether to continue to run this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // Setting this attribute to true cancels the job as soon as the build is marked as failing.
    CancelOnBuildFailing bool `json:"cancel_on_build_failing,omitempty"`

    // A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.
    ConcurrencyGroup int `json:"concurrency_group,omitempty"`

    // A list of step keys that this step depends on. This step will only run after the named steps have completed. See managing step dependencies for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // An array of values to be used in the matrix expansion.
    Matrix []string `json:"matrix,omitempty"`

    // Allow specified non-zero exit statuses not to fail the build.
    SoftFail []interface{} `json:"soft_fail,omitempty"`

}
func (s *stepBuilder) AddCommandStep(step *CommandStepArgs) *stepBuilder {
    if step == nil {
        s.Steps = append(s.Steps, "")
        return s
    }
    s.Steps = append(s.Steps, step)
    return s
}

type WaitStepArgs struct {
    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    If string `json:"if,omitempty"`

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See [managing step dependencies](https://buildkite.com/docs/pipelines/dependencies) for more information.
    DependsOn []string `json:"depends_on,omitempty"`

    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    AllowDependencyFailure bool `json:"allow_dependency_failure,omitempty"`

    // When providing options for the wait step, you will need to set this value to "~".
    Wait string `json:"wait,omitempty"`

    // Run the next step, even if the previous step has failed.
    ContinueOnFailure bool `json:"continue_on_failure,omitempty"`

}
func (s *stepBuilder) AddWaitStep(step *WaitStepArgs) *stepBuilder {
    if step == nil {
        s.Steps = append(s.Steps, "wait")
        return s
    }
    s.Steps = append(s.Steps, step)
    return s
}

func (s *stepBuilder) Print() error {
	jsonBytes, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
	    return err
	}

    fmt.Println(string(jsonBytes))
    return nil
}
