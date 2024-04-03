#!/bin/bash

set -eo pipefail

echo "--- :pipeline: generating pipeline"
cd .buildkite
npm install
npm link buildkite-pipline-sdk
npm run build

echo "--- :pipeline_upload: uploading pipeline.json"
buildkite-agent pipeline upload pipeline.json