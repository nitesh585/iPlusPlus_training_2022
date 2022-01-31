### VMs Vs Containers

![vm_vs_container](https://wiki.geant.org/download/attachments/184549510/containers-vs-virtual-machines.jpg?version=1&modificationDate=1607437321973&api=v2)
<b>VMs</b> solve infrastructure problems by letting organizations get more out of servers and facilitate limited workload portability. It runs on top of an emulating software called the hypervisor which sit between the hardware and the virtual machine.Each virtual machine runs its own guest operating system. They are less agile and have low portability than containers.

<b>Containers</b> solve application problems by improving DevOps, enabling microservices, increasing portability, and further improving resource utilization. It sits on the top of a physical server and its host operating system. They share a common operating system that requires care and feeding for bug fixes and patches. They are more agile and have high portability than virtual machines.

### What are the pros and cons of virtualization?

#### VM Pros:

- Decades of virtualization expertise enables access to a robust set of VM management and security tools
- VMs offer the ability to run multiple applications requiring different OSs on a single piece of infrastructure
- VMs emulate an entire computing environment, including all OS resources
- VMs simplify the portability and migration between on-premises and cloud-based platforms
- There is a vast, established VM ecosystem and marketplace with industry leaders such as VMware

#### VM Cons:

- VM images typically consume gigabytes and thus take longer to backup or migrate between platforms Because they encapsulate the entire server including OS, a physical server can support fewer VMs than containers
- VM Spin-up time can take minutes

### What are the pros and cons of containers?

#### Container Pros:

- Containers are more lightweight than VMs, as their images are measured in megabytes rather than gigabytes
- Containers require fewer IT resources to deploy, run, and manage Containers spin up in milliseconds Since their order of magnitude is smaller
- A single system can host many more containers as compared to VMs

#### Container Cons:

- All containers must run atop the same OS – no mix and match of OSs or versions
- Containers may be less secure than VMs since the underlying OS is shared Containers are a newer technology, and the ecosystem is still evolving

### Docker

![Doccker](https://assets.peak.ai/app/uploads/2020/02/13120159/2-what-are-docker-containers.png)
Docker is an open source containerization platform. It enables developers to package applications into containers—standardized executable components combining application source code with the operating system (OS) libraries and dependencies required to run that code in any environment. Containers simplify delivery of distributed applications,

### Docker Architecture:

Docker uses a client-server architecture. The Docker client talks to the Docker daemon, which does the heavy lifting of building, running, and distributing your Docker containers. The Docker client and daemon communicate using a REST API, over UNIX sockets or a network interface. Another Docker client is Docker Compose, that lets you work with applications consisting of a set of containers.
![Docker architecture](https://docs.docker.com/engine/images/architecture.svg)

#### The Docker daemon

The Docker daemon (dockerd) listens for Docker API requests and manages Docker objects such as images, containers, networks, and volumes. A daemon can also communicate with other daemons to manage Docker services.

#### The Docker client

The Docker client (docker) is the primary way that many Docker users interact with Docker. When you use commands such as docker run, the client sends these commands to dockerd, which carries them out. The docker command uses the Docker API. The Docker client can communicate with more than one daemon.

#### Docker Desktop

Docker Desktop is an easy-to-install application for your Mac or Windows environment that enables you to build and share containerized applications and microservices. Docker Desktop includes the Docker daemon (dockerd), the Docker client (docker), Docker Compose, Docker Content Trust, Kubernetes, and Credential Helper. For more information, see Docker Desktop.

#### Docker registries

A Docker registry stores Docker images. Docker Hub is a public registry that anyone can use, and Docker is configured to look for images on Docker Hub by default. You can even run your own private registry.

When you use the docker pull or docker run commands, the required images are pulled from your configured registry. When you use the docker push command, your image is pushed to your configured registry.

#### Docker objects

When you use Docker, you are creating and using images, containers, networks, volumes, plugins, and other objects. This section is a brief overview of some of those objects.

##### Images

An image is a read-only template with instructions for creating a Docker container. Often, an image is based on another image, with some additional customization. For example, you may build an image which is based on the ubuntu image, but installs the Apache web server and your application, as well as the configuration details needed to make your application run.

##### Containers

A container is a runnable instance of an image. You can create, start, stop, move, or delete a container using the Docker API or CLI. You can connect a container to one or more networks, attach storage to it, or even create a new image based on its current state.

By default, a container is relatively well isolated from other containers and its host machine. You can control how isolated a container’s network, storage, or other underlying subsystems are from other containers or from the host machine.

### Must Read Resources:

- [Why use containers vs. VMs? | VMware](https://www.vmware.com/topics/glossary/content/vms-vs-containers.html#:~:text=Container%20Pros%3A%201%20Containers%20are%20more%20lightweight%20than,host%20many%20more%20containers%20as%20compared%20to%20VMs)
- [Docker: Overview](https://docs.docker.com/get-started/overview/)
- [Docker: Get Started](https://docs.docker.com/get-started/)
- [Docker: Get Docker](https://docs.docker.com/get-docker/)
- [Docker: GoLang build images](https://docs.docker.com/language/golang/build-images/)
