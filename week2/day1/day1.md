## Docker - Intermediate

### Dockerfile reference

Docker can build images automatically by reading the instructions from a `Dockerfile`. A Dockerfile is a text document that contains all the commands a user could call on the command line to assemble an image. Using `docker build` users can create an automated build that executes several command-line instructions in succession.

Here is the format of the Dockerfile:

```js
# Comment
INSTRUCTION arguments
```

The instruction is not case-sensitive. However, convention is for them to be UPPERCASE to distinguish them from arguments more easily.

#### .dockerignore file

Before the docker CLI sends the context to the docker daemon, it looks for a file named .dockerignore in the root directory of the context. If this file exists, the CLI modifies the context to exclude files and directories that match patterns in it. This helps to avoid unnecessarily sending large or sensitive files and directories to the daemon and potentially adding them to images using ADD or COPY.

#### Instructions

- <b>FROM:</b> The FROM instruction initializes a new build stage and sets the Base Image for subsequent instructions. As such, a valid `Dockerfile` must start with a FROM instruction. The image can be any valid image – it is especially easy to start by pulling an image from the Public Repositories.
- <b>RUN:</b> The RUN instruction will execute any commands in a new layer on top of the current image and commit the results. The resulting committed image will be used for the next step in the Dockerfile.
- <b>CMD:</b> The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an `ENTRYPOINT` instruction as well.
- <b>LABEL:</b> The LABEL instruction adds metadata to an image. A LABEL is a `key-value pair`. To include spaces within a LABEL value, use quotes and backslashes as you would in command-line parsing.
- <b>EXPOSE:</b> The EXPOSE instruction informs Docker that the container listens on the specified network ports at runtime. You can specify whether the port listens on `TCP or UDP`, and the default is TCP if the protocol is not specified.
- <b>ENV:</b> The ENV instruction sets the environment variable key to the value. This value will be in the environment for all subsequent instructions in the build stage and can be replaced inline in many as well.
- <b>ADD:</b> The ADD instruction copies new files, directories or remote file URLs from src and adds them to the filesystem of the image at the path dest. Multiple src resources may be specified but if they are files or directories, their paths are interpreted as relative to the source of the context of the build.
- <b>COPY:</b> The COPY instruction copies new files or directories from src and adds them to the filesystem of the container at the path dest. Multiple src resources may be specified but the paths of files and directories will be interpreted as relative to the source of the context of the build.
- <b>ENTRYPOINT:</b> An ENTRYPOINT allows you to configure a container that will run as an executable.
- <b>VOLUME:</b> The VOLUME instruction creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers. The value can be a JSON array, `VOLUME ["/var/log/"]`, or a plain string with multiple arguments, such as `VOLUME /var/log` or `VOLUME /var/log /var/db`
- <b>USER:</b> The USER instruction sets the user name (or UID) and optionally the user group (or GID) to use when running the image and for any `RUN`, `CMD` and `ENTRYPOINT` instructions that follow it in the Dockerfile.
- <b>WORKDIR:</b> The WORKDIR instruction sets the working directory for any `RUN, CMD, ENTRYPOINT, COPY and ADD` instructions that follow it in the Dockerfile. If the WORKDIR doesn’t exist, it will be created even if it’s not used in any subsequent `Dockerfile` instruction.
- <b>ARG:</b> The ARG instruction defines a variable that users can pass at build-time to the builder with the `docker build` command using the `--build-arg <varname>=<value>` flag. If a user specifies a build argument that was not defined in the Dockerfile, the build outputs a warning.
- <b>ONBUILD:</b> The ONBUILD instruction adds to the image a trigger instruction to be executed at a later time, when the image is used as the base for another build. The trigger will be executed in the context of the downstream build, as if it had been inserted immediately after the FROM instruction in the downstream Dockerfile.
- <b>STOPSIGNAL:</b> The STOPSIGNAL instruction sets the system call signal that will be sent to the container to exit.
- <b>HEALTHCHECK:</b> The HEALTHCHECK instruction tells Docker how to test a container to check that it is still working. This can detect cases such as a web server that is stuck in an infinite loop and unable to handle new connections, even though the server process is still running.
- <b>SHELL:</b> The SHELL instruction allows the default shell used for the shell form of commands to be overridden.

### DockerHub

![Docker Hub](https://www.unixtutorial.org/images/software/docker-hub.png)
[Docker Hub](https://hub.docker.com/) is a registry service on the cloud that allows you to download Docker images that are built by other communities. You can also upload your own Docker built images to Docker hub.

#### Push Image

- First of all, we will create a repository in DockerHub that name is `docker-demo`. Make public the repository.
- After creating a repository in DockerHub, we need to have an application that run on the web environment. We will create its image and will push them to DockerHub.
- Create `Dockerfile` in project directory and set configurations.
- After doing this setup and application is ready. we will just build docker image with docker command.
  ```sh
  docker build -t docker-demo .
  ```
- login to docker in terminal using `docker login`
- After build docker image, we need to tag our image. we can use tag command.
  ```sh
  docker tag <image-id> <docker-hub-username>/docker-demo
  ```
- Pushing your image to DockerHub is easy after tagged image. we can type push command like below.
  ```sh
  docker push <docker-hub-username>/docker-demo
  ```

After uploading your image to DockerHub, you will see immediately your image in the repository.

#### Pull Image

Pulling an image from DockerHub is easy. First of all, you need to go to your profile and get the repository name. I got the name docker-demo. Then you just type below command to your terminal to get latest image from DockerHub.

```sh
docker pull <docker-hub-username>/docker-demo:latest
```

#### Running Application

You need to run this installed image with run command.

```sh
docker run <docker-hub-username>/docker-demo
```

### Must read resources:

- [Dockerfile reference](https://docs.docker.com/engine/reference/builder/#dockerfile-examples)
- [Docker: Command Line Interface](https://docs.docker.com/engine/reference/run/)
- [Push and Pull container images to docker hub](https://techtutorialsite.com/docker-push-images-to-dockerhub/)
- [Docker command cheat sheet-1](https://dockerlabs.collabnix.com/docker/cheatsheet/)
- [Docker command cheat sheet-2](https://www.educba.com/docker-commands-cheat-sheet/)
- [Docker command cheat sheet-3](https://www.docker.com/sites/default/files/d8/2019-09/docker-cheat-sheet.pdf)
