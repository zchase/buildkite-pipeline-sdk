# Buildkite Pipeline SDK

This is an experimental approach to generating SDKs for defining [Buildkite Dynamic Pipelines](https://buildkite.com/docs/pipelines/defining-steps#dynamic-pipelines) from a YAML schema.

The purpose of this repo is to demonstrate the technique and provide simple examples. You should not use the code in this repo for any sort of production use.

## Usage

The SDKs are intended to be used locally but you __could__ use the Go SDK on any machine since the repo is public. The TypeScript isn't published to NPM so it can only be used after building/install locally.

### Build and Install

To build and install the SDKs, you just need to run the build and install script:

```
$ ./scripts/build_and_install.sh
```

### Try it out

To try out the SDKs you can create a new project and use the files in the `examples` directory as a guide.