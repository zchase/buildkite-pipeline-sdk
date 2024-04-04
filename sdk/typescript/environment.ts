// This file is auto-generated please do not edit

class Environment {
    // The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).
    public repoMirror(): string {
        return process.env.BUILDKITE_REPO_MIRROR;
    }
    // The name of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public buildCreator(): string {
        return process.env.BUILDKITE_BUILD_CREATOR;
    }
    // The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).
    public commandEval(): boolean {
        return Boolean(process.env.BUILDKITE_COMMAND_EVAL);
    }
    // The url for this build on Buildkite.
    public buildUrl(): string {
        return process.env.BUILDKITE_BUILD_URL;
    }
    // The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).
    public gitSubmodules(): boolean {
        return Boolean(process.env.BUILDKITE_GIT_SUBMODULES);
    }
    // The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public artifactPaths(): string[] {
        return process.env.BUILDKITE_ARTIFACT_PATHS.split(";");
    }
    // The path to the agent config file.
    public configPath(): string {
        return process.env.BUILDKITE_CONFIG_PATH;
    }
    // The ID of the source code provider for the pipeline's repository.
    public pipelineProvider(): string {
        return process.env.BUILDKITE_PIPELINE_PROVIDER;
    }
    // A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public refspec(): string {
        return process.env.BUILDKITE_REFSPEC;
    }
    // The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public artifactUploadDestination(): string {
        return process.env.BUILDKITE_ARTIFACT_UPLOAD_DESTINATION;
    }
    // The number of the build that triggered this build or `""` if the build was not triggered from another build.
    public triggeredFromBuildNumber(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER;
    }
    // Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public s3SseEnabled(): boolean {
        return Boolean(process.env.BUILDKITE_S3_SSE_ENABLED);
    }
    // The path to the directory containing the `buildkite-agent` binary.
    public binPath(): string {
        return process.env.BUILDKITE_BIN_PATH;
    }
    // The label/name of the current job.
    public label(): string {
        return process.env.BUILDKITE_LABEL;
    }
    // The repository URL of the pull request or `""` if not a pull request.
    public pullRequestRepo(): string {
        return process.env.BUILDKITE_PULL_REQUEST_REPO;
    }
    // The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).
    public shell(): string {
        return process.env.BUILDKITE_SHELL;
    }
    // The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.
    public agentEndpoint(): string {
        return process.env.BUILDKITE_AGENT_ENDPOINT;
    }
    // The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as "passed".
    public timeout(): number {
        return Number(process.env.BUILDKITE_TIMEOUT);
    }
    // The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
    public buildAuthorEmail(): string {
        return process.env.BUILDKITE_BUILD_AUTHOR_EMAIL;
    }
    // How many times this job has been retried.
    public retryCount(): number {
        return Number(process.env.BUILDKITE_RETRY_COUNT);
    }
    // The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public githubDeploymentEnvironment(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT;
    }
    // A colon separated list of non-private team slugs that the user who unblocked the build belongs to.
    public unblockerTeams(): string[] {
        return process.env.BUILDKITE_UNBLOCKER_TEAMS.split(":");
    }
    // The command that will be run for the job.
    public command(): string {
        return process.env.BUILDKITE_COMMAND;
    }
    // The notification email of the user who created the build. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public buildCreatorEmail(): string {
        return process.env.BUILDKITE_BUILD_CREATOR_EMAIL;
    }
    // A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.
    public ignoredEnv(): string[] {
        return process.env.BUILDKITE_IGNORED_ENV.split("");
    }
    // The name of the agent that ran the job.
    public agentName(): string {
        return process.env.BUILDKITE_AGENT_NAME;
    }
    // The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the 'buildkite-agent artifact upload' command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the 'environment', 'pre-checkout' or 'pre-command' hooks.
