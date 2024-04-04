#!/bin/bash

set -eo pipefail

echo "--- :pipeline: generating pipeline"
cd .buildkite
npm install
npm link buildkite-pipeline-sdk
npm run build

echo "--- :pipeline_upload: uploading pipeline.json"
buildkite-agent pipeline upload pipeline.json