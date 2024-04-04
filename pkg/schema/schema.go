package schema

var Schema = PipelineSchema{
	Name: "Buildkite Pipeline Schema",
	Types: NewPropertyMap(map[string]Property{
		"soft_fail_option": {
			Type: schemaObject,
			Items: ObjectArrayProperty{
				Type: schemaNumber,
			},
			Description: "Allow specified non-zero exit statuses not to fail the build.",
		},
		"retry_options": {
			Type:        schemaObject,
			Description: "The conditions for retrying this step.",
			Properties: NewPropertyMap(map[string]Property{
				"automatic": {
					Type:        schemaBoolean,
					Description: "Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.",
				},
				"manual": {
					Type:        schemaBoolean,
					Description: "Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.",
				},
			}),
		},
		"string_record": {
			Type:        schemaObject,
			Description: "A generic string map.",
			Items: ObjectArrayProperty{
				Type: schemaString,
			},
		},
		"plugin_map": {
			Type:        schemaObject,
			Description: "An array of plugins for this step.",
			Items: ObjectArrayProperty{
				TypeRef: "string_record",
			},
		},
	}),
	Environment: NewEnvironmentVariableMap(map[string]EnvironmentVariable{
		"BUILDKITE_AGENT_ACCESS_TOKEN": {
			Type:        schemaString,
			Description: "The agent session token for the job. The variable is read by the agent `artifact` and `meta-data` commands.",
		},
		"BUILDKITE_AGENT_DEBUG": {
			Type:        schemaBoolean,
			Description: "The value of the debug [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).",
		},
		"BUILDKITE_AGENT_DISCONNECT_AFTER_JOB": {
			Type:        schemaBoolean,
			Description: "The value of the `disconnect-after-job` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_AGENT_DISCONNECT_AFTER_IDLE_TIMEOUT": {
			Type:        schemaNumber,
			Description: "The value of the `disconnect-after-idle-timeout` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_AGENT_ENDPOINT": {
			Type:        schemaString,
			Description: "The value of the `endpoint` [agent configuration option](/docs/agent/v3/configuration). This is set as an environment variable by the bootstrap and then read by most of the `buildkite-agent` commands.",
		},
		"BUILDKITE_AGENT_EXPERIMENT": {
			Type:        schemaArray,
			Description: "A list of the [experimental agent features](/docs/agent/v3#experimental-features) that are currently enabled. The value can be set using the `--experiment` flag on the [`buildkite-agent start` command](/docs/agent/v3/cli-start#starting-an-agent) or in your [agent configuration file](/docs/agent/v3/configuration).",
			Items: ObjectArrayProperty{
				Type: schemaString,
			},
		},
		"BUILDKITE_AGENT_HEALTH_CHECK_ADDR": {
			Type:        schemaString,
			Description: "The value of the `health-check-addr` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_AGENT_ID": {
			Type:        schemaString,
			Description: "The UUID of the agent.",
		},
		"BUILDKITE_AGENT_META_DATA": {
			Dynamic:     true,
			Type:        schemaString,
			Description: "The value of each agent tag. The tag name is appended to the end of the variable name. They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file. The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.",
		},
		"BUILDKITE_AGENT_NAME": {
			Type:        schemaString,
			Description: "The name of the agent that ran the job.",
		},
		"BUILDKITE_AGENT_PID": {
			Type:        schemaNumber,
			Description: "The process ID of the agent.",
		},
		"BUILDKITE_ARTIFACT_PATHS": {
			Type:        schemaArray,
			Description: "The artifact paths to upload after the job, if any have been specified. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
			Items: ObjectArrayProperty{
				Type:      schemaString,
				Delimiter: ";",
			},
		},
		"BUILDKITE_ARTIFACT_UPLOAD_DESTINATION": {
			Type:        schemaString,
			Description: "The path where artifacts will be uploaded. This variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). It can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.",
		},
		"BUILDKITE_BIN_PATH": {
			Type:        schemaString,
			Description: "The path to the directory containing the `buildkite-agent` binary.",
		},
		"BUILDKITE_BRANCH": {
			Type:        schemaString,
			Description: "The branch being built. Note that for manually triggered builds, this branch is not guaranteed to contain the commit specified by `BUILDKITE_COMMIT`.",
		},
		"BUILDKITE_BUILD_CHECKOUT_PATH": {
			Type:        schemaString,
			Description: "The path where the agent has checked out your code for this build. This variable is read by the bootstrap when the agent is started, and can only be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_BUILD_AUTHOR": {
			Type:        schemaString,
			Description: "The name of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.",
		},
		"BUILDKITE_BUILD_AUTHOR_EMAIL": {
			Type:        schemaString,
			Description: "The notification email of the user who authored the commit being built. May be **[unverified](#unverified-commits)**. This value can be blank in some situations, including builds manually triggered using API or Buildkite web interface.",
		},
		"BUILDKITE_BUILD_CREATOR": {
			Type: schemaString,
			Description: `The name of the user who created the build. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`,
		},
		"BUILDKITE_BUILD_CREATOR_EMAIL": {
			Type: schemaString,
			Description: `The notification email of the user who created the build. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`,
		},
		"BUILDKITE_BUILD_CREATOR_TEAMS": {
			Type: schemaArray,
			Items: ObjectArrayProperty{
				Type:      schemaString,
				Delimiter: ":",
			},
			Description: `A colon separated list of non-private team slugs that the build creator belongs to. The value differs depending on how the build was created:

- **Buildkite dashboard:** Set based on who manually created the build.
- **GitHub webhook:** Set from the  **[unverified](#unverified-commits)** HEAD commit.
- **Webhook:** Set based on which user is attached to the API Key used.`,
		},
		"BUILDKITE_BUILD_ID": {
			Type:        schemaString,
			Description: "The UUID of the build.",
		},
		"BUILDKITE_BUILD_NUMBER": {
			Type:        schemaNumber,
			Description: "The build number. This number increases by 1 with every build, and is guaranteed to be unique within each pipeline.",
		},
		"BUILDKITE_BUILD_PATH": {
			Type:        schemaString,
			Description: "The value of the `build-path` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_BUILD_URL": {
			Type:        schemaString,
			Description: "The url for this build on Buildkite.",
		},
		"BUILDKITE_CANCEL_GRACE_PERIOD": {
			Type:        schemaNumber,
			Description: "The value of the `cancel-grace-period` [agent configuration option](/docs/agent/v3/configuration) in seconds.",
		},
		"BUILDKITE_CANCEL_SIGNAL": {
			Type:        schemaString,
			Description: "The value of the `cancel-signal` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_CLEAN_CHECKOUT": {
			Type:        schemaBoolean,
			Description: "Whether the build should perform a clean checkout. The variable is read during the default checkout phase of the bootstrap and can be overridden in `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_CLUSTER_ID": {
			Type:        schemaString,
			Description: "The UUID value of the cluster, but only if the build has an associated `cluster_queue`. Otherwise, this environment variable is not set.",
		},
		"BUILDKITE_COMMAND": {
			Type:        schemaString,
			Description: "The command that will be run for the job.",
		},
		"BUILDKITE_COMMAND_EVAL": {
			Type:        schemaBoolean,
			Description: "The opposite of the value of the `no-command-eval` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_COMMAND_EXIT_STATUS": {
			Type:        schemaNumber,
			Description: "The exit code from the last command run in the command hook.",
		},
		"BUILDKITE_COMMIT": {
			Type:        schemaString,
			Description: "The git commit object of the build. This is usually a 40-byte hexadecimal SHA-1 hash, but can also be a symbolic name like `HEAD`.",
		},
		"BUILDKITE_CONFIG_PATH": {
			Type:        schemaString,
			Description: "The path to the agent config file.",
		},
		"BUILDKITE_ENV_FILE": {
			Type:        schemaString,
			Description: "The path to the file containing the job's environment variables.",
		},
		"BUILDKITE_GIT_CLEAN_FLAGS": {
			Type:        schemaString,
			Description: "The value of the `git-clean-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_GIT_CLONE_FLAGS": {
			Type:        schemaString,
			Description: "The value of the `git-clone-flags` [agent configuration option](/docs/agent/v3/configuration). The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_GIT_SUBMODULES": {
			Type:        schemaBoolean,
			Description: "The opposite of the value of the `no-git-submodules` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_GITHUB_DEPLOYMENT_ID": {
			Type:        schemaString,
			Description: "The GitHub deployment ID. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).",
		},
		"BUILDKITE_GITHUB_DEPLOYMENT_ENVIRONMENT": {
			Type:        schemaString,
			Description: "The name of the GitHub deployment environment. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).",
		},
		"BUILDKITE_GITHUB_DEPLOYMENT_TASK": {
			Type:        schemaString,
			Description: "The name of the GitHub deployment task. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).",
		},
		"BUILDKITE_GITHUB_DEPLOYMENT_PAYLOAD": {
			Type:        schemaString,
			Description: "The GitHub deployment payload data as serialized JSON. Only available on builds triggered by a [GitHub Deployment](https://developer.github.com/v3/repos/deployments/).",
		},
		"BUILDKITE_GROUP_ID": {
			Type:        schemaString,
			Description: "The UUID of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.",
		},
		"BUILDKITE_GROUP_KEY": {
			Type:        schemaString,
			Description: "The value of the `key` attribute of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.",
		},
		"BUILDKITE_GROUP_LABEL": {
			Type:        schemaString,
			Description: "The label/name of the [group step](https://buildkite.com/docs/pipelines/group-step) the job belongs to. This variable is only available if the job belongs to a group step.",
		},
		"BUILDKITE_HOOKS_PATH": {
			Type:        schemaString,
			Description: "The value of the `hooks-path` [agent configuration option](https://buildkite.com/docs/agent/v3/configuration).",
		},
		"BUILDKITE_IGNORED_ENV": {
			Type: schemaArray,
			Items: ObjectArrayProperty{
				Type: schemaString,
			},
			Description: "A list of environment variables that have been set in your pipeline that are protected and will be overridden, used internally to pass data from the bootstrap to the agent.",
		},
		"BUILDKITE_JOB_ID": {
			Type:        schemaString,
			Description: "The internal UUID Buildkite uses for this job.",
		},
		"BUILDKITE_JOB_LOG_TMPFILE": {
			Type:        schemaString,
			Description: "The path to a temporary file containing the logs for this job. Requires enabling the `enable-job-log-tmpfile` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_LABEL": {
			Type:        schemaString,
			Description: "The label/name of the current job.",
		},
		"BUILDKITE_LAST_HOOK_EXIT_STATUS": {
			Type:        schemaNumber,
			Description: "The exit code of the last hook that ran, used internally by the hooks.",
		},
		"BUILDKITE_LOCAL_HOOKS_ENABLED": {
			Type:        schemaBoolean,
			Description: "The opposite of the value of the `no-local-hooks` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_MESSAGE": {
			Type:        schemaString,
			Description: "The message associated with the build, usually the commit message. The value is empty when a message is not set. For example, when a user triggers a build from the Buildkite dashboard without entering a message, the variable returns an empty value.",
		},
		"BUILDKITE_ORGANIZATION_SLUG": {
			Type:        schemaString,
			Description: "The organization name on Buildkite as used in URLs.",
		},
		"BUILDKITE_PARALLEL_JOB": {
			Type:        schemaNumber,
			Description: "The index of each parallel job created from a parallel build step, starting from 0. For a build step with `parallelism: 5`, the value would be 0, 1, 2, 3, and 4 respectively.",
		},
		"BUILDKITE_PARALLEL_JOB_COUNT": {
			Type:        schemaNumber,
			Description: "The total number of parallel jobs created from a parallel build step. For a build step with `parallelism: 5`, the value is 5.",
		},
		"BUILDKITE_PIPELINE_DEFAULT_BRANCH": {
			Type:        schemaString,
			Description: "The default branch for this pipeline.",
		},
		"BUILDKITE_PIPELINE_NAME": {
			Type:        schemaString,
			Description: "The displayed pipeline name on Buildkite.",
		},
		"BUILDKITE_PIPELINE_PROVIDER": {
			Type:        schemaString,
			Description: "The ID of the source code provider for the pipeline's repository.",
		},
		"BUILDKITE_PIPELINE_SLUG": {
			Type:        schemaString,
			Description: "The pipeline slug on Buildkite as used in URLs.",
		},
		"BUILDKITE_PIPELINE_TEAMS": {
			Type: schemaArray,
			Items: ObjectArrayProperty{
				Type:      schemaString,
				Delimiter: ":",
			},
			Description: "A colon separated list of the pipeline's non-private team slugs.",
		},
		"BUILDKITE_PLUGIN_CONFIGURATION": {
			Type:        schemaString,
			Description: "A JSON string holding the current plugin's configuration (as opposed to all the plugin configurations in the `BUILDKITE_PLUGINS` environment variable).",
		},
		"BUILDKITE_PLUGIN_NAME": {
			Type:        schemaString,
			Description: "The current plugin's name, with all letters in uppercase and any spaces replaced with underscores.",
		},
		"BUILDKITE_PLUGINS": {
			Type:        schemaString,
			Description: "A JSON object containing a list plugins used in the step, and their configuration.",
		},
		"BUILDKITE_PLUGINS_ENABLED": {
			Type:        schemaBoolean,
			Description: "The opposite of the value of the `no-plugins` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_PLUGINS_PATH": {
			Type:        schemaString,
			Description: "The value of the `plugins-path` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_PLUGIN_VALIDATION": {
			Type:        schemaBoolean,
			Description: "Whether to validate plugin configuration and requirements. The value can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks, or in a `pipeline.yml` file. It can also be enabled using the `no-plugin-validation` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_PULL_REQUEST": {
			Type:        schemaNumber,
			Description: "The number of the pull request or `false` if not a pull request.",
		},
		"BUILDKITE_PULL_REQUEST_BASE_BRANCH": {
			Type:        schemaString,
			Description: "The base branch that the pull request is targeting or `\"\"` if not a pull request.",
		},
		"BUILDKITE_PULL_REQUEST_DRAFT": {
			Type:        schemaBoolean,
			Description: "Set to `true` when the pull request is a draft. This variable is only available if a build contains a draft pull request.",
		},
		"BUILDKITE_PULL_REQUEST_REPO": {
			Type:        schemaString,
			Description: "The repository URL of the pull request or `\"\"` if not a pull request.",
		},
		"BUILDKITE_REBUILT_FROM_BUILD_ID": {
			Type:        schemaString,
			Description: "The UUID of the original build this was rebuilt from or `\"\"` if not a rebuild.",
		},
		"BUILDKITE_REBUILT_FROM_BUILD_NUMBER": {
			Type:        schemaString,
			Description: "The UUID of the original build this was rebuilt from or `\"\"` if not a rebuild.",
		},
		"BUILDKITE_REFSPEC": {
			Type:        schemaString,
			Description: "A custom refspec for the buildkite-agent bootstrap script to use when checking out code. This variable can be modified by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_REPO": {
			Type:        schemaString,
			Description: "The repository of your pipeline. This variable can be set by exporting the environment variable in the `environment` or `pre-checkout` hooks.",
		},
		"BUILDKITE_REPO_MIRROR": {
			Type:        schemaString,
			Description: "The path to the shared git mirror. Introduced in [v3.47.0](https://github.com/buildkite/agent/releases/tag/v3.47.0).",
		},
		"BUILDKITE_RETRY_COUNT": {
			Type:        schemaNumber,
			Description: "How many times this job has been retried.",
		},
		"BUILDKITE_S3_ACCESS_KEY_ID": {
			Type:        schemaString,
			Description: "The access key ID for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, and during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.",
		},
		"BUILDKITE_S3_ACCESS_URL": {
			Type:        schemaString,
			Description: "The access URL for your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket), if you are using a proxy. The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.",
		},
		"BUILDKITE_S3_ACL": {
			Type: schemaString,
			Description: `The Access Control List to be set on artifacts being uploaded to your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the 'buildkite-agent artifact upload' command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the 'environment', 'pre-checkout' or 'pre-command' hooks.

Must be one of the following values which map to [S3 Canned ACL grants](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl).`,
		},
		"BUILDKITE_S3_DEFAULT_REGION": {
			Type:        schemaString,
			Description: "The region of your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.",
		},
		"BUILDKITE_S3_SECRET_ACCESS_KEY": {
			Type:        schemaString,
			Description: "The secret access key for your S3 IAM user, for use with [private S3 buckets](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks. Do not print or export this variable anywhere except your agent hooks.",
		},
		"BUILDKITE_S3_SSE_ENABLED": {
			Type:        schemaBoolean,
			Description: "Whether to enable encryption for the artifacts in your [private S3 bucket](/docs/agent/v3/cli-artifact#using-your-private-aws-s3-bucket). The variable is read by the `buildkite-agent artifact upload` command, as well as during the artifact upload phase of [command steps](/docs/pipelines/command-step#command-step-attributes). The value can only be set by exporting the environment variable in the `environment`, `pre-checkout` or `pre-command` hooks.",
		},
		"BUILDKITE_SHELL": {
			Type:        schemaString,
			Description: "The value of the `shell` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_SOURCE": {
			Type:        schemaString,
			Description: "The source of the event that created the build.",
		},
		"BUILDKITE_SSH_KEYSCAN": {
			Type:        schemaBoolean,
			Description: "The opposite of the value of the `no-ssh-keyscan` [agent configuration option](/docs/agent/v3/configuration).",
		},
		"BUILDKITE_STEP_ID": {
			Type:        schemaString,
			Description: "A unique string that identifies a step.",
		},
		"BUILDKITE_STEP_KEY": {
			Type:        schemaString,
			Description: "The value of the `key` [command step attribute](/docs/pipelines/command-step#command-step-attributes), a unique string set by you to identify a step.",
		},
		"BUILDKITE_TAG": {
			Type:        schemaString,
			Description: "The name of the tag being built, if this build was triggered from a tag.",
		},
		"BUILDKITE_TIMEOUT": {
			Type:        schemaNumber,
			Description: "The number of minutes until Buildkite automatically cancels this job, if a timeout has been specified, otherwise it `false` if no timeout is set. Jobs that time out with an exit status of 0 are marked as \"passed\".",
		},
		"BUILDKITE_TRACING_BACKEND": {
			Type: schemaString,
			Description: `Set to "datadog" to send metrics to the [Datadog APM](https://docs.datadoghq.com/tracing/) using 'localhost:8126', or 'DD_AGENT_HOST:DD_AGENT_APM_PORT'.

Also available as a [buildkite agent configuration option.](/docs/agent/v3/configuration#configuration-settings)`,
		},
		"BUILDKITE_TRIGGERED_FROM_BUILD_ID": {
			Type:        schemaString,
			Description: "The UUID of the build that triggered this build. This will be empty if the build was not triggered from another build.",
		},
		"BUILDKITE_TRIGGERED_FROM_BUILD_NUMBER": {
			Type:        schemaString,
			Description: "The number of the build that triggered this build or `\"\"` if the build was not triggered from another build.",
		},
		"BUILDKITE_TRIGGERED_FROM_BUILD_PIPELINE_SLUG": {
			Type:        schemaString,
			Description: "The slug of the pipeline that was used to trigger this build or `\"\"` if the build was not triggered from another build.",
		},
		"BUILDKITE_UNBLOCKER": {
			Type:        schemaString,
			Description: "The name of the user who unblocked the build.",
		},
		"BUILDKITE_UNBLOCKER_EMAIL": {
			Type:        schemaString,
			Description: "The notification email of the user who unblocked the build.",
		},
		"BUILDKITE_UNBLOCKER_ID": {
			Type:        schemaString,
			Description: "The UUID of the user who unblocked the build.",
		},
		"BUILDKITE_UNBLOCKER_TEAMS": {
			Type: schemaArray,
			Items: ObjectArrayProperty{
				Type:      schemaString,
				Delimiter: ":",
			},
			Description: "A colon separated list of non-private team slugs that the user who unblocked the build belongs to.",
		},
	}),
	Steps: NewStepMap(map[string]Step{
		"WAIT_STEP": {
			EmptyArgsValue: "wait",
			Description:    "A wait step waits for all previous steps to have successfully completed before allowing following jobs to continue.",
			Properties: NewPropertyMap(map[string]Property{
				"wait": {
					Type:        schemaString,
					Description: "When providing options for the wait step, you will need to set this value to \"~\".",
				},
				"continue_on_failure": {
					Type:        schemaBoolean,
					Description: "Run the next step, even if the previous step has failed.",
				},
				"if": {
					Type:        schemaString,
					Description: "A boolean expression that omits the step when false. See Using conditionals for supported expressions.",
				},
				"depends_on": {
					Type:        schemaArray,
					Description: "A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See [managing step dependencies](https://buildkite.com/docs/pipelines/dependencies) for more information.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"allow_dependency_failure": {
					Type:        schemaBoolean,
					Description: "Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.",
				},
			}),
		},
		"COMMAND_STEP": {
			Description: "A command step runs one or more shell commands on one or more agents.",
			Properties: NewPropertyMap(map[string]Property{
				"commands": {
					Type:        schemaArray,
					Description: "The shell command to run during this step.",
					Required:    true,
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"agents": {
					Type:        schemaObject,
					Description: "A map of agent tag keys to values to target specific agents for this step.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"allow_dependency_failure": {
					Type:        schemaBoolean,
					Description: "Whether to continue to run this step if any of the steps named in the depends_on attribute fail.",
				},
				"artifact_paths": {
					Type:        schemaArray,
					Description: "The glob path or paths of artifacts to upload from this step.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"branches": {
					Type:        schemaString,
					Description: "The branch pattern defining which branches will include this step in their builds.",
				},
				"cancel_on_build_failing": {
					Type:        schemaBoolean,
					Description: "Setting this attribute to true cancels the job as soon as the build is marked as failing.",
				},
				"concurrency": {
					Type:        schemaNumber,
					Description: "The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.",
				},
				"concurrency_group": {
					Type:        schemaNumber,
					Description: "A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.",
				},
				"depends_on": {
					Type:        schemaArray,
					Description: "A list of step keys that this step depends on. This step will only run after the named steps have completed. See managing step dependencies for more information.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"env": {
					Type:        schemaObject,
					Description: "A map of environment variables for this step.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"if": {
					Type:        schemaString,
					Description: "A boolean expression that omits the step when false. See Using conditionals for supported expressions.",
				},
				"key": {
					Type: schemaString,
					Description: `A unique string to identify the step. The value is available in the BUILDKITE_STEP_KEY environment variable.
Keys can not have the same pattern as a UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).`,
				},
				"label": {
					Type:        schemaString,
					Description: "The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.",
				},
				// TODO: Fully implement this
				"martrix": {
					Type:        schemaArray,
					Description: "An array of values to be used in the matrix expansion.",
					Items: ObjectArrayProperty{
						Type: schemaString,
					},
				},
				"parallelism": {
					Type:        schemaNumber,
					Description: "The number of parallel jobs that will be created based on this step.",
				},
				"plugins": {
					TypeRef: "plugin_map",
				},
				"priority": {
					Type:        schemaNumber,
					Description: "Adjust the priority for a specific job, as a positive or negative integer.",
				},
				"rety": {
					TypeRef: "retry_options",
				},
				// Enforcing folks to put in a reason for when you skip a pipeline seems like
				// a reasonable thing to enforce for now. It should make it clearer as to why
				// a step was skipped to folks who are not knowledgable about a pipeline.
				"skip": {
					Type:        schemaString,
					Description: "Whether to skip this step or not. Passing a string (with a 70-character limit) provides a reason for skipping this command. Passing an empty string is equivalent to false. Note: Skipped steps will be hidden in the pipeline view by default, but can be made visible by toggling the 'Skipped jobs' icon.",
				},
				"soft_fail_all": {
					Type:        schemaBoolean,
					Description: "Allow all non-zero exit statuses not to fail the build.",
				},
				"soft_fail": {
					Type:        schemaArray,
					Description: "Allow specified non-zero exit statuses not to fail the build.",
					Items: ObjectArrayProperty{
						TypeRef: "soft_fail_option",
					},
				},
				"timeout_in_minutes": {
					Type:        schemaNumber,
					Description: " The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.",
				},
			}),
		},
	}),
}
