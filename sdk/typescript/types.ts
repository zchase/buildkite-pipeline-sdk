// This file is auto-generated please do not edit

// Allow specified non-zero exit statuses not to fail the build.
export type SoftFailOption = Record<string, number>

// The conditions for retrying this step.
export enum RetryOptions {
    AUTOMATIC = "automatic",
    MANUAL = "manual",
}

// A generic string map.
export type StringRecord = Record<string, string>

// An array of plugins for this step.
export type PluginMap = Record<string, StringRecord>
