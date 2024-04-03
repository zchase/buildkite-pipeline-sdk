import * as bk from "buildkite-pipline-sdk";
import * as fs from "fs";

let pipeline = bk.stepBuilder
    .addCommandStep({
        commands: [ "echo \"Hello World!\"" ],
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

fs.writeFileSync("pipeline.json", JSON.stringify(pipeline.write(), null, 4));
