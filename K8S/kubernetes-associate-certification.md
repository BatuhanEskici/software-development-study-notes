# Kubernetes Certified (KCNA) + Hands On Labs + Practice Exams

## Section 1: Course and Cloud Native Introduction

### 1. What is the meaning of 'Cloud Native'?

It applies to all cloud offerings, including Public, Private, and Hybrid Clouds.

**Public Cloud** 🌐
The most common model is where computing resources are owned and operated by a third-party provider and delivered over the public internet.

* **Meaning:** Services (like computing power, storage, and networking) are offered by major providers (e.g., AWS, Google Cloud, Microsoft Azure) to the general public or large organizations.
* **Context for Cloud Native:** Applications are built to be highly scalable and resilient across the provider's global infrastructure. You utilize fully managed services (such as serverless functions, managed Kubernetes, and managed databases) offered by the vendor.
* **Advantage:** Infinite scalability, pay-as-you-go pricing, and minimal operational overhead for the underlying infrastructure.

**Private Cloud** 🔒
A computing environment used exclusively by one business or organization.

* **Meaning:** The infrastructure (servers, network, storage) is located in a company's internal data center or managed by a third party solely for that organization. The organization has exclusive control and ownership.
* **Context for Cloud Native:** Organizations build their own internal platforms (often using technologies like OpenStack, VMware, or OpenShift) to mimic the service delivery model of a public cloud. They deploy cloud-native tools (like Kubernetes and Prometheus) themselves.
* **Advantage:** Maximum control over security, data governance, and regulatory compliance.

**Hybrid Cloud** 🤝
A mixed computing environment that connects a public cloud infrastructure with a private cloud infrastructure.

* **Meaning:** It's a blend of two or more distinct cloud infrastructures. Data and applications can move between the two environments, providing flexibility.
* **Context for Cloud Native:** Applications are often designed to run across both environments. For instance, sensitive customer data may be kept in the Private Cloud (on-premises), while burst capacity or less sensitive data processing is handled by the Public Cloud. Cloud-native tools ensure consistency across both environments.
* **Advantage:** Combines the security and control of a private cloud with the scalability and cost-efficiency of a public cloud.

***

### General Cloud Native Facts

| Topic | Answer |
| :--- | :--- |
| **1. What does CNCF stand for?** | **C**loud **N**ative **C**omputing **F**oundation |
| **2. Which project was the first-ever project submission to the CNCF?** | **Kubernetes** |
| **3. What is the significance of the Linux Foundation in relation to Cloud Native?** | It is responsible for hosting significant projects like the Linux Kernel, Kubernetes, and the CNCF |
| **4. When was the Linux Foundation founded?** | 2000 |
| **5. Cloud Native Architecture refers to?** | Applications built using Cloud Native best practices to run across all Cloud Native Systems |
| **6. What distinguishes Cloud Native from a legacy deployment in terms of application development?** | Applications designed for **high availability and fault tolerance**; Applications leveraging the **infrastructure** provided by cloud service providers; Applications built using **infrastructure-as-code** tools like Terraform for flexible and vendor-agnostic management |
| **7. Who was the original developer of Kubernetes?** | **Google** |

***

## Section 2: Cloud Native Architecture Fundamentals

### Cloud Native Architecture Fundamentals
The goal of Cloud Native Architecture is to effectively design applications and infrastructure.

**Best Practices:**
* Application Availability
* Cost Management
* Efficiency
* Reliability

These techniques enable loosely coupled systems that are resilient, manageable, and observable. Combined with robust automation, they allow engineers to make high-impact changes frequently and predictably with minimal toil.

![ss-1](./images/ss-1.png)
![ss-2](./images/ss-2.png)
![ss-3](./images/ss-3.png)

#### Characteristics of Cloud Native Applications
Cloud Native Applications harness the power of the cloud to provide increased resilience, agility, operability, and observability. Let's dive a bit deeper into these characteristics.

**Resiliency**
Resilient applications are designed to withstand failures and continue to function or recover quickly. They typically make use of patterns such as redundancy, failover, and graceful degradation. Self-healing, where systems automatically detect and recover from failure, is a key aspect of resilient cloud-native applications. Kubernetes, for instance, has a built-in self-healing mechanism where it maintains a desired number of pod replicas and replaces failed instances.

**Agility**
Agility in the context of cloud-native applications refers to the ability to quickly build, modify, and deploy applications. Agile practices such as microservices and continuous delivery pipelines, backed by automation, promote rapid iteration and responsiveness to change.