// 
// Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).
    public s3Acl(): string {
        return process.env.BUILDKITE_S3_ACL;
    }
    // The exit code of the last hook that ran, used internally by the hooks.
    public lastHookExitStatus(): number {
        return Number(process.env.BUILDKITE_LAST_HOOK_EXIT_STATUS);
    }
    // The UUID of the original build this was rebuilt from or `""` if not a rebuild.
    public rebuiltFromBuildNumber(): string {
        return process.env.BUILDKITE_REBUILT_FROM_BUILD_NUMBER;
    }
    // The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).
    public localHooksEnabled(): boolean {
        return Boolean(process.env.BUILDKITE_LOCAL_HOOKS_ENABLED);
    }
    // The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.
    public s3SecretAccessKey(): string {
        return process.env.BUILDKITE_S3_SECRET_ACCESS_KEY;
    }
    // The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public githubDeploymentId(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_ID;
    }
    // The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).
    public pluginsPath(): string {
        return process.env.BUILDKITE_PLUGINS_PATH;
    }
    // The value of the debug [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
    public agentDebug(): boolean {
        return Boolean(process.env.BUILDKITE_AGENT_DEBUG);
    }
    // The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public cancelSignal(): string {
        return process.env.BUILDKITE_CANCEL_SIGNAL;
    }
    // A colon separated list of the pipeline's non-private team slugs.
    public pipelineTeams(): string[] {
        return process.env.BUILDKITE_PIPELINE_TEAMS.split(":");
    }
    // Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).
    public pluginValidation(): boolean {
        return Boolean(process.env.BUILDKITE_PLUGIN_VALIDATION);
    }
    // The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public gitCloneFlags(): string {
        return process.env.BUILDKITE_GIT_CLONE_FLAGS;
    }
    // The internal UUID Buildkite uses for this job.
    public jobId(): string {
        return process.env.BUILDKITE_JOB_ID;
    }
    // Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.
    public cleanCheckout(): boolean {
        return Boolean(process.env.BUILDKITE_CLEAN_CHECKOUT);
    }
    // The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).
    public jobLogTmpfile(): string {
        return process.env.BUILDKITE_JOB_LOG_TMPFILE;
    }
    // The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.
    public pluginName(): string {
        return process.env.BUILDKITE_PLUGIN_NAME;
    }
    // The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.
    public clusterId(): string {
        return process.env.BUILDKITE_CLUSTER_ID;
    }
    // The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.
    public triggeredFromBuildId(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_ID;
    }
    // The slug of the pipeline that was used to trigger this build or `""` if the build was not triggered from another build.
    public triggeredFromBuildPipelineSlug(): string {
        return process.env.BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG;
    }
    // The process ID of the agent.
    public agentPid(): number {
        return Number(process.env.BUILDKITE_AGENT_PID);
    }
    // The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public s3DefaultRegion(): string {
        return process.env.BUILDKITE_S3_DEFAULT_REGION;
    }
    // The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.
    public buildAuthor(): string {
        return process.env.BUILDKITE_BUILD_AUTHOR;
    }
    // A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:
