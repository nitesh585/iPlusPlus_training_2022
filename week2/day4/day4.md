## Docker compose

![docker compose](https://quintagroup.com/cms/technology/Images/docker-compose-button.jpg)
Docker Compose is used to run multiple containers as a single service. For example, suppose you had an application which required NGNIX and MySQL, you could create one file which would start both the containers as a service without the need to start each one separately.

### Docker Compose ─ Installation

The following steps need to be followed to get Docker Compose up and running.

<b>Step 1</b> − Download the necessary files from github using the following command −

```sh
curl -L "https://github.com/docker/compose/releases/download/1.10.0-rc2/dockercompose
   -$(uname -s) -$(uname -m)" -o /home/demo/docker-compose
```

The above command will download the latest version of Docker Compose which at the time of writing this article is 1.10.0-rc2. It will then store it in the directory `/home/demo/`.

<b>Step 2</b> − Next, we need to provide execute privileges to the downloaded Docker Compose file, using the following command −

```
chmod +x /home/demo/docker-compose
```

We can then use the following command to see the compose version.

Syntax

```sh
docker-compose version
```

### Example of running MySQL using docker compose

create `docker-compose.yaml` file.

```docker
# docker-compose.yaml
# Use root/example as user/password credentials
version: '3.1'

services:

  db:
    image: mysql
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example

  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```

Run `docker-compose -f docker-compose.yml up`, wait for it to initialize completely, and visit `http://localhost:8080`.

## Must read resources

- [Docker compose: Overview](https://docs.docker.com/compose/)
- [Docker Compose vs Multi-Stage Build](https://stackoverflow.com/questions/44833242/docker-compose-vs-multi-stage-build)
- [Docker hub: Mysql](https://hub.docker.com/_/mysql)
- [Make Jenkins run docker without sudo](https://stackoverflow.com/questions/39638772/make-jenkins-run-docker-without-sudo)
- [Setup Docker Containers as Build Agents for Jenkins](slaves)
- [Jenkins: docker pipeline](https://www.jenkins.io/doc/book/pipeline/docker/)
