+++
title = "Put some Swagger in your step"
outputs = ["Reveal"]
+++
# Let's talk about Swagger

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
# 5 Components

---
### 1. OpenAPI Specification
- {{% fragment %}}YAML or JSON{{% /fragment %}}
- {{% fragment %}}Documents all facets of your API{{% /fragment %}}
    - {{% fragment %}}URLs & REST verbs{{% /fragment %}}
    - {{% fragment %}}Parameters{{% /fragment %}}
    - {{% fragment %}}Return values/errors{{% /fragment %}}
    - {{% fragment %}}Models/DTOs{{% /fragment %}}
    - {{% fragment %}}In-line Examples{{% /fragment %}}
- {{% fragment %}}Achievement: Create Full MedClaims spec file [IDE_DEMO] {{% /fragment %}}

---
### 2. Gradle plugin
- {{% fragment %}}Generate spec files from existing projects{{% /fragment %}}
- {{% fragment %}}Achievement: Full Records service spec{{% /fragment %}}
    
---
### 3. SwaggerUI
- {{% fragment %}}Browser UI for OAS spec files {{% /fragment %}}
- {{% fragment %}}Achievement: Browse Reports Microservice API[BROWSER DEMO]{{% /fragment %}}

---
### 4. Swagger Codegen
- {{% fragment %}}Generate client/server stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} JAXRS, Jersey, RESTEasy, Spring {{% /fragment %}}
    - {{% fragment %}} Typescript/Angular {{% /fragment %}}
    - {{% fragment %}} No struts support :( {{% /fragment %}}
- {{% fragment %}}Achievement: Generate MedClaims API that matches *live* implementation{{% /fragment %}}
    {{% fragment %}}with all relevant Servers/Clients[IDE DEMO]{{% /fragment %}}

---
### 5. SwaggerHub
- {{% fragment %}}Ties all of these projects together{{% /fragment %}}
- {{% fragment %}}Source of truth for API documentation{{% /fragment %}}
- {{% fragment %}}Achievement: Browse MedClaims & Reports from the same hub[BROWSER DEMO]{{% /fragment %}}

---
### The Dream
- {{% fragment %}}Each service generates/publishes OAS specs to SwaggerHub{{% /fragment %}}
- {{% fragment %}}Front-end & Back-end engineers start *and continue* with API spec{{% /fragment %}}
- {{% fragment %}}We only write/edit logic code.{{% /fragment %}}
