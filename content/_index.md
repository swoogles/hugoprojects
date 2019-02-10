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
{{<mermaid align="left">}}
graph LR;
  A[Hard edge] -->|Link text| B(Round edge)
  B --> C{Decision}
  C -->|One| D[Result one]
  C -->|Two| E[Result two]
{{< /mermaid >}}

---
{{<mermaid>}}
  graph TD
  Start --> Stop
{{</mermaid>}}

---
{{<mermaid>}}
graph TB
subgraph 
CdiRequest-->FacilityConfigLogic
CdiRequest-->InjectedDependency
end
subgraph 
StaticRequest-->FacilityConfigLogic
StaticRequest-->StaticDependency
end
{{</mermaid>}}
