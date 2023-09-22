# Dash Days Portable CI -- Taskfile

## Getting Started

NOTE: This repo is just a Proof of Concept.

This repo is an example of using Taskfile instead of Makefile to run reusable, discrete steps the same in CI and local-dev. This repo makes use of remote tasks that define common actions and use them in the repo Taskfile.yml to spin up a k3d cluster in AWS, build and publish a zarf package, deploy that package onto the cluster and tear it back down.

To use locally, assuming you already have your AWS credentials set up, all you need to do is clone the repo and run:

```bash
task k3d-full -y
```

## Notes and Gotchas

Remote task file includes are currently in experimental mode having been added ~ a month ago. This comes with some gotchas. Namely, caching and git nodes are not supported yet. Additionally, the -y flag in the tool to answer yes to any prompts currently is not respected by the confirmation when pulling in remote files. It should be addressed in the next iteration. Workaround is ```yes|task```. See [here](https://github.com/go-task/task/issues/1317#issuecomment-1721463929) for details.
