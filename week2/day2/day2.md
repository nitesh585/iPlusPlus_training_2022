## Continue Docker and Intro of CI/CD and CD

![docker multi-stage build](https://image.slidesharecdn.com/multi-stagebuild-170525164324/95/docker-multistage-build-8-638.jpg?cb=1502176649)

With multi-stage builds, you use multiple `FROM` statements in your Dockerfile. Each `FROM` instruction can use a different base, and each of them begins a new stage of the build. You can selectively copy artifacts from one stage to another, leaving behind everything you don’t want in the final image. example,

```dockerfile
FROM golang:1.16
WORKDIR /go/src/github.com/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html
COPY app.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/alexellis/href-counter/app ./
CMD ["./app"]
```

You only need the single Dockerfile. You don’t need a separate build script, either. Just run docker build.

```sh
docker build -t alexellis2/href-counter:latest .
```

The `COPY --from=0` line copies just the built artifact from the previous stage into this new stage.

The end result is the same tiny production image as before, with a significant reduction in complexity

### CI/CD and CD

![CI/CD and CD](https://www.redhat.com/cms/managed-files/styles/wysiwyg_full_width/s3/ci-cd-flow-desktop.png?itok=2EX0MpQZ)

The <b>CI</b> in CI/CD always refers to continuous integration, which is an automation process for developers. Successful CI means new code changes to an app are regularly built, tested, and merged to a shared repository. It’s a solution to the problem of having too many branches of an app in development at once that might conflict with each other.

<b>CD</b> in CD/CD referes to continuous delivery usually means a developer’s changes to an application are automatically bug tested and uploaded to a repository (like GitHub or a container registry), where they can then be deployed to a live production environment by the operations team. It’s an answer to the problem of poor visibility and communication between dev and business teams. To that end, the purpose of continuous delivery is to ensure that it takes minimal effort to deploy new code.

<b>CD</b> (Continuous deployment) can refer to automatically releasing a developer’s changes from the repository to production, where it is usable by customers. It addresses the problem of overloading operations teams with manual processes that slow down app delivery. It builds on the benefits of continuous delivery by automating the next stage in the pipeline.

![CI/CD pipeline](https://www.plutora.com/wp-content/uploads/2019/03/cicd-pipeline-1024x355.png)

### Must read resources:

- [Multi-stage build](https://docs.docker.com/develop/develop-images/multistage-build/)
- [Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [Best practices for using Docker Hub for CI/CD](https://docs.docker.com/ci-cd/best-practices/)
- [CI/CD pipeline](https://www.edureka.co/blog/ci-cd-pipeline/)
