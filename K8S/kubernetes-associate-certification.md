# Kubernetes & Cloud Native Associate (KCNA) Study Guide

## Section 1: Introduction to Cloud Native Principles

### 1. Defining 'Cloud Native' Concepts

Cloud Native applies to Public, Private, and Hybrid Clouds.

**Public Cloud 🌐:** Vendors (e.g., **AWS, GCP, Azure**) offer compute, storage, and networking to the public.

* Apps leverage managed services (serverless, K8s) for global scalability and resilience.
* **Benefits:** Infinite scalability, pay-as-you-go pricing, and reduced infrastructure overhead.

**Private Cloud 🔒:** On-premises infrastructure grants organizations exclusive control and ownership.

* Internal platforms (e.g., **OpenStack, VMware**) emulate public clouds using self-deployed K8s and Prometheus.
* **Benefits:** Maximum control over security, governance, and compliance.

**Hybrid Cloud 🤝:** Links multiple infrastructures, enabling flexible data and app movement.

* Apps span environments, keeping sensitive data private while using public clouds for burst capacity.
* **Benefits:** Merges private cloud security with public cloud scalability and efficiency.

***

### Foundational Cloud Native Facts

| Topic | Answer |
| :--- | :--- |
| **1. What does CNCF stand for?** | **C**loud **N**ative **C**omputing **F**oundation |
| **2. Which project was the first-ever project submission to the CNCF?** | **Kubernetes** |
| **3. What is the significance of the Linux Foundation in relation to Cloud Native?** | Hosts major projects like the Linux Kernel, Kubernetes, and the CNCF |
| **4. When was the Linux Foundation founded?** | 2000 |
| **5. Cloud Native Architecture refers to?** | Apps built with best practices to run across Cloud Native systems |
| **6. What distinguishes Cloud Native from a legacy deployment?** | Apps designed for high availability, leveraging cloud infrastructure and IaC (e.g., Terraform) |
| **7. Who was the original developer of Kubernetes?** | **Google** |

***

## Section 2: Architecture Principles and Practices

### Core Architecture Goals

Cloud Native Architecture aims to design robust applications and supporting infrastructure.

**Architecture Goals:**

* Application Availability
* Cost Management
* Efficiency
* Reliability

These principles enable resilient, manageable, loosely coupled systems. Automation ensures frequent, predictable changes.

#### Characteristics of Cloud Native Applications

Cloud Native apps use the cloud for resilience, agility, operability, and observability (RAOO).

* **Resiliency:** Apps recover quickly using redundancy and failover. Kubernetes' self-healing maintains Pod replicas and replaces failures.
* **Agility:** The speed of building and deploying. Microservices and CD pipelines drive rapid iteration.
* **Operability:** Simplifies management through automation and IaC (e.g., **Terraform**) for monitoring and configuration.
* **Observability:** Uses outputs to understand system state. Logging, monitoring, and tracing are its three essential pillars.

### Cloud Native Automation and Scaling Practices

#### Self Healing

Kubernetes ensures availability by replacing failed containers, rescheduling on node failure, and maintaining the desired state.

#### Application Automation

Kubernetes uses progressive rollouts and health monitoring to prevent failures, automatically rolling back if issues occur.

* **Terraform:** Defines infrastructure as code, declaring the desired state. It is provider-agnostic and creates K8s cluster components in the pre-cluster stage.
* **Ansible:** Manages server configurations via agentless **YAML Playbooks** over SSH. It prepares nodes or deploys cluster configs in the intra/extra-cluster stage.
* **Knative:** Simplifies serverless development via autoscaling and eventing. It scales to zero when idle and manages asynchronous event communication.

#### Autoscaling Types

* **Reactive Autoscaling:** Scales immediately when thresholds (e.g., CPU) are crossed, responding after spikes occur.
* **Scheduled Autoscaling:** Scales via predefined rules for predictable load patterns, like morning peaks.
* **Predictive Autoscaling:** Uses historical data and ML to forecast load, scaling before spikes to minimize lag.
* **Vertical Autoscaling:** Scales capacity (RAM/CPU) of existing machines; also known as "scaling up/down."
* **Horizontal Autoscaling:** Scales by adding or removing instances (Pods/VMs); also known as "scaling out/in."

