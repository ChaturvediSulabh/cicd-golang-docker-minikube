## Introduction

create a simple HTTP service that stores and returns configurations that satisfy certain conditions. Then, the service should be automatically deployed to kubernetes.

### Endpoints


| Name   | Method      | URL
| ---    | ---         | ---
| List   | `GET`       | `/configs`
| Create | `POST`      | `/configs`
| Get    | `GET`       | `/configs/{name}`
| Update | `PUT/PATCH` | `/configs/{name}`
| Delete | `DELETE`    | `/configs/{name}`
| Query  | `GET`       | `/search?name={config_name}&data.{key}={value}`


#### Schema

- **Config**
  - Name (string)
  - Data (key:value pairs)
  
  
#### How To ?
- Golang as I find it easy to understand and moreover, I've been writing only Golang since, last 6 months. 
- Docker (containerze the app (build image and publish to docker registry))
- Minikube Cluster (I had this on my machine pre-installed)
- [Drone CI](https://drone.io)
   - Easy to use/adapt
   - Built in plugins (Docker, Kubernetes, AWS, GCP, Openstack ... )
   - OSS and community driven
   - You can write your own plugins
   - You'll need to create an account on Drone CI and then, sync and activate your GitHUB repository. Drone CI will read from .drone.yml.
   - In .drone.yml you will some secrets - In Drone CI go to settings of your synchronized repos and setup these secrets. The secrets are straight forward  like Docker Username, Docker password, Kubernetes Server (minikube url), Kubernetes CA.pem (```echo $(cat ~/.minikube/ca.pem | base64)```), kubernetes token (```echo $(kubectl get secret $(kubectl get serviceaccount default -o jsonpath='{.secrets[0].name}') -o jsonpath='{.data.token}' | base64 --decode )```)
   - You can read more about drone ci plugins that I've used by clicking following respective URL
     - [docker plugin](http://plugins.drone.io/drone-plugins/drone-docker/)
     - [kubernetes plugin](https://github.com/vallard/drone-kube/blob/master/DOCS.md)
   - I have used NodePort `service`in  `./kubernetes/manifests/deployment.yml` to expose my deployment externally.
