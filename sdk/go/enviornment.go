// This file is auto-generated please do not edit

package buildkite

import (
	"os"
	"strings"
)

type environment struct{}
// The path to the directory containing the `buildkite-agent` binary.
func (e environment) BinPath() string {
    str := os.Getenv("BUILDKITE_BIN_PATH")
    return str
}

// The label/name of the current job.
func (e environment) Label() string {
    str := os.Getenv("BUILDKITE_LABEL")
    return str
}

// The repository URL of the pull request or `""` if not a pull request.
func (e environment) PullRequestRepo() string {
    str := os.Getenv("BUILDKITE_PULL_REQUEST_REPO")
    return str
}

// The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) Shell() string {
    str := os.Getenv("BUILDKITE_SHELL")
    return str
}

// Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
func (e environment) S3SseEnabled() bool {
    str := os.Getenv("BUILDKITE_S3_SSE_ENABLED")
    return ParseStringToBool(str)
}

// The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as "passed".
func (e environment) Timeout() int {
    str := os.Getenv("BUILDKITE_TIMEOUT")
    return ParseStringToInt(str)
}

// The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
func (e environment) BuildAuthorEmail() string {
    str := os.Getenv("BUILDKITE_BUILD_AUTHOR_EMAIL")
    return str
}

// How many times this job has been retried.
func (e environment) RetryCount() int {
    str := os.Getenv("BUILDKITE_RETRY_COUNT")
    return ParseStringToInt(str)
}

// The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
func (e environment) GithubDeploymentEnvironment() string {
    str := os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT")
    return str
}

// A colon separated list of non-private team slugs that the user who unblocked the build belongs to.
func (e environment) UnblockerTeams() string {
    str := os.Getenv("BUILDKITE_UNBLOCKER_TEAMS")
    return str
}

// The command that will be run for the job.
func (e environment) Command() string {
    str := os.Getenv("BUILDKITE_COMMAND")
    return str
}

// The notification email of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
func (e environment) BuildCreatorEmail() string {
    str := os.Getenv("BUILDKITE_BUILD_CREATOR_EMAIL")
    return str
}

// The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.
func (e environment) AgentEndpoint() string {
    str := os.Getenv("BUILDKITE_AGENT_ENDPOINT")
    return str
}

// A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.
func (e environment) IgnoredEnv() string {
    str := os.Getenv("BUILDKITE_IGNORED_ENV")
    return str
}

// The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the 'buildkite-agent artifact upload' command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the 'environment', 'pre-checkout' or 'pre-command' hooks.
// 
// Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).
func (e environment) S3Acl() string {
    str := os.Getenv("BUILDKITE_S3_ACL")
    return str
}

// The exit code of the last hook that ran, used internally by the hooks.
func (e environment) LastHookExitStatus() int {
    str := os.Getenv("BUILDKITE_LAST_HOOK_EXIT_STATUS")
    return ParseStringToInt(str)
}

// The UUID of the original build this was rebuilt from or `""` if not a rebuild.
func (e environment) RebuiltFromBuildNumber() string {
    str := os.Getenv("BUILDKITE_REBUILT_FROM_BUILD_NUMBER")
    return str
}

// The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) LocalHooksEnabled() bool {
    str := os.Getenv("BUILDKITE_LOCAL_HOOKS_ENABLED")
    return ParseStringToBool(str)
}

// The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.
func (e environment) S3SecretAccessKey() string {
    str := os.Getenv("BUILDKITE_S3_SECRET_ACCESS_KEY")
    return str
}

// The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
func (e environment) GithubDeploymentId() string {
    str := os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_ID")
    return str
}

// The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) PluginsPath() string {
    str := os.Getenv("BUILDKITE_PLUGINS_PATH")
    return str
}

// The name of the agent that ran the job.
func (e environment) AgentName() string {
    str := os.Getenv("BUILDKITE_AGENT_NAME")
    return str
}

// The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) CancelSignal() string {
    str := os.Getenv("BUILDKITE_CANCEL_SIGNAL")
    return str
}

