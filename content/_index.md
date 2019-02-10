+++
title = "My presentation"
outputs = ["Reveal"]
+++

# Static Shock
or, how a single request can generate 800,000 exceptions.

---
# Troubling Symptoms
- {{% fragment %}}DB Degradation{{% /fragment %}}
- {{% fragment %}}Ram Exhaustion{{% /fragment %}}
- {{% fragment %}}Full Hard Drives!{{% /fragment %}}


---
# Analyzing (Turn into separate slides)
- {{% fragment %}}"I see something about Consent" - Austin{{% /fragment %}}
- {{% fragment %}}Errors in FacilityConfig{{% /fragment %}}
- {{% fragment %}}Can't get Facility null{{% /fragment %}}

---
{{<mermaid>}}
graph TB
subgraph 
CdiRequest
CdiRequest-->InjectedDependencyA
CdiRequest-->InjectedDependencyB
InjectedDependencyB-->FacilityConfigLogic
end
subgraph 
StaticRequest-->StaticDependencyA
StaticRequest-->StaticDependencyB
StaticDependencyB-->FacilityConfigLogic
end
{{</mermaid>}}

---
{{<mermaid align="left">}}
graph LR;
  A[Hard edge] -->|Link text| B(Round edge)
  B --> C{Decision}
  C -->|One| D[Result one]
  C -->|Two| E[Result two]
{{< /mermaid >}}

