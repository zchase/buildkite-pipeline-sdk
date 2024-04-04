import * as bk from "buildkite-pipeline-sdk";

let pipeline = bk.stepBuilder
    .addCommandStep({
        commands: [ `echo "Hello World!"` ],
    });

const branchName = bk.environment.branch();
if (branchName === "main") {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "I am on the main branch"` ],
    })
} else {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "This isn't the main its the ${branchName} branch"` ],
    })
}

pipeline.write();
