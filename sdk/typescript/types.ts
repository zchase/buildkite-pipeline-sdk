// This file is auto-generated please do not edit

// Allow specified non-zero exit statuses not to fail the build.
export type SoftFailOption = Record<string, number>

// A generic string map.
export type StringRecord = Record<string, string>

// An array of plugins for this step.
export type PluginMap = Record<string, StringRecord>

export interface RetryOptions {
    // Whether to allow a job to retry automatically. This field accepts a boolean value, individual retry conditions, or a list of multiple different retry conditions.
    automatic?: boolean;
    // Whether to allow a job to be retried manually. This field accepts a boolean value, or a single retry condition.
    manual?: boolean;
}