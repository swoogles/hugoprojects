+++
title = "Put some Swagger in your step"
outputs = ["Reveal"]
+++
# Let's talk about swagger

---
Document your APIs
    {{% fragment %}}in a way that we can actually trust{{% /fragment %}}

---
### Do you trust that Engineers will...
- {{% fragment %}}Find that related wiki page?{{% /fragment %}}
- {{% fragment %}}*Read* that related wiki page?{{% /fragment %}}
- {{% fragment %}}__*Maintain*__ that related wiki page?!{{% /fragment %}}

---
![Example image5](/images/Swagger-logo.png)


---
### Components
- OpenAPI Specification
- Gradle Plugin
- UI
- Codegen
- Editor
- Hub

---
### OpenAPI Specification
- YAML or JSON
- {{% fragment %}}Documents all facets of your API{{% /fragment %}}
    - {{% fragment %}}URL{{% /fragment %}}
    - {{% fragment %}}Parameters{{% /fragment %}}
    - {{% fragment %}}Success/Error return values{{% /fragment %}}
    - {{% fragment %}}Models/Entities/DTOs{{% /fragment %}}
    - {{% fragment %}}Examples{{% /fragment %}}
- {{% fragment %}}Achievement: Create Full MedClaims spec file: [IDE_DEMO] {{% /fragment %}}

---
### Gradle plugin
- {{% fragment %}}Generate spec files from existing projects{{% /fragment %}}
- {{% fragment %}}Achievement: Full Records service spec [BROWSER DEMO]{{% /fragment %}}
    
---
### SwaggerUI
- {{% fragment %}}Browser UI for OAS spec files {{% /fragment %}}
- {{% fragment %}}Achievement: Browse Reports Microservice API via local spec file[BROWSER DEMO]{{% /fragment %}}

---
### Swagger Codegen
- {{% fragment %}}Generate client/server stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} JAXRS, Jersey, RESTEasy, Spring {{% /fragment %}}
    - {{% fragment %}} Typescript/Angular {{% /fragment %}}
    - {{% fragment %}} No struts support :( {{% /fragment %}}
- {{% fragment %}}Achievement: Generate MedClaims API that matches *live* implementation{{% /fragment %}}
    {{% fragment %}}with all relevant Server/Client implementations[IDE DEMO]{{% /fragment %}}

---
### SwaggerHub
- {{% fragment %}}Ties all of these projects together{{% /fragment %}}
- {{% fragment %}}Source of truth for API documentation{{% /fragment %}}
- {{% fragment %}}Achievement: Browse MedClaims & Reports from the same hub[BROWSER DEMO]{{% /fragment %}}

---
### The Dream
- {{% fragment %}}Each service generates/publishes OAS specs to SwaggerHub{{% /fragment %}}
- {{% fragment %}}Front-end & Back-end engineers start *and continue* with API spec{{% /fragment %}}
- {{% fragment %}}We only write/edit logic code.{{% /fragment %}}