// A colon separated list of the pipeline's non-private team slugs.
func (e environment) PipelineTeams() string {
    str := os.Getenv("BUILDKITE_PIPELINE_TEAMS")
    return str
}

// Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) PluginValidation() bool {
    str := os.Getenv("BUILDKITE_PLUGIN_VALIDATION")
    return ParseStringToBool(str)
}

// The value of the debug [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
func (e environment) AgentDebug() bool {
    str := os.Getenv("BUILDKITE_AGENT_DEBUG")
    return ParseStringToBool(str)
}

// The internal UUID Buildkite uses for this job.
func (e environment) JobId() string {
    str := os.Getenv("BUILDKITE_JOB_ID")
    return str
}

// Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.
func (e environment) CleanCheckout() bool {
    str := os.Getenv("BUILDKITE_CLEAN_CHECKOUT")
    return ParseStringToBool(str)
}

// The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) JobLogTmpfile() string {
    str := os.Getenv("BUILDKITE_JOB_LOG_TMPFILE")
    return str
}

// The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.
func (e environment) PluginName() string {
    str := os.Getenv("BUILDKITE_PLUGIN_NAME")
    return str
}

// The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) GitCloneFlags() string {
    str := os.Getenv("BUILDKITE_GIT_CLONE_FLAGS")
    return str
}

// The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.
func (e environment) TriggeredFromBuildId() string {
    str := os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_ID")
    return str
}

// The slug of the pipeline that was used to trigger this build or `""` if the build was not triggered from another build.
func (e environment) TriggeredFromBuildPipelineSlug() string {
    str := os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG")
    return str
}

// The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.
func (e environment) ClusterId() string {
    str := os.Getenv("BUILDKITE_CLUSTER_ID")
    return str
}

// The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
func (e environment) S3DefaultRegion() string {
    str := os.Getenv("BUILDKITE_S3_DEFAULT_REGION")
    return str
}

// The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
func (e environment) BuildAuthor() string {
    str := os.Getenv("BUILDKITE_BUILD_AUTHOR")
    return str
}

// A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
func (e environment) BuildCreatorTeams() string {
    str := os.Getenv("BUILDKITE_BUILD_CREATOR_TEAMS")
    return str
}

// The process ID of the agent.
func (e environment) AgentPid() int {
    str := os.Getenv("BUILDKITE_AGENT_PID")
    return ParseStringToInt(str)
}

// The path to the file containing the job's environment variables.
func (e environment) EnvFile() string {
    str := os.Getenv("BUILDKITE_ENV_FILE")
    return str
}

// The value of the `hooks-path` [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
func (e environment) HooksPath() string {
    str := os.Getenv("BUILDKITE_HOOKS_PATH")
    return str
}

// The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.
func (e environment) AgentAccessToken() string {
    str := os.Getenv("BUILDKITE_AGENT_ACCESS_TOKEN")
    return str
}

// The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) AgentDisconnectAfterJob() bool {
    str := os.Getenv("BUILDKITE_AGENT_DISCONNECT_AFTER_JOB")
    return ParseStringToBool(str)
}

// The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.
func (e environment) Branch() string {
    str := os.Getenv("BUILDKITE_BRANCH")
    return str
}

// The UUID of the user who unblocked the build.
func (e environment) UnblockerId() string {
    str := os.Getenv("BUILDKITE_UNBLOCKER_ID")
    return str
}

// The value of the `key` attribute of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
func (e environment) GroupKey() string {
    str := os.Getenv("BUILDKITE_GROUP_KEY")
    return str
}

// The exit code from the last command run in the command hook.
func (e environment) CommandExitStatus() int {
    str := os.Getenv("BUILDKITE_COMMAND_EXIT_STATUS")
    return ParseStringToInt(str)
}

// The message associated with the build, usually the commit message. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.
func (e environment) Message() string {
    str := os.Getenv("BUILDKITE_MESSAGE")
    return str
}

// Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.
func (e environment) PullRequestDraft() bool {
    str := os.Getenv("BUILDKITE_PULL_REQUEST_DRAFT")
    return ParseStringToBool(str)
}

