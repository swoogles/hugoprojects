+++
title = "Put some swagger in your step"
outputs = ["Reveal"]
+++
# Let's talk about swagger

---
### Swagger Components
- OpenAPI Specification (Formerly SwaggerSpec)
- SwaggerUI
- SwaggerCodegen
- SwaggerEditor
- SwaggerHub

---
### OpenAPI Specification
- YAML or JSON
- {{% fragment %}}Describes all facets of your API{{% /fragment %}}
    - {{% fragment %}}URL{{% /fragment %}}
    - {{% fragment %}}Parameters{{% /fragment %}}
    - {{% fragment %}}Success/Error return values{{% /fragment %}}
    - {{% fragment %}}Models{{% /fragment %}}
    - {{% fragment %}}Examples{{% /fragment %}}
- {{% fragment %}}Achievement: Completely re-implement MedClaims via spec file: $URL_GOES_HERE {{% /fragment %}}
    
---
### SwaggerUI
- Turns OAS spec files into browsable UI
- Achievement: Browse Reports Microservice API via local spec file

---
### Swagger Codegen
- {{% fragment %}}Generate server stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} JAXRS {{% /fragment %}}
    - {{% fragment %}} RESTEasy{{% /fragment %}}
    - {{% fragment %}} Spring {{% /fragment %}}
- {{% fragment %}}Generate client stubs from OAS spec{{% /fragment %}}
    - {{% fragment %}} Typescript/Angular {{% /fragment %}}
    - {{% fragment %}} No struts support :/ {{% /fragment %}}
- {{% fragment %}}Really cool{{% /fragment %}}
- {{% fragment %}}Really powerful{{% /fragment %}}


---
### SwaggerEditor
- Immediate OAS violation reporting
- Auto-complete

---
### SwaggerHub
- Ties all of these projects together
- Source of truth for API interfaces
- 