**Operability**
Operability encompasses the ease of deploying, running, and managing applications. Cloud-native applications are designed to be easily monitored, configured, and maintained. They typically leverage automation and Infrastructure as Code (IaC) tools like Terraform to streamline operations and minimize toil.

**Observability**
Observability is the ability to understand the internal state of your system based on the outputs it generates. It's a critical component in diagnosing issues and understanding how an application behaves in the wild. Logging, monitoring, and tracing (collectively known as the 'three pillars of observability') are vital practices to understand the state and performance of cloud-native applications.

_It is important as part of your studies especially if you’re planning on taking the KCNA exam to remember these characteristics! A friendly anagram for this is “RAOO: Racoons are Often Observant” - Raccoons Are Often Observant = Resiliency, Agility, Operability, Observability_

### Cloud Native Practices

#### Self Healing

![ss-4](./images/ss-4.png)
![ss-5](./images/ss-5.png)
![ss-6](./images/ss-6.png)

**Self Healing:** replica count decreases, Kubernetes will automatically attempt to replace the failed replica

#### Application Automation

![ss-7](./images/ss-7.png)
![ss-8](./images/ss-8.png)

**Terraform**
Terraform is an open-source Infrastructure as Code (IaC) tool that allows you to define, provision, and manage cloud infrastructure using declarative configuration files. It supports multiple cloud providers (AWS, Azure, GCP, etc.) and automates the creation and modification of resources like virtual machines, networks, and storage. Terraform tracks the state of your infrastructure and enables reproducible and consistent environments.

**In Kubernetes context:** Terraform can be used to provision the underlying infrastructure (e.g., cloud VMs, networking, load balancers) where your Kubernetes cluster will run.

**Ansible**
Ansible is an open-source automation tool used primarily for configuration management, application deployment, and orchestration. It uses simple YAML playbooks to automate tasks such as installing software, configuring servers, and managing services. Ansible is agentless and works over SSH, making it easy to manage and configure existing servers or nodes.

**In Kubernetes context:** Ansible can be used to configure Kubernetes nodes, deploy applications onto the cluster, or automate operational tasks within Kubernetes environments.

![ss-9](./images/ss-9.png)
![ss-10](./images/ss-10.png)
![ss-11](./images/ss-11.png)
![ss-12](./images/ss-12.png)

#### Autoscaling

![ss-13](./images/ss-13.png)
![ss-14](./images/ss-14.png)
![ss-15](./images/ss-15.png)
![ss-16](./images/ss-16.png)
![ss-17](./images/ss-17.png)
![ss-18](./images/ss-18.png)

#### Vertical Autoscaling (Scaling Up/Down)

**Definition:** Adding more power (CPU, RAM, etc.) to a single machine or server.

**Daily Example:** Imagine you’re working on your laptop. One day, you realize it’s too slow for video editing. So, you upgrade it by:

* Adding more RAM
* Replacing the processor with a faster one
* You're still using one laptop—just making it more powerful. That's vertical scaling.

**In Cloud Context:** If a web app is struggling with traffic, you might upgrade the server from 2 CPUs to 8 CPUs and add more RAM.

#### Horizontal Autoscaling (Scaling Out/In)
**Definition:** Adding more machines/instances to handle the load, rather than upgrading a single one.

**Daily Example:** Suppose you’re running a pizza shop. At first, you cook alone. But during peak hours, you hire 2 more cooks to keep up. You don’t become stronger; you just add more people to share the work.

**In Cloud Context:** If your app is overloaded, you launch 3 more instances (servers) to distribute the traffic among them.

**Summary Comparison:**
| Feature | Advantage | Disadvantage |
| :--- | :--- | :--- |
| **Insertion/Deletion** | **Very Fast.** To add a new item to the list, you just need to change the note of the previous item. You don't have to shift the entire list. | **Slow Access.** To find the 5th book, you are forced to check the 1st, 2nd, 3rd, and 4th books one by one. You cannot jump directly to the 5th spot. |

![ss-19](./images/ss-19.png)

#### Real-World Daily Life Examples:

    1. Coffee Shop Analogy
    Without KEDA (just HPA):

    You check how many customers are in the shop (CPU usage), and add more baristas if it's crowded.

    But you never close the shop, even at night—there’s always 1 barista on duty, even if no one's around.

    With KEDA:

    You listen for events like online coffee orders or mobile app pings.

    If no one orders coffee for an hour, you close the shop (scale to 0).

    If 50 orders come in at once, you quickly hire 5 baristas (scale up pods instantly).

    2. Email Campaign Service
    Suppose you send bulk emails through a queue (like RabbitMQ).

    You use KEDA to watch the queue.

    💤 If the queue is empty → scale to zero (no pods running).

    🔥 If the queue fills up with 10,000 emails → scale up to 10 pods to handle load.

    3. Online Food Delivery (e.g. Getir, Yemeksepeti)
    During lunch/dinner rush:

    More users open the app → KEDA sees an event spike (API traffic, order queue).

    It launches more pods to handle traffic.

    After midnight:

    No one is ordering → KEDA scales the pods down to zero to save resources.

    How KEDA Works (Simplified)
    You define a ScaledObject in Kubernetes.

    Inside it, you set:

    The target deployment

    The event source (e.g., Kafka topic, queue size, custom metric)

    Scaling rules (e.g., when >100 messages, add 3 pods)

    KEDA listens and adjusts pod count accordingly.

