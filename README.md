# golang-basic-backend
Basic backend web project for learning GOLANG

## Introduction
This project is made as part of the learning curve presented by course Aprende lenguaje GO (GOLANG) desde 0 (Versión GO 1.20) on udemy. The objective is to develop and deploy a serverless web API for handling message board entities close to the social network "Twitter"

---
# Warning!
After starting the development of the mongoDB component, the course's coupled design becomes so incompatible with my moddifications, I currently have not the time, nor the willpower to redesign this project in any way to make it better, so onwards expect a lot of coupling and bad practices... I'm sorry for what you are going to see going forward, but getting to the end takes priority here. 

---

## Tools
| Name | Website | Purpose |
|------|---------|---------|
| AWS Lambda | aws.amazon.com | Serverless deployment |
| AWS API Gateway | aws.amazon.com | API Gateway for public access to private APIs |
| AWS S3 | aws.amazon.com | Cloud storage solution for image persistence |
| MongoDB | mongodb.com | Json based file database |

---
## Packages
### AWS
General package for AWS implementation of the handler for the whole back end features
### DB
Package that contains implementations of the databanager interface for especific technologies
### JWT
Package that contains types and functions relates with json web token management and operations
### Shared
Package for shared functionalities and models that can be used around the whole solution
### Interfaces
Package containing only declarations of interfaces used on every other package of the solution