**KEDA (Kubernetes Event-driven Autoscaling):**

KEDA scales workloads using external metrics like queue depth or Kafka, beyond standard CPU usage.

**How KEDA Works:**

You define a **ScaledObject** in Kubernetes, specifying:

* Target deployment
* **Event source** (e.g., Kafka, queue size)
* Scaling rules

KEDA monitors the event source and adjusts the Pod count accordingly.

#### Serverless Computing

Serverless runs on K8s without developers managing operations, scaling, or resource provisioning.

* **Abstraction:** Users define code and resources, bypassing direct K8s primitives (Pods, Services).
* **Scaling to Zero:** Apps scale automatically with traffic, down to zero when idle to save costs.
* **Event-Driven:** Apps are triggered by events (e.g., HTTP, queues) instead of running constantly.
* **Underlying Technology:** Built on K8s projects like **Knative** and **KEDA**.

#### Serverless Ecosystem Concepts

* **AWS Lambda:** AWS's FaaS executes code in response to events without server management; the original public cloud serverless model.
* **Knative:** Extends K8s with auto-scaling (to zero) and eventing, turning clusters into FaaS platforms.
* **OpenFaaS:** Simplifies deploying serverless functions on K8s by packaging existing code as containers.
* **CloudEvents:** A standard for describing event data, ensuring interoperability across serverless platforms.

### Key Pillars of Cloud Native

* **Microservices Architecture:** Apps are split into loosely coupled, independent components for agility and resilience.
* **Containerization:** Containers package apps and dependencies for uniform execution, ensuring isolation and consistency.
* **DevOps:** Merges development and operations for faster, reliable delivery through automation and collaboration.
* **Continuous Delivery (CD):** Automates building and testing for production, speeding cycles and reducing risk.

### Community Structure and Governance

* **CNCF (Cloud Native Computing Foundation):** Hosts and supports projects like **Kubernetes, Envoy, and Prometheus**.
* **Mission:** To make Cloud Native Computing **Ubiquitous**.

**Graduated projects:** <https://www.cncf.io/projects/>

#### Cloud Native Terminology Acronyms

* **TOC:** Technical Oversight Committee
* **SIG:** Special Interest Groups
* **TAG:** Technical Advisory Groups

#### TAG: Technical Advisory Groups

* Provide guidance in domains like Storage, Security, Networking, and Observability.
* Support new projects through Sandbox onboarding.
* Review projects transitioning from Sandbox to Incubation.

### Cloud Native Professional Roles (Personas)

* **DevOps:** Automates CI/CD pipelines and manages configuration tools like Helm or ArgoCD.
* **Site Reliability Engineer (SRE):** Focuses on production reliability and performance, using automation to define SLOs/SLIs.
* **CloudOps Engineer:** Manages cloud infrastructure (AWS, Azure, GCP), optimizing spend, networking, and security.
* **Security Engineer:** Designs security policies (RBAC, auditing) and manages vulnerability scanning and secrets.
* **DevSecOps Engineer:** Integrates security directly into CI/CD pipelines for continuous testing.
* **Full Stack Developer:** Builds all layers of a feature, focusing on app logic within containerized environments.
* **Cloud Architect:** Designs long-term cloud strategy, selecting services and defining networking blueprints.
* **Data Engineer:** Manages data pipelines and storage (e.g., Kafka) inside K8s, ensuring data quality.
* **FinOps Engineer:** Drives financial efficiency by monitoring K8s resource utilization and spending.
* **Machine Learning Engineer:** Deploys ML models (e.g., Kubeflow) as scalable services with proper allocation.
* **Data Scientist:** Builds predictive models and develops statistical insights from prepared data.

### Essential Open Standards

Open standards ensure interoperability and portability, preventing vendor lock-in.

#### Key Open Standards and Specifications