#### Serverless

Serverless refers to using Kubernetes as the underlying infrastructure to run serverless workloads. It enables developers to execute code without managing the Kubernetes cluster operations, scaling, or resource provisioning for individual functions or services.

**Abstraction:** The user only defines their code (function or service) and required resources (like memory). They don't interact with Kubernetes primitives like Pods, Deployments, or Services.

**Automatic Scaling to Zero:** The platform automatically scales the application based on incoming traffic. When traffic stops, the service scales down to zero running instances, conserving resources and cost.

**Event-Driven:** Applications are often triggered by events (e.g., messages on a queue, HTTP requests) rather than running continuously.

**Underlying Technology:** It is typically implemented using projects built on top of Kubernetes, such as:

**Knative:** A leading framework that provides middleware components for building, deploying, and managing modern serverless workloads on Kubernetes.

**KEDA (Kubernetes Event-driven Autoscaling):** Allows Kubernetes workloads to be scaled based on external metrics (like a Kafka queue depth) rather than just CPU usage.

![ss-20](./images/ss-20.png)
![ss-21](./images/ss-21.png)
![ss-22](./images/ss-22.png)
![ss-23](./images/ss-23.png)

### Key Pillars of Cloud Native Architecture

**1. Microservices Architecture**
Microservices architecture involves breaking down the application into loosely coupled, independently deployable components, each focusing on a single responsibility. This design enables agility, scalability, and resilience as each microservice can be developed, scaled, and managed independently.

**2. Containerisation**
Containerisation involves encapsulating an application with its dependencies into a container, which can run uniformly across different environments. It facilitates isolation, consistency, and efficiency, making applications easier to build, deploy, and manage.

**3. DevOps**
DevOps is a collaborative approach that combines software development (Dev) and IT operations (Ops) to enhance the efficiency, reliability, and speed of software delivery. By fostering a culture of excellence, DevOps emphasises automation, monitoring, and collaboration across development and operations teams.

**4. Continuous Delivery (CD)**
Continuous Delivery is a practice where code changes are automatically built, tested, and prepared for a release to production. CD accelerates the release cycle, enhances productivity, and reduces the risk, complexity, and downtime of application deployment.

In essence, building cloud-native applications is a strategy that promotes resilience, agility, operability, and observability by leveraging modern technological practices.

With a clear understanding of these characteristics and key pillars, organisations can fully exploit the advantages of cloud-native architectures.

This is also an important area for your studies especially with reference to the KCNA exam. A friendly acronym for this is Morning Coffee Delivers Caffeine Delight = Microservices, Containerisation, DevOps, Continuous Delivery.

### Community and Governance

CNCF is responsible for hosting, support, oversight and direction of cloud native projects

Most well known being Kubernetes, but also including the likes of Envoy and Prometheus

**Their mission:** To make Cloud Native Computing, Ubiquitous

![ss-24](./images/ss-24.png)
![ss-25](./images/ss-25.png)

**Graduated projects:**
https://www.cncf.io/projects/

![ss-26](./images/ss-26.png)
![ss-27](./images/ss-27.png)

### Cloud Native Personas

#### DevOps Engineer
In the Kubernetes context, a DevOps Engineer is a specialized role focused on automating, managing, and maintaining the container orchestration platform and the continuous delivery pipelines that run applications on it.

They essentially bridge the gap between development (building/testing code) and operations (deploying/monitoring the infrastructure).

