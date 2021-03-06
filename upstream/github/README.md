# Introduction

This plugin implements github/github entreprise configuration to build an SCM environment for GIT repositories upstream.

It has been implemented as REST API. See ...(TBD) for FORJJ REST API description.

Depending on tasks, the driver will concretely do several things described below:

## Create task

`Create` properly configure the server side to have at least the `infra` repository created.

It will create a `<infra>/apps/upstream/github.yaml` which describes how github must be configured properly.

On github server, through Github API, the driver will ensure element requested will exist. So it can:
- Create/Configure the organization (with members)
- Create/Configure the `infra` repository (with members)

The creation of the initial repositories/organization/... are done automatically at create time by the maintain task described below.

As soon as `<infra>/apps/upstream/github.yaml` exists, create will fails and ask to use update instead.
**Warning!!!** If the server side has already been configured, create will create the file and start the maintain task, as usual. And then Maintain will change anything to comply the newly created `github.yaml`

## Update task

Update mainly do update in the local `infra` repo and reports file updated to forjj. (The flow must be configured to push to the right place.)
It may (TBD) do the following:
- Update the list of repositories and rights
- Update github configuration data about each repositories in `<organization>-infra/repos/<repo>/`

## Maintain task
This will ensure the SCM server side is properly configured and really update the server:

- update organization/rights
- update repositories/rights
- create/remove repositories
  By default it initializes the repository with a README.md
  But if a migrate upstream has been identified, it will clone it and push it as a new repo in github.

Usually this tasks a 'Ops' task that has the appropriate right to update the infra. And it should be update by an automated job, like a jenkins job or a hubot task.

This does mean you no one approve this to be done. It depends on the flow you choose.

In the github PR flow, a jenkins job should be started as soon as the Ops infra team accept the update through a PR and merge it.

# Local remotes

Depending on the flow the team want to use, local git remote can have 'origin' + 'upstream'

# github flows

The way you use github depends on your need.
A flow usually means that you have CI/CHAT/SCM/... tools integrated to realize your flow.
So a real flow should be implemented technically by apply a flow to your repository. This could be part of the forjj link task. But this part still under reflexion and needs to be shortly implemented.

The idea could be to implement the flow by calling a 'flow create/update/maintain' that a specific github or other app driver will do.
For example, for a PR with github/jenkins, who will take care of configuring them (jenkins/github) ? may be both but not completely sure.
For now, the jenkins-ci discussed a little about this flow implementation details as what we need to see on jenkins side in term of jobs.

See [jenkins-ci README](https://github.com/forj-oss/forjj-contribs/ci/jenkins-ci/README.md)

This will be analyzed and a proposal will be posted later.


The FORJ Team
