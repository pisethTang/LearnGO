# Prerequisites

- Docker containerizations. 


# Kubernetes

From my understanding so far, Kubernetes and Docker are virtualization software and so they share a lot of similarities except for the fact that Kubernetes are usually used in more enterprise-level environments (when dealing with a large volume of continuous stream of data). 


## History
The goal (as it has been this entire time) for tech companies is to deploy their apps to their clients as quickly, efficiently and safely as possible. 


Monolith (monorepo -- one codebase) --> microservice (each service can be written in different languages and each can be scaled individually & independently.)

virtualization (taking a slice of major components to run an application) as a means to maximize CPU utilization. 

Cloud computing came into the scene due to many factors. One primary factor is that small companies or startups were thinking of how they can deploy their applications to clients. Large enterprises (like IBM, Oracle, etc) had already been doing that (although quite pricey and expensive.)


Cloud providers: AWS, Azure, GCP, Exoscale (for startups).

- Kubernetes (K8s): orchestrator engine to manage containerized applications (but it's actually even more than this -- extensible.). K8s can manage other resources (not just containerized apps, it can manage other K8s)
- automatic scheduling
- mix of controllers to run applications. 




# Brief
DC -> virtualization -> Clouds -> Containers (-> WASM)



# Objectives
1. Application
2. Deployment services: ingress
3. Gateway api (service manager)
4. All K8 constructs. 
5. Create clusters (self-managed and Exoscale-managed)







# References
1. [Learn Kubernetes in 6 Hours – Full Course with Real-World Project
](https://www.youtube.com/watch?v=_4uQI4ihGVU&t=2666s)
2. [DevOps Directive](https://www.youtube.com/@DevOpsDirective)







