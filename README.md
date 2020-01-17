# Project Layout
API is designed with Clean Architecture.  

![Clean Architecture](https://miro.medium.com/max/772/1*B7LkQDyDqLN3rRSrNYkETA.jpeg) 

# Directory
Directory structre of api follows Standard Go Project Layout.  
https://github.com/golang-standards/project-layout

```
├── api - OpenAPI/Swagger specs, JSON schema files, protocol definition files.
├── build - Packaging and Continuous Integration.
├── cmd - Main applications for this project.
├── configs - Configuration files.
├── deployments - Confiuration files for deployment
├── internal - Private application and library code.
│   ├── domain - Domain layer
│   ├── infrastracture - Infrastracture layer
│   ├── interface - Interface layer
│   ├── registry - DI container
│   └── usecase - Usecase files
├── pkg - Library code that's ok to use by external applications.
└── test - Testing files. 
```