##### Core Responsibilities
The DevOps Engineer's primary goal in a Kubernetes environment is to ensure applications can be deployed reliably and quickly, and that the underlying cluster is stable, secure, and cost-efficient.

    1. Infrastructure Management (The "Ops" Side)
    Cluster Provisioning & Maintenance: Setting up, securing, and upgrading Kubernetes clusters (e.g., using tools like Terraform, Ansible, or cloud-specific services like EKS/GKE/AKS).

    Networking and Storage: Configuring service meshes (like Istio/Envoy), ingress controllers (like NGINX), and persistent storage solutions (like volume provisioners).

    Resource Optimization: Managing resource requests and limits, implementing Horizontal Pod Autoscalers (HPA) and Cluster Autoscalers to balance cost and performance.

    2. Automation and CI/CD (The "Dev" Side)
    Pipeline Automation: Designing and maintaining Continuous Integration (CI) and Continuous Deployment (CD) pipelines (e.g., using GitLab CI, Jenkins, or ArgoCD) to automate the building, testing, and deployment of container images.

    Configuration as Code (IaC): Managing all Kubernetes configurations (Deployments, Services, ConfigMaps, etc.) using declarative tools like Helm or Kustomize.

    GitOps: Implementing GitOps workflows, where the entire state of the cluster is defined in Git and automatically reconciled (managed) by tools like ArgoCD or Flux.

    3. Monitoring and Observability
    Logging and Metrics: Setting up and managing observability stacks (Prometheus, Grafana, EFK/Loki) to monitor cluster health, application performance, and identify resource bottlenecks.

    Alerting: Defining and tuning alerts to notify teams of critical cluster failures or application performance degradation.

    In short, the Kubernetes DevOps Engineer is the architect and mechanic of the automated system that gets code from a developer's laptop into a production-ready container running seamlessly inside the cluster.

![ss-28](./images/ss-28.png)
![ss-29](./images/ss-29.png)

#### Site Reliability Engineer (SRE)

The Site Reliability Engineer (SRE) in the Kubernetes context is responsible for treating the Kubernetes cluster itself as a system that must be highly reliable, scalable, and observable. Their role is centered on applying software engineering principles to operations tasks to ensure the stability and performance of the platform that hosts the applications.

While a DevOps engineer often focuses on automation and delivery to the cluster, the SRE focuses on the reliability and maintenance of the cluster.

##### Core Responsibilities of an SRE on Kubernetes

    1. Defining and Enforcing Reliability
    SLOs and SLIs: Defining and tracking ​​Service Level Agreement (SLA), Service Level Indicators (SLIs) and Service Level Objectives (SLOs) for the Kubernetes control plane and core services (e.g., API server uptime, etcd latency).

    Capacity Planning: Forecasting future cluster needs (CPU, memory, storage) and planning horizontal scaling to prevent resource exhaustion.

    Incident Response: Designing and optimizing the process for responding to major incidents affecting the cluster's health.

    2. Operational Tooling and Automation
    Platform Engineering: Writing code and developing tools to automate routine operational tasks (e.g., node health checks, certificate rotation, cluster upgrades).

    Chaos Engineering: Designing and running controlled experiments (using tools like Chaos Mesh or Gremlin) to intentionally inject failure into the cluster to test its resilience and verify that automated failover mechanisms work.

    Security Automation: Automating security patching and compliance checks for the cluster and worker nodes.

    3. Observability and Monitoring
    Advanced Monitoring: Implementing and maintaining the core monitoring stack (Prometheus, Grafana, Alertmanager) not just for applications, but for the entire control plane (kube-apiserver, kube-controller-manager, scheduler, etcd).

    Logging and Tracing: Ensuring robust logging and tracing pipelines (e.g., using Loki, Jaeger) are in place to quickly diagnose systemic issues and latency spikes within the Kubernetes network.

    In short, the SRE is the guardian of the cluster's uptime, stability, and future performance, using data and engineering to reduce human toil and operational risk.

![ss-30](./images/ss-30.png)

#### CloudOps Engineer

A CloudOps Engineer in the Kubernetes context is a role focused on managing the underlying cloud infrastructure that hosts the Kubernetes cluster. They ensure the cluster has the necessary resources, stability, and cost efficiency within the chosen cloud provider (AWS, Azure, GCP, etc.).

Their primary job is to master the cloud provider's tools to support the Kubernetes platform, bridging the gap between traditional cloud resource management and the requirements of container orchestration.

