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
    subgraph 
    StaticRequest-->StaticDependencyA
    StaticRequest-->StaticDependencyB
    StaticDependencyB-->StaticConfigAccess
    StaticConfigAccess-->threadLocalConnection
    StaticConfigAccess-.->threadLocalCdiConnection
  end

  subgraph 
    CdiRequest
    CdiRequest-->InjectedDependencyA
    CdiRequest-->InjectedDependencyB
  end

  InjectedDependencyB-->StaticConfigAccess
  InjectedConfigAccess-->threadLocalConnection
end
{{</mermaid>}}



---
### Intended Connection Behavior
{{<mermaid>}}
sequenceDiagram
  participant CP as ConnectionPool
  participant CDI as CdiRequest
  participant STAT as StaticRequest
  STAT->>CP: Give me a connection
  CP-->>STAT: Here's connectionA
  activate STAT
  CDI->>CP: Give me a connection
  CP-->>CDI: Here's connectionB
  activate CDI
  CDI-->STAT: Non-conflicting action
  STAT->>CP: Done with connectionA
  deactivate STAT
  CDI->>CP: Done with connectionA
  deactivate CDI
{{</mermaid>}}






---
Junk drawer from here on out.

---
### Sequence Diagram Junk
{{<mermaid>}}
sequenceDiagram
    Alice->John: Hello John, how are you?
    Note over Alice,John: A typical interaction
{{</mermaid>}}

---
{{<mermaid>}}
graph LR
  id1(Start)-->id2(Stop)
  style id1 fill:#f9f,stroke:#333,stroke-width:8px
  style id2 fill:#ccf,stroke:#f66,stroke-width:8px,stroke-dasharray: 5, 5
{{</mermaid>}}

---
{{<mermaid align="left">}}
graph LR;
  A[Hard edge] -->|Link text| B(Round edge)
  B --> C{Decision}
  C -->|One| D[Result one]
  C -->|Two| E[Result two]
{{< /mermaid >}}

