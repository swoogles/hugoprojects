+++
title = "Put some Swagger in your step"
outputs = ["Reveal"]
+++
# Let's talk about swagger

---
Document your APIs
    {{% fragment %}}in a way that we can actually trust{{% /fragment %}}

---
### Do you trust that...
- {{% fragment %}}Engineers will find that related wiki page?{{% /fragment %}}
- {{% fragment %}}Engineers will *read* that related wiki page?{{% /fragment %}}
- {{% fragment %}}This distant page will stay in sync as code changes?{{% /fragment %}}

---
## Enter Swagger

---
### Components
- OpenAPI Specification (Formerly SwaggerSpec)
- Gradle Plugin
- UI
- Codegen
- Editor
- Hub

---
### OpenAPI Specification
- YAML or JSON
- {{% fragment %}}Describes all facets of your API{{% /fragment %}}
    - {{% fragment %}}URL{{% /fragment %}}
    - {{% fragment %}}Parameters{{% /fragment %}}
    - {{% fragment %}}Success/Error return values{{% /fragment %}}
    - {{% fragment %}}Models/Entities/DTOs{{% /fragment %}}
    - {{% fragment %}}Examples{{% /fragment %}}
- {{% fragment %}}Achievement: Full MedClaims description via spec file: [IDE_DEMO] {{% /fragment %}}

---
### Gradle plugin
- {{% fragment %}}Generate spec files from existing projects{{% /fragment %}}
- {{% fragment %}}Achievement: Generate Records service spec [BROWSER DEMO]{{% /fragment %}}
    
---
### SwaggerUI
- {{% fragment %}}Browser UI for OAS spec files {{% /fragment %}}
- {{% fragment %}}Achievement: Browse Reports Microservice API via local spec file[BROWSER DEMO]{{% /fragment %}}

---
### Swagger Codegen
- {{% fragment %}}Generate server stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} JAXRS, Jersey, RESTEasy, Spring {{% /fragment %}}
    - {{% fragment %}} No struts support :( {{% /fragment %}}
- {{% fragment %}}Generate client stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} Typescript/Angular {{% /fragment %}}
- {{% fragment %}}Awesome{{% /fragment %}}
- {{% fragment %}}Achievement: Generate all Server/Client types that CMT cares about [IDE DEMO]{{% /fragment %}}
- {{% fragment %}}Achievement: Generate MedClaims server that matches *live* implementation[IDE DEMO]{{% /fragment %}}

---
### SwaggerHub
- {{% fragment %}}Ties all of these projects together{{% /fragment %}}
- {{% fragment %}}Source of truth for API documentation{{% /fragment %}}
- {{% fragment %}}Achievement: Browse MedClaims & Reports from the same hub{{% /fragment %}}

---
### The Dream
- {{% fragment %}}Each service generates/publishes OAS specs to SwaggerHub{{% /fragment %}}
- {{% fragment %}}Front-end & Back-end engineers start *and continue* with API spec{{% /fragment %}}
- {{% fragment %}}We only write logic code.{{% /fragment %}}