// 
// - **Buildkite dashboard:** Set based on who manually created the build.
// - **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
// - **Webhook:** Set based on which user is attached to the API Key used.
    public buildCreatorTeams(): string[] {
        return process.env.BUILDKITE_BUILD_CREATOR_TEAMS.split(":");
    }
    // The exit code from the last command run in the command hook.
    public commandExitStatus(): number {
        return Number(process.env.BUILDKITE_COMMAND_EXIT_STATUS);
    }
    // The path to the file containing the job's environment variables.
    public envFile(): string {
        return process.env.BUILDKITE_ENV_FILE;
    }
    // The value of the `hooks-path` [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).
    public hooksPath(): string {
        return process.env.BUILDKITE_HOOKS_PATH;
    }
    // The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.
    public agentAccessToken(): string {
        return process.env.BUILDKITE_AGENT_ACCESS_TOKEN;
    }
    // The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).
    public agentDisconnectAfterJob(): boolean {
        return Boolean(process.env.BUILDKITE_AGENT_DISCONNECT_AFTER_JOB);
    }
    // The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.
    public branch(): string {
        return process.env.BUILDKITE_BRANCH;
    }
    // The UUID of the user who unblocked the build.
    public unblockerId(): string {
        return process.env.BUILDKITE_UNBLOCKER_ID;
    }
    // The value of the `key` attribute of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public groupKey(): string {
        return process.env.BUILDKITE_GROUP_KEY;
    }
    // The message associated with the build, usually the commit message. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.
    public message(): string {
        return process.env.BUILDKITE_MESSAGE;
    }
    // Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.
    public pullRequestDraft(): boolean {
        return Boolean(process.env.BUILDKITE_PULL_REQUEST_DRAFT);
    }
    // The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public repo(): string {
        return process.env.BUILDKITE_REPO;
    }
    // A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).
    public agentExperiment(): string[] {
        return process.env.BUILDKITE_AGENT_EXPERIMENT.split("");
    }
    // A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).
    public pluginConfiguration(): string {
        return process.env.BUILDKITE_PLUGIN_CONFIGURATION;
    }
    // The UUID of the original build this was rebuilt from or `""` if not a rebuild.
    public rebuiltFromBuildId(): string {
        return process.env.BUILDKITE_REBUILT_FROM_BUILD_ID;
    }
    // The UUID of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public groupId(): string {
        return process.env.BUILDKITE_GROUP_ID;
    }
    // The value of each agent tag. The tag name is appended to the end of the variable name. They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file. The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.
    public agentMetaData(...strs: string[]): string {
        return process.env[strs.join("_").toUpperCase()];
    }
    // The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public githubDeploymentTask(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_TASK;
    }
    // A unique string that identifies a step.
    public stepId(): string {
        return process.env.BUILDKITE_STEP_ID;
    }
    // The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public buildCheckoutPath(): string {
        return process.env.BUILDKITE_BUILD_CHECKOUT_PATH;
    }
    // The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.
    public gitCleanFlags(): string {
        return process.env.BUILDKITE_GIT_CLEAN_FLAGS;
    }
    // The name of the tag being built, if this build was triggered from a tag.
    public tag(): string {
        return process.env.BUILDKITE_TAG;
    }
    // The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.
    public parallelJobCount(): number {
        return Number(process.env.BUILDKITE_PARALLEL_JOB_COUNT);
    }
    // The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.
    public commit(): string {
        return process.env.BUILDKITE_COMMIT;
    }
    // The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).
    public sshKeyscan(): boolean {
        return Boolean(process.env.BUILDKITE_SSH_KEYSCAN);
    }
    // The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.
    public parallelJob(): number {
        return Number(process.env.BUILDKITE_PARALLEL_JOB);
    }
    // The base branch that the pull request is targeting or `""` if not a pull request.
    public pullRequestBaseBranch(): string {
        return process.env.BUILDKITE_PULL_REQUEST_BASE_BRANCH;
    }
    // The displayed pipeline name on Buildkite.
    public pipelineName(): string {
        return process.env.BUILDKITE_PIPELINE_NAME;
    }
    // The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public s3AccessKeyId(): string {
        return process.env.BUILDKITE_S3_ACCESS_KEY_ID;
    }
    // The organization name on Buildkite as used in URLs.
    public organizationSlug(): string {
        return process.env.BUILDKITE_ORGANIZATION_SLUG;
    }
    // The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).
    public githubDeploymentPayload(): string {
        return process.env.BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD;
    }
    // The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.
    public buildNumber(): number {
        return Number(process.env.BUILDKITE_BUILD_NUMBER);
    }
    // The default branch for this pipeline.
    public pipelineDefaultBranch(): string {
        return process.env.BUILDKITE_PIPELINE_DEFAULT_BRANCH;
    }
    // The UUID of the build.
    public buildId(): string {
        return process.env.BUILDKITE_BUILD_ID;
    }
    // A JSON object containing a list plugins used in the step, and their configuration.
    public plugins(): string {
        return process.env.BUILDKITE_PLUGINS;
    }
    // The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).
    public buildPath(): string {
        return process.env.BUILDKITE_BUILD_PATH;
    }
    // The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).
    public pluginsEnabled(): boolean {
        return Boolean(process.env.BUILDKITE_PLUGINS_ENABLED);
    }
    // The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).
    public agentHealthCheckAddr(): string {
        return process.env.BUILDKITE_AGENT_HEALTH_CHECK_ADDR;
    }
    // The name of the user who unblocked the build.
    public unblocker(): string {
        return process.env.BUILDKITE_UNBLOCKER;
    }
    // The pipeline slug on Buildkite as used in URLs.
    public pipelineSlug(): string {
        return process.env.BUILDKITE_PIPELINE_SLUG;
    }
    // Set to "datadog" to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using 'localhost:8126', or 'DD_AGENT_HOST:DD_AGENT_APM_PORT'.
// 
// Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)
    public tracingBackend(): string {
        return process.env.BUILDKITE_TRACING_BACKEND;
    }
    // The label/name of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.
    public groupLabel(): string {
        return process.env.BUILDKITE_GROUP_LABEL;
    }
    // The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.
    public stepKey(): string {
        return process.env.BUILDKITE_STEP_KEY;
    }
    // The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.
    public s3AccessUrl(): string {
        return process.env.BUILDKITE_S3_ACCESS_URL;
    }
    // The number of the pull request or `false` if not a pull request.
    public pullRequest(): number {
        return Number(process.env.BUILDKITE_PULL_REQUEST);
    }
    // The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).
    public agentDisconnectAfterIdleTimeout(): number {
        return Number(process.env.BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT);
    }
    // The UUID of the agent.
    public agentId(): string {
        return process.env.BUILDKITE_AGENT_ID;
    }
    // The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.
    public cancelGracePeriod(): number {
        return Number(process.env.BUILDKITE_CANCEL_GRACE_PERIOD);
    }
    // The source of the event that created the build.
    public source(): string {
        return process.env.BUILDKITE_SOURCE;
    }
    // The notification email of the user who unblocked the build.
    public unblockerEmail(): string {
        return process.env.BUILDKITE_UNBLOCKER_EMAIL;
    }
}


export default Environment;