// A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).
func (e environment) AgentExperiment() string {
    str := os.Getenv("BUILDKITE_AGENT_EXPERIMENT")
    return str
}

// A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).
func (e environment) PluginConfiguration() string {
    str := os.Getenv("BUILDKITE_PLUGIN_CONFIGURATION")
    return str
}

// The UUID of the original build this was rebuilt from or `""` if not a rebuild.
func (e environment) RebuiltFromBuildId() string {
    str := os.Getenv("BUILDKITE_REBUILT_FROM_BUILD_ID")
    return str
}

// The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) Repo() string {
    str := os.Getenv("BUILDKITE_REPO")
    return str
}

// The value of each agent tag. The tag name is appended to the end of the variable name. They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file. The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.
func (e environment) AgentMetaData(strs ...string) string {
    envKey := strings.ToUpper(strings.Join(strs, "_"))
str := os.Getenv(envKey)

    return str
}

// The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
func (e environment) GithubDeploymentTask() string {
    str := os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_TASK")
    return str
}

// A unique string that identifies a step.
func (e environment) StepId() string {
    str := os.Getenv("BUILDKITE_STEP_ID")
    return str
}

// The UUID of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
func (e environment) GroupId() string {
    str := os.Getenv("BUILDKITE_GROUP_ID")
    return str
}

// The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) GitCleanFlags() string {
    str := os.Getenv("BUILDKITE_GIT_CLEAN_FLAGS")
    return str
}

// The name of the tag being built, if this build was triggered from a tag.
func (e environment) Tag() string {
    str := os.Getenv("BUILDKITE_TAG")
    return str
}

// The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.
func (e environment) ParallelJobCount() int {
    str := os.Getenv("BUILDKITE_PARALLEL_JOB_COUNT")
    return ParseStringToInt(str)
}

// The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.
func (e environment) Commit() string {
    str := os.Getenv("BUILDKITE_COMMIT")
    return str
}

// The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) BuildCheckoutPath() string {
    str := os.Getenv("BUILDKITE_BUILD_CHECKOUT_PATH")
    return str
}

// The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.
func (e environment) ParallelJob() int {
    str := os.Getenv("BUILDKITE_PARALLEL_JOB")
    return ParseStringToInt(str)
}

// The base branch that the pull request is targeting or `""` if not a pull request.
func (e environment) PullRequestBaseBranch() string {
    str := os.Getenv("BUILDKITE_PULL_REQUEST_BASE_BRANCH")
    return str
}

// The displayed pipeline name on Buildkite.
func (e environment) PipelineName() string {
    str := os.Getenv("BUILDKITE_PIPELINE_NAME")
    return str
}

// The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
func (e environment) S3AccessKeyId() string {
    str := os.Getenv("BUILDKITE_S3_ACCESS_KEY_ID")
    return str
}

// The organization name on Buildkite as used in URLs.
func (e environment) OrganizationSlug() string {
    str := os.Getenv("BUILDKITE_ORGANIZATION_SLUG")
    return str
}

// The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
func (e environment) GithubDeploymentPayload() string {
    str := os.Getenv("BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD")
    return str
}

// The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.
func (e environment) BuildNumber() int {
    str := os.Getenv("BUILDKITE_BUILD_NUMBER")
    return ParseStringToInt(str)
}

// The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) SshKeyscan() bool {
    str := os.Getenv("BUILDKITE_SSH_KEYSCAN")
    return ParseStringToBool(str)
}

// The UUID of the build.
func (e environment) BuildId() string {
    str := os.Getenv("BUILDKITE_BUILD_ID")
    return str
}

// A JSON object containing a list plugins used in the step, and their configuration.
func (e environment) Plugins() string {
    str := os.Getenv("BUILDKITE_PLUGINS")
    return str
}

// The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) BuildPath() string {
    str := os.Getenv("BUILDKITE_BUILD_PATH")
    return str
}

// The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) PluginsEnabled() bool {
    str := os.Getenv("BUILDKITE_PLUGINS_ENABLED")
    return ParseStringToBool(str)
}

// The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) AgentHealthCheckAddr() string {
    str := os.Getenv("BUILDKITE_AGENT_HEALTH_CHECK_ADDR")
    return str
}