##### Core Responsibilities

    1. Cloud Infrastructure Management
    Provisioning: Setting up the foundational compute (VMs/worker nodes), networking (VPCs, subnets, load balancers), and storage (EBS, S3, Azure Disks) resources needed for the Kubernetes cluster.

    Cost Optimization: Monitoring and managing the cloud bill. This includes selecting the most cost-effective machine types, utilizing reserved instances, and ensuring idle resources are scaled down or terminated.

    Cloud Security: Configuring and enforcing cloud-native security policies, such as IAM roles and policies, security groups, and virtual firewalls, which Kubernetes itself relies on.

    2. Networking and Identity
    External Access: Configuring cloud load balancers and domain name services (DNS) to properly route external traffic to the Kubernetes Ingress Controllers.

    Identity and Access Management (IAM): Integrating the Kubernetes cluster with the cloud provider's IAM system (e.g., AWS IAM, Azure AD) to handle authentication and authorization for services and users.

    3. Cluster Lifecycle Support
    Integration: Ensuring cloud-native services (like managed databases, caching layers, or cloud storage) are securely and efficiently integrated with applications running inside the cluster.

    Backup and Disaster Recovery: Implementing cloud-based backup solutions for persistent volumes (PVs) and Kubernetes etcd data to ensure business continuity.

    In essence, the CloudOps Engineer is the expert on the cloud platform's controls, ensuring the foundation upon which the DevOps and SRE teams build and maintain the applications is robust, secure, and budget-friendly.

#### Security Engineer

In the Kubernetes context, the Security Engineer is focused on securing the entire container ecosystem, from the underlying infrastructure to the application running inside the container. They implement security policies as code and ensure all components meet organizational compliance and threat mitigation standards.

They operate across the entire stack, working closely with SRE and DevOps teams to integrate security early into the development and deployment lifecycle (DevSecOps).

##### Core Responsibilities

    1. Cluster and Platform Hardening
    API Server Security: Securing the Kubernetes API server, enforcing strong authentication (e.g., RBAC), and network policies to restrict administrative access.

    Network Segmentation: Implementing strict Network Policies to control traffic flow between Pods (microservices), ensuring that only necessary services can communicate with each other.

    Secrets Management: Managing sensitive data (passwords, tokens, keys) using solutions like Vault or Kubernetes Secrets with robust encryption and access control.

    2. Supply Chain Security (Images and Pipelines)
    Image Scanning: Integrating automated vulnerability scanning tools (e.g., Clair, Trivy) into the CI/CD pipeline to analyze container images and block deployments if they contain known security flaws or outdated components.

    Registry Security: Ensuring the container image registry (e.g., Docker Hub, GitLab Registry) is secure and that images are digitally signed and verified before deployment.

    3. Runtime Protection
    Runtime Security Tools: Deploying and managing tools (like Falco or other behavioral monitoring systems) to detect and respond to suspicious activity inside running containers (e.g., unexpected shell execution, file system changes).

    Pod Security Standards (PSS) / Admission Controllers: Implementing and enforcing policies through Admission Controllers (like Kyverno or OPA Gatekeeper) to ensure new Pods adhere to security baselines before they are allowed to run (e.g., preventing containers from running as root).

    In short, the Security Engineer's goal is to embed security directly into the Kubernetes platform, making it difficult for vulnerabilities to be exploited and ensuring compliance is automated.

#### DevSecOps Engineer

A DevSecOps Engineer in the Kubernetes context is a specialized role focused on integrating security practices, tools, and automation directly into the entire Continuous Integration/Continuous Delivery (CI/CD) pipeline and the container lifecycle.

Their primary goal is to shift security "left," ensuring that security measures are addressed early and continuously—from the moment code is written to when it's running in production within the Kubernetes cluster.

##### Core Responsibilities

    1. Pipeline Automation and Security Gates
    SAST/DAST Integration: Embedding Static Application Security Testing (SAST) and Dynamic Application Security Testing (DAST) tools directly into the CI pipeline (e.g., in GitLab CI, Jenkins) to automatically scan code and running services for vulnerabilities before deployment.

    Vulnerability Scanning: Automating the scanning of all Docker container images (using tools like Trivy or Clair) and configuring the CI/CD pipeline to automatically block images with known Common Vulnerabilities and Exposures (CVEs) from being pushed to the registry or deployed.

    Compliance Checks: Enforcing security and compliance checks on infrastructure-as-code files (like Helm charts or Kubernetes YAMLs) using tools like Checkov or Kube-bench.

    2. Kubernetes Hardening as Code
    Admission Controllers: Working with Security Engineers to configure and manage Admission Controllers (like OPA Gatekeeper or Kyverno). These tools enforce security policies before a resource (like a Pod) is allowed into the cluster, ensuring all deployments comply with established Pod Security Standards (PSS).

    Secrets Management: Automating the secure injection of secrets (API keys, credentials) into Pods using solutions like HashiCorp Vault or cloud-native secrets managers, eliminating the storage of plain-text secrets in Git or ConfigMaps.

    3. Runtime Protection
    Security Observability: Configuring runtime security monitoring tools (like Falco) to detect suspicious system calls or unexpected behavior inside running containers.

    Alerting and Response: Integrating runtime detection tools with alerting systems to notify security teams immediately if a deployed application is compromised or begins acting maliciously.

    In essence, the DevSecOps Engineer ensures that the speed and automation benefits brought by DevOps are never achieved at the expense of security.