* **Containerization Standards**
    * **Open Container Initiative (OCI):** Governs **Image Format** (packaging) and **Runtime Specification** (execution).
    * **Container Runtime Interface (CRI):** Defines the API between the **kubelet** and the container runtime (e.g., **containerd**).
* **Networking and Storage**
    * **Container Network Interface (CNI):** Ensures networking solutions (e.g., Calico) integrate consistently with K8s.
    * **Container Storage Interface (CSI):** Allows storage providers (e.g., AWS EBS) to create universal drivers for K8s.
* **Service Mesh and Observability**
    * **OpenTelemetry (OTel):** Standard APIs for collecting telemetry (metrics, logs, traces) across microservices.

## Section 3: Containerization with Docker

### Evolution of Containers

**Mainframe (1960s–1970s):** Introduced workload isolation via **LPARs** and **Time-sharing**, the ancestors of containers.

**chroot (1979):** Unix command providing partial filesystem isolation by changing the root directory. It lacked CPU, memory, and user isolation.

**FreeBSD Jails (2000):** Combined **chroot** with process and network isolation, giving each jail its own IP and process tree.

**Virtual Machines (1990s–2000s):** VMs use a hypervisor to run a **full OS** on virtual hardware. While secure, they are slow and resource-heavy.

**Linux Advancements — The Foundation:** True containerization was enabled by Linux **Namespaces** and **cgroups**.

* **Namespaces:** Isolate user IDs, PIDs, network, filesystem, hostname, and IPC.
* **cgroups (Control Groups, 2006):** Limits resource usage (CPU, Memory, I/O). **Namespaces isolate; cgroups control.**

**Docker (2013) — The Usability Revolution:** Docker standardized containers using Namespaces, cgroups, and **Union filesystems**, making them portable and easy to share.

### Docker Environment Setup

#### Docker Commands and Kubernetes Node Status

`docker run -it <image-name> <default-command>`:

* `-i`: Keep STDIN open.
* `-t`: Allocate a terminal (TTY).
* `-it`: Run interactively with a terminal.

![docker-run](./images/docker-run.png)

`exit`: Exit the container.

`kubectl get nodes`: Queries the API server for cluster nodes.

* NAME: Node identifier.
* STATUS: Readiness (Ready / NotReady / Unknown).
* ROLES: Node role (control-plane, worker).
* AGE: Time in the cluster.
* VERSION: Kubernetes version.

![kubectl-get-nodes](./images/kubectl-get-nodes.png)

### Container Image Fundamentals

**Container Image:** A bundle of software and dependencies. Use **OCI Compliant Container Image** instead of "Docker Image."

<https://opencontainers.org>

**Container Image vs Container:** An **Image** is the static bundle; a **Container** is its running instance.

**Container Registry:** <https://hub.docker.com>

**Image Tag:** Distinguishes image versions (e.g., `1.0.0`, `latest`).

`docker pull <image-name>:<tag>`: Pull the image from a registry.

![docker-pull](./images/docker-pull.png)

To specify a registry: `docker pull docker.io/wernight/funbox:latest`

![docker-pull-2](./images/docker-pull-2.png)

![docker-pull-3](./images/docker-pull-3.png)

* A **union file system** merges layers into one view with a writable top layer.
* Changes are only written to this thin top layer.
* Deleting a file removes its reference in the top layer; it remains in the immutable lower layers.
* Sharing immutable layers makes running multiple containers space-efficient.

**Digest:** A secure, unique identifier for the image (checksum from the registry).

`docker manifest inspect <image-name>:` Show the registry manifest.

![docker-manifest](./images/docker-manifest.png)

`docker save <image-name> -o <file-name>`: Saves images to a portable tar archive.

![docker-save](./images/docker-save.png)

* `-o`: Writes output to a file instead of streaming to stdout.
* The .tar archive is portable; `docker load -i` recreates the image locally.

`tar xvf funbox.tar`: Extracts `funbox.tar` into the current directory.

![docker-save-2](./images/docker-save-2.png)

* Modern Docker versions default to creating **OCI-compliant archives**.
* Layers are stored as blobs named by SHA digests instead of `layer.tar` files.

