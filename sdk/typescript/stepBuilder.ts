// This file is auto-generated please do not edit
import {
    PluginMap,
    RetryOptions,
} from "./types";

import * as fs from "fs";
interface WaitStepArgs {
    // When providing options for the wait step, you will need to set this value to "~".
    wait?: string;

    // Run the next step, even if the previous step has failed.
    continueOnFailure?: boolean;

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;

    // A list of step keys that this step depends on. This step will only proceed after the named steps have completed. See [managing step dependencies](https://buildkite.com/docs/pipelines/dependencies) for more information.
    dependsOn?: string[];

    // Whether to continue to proceed past this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;

}

interface CommandStepArgs {
    // The number of parallel jobs that will be created based on this step.
    parallelism?: number;

    // Setting this attribute to true cancels the job as soon as the build is marked as failing.
    cancelOnBuildFailing?: boolean;

    // A map of environment variables for this step.
    env?: Record<string, string>;

    // Whether to skip this step or not. Passing a string (with a 70-character limit) provides a reason for skipping this command. Passing an empty string is equivalent to false. Note: Skipped steps will be hidden in the pipeline view by default, but can be made visible by toggling the 'Skipped jobs' icon.
    skip?: string;

    // Allow all non-zero exit statuses not to fail the build.
    softFailAll?: boolean;

    // The maximum number of jobs created from this step that are allowed to run at the same time. If you use this attribute, you must also define a label for it with the concurrency_group attribute.
    concurrency?: number;

    // An array of plugins for this step.
    plugins?: PluginMap;

    // A map of agent tag keys to values to target specific agents for this step.
    agents?: Record<string, string>;

    // An array of values to be used in the matrix expansion.
    martrix?: string[];

    // The glob path or paths of artifacts to upload from this step.
    artifactPaths?: string[];

    // A unique string to identify the step. The value is available in the BUILDKITE_STEP_KEY environment variable.
    // Keys can not have the same pattern as a UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
    key?: string;

    // A unique name for the concurrency group that you are creating. If you use this attribute, you must also define the concurrency attribute.
    concurrencyGroup?: number;

    //  The maximum number of minutes a job created from this step is allowed to run. If the job exceeds this time limit, or if it finishes with a non-zero exit status, the job is automatically canceled and the build fails. Jobs that time out with an exit status of 0 are marked as passed.
    timeoutInMinutes?: number;

    // The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
    label?: string;

    // A list of step keys that this step depends on. This step will only run after the named steps have completed. See managing step dependencies for more information.
    dependsOn?: string[];

    // A boolean expression that omits the step when false. See Using conditionals for supported expressions.
    if?: string;

    // Adjust the priority for a specific job, as a positive or negative integer.
    priority?: number;

    // Whether to continue to run this step if any of the steps named in the depends_on attribute fail.
    allowDependencyFailure?: boolean;

    // The conditions for retrying this step.
    rety?: RetryOptions;

    // The shell command to run during this step.
    commands: string[];

    // Allow specified non-zero exit statuses not to fail the build.
    softFail?: any[];

    // The branch pattern defining which branches will include this step in their builds.
    branches?: string;

}


class StepBuilder {
	private steps: any[] = [];

	public write() {
		fs.writeFileSync("pipeline.json", JSON.stringify({ steps: this.steps }, null, 4));
	}
    // A wait step waits for all previous steps to have successfully completed before allowing following jobs to continue.
    public addWaitStep(args?: WaitStepArgs): this {
        if (args === undefined) {
            this.steps.push("wait")
        }
        this.steps.push({ ...args });
        return this;
    }
    // A command step runs one or more shell commands on one or more agents.
    public addCommandStep(args: CommandStepArgs): this {
        if (args === undefined) {
            this.steps.push("")
        }
        this.steps.push({ ...args });
        return this;
    }
}
export default StepBuilder;