![ss-31](./images/ss-31.png)

#### Full Stack Developer

A Fullstack Developer in the Kubernetes context is an individual responsible for working across the entire application stack, from the frontend user interface (UI) down through the backend business logic and all the way to defining the deployment configuration for Kubernetes.

They possess a broad understanding of both application development and containerization concepts, allowing them to own a feature end-to-end.

##### Core Responsibilities

The Fullstack Developer's role in a Kubernetes environment requires skills that bridge traditional development and modern cloud-native deployment.

    1. Application Development (Frontend & Backend)
    Frontend: Building and maintaining the UI (using frameworks like React, Vue, or Angular).

    Backend: Writing the API and business logic (using languages like Go, Python, or Node.js).

    Containerization: Writing the Dockerfile for both the frontend and backend services to ensure they can be packaged and run as standard containers.

    2. Infrastructure Interface
    Configuration Files: Writing and maintaining the Kubernetes manifest files (YAML) or configuration templates (Helm charts, Kustomize) for their specific service. This includes defining:

    Deployments: How their service is rolled out.

    Services: How their service is exposed within the cluster.

    Ingress Rules: How external traffic reaches their service.

    Local Development: Ensuring their local development environment (often using tools like Minikube or Docker Compose) closely mirrors the Kubernetes production environment for reliable testing.

    3. CI/CD Integration
    Pipeline Input: Working with DevSecOps/DevOps engineers to define the steps needed to build and test their service's container image, ensuring the continuous integration (CI) pipeline functions correctly with their codebase.

    In short, a Fullstack Developer in this context isn't just coding features; they are defining how their features are packaged, deployed, and managed by the orchestration system. They own the feature from the first line of code to its successful running state in the Kubernetes cluster.

![ss-32](./images/ss-32.png)

#### Cloud Architect
A Cloud Architect in the Kubernetes context is a high-level strategist and designer responsible for determining how an organization uses its cloud platform to host and scale the Kubernetes environment. They focus on the big picture: security, cost, compliance, and multi-region strategy.

They are the primary decision-makers regarding which cloud services are utilized, how they are connected, and how the entire Kubernetes ecosystem aligns with the business goals.

##### Core Responsibilities

    The Cloud Architect works above the cluster layer, designing the robust, scalable foundation upon which the SRE and DevOps teams build.

    1. Strategic Design and Decision-Making
    Platform Choice: Deciding which managed Kubernetes offering to use (e.g., AWS EKS, GCP GKE, Azure AKS) or if a self-managed solution is necessary.

    Multi-Cloud/Hybrid Strategy: Designing solutions that span multiple cloud providers or integrate seamlessly with on-premises infrastructure (hybrid cloud), ensuring Kubernetes is portable.

    Networking Architecture: Designing the overall Virtual Private Cloud (VPC), subnet, and firewall strategy to ensure secure and efficient network communication for the cluster.

    2. Governance, Security, and Compliance
    Cost Management: Selecting appropriate cloud resources and instance types to meet performance goals while adhering to budget constraints, often by defining a comprehensive FinOps strategy.

    Cloud IAM Integration: Defining the security model for user and service access, ensuring Kubernetes Role-Based Access Control (RBAC) is correctly mapped to the cloud provider's Identity and Access Management (IAM) system.

    Compliance: Ensuring the cloud environment and the Kubernetes cluster meet industry-specific regulatory requirements (e.g., GDPR, HIPAA, SOC 2).

    3. High Availability and Disaster Recovery (DR)
    Designing the topology for the cluster to ensure high availability across multiple availability zones or regions, minimizing downtime risks.

    Establishing comprehensive cloud-native backup and recovery procedures for the cluster and its persistent data.

![ss-33](./images/ss-33.png)

#### Data Engineer
A Data Engineer in the Kubernetes context is primarily responsible for designing, building, and maintaining the infrastructure and pipelines that handle large-scale data movement, processing, and storage using cloud-native tools.

Their focus is on ensuring data is reliable, accessible, and efficiently processed by applications running within the cluster.