// The default branch for this pipeline.
func (e environment) PipelineDefaultBranch() string {
    str := os.Getenv("BUILDKITE_PIPELINE_DEFAULT_BRANCH")
    return str
}

// The pipeline slug on Buildkite as used in URLs.
func (e environment) PipelineSlug() string {
    str := os.Getenv("BUILDKITE_PIPELINE_SLUG")
    return str
}

// Set to "datadog" to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using 'localhost:8126', or 'DD_AGENT_HOST:DD_AGENT_APM_PORT'.
// 
// Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)
func (e environment) TracingBackend() string {
    str := os.Getenv("BUILDKITE_TRACING_BACKEND")
    return str
}

// The label/name of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
func (e environment) GroupLabel() string {
    str := os.Getenv("BUILDKITE_GROUP_LABEL")
    return str
}

// The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.
func (e environment) StepKey() string {
    str := os.Getenv("BUILDKITE_STEP_KEY")
    return str
}

// The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
func (e environment) S3AccessUrl() string {
    str := os.Getenv("BUILDKITE_S3_ACCESS_URL")
    return str
}

// The name of the user who unblocked the build.
func (e environment) Unblocker() string {
    str := os.Getenv("BUILDKITE_UNBLOCKER")
    return str
}

// The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) AgentDisconnectAfterIdleTimeout() int {
    str := os.Getenv("BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT")
    return ParseStringToInt(str)
}

// The UUID of the agent.
func (e environment) AgentId() string {
    str := os.Getenv("BUILDKITE_AGENT_ID")
    return str
}

// The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.
func (e environment) CancelGracePeriod() int {
    str := os.Getenv("BUILDKITE_CANCEL_GRACE_PERIOD")
    return ParseStringToInt(str)
}

// The source of the event that created the build.
func (e environment) Source() string {
    str := os.Getenv("BUILDKITE_SOURCE")
    return str
}

// The notification email of the user who unblocked the build.
func (e environment) UnblockerEmail() string {
    str := os.Getenv("BUILDKITE_UNBLOCKER_EMAIL")
    return str
}

// The number of the pull request or `false` if not a pull request.
func (e environment) PullRequest() int {
    str := os.Getenv("BUILDKITE_PULL_REQUEST")
    return ParseStringToInt(str)
}

// The name of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
func (e environment) BuildCreator() string {
    str := os.Getenv("BUILDKITE_BUILD_CREATOR")
    return str
}

// The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) CommandEval() bool {
    str := os.Getenv("BUILDKITE_COMMAND_EVAL")
    return ParseStringToBool(str)
}

// The url for this build on Buildkite.
func (e environment) BuildUrl() string {
    str := os.Getenv("BUILDKITE_BUILD_URL")
    return str
}

// The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).
func (e environment) GitSubmodules() bool {
    str := os.Getenv("BUILDKITE_GIT_SUBMODULES")
    return ParseStringToBool(str)
}

// The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).
func (e environment) RepoMirror() string {
    str := os.Getenv("BUILDKITE_REPO_MIRROR")
    return str
}

// The path to the agent config file.
func (e environment) ConfigPath() string {
    str := os.Getenv("BUILDKITE_CONFIG_PATH")
    return str
}

// The ID of the source code provider for the pipeline's repository.
func (e environment) PipelineProvider() string {
    str := os.Getenv("BUILDKITE_PIPELINE_PROVIDER")
    return str
}

// A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) Refspec() string {
    str := os.Getenv("BUILDKITE_REFSPEC")
    return str
}

// The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
func (e environment) ArtifactUploadDestination() string {
    str := os.Getenv("BUILDKITE_ARTIFACT_UPLOAD_DESTINATION")
    return str
}

// The number of the build that triggered this build or `""` if the build was not triggered from another build.
func (e environment) TriggeredFromBuildNumber() string {
    str := os.Getenv("BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER")
    return str
}

// The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
func (e environment) ArtifactPaths() string {
    str := os.Getenv("BUILDKITE_ARTIFACT_PATHS")
    return str
}