* `blobs/`: Contains all image layers and config blobs.
* `manifest.json`: Describes image layers.
* `index.json`: Entry point for OCI format.
* `oci-layout`: Specifies OCI version.

![docker-manifest-2](./images/docker-manifest-2.png)

### Managing Running Containers

* Docker uses **containerd**, a CNCF-graduated project.
* **runc** is the OCI reference implementation for container runtimes.

![docker-version](./images/docker-version.png)

`docker run -it spurin/funbox --rm`: Removes the container automatically upon exit.

![docker-run-2](./images/docker-run-2.png)
![docker-run-3](./images/docker-run-3.png)

`docker ps`: Show running containers (`-a` for all).

![docker-ps](./images/docker-ps.png)

Override default command: `docker run -it spurin/funbox nyancat`

![docker-user](./images/docker-user.png)
![docker-user-2](./images/docker-user-2.png)

Containers should ideally run as **non-root users**.

`docker rm <container-name>`: Remove the container.

### Container Networking and Persistent Data (Volumes)

`docker run -d --rm <image>`: `-d` detaches the container to run in the background.

![docker-run-4](./images/docker-run-4.png)

`docker stop <container-id>`: Stops the container.

![docker-stop](./images/docker-stop.png)

`docker run -d --rm -P <container>`: Publishes exposed ports to random host ports.

![docker-run-5](./images/docker-run-5.png)

![docker-run-expose](./images/docker-run-expose.png)
![nginx-curl](./images/nginx-curl.png)
![nginx-browser](./images/nginx-browser.png)

`docker run -d --rm -p 1235:80 <container-name>`: Specify a static port.

![docker-run-expose-2](./images/docker-run-expose-2.png)
![nginx-curl-2](./images/nginx-curl-2.png)

`docker exec -it <container-id> <command>`: Runs a command inside a running container.

![docker-exec](./images/docker-exec.png)

`find / -name "index.html" 2>/dev/null`: Searches for files while hiding "Permission denied" errors.

![cmd-find](./images/cmd-find.png)

`echo hello > /usr/share/nginx/html/index.html`: Overwrites the default Nginx page.

![nginx-content](./images/nginx-content.png)

`docker run -d --rm -p <port> -v <local-path>:<container-path>`: **Volumes** map host folders into containers.

![docker-path](./images/docker-path.png)
![docker-volume](./images/docker-volume.png)

### Creating Container Images

A **Dockerfile** contains instructions to build a container image.

![dockerfile](./images/dockerfile.png)

![dockerfile-2](./images/dockerfile-2.png)

* `FROM`: Specifies the Parent Image for the build.
* `LABEL`: Key-value metadata for Docker objects.

`docker build <dir> -t <tag>`: Builds an image from the Dockerfile.

![docker-build](./images/docker-build.png)
![docker-build-2](./images/docker-build-2.png)

* `RUN`: Executes commands during build, creating new layers.
* `WORKDIR`: Sets the working directory for subsequent instructions.

#### Multiple Stage Builds

* `FROM <image> AS <alias>`: Names a build stage for multi-stage builds.
* `COPY --from=<alias>`: Copies files from a previous stage into the final image.

`RUN adduser -g "John Doe" -s /usr/sbin/nologin -D john`: Adds a non-root user.

* `CMD`: Defines the default command for the container.
* `&&`: Chains commands; the next runs only if the previous succeeds.
* `USER john`: Switches to the specified user.
* `ENTRYPOINT`: Defines the main command that always runs at startup.

`docker buildx build --platform linux/amd64,linux/arm64 -t <tag>`: CLI plugin for **BuildKit**, enabling multi-platform builds and better caching.

`docker system prune`: Removes unused containers, networks, images, and volumes.

## Section 4: Container Orchestration

Container orchestration automates and coordinates multiple containers at scale.

### Features Of Container Orchestration