##### Core Responsibilities

    1. Data Pipeline Orchestration
    Workflow Management: Deploying and managing data orchestration tools like Apache Airflow, Prefect, or Kubeflow on top of Kubernetes. They containerize data processing jobs and manage their dependencies, scheduling, and scaling within the cluster.

    Streaming and Batch Processing: Implementing and maintaining frameworks like Spark, Kafka, or Flink (often deployed via Kubernetes Operators) to handle real-time (streaming) and large-volume (batch) data processing tasks.

    2. Storage and Persistence
    Data Persistence Strategy: Deciding how data should be stored. This often involves integrating Kubernetes with external, high-performance cloud storage services (like AWS S3 or Google Cloud Storage) or dedicated cluster-internal systems (like distributed file systems or managed databases).

    Volume Management: Working with persistent volumes (PVs) and persistent volume claims (PVCs) for applications that require state, ensuring data is backed up and resilient.

    3. Monitoring and Efficiency
    Resource Optimization: Tuning data processing jobs (e.g., Spark executors) to efficiently use the allocated CPU and memory resources provided by the Kubernetes cluster, preventing unnecessary resource consumption.

    Data Quality and Observability: Setting up monitoring and alerting for data pipelines to ensure data integrity, timely processing, and to quickly detect data-related failures.

    In short, the Data Engineer uses Kubernetes as the foundational platform to deploy and scale the complex, data-intensive workloads that power the organization's analytics and data-driven applications.

![ss-34](./images/ss-34.png)

#### FinOps Engineer

FinOps or Cloud Financial Management is a practice that aims to bring financial accountability to the variable spend model of the cloud, enabling distributed teams to make business trade-offs between speed, cost, and quality.

FinOps Engineers can help a company better understand and control cloud costs, making cloud expenses more predictable and manageable.

This role intersects between the financial, business, and IT departments within a company. It's a multidisciplinary field requiring knowledge of cloud computing, financial management, and often a solid understanding of various business operations. The role demands an understanding of cost allocation, cost optimisation strategies, and the ability to create financial models to predict future spend.

##### Typical skills of a FinOps Engineer may include:

* Cloud Cost Management and Optimisation
* Financial Reporting and Analysis
* Budgeting and Forecasting
* Knowledge of various Cloud Service Providers
* Communication and Collaboration

#### Machine Learning Engineer
Machine Learning Engineers are computer programmers, but their focus goes beyond specifically programming machines to perform specific tasks. They create programs that will enable machines to take actions without being specifically directed to perform those tasks. This kind of artificial intelligence is known as machine learning.

Machine Learning Engineers utilise large amounts of data and let the machine process that data, make decisions, and learn from those decisions. Over time, as the machine processes more data, it is able to make more accurate predictions and smarter decisions, hence learning.

##### Typical skills of a Machine Learning Engineer may include:

* Programming (Python, R, Java)
* Data Modelling and Evaluation
* Machine Learning Algorithms
* Advanced Mathematics
* Distributed Computing

#### Data Scientist

Data Scientists analyse and interpret complex digital data to assist a company in its decision-making processes. They apply strong business acumen, coupled with an in-depth understanding of data, to provide insights and recommendations to the business.

A Data Scientist's role can involve predicting future demands, identifying business trends and patterns, and helping senior management make informed business decisions. They use a range of tools and methodologies to handle, interpret, and model data.

##### Typical skills of a Data Scientist may include:

* Programming (Python, R)
* Statistics and Machine Learning
* Data Wrangling and Data Visualization
* Big Data Platforms (like Hadoop and Spark)
* SQL Database/Coding

_These roles, along with the ones covered previously, comprise some of the essential tech roles in the industry today. Each role has its unique set of responsibilities and skills required, yet they often overlap and collaborate to drive the success of the organisation's objectives. Understanding the differences and similarities between these roles is crucial for effective team building and project management in any tech-oriented organisation._

### Open Standards

Open standards in the Kubernetes context refer to the publicly available, non-proprietary specifications and protocols that govern how different components of the container ecosystem interact. These standards ensure interoperability, portability, and prevent vendor lock-in.

Kubernetes itself is built on a foundation of open standards and is governed by the Cloud Native Computing Foundation (CNCF), which promotes many of the key standards used within the ecosystem.

#### Key Open Standards and Specifications

    1. Containerization Standards
    These standards ensure that any container engine can run any container image, and that Kubernetes can reliably manage them.

    Open Container Initiative (OCI): This is perhaps the most critical standard. It governs the specification for two core things:

    Image Format Specification: Defines how a container image must be built and packaged (e.g., Docker images are now OCI images).

    Runtime Specification: Defines how a container runtime (like runc) must execute the container.

    Container Runtime Interface (CRI): This is a Kubernetes standard. It defines the API interface between the kubelet (the node agent) and the actual container runtime (like containerd or CRI-O).

    2. Networking and Storage
    These standards ensure that networking and persistent storage solutions can be easily plugged into any Kubernetes cluster.

    Container Network Interface (CNI): This is a standard specification for writing network plugins. It ensures that various networking solutions (like Calico, Flannel, or Cilium) can be consistently integrated with Kubernetes.

    Container Storage Interface (CSI): This is a standard specification for exposing arbitrary storage systems (like AWS EBS, NFS, or local storage) to container orchestration systems. It allows any storage provider to create a driver that works across different cloud environments.

    3. Service Mesh and Observability
    OpenTracing/OpenTelemetry (OTel): This is a standard set of APIs, libraries, and instrumentation tools for generating, collecting, and exporting telemetry data (metrics, logs, and traces). This ensures consistent observability across microservices regardless of the programming language or vendor.

    By adhering to these open standards, Kubernetes maintains its strong position as a portable and extensible platform.


