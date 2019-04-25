[![Build Status](https://cloud.drone.io/api/badges/ChaturvediSulabh/CICDWithGolangDockerMinikube/status.svg)](https://cloud.drone.io/ChaturvediSulabh/CICDWithGolangDockerMinikube)
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