1. **Automated Deployment & Scheduling**: Schedules containers based on resource availability and constraints.
2. **Scaling & Resource Management**: Adjusts container counts or resources automatically while enforcing limits.
3. **Self-Healing**: Uses health checks to restart failed or replace unhealthy containers.
4. **Service Discovery & Load Balancing**: Ensures stable network identities and load balances traffic.
5. **Config & Secret Management**: Manages environment variables, configs, and secrets externally.
6. **High Availability & Fault Tolerance**: Distributes workloads across nodes to avoid single points of failure.
7. **Rolling Updates & Rollbacks**: Enables zero-downtime updates and easy rollbacks.
8. **Networking & Security**: Controls container traffic, network policies, and RBAC.
9. **Observability**: Integrates monitoring and logging for easier troubleshooting.
10. **Storage Orchestration**: Manages persistent volumes and dynamic provisioning for stateful applications.
11. **Extensibility**: Supports custom resources, operators, and CI/CD.

### Kubernetes Architecture

Kubernetes uses a client-server architecture with a **Control Plane** (brain) and **Worker Nodes** (execution).

#### Control Plane

*Manages the cluster state and decision-making.*

* **API Server (kube-apiserver):** Entry point for all cluster commands; validates and processes API requests.
* **etcd:** Highly available key-value store for all cluster data; the "Source of Truth."
* **Scheduler (kube-scheduler):** Assigns new Pods to nodes based on resource availability and constraints.
* **Controller Manager (kube-controller-manager):** Handles routine tasks like node replacement and maintaining Pod replicas.
* **Cloud Controller Manager (CCM):** Links the cluster to cloud provider APIs (AWS, GCP, etc.) to manage Load Balancers and storage.

#### Worker Node

*Executes the workloads.*

* **Kubelet:** Node agent running on every node to ensure Pod containers are running via the **CRI**.
    * **Static Pods:** Control Plane Kubelets watch `/etc/kubernetes/manifests` to launch core components (e.g., API Server) automatically.
* **Kube-proxy:** Maintains network rules for Pod communication, working alongside the **CNI**.
* **Container Runtime:** Software that executes containers.
    * **High-Level Runtime:** Manages container lifecycles, images, and networking (e.g., **containerd**).
    * **Low-Level Runtime:** Interfaces with the Linux kernel to create containers (e.g., **runc**).

**The Kubernetes Hierarchy:**
**Cluster** → **Node** (Machine) → **Pod** (Deployment Unit) → **Container** (App).

**runc (Low-Level Runtime)**
runc creates and runs containers on Linux by interacting with the kernel.

* **Stack:** Kubernetes → CRI → Container Runtime (containerd) → **runc** → Linux Kernel.
* **Key Features:**
    * Uses **namespaces** (isolation) and **cgroups** (limits).
    * Follows the **OCI** runtime spec.
    * Runs containers without a persistent daemon.

### Core Kubernetes Objects

* **Pod:** The smallest deployable unit, holding one or more containers sharing network and storage.
* **Deployment:** Manages identical Pod replicas and handles rolling updates.
* **Service:** Acts as a load balancer, providing a stable IP or DNS for Pod access.

### CRD (Custom Resource Definition)

Allows extending Kubernetes with custom resource types beyond built-in objects.

**What does a CRD do?**

* Create custom resources using `kubectl`.
* Store them in `etcd`.
* Manage objects via the Kubernetes API.

Example: Extend K8s to recognize objects like **Database**.

**Basic CRD example:**

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: databases.example.com
spec:
  group: example.com
  scope: Namespaced
  names:
    plural: databases
    singular: database
    kind: Database
  versions:
    - name: v1
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          properties:
            spec:
              type: object
              properties:
                engine:
                  type: string
                size:
                  type: string
```

**Why are CRDs useful?**

* **Operators:** Combine CRDs with custom controllers to automate complex applications (e.g., ArgoCD).
* **Infrastructure automation:** Extend K8s to manage external resources like SSL certificates.

#### Key Considerations

* CRDs are **cluster-scoped**, but their resources can be namespaced.
* CRDs provide definitions; logic resides in the **controller** or **operator**.
* Stored in **etcd** and managed via the API server.