![ss-35](./images/ss-35.png)
![ss-36](./images/ss-36.png)
![ss-37](./images/ss-37.png)
![ss-38](./images/ss-38.png)
![ss-39](./images/ss-39.png)
![ss-40](./images/ss-40.png)

## Section 3: Containers with Docker

![ss-41](./images/ss-41.png)

### Mainframe Virtualisation

![ss-42](./images/ss-42.png)

### Chroot

![ss-43](./images/ss-43.png)

### FreeBSD Jails

![ss-44](./images/ss-44.png)

### Linux Advancements

![ss-45](./images/ss-45.png)
![ss-46](./images/ss-46.png)

### Virtual Machine Era

![ss-47](./images/ss-47.png)

### Docker

![ss-48](./images/ss-48.png)

### Traditional Docker

![ss-49](./images/ss-49.png)

### Docker Desktop

![ss-50](./images/ss-50.png)

#### docker run

The command docker run -it ubuntu bash uses Docker to start a new Ubuntu container and immediately connects you to that container's command line (bash shell).

This command is typically used to quickly create an isolated environment for testing or debugging.

![ss-51](./images/ss-51.png)

This command provides a temporary and isolated environment. When you exit the container (by typing exit), the container is stopped, and any changes you made inside it (like creating files) are lost (unless you used a volume).

This makes it ideal for testing a Linux command, experimenting with software, or debugging in a clean, isolated environment.

#### kubectl get nodes

![ss-52](./images/ss-52.png)

### Container Images

![ss-53](./images/ss-53.png)
![ss-54](./images/ss-54.png)
![ss-55](./images/ss-55.png)
![ss-56](./images/ss-56.png)

**Dockerhub:** Enterprise-grade Docker images with built-in security, compliance, and continuous updates. Minimize vulnerabilities and deploy with confidence. We can push our docker images to container registry.

* https://hub.docker.com/
* https://hub.docker.com/r/wernight/funbox

![ss-57](./images/ss-57.png)
![ss-58](./images/ss-58.png)
![ss-59](./images/ss-59.png)
![ss-60](./images/ss-60.png)
![ss-61](./images/ss-61.png)
![ss-62](./images/ss-62.png)
![ss-63](./images/ss-63.png)
![ss-64](./images/ss-64.png)

**digest:** hash of the image, will be created when the image pushed to the registry

**image id:** checksum based on the local container image used by docker for this image

![ss-65](./images/ss-65.png)
![ss-66](./images/ss-66.png)
![ss-67](./images/ss-67.png)
![ss-68](./images/ss-68.png)
![ss-69](./images/ss-69.png)

### Running Containers

Docker client and server running independently

![ss-70](./images/ss-70.png)
![ss-71](./images/ss-71.png)

**docker ps:** list containers

**docker ps -a:** show all containers (default shows just running)

![ss-72](./images/ss-72.png)

**docker rm <container_name>:** remove a container by name

**docker run –rm:** remove container after exit

**containerd:** container engine used by Docker

**docker run -it –rm <image> <command>:** override the default command in a Docker container

**-it:** runs the container in interactive mode with a terminal

**docker –help:** list all the parameters available to the Docker command

**The purpose of running a container as a non-root user:** to improve security by limiting privileges

### Container Network Services and Volumes

**docker run -d –rm nginx:** the container will be working in the background, output we get a container id

![ss-72](./images/ss-72.png)
![ss-73](./images/ss-73.png)
![ss-74](./images/ss-74.png)
![ss-75](./images/ss-75.png)
![ss-76](./images/ss-76.png)
![ss-77](./images/ss-77.png)
![ss-78](./images/ss-78.png)
![ss-79](./images/ss-79.png)

**specify the port**

![ss-80](./images/ss-80.png)

**docker exec:** run a command in a running container

![ss-81](./images/ss-81.png)
![ss-82](./images/ss-82.png)

### Volume:

![ss-83](./images/ss-83.png)