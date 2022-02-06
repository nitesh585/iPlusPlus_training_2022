## Pipeline in Jenkins

![Jenkins pipeline](https://miro.medium.com/max/1200/1*iKuaNfxgZSTe_J2x3PYRUg.png)
Jenkins Pipeline is a collection of jobs or events that brings the software from version control into the hands of the end users by using automation tools. It is used to incorporate continuous delivery in our software development workflow.

### JenkinsFile

Jenkins Pipeline can be defined by a text file called JenkinsFile. You can implement pipeline as code using JenkinsFile, and this can be defined by using a DSL (Domain Specific Language). With the help of JenkinsFile, you can write the steps required for running a Jenkins Pipeline.

### Create pipeline of Go

#### Global configuration

- In the Jenkins global tool configuration settings (Manage Jenkins → Global Tool Configuration), find the "Go" section, click "Go Installations…" and "Add Go".
- Enter a name, e.g. "Go 1.15" — the name itself has no significance, but will be displayed to users during Freestyle job configuration, or is what you need to enter as the name in a Pipeline
- Either select "Install automatically" and select the desired Go version from the drop-down list or specify the installation directory manually

#### Per-job configuration

##### Pipeline

As with any other type of Tool Installer, you can use the tool step, in this case with the go tool type.

```groovy
pipeline {
    agent any
    tools {
        // Install the Maven version configured as "M3" and add it to the path.
        go "go"
    }
    stages {
        stage('version') {
            steps {
                sh "go version"
            }
        }
    }
}
```

# gitlab CI

Use GitLab CI/CD to catch bugs and errors early in the development cycle. Ensure that all the code deployed to production complies with the code standards you established for your app.

GitLab CI/CD can automatically build, test, deploy, and monitor your applications by using Auto DevOps.

# gitlab runner

![gitlab runner](https://www.onyxpoint.com/images/gitlab-runner-pipeline.png)
GitLab Runner is an application that works with GitLab CI/CD to run jobs in a pipeline.
You can install GitLab Runner on several different supported operating systems. Other operating systems may also work, as long as you can compile a Go binary on them.

GitLab Runner can also run inside a Docker container or be deployed into a Kubernetes cluster.

### Values of Swiggy

[![Black Widow](https://img.youtube.com/vi/TjEK0MbJUz4/sddefault.jpg)](https://www.youtube.com/watch?v=TjEK0MbJUz4)

### must read resources

- [Jenkins: Getting Started with Pipelines ](https://www.jenkins.io/pipeline/getting-started-pipelines/)

- [Jenkins: Creating your first Pipeline ](https://www.jenkins.io/doc/pipeline/tour/hello-world/)

- [Jenkins: Pipeline Syntax](https://www.jenkins.io/doc/book/pipeline/syntax/)

- [Jenkins: Blue Ocean](https://www.jenkins.io/projects/blueocean/about/)

- [GitLab CI](https://docs.gitlab.com/ee/ci/)

- [Gitlab Runner](https://docs.gitlab.com/runner/)
