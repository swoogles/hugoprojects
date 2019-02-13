+++
title = "My presentation"
outputs = ["Reveal"]
+++

# Static Shock
or, how a single request can generate 800,000 exceptions.

---
## Troubling Symptoms
- {{% fragment %}}DB Degradation{{% /fragment %}}
- {{% fragment %}}Ram Exhaustion{{% /fragment %}}
- {{% fragment %}}Full Hard Drives!{{% /fragment %}}


---
{{% section %}}
# Analyzing

---
### "I see something about Consent" - Austin

---
### Errors in FacilityConfig

---
### Can't get Facility null

{{% /section %}}



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

  InjectedDependencyA
  subgraph 
    CdiRequest
    CdiRequest-->InjectedDependencyB
    CdiRequest-->InjectedDependencyA
  end

  InjectedDependencyB-->StaticConfigAccess
  InjectedConfigAccess-->threadLocalConnection
end
{{</mermaid>}}

---
{{<mermaid>}}
graph TB
  CensusResource-->Encounters
  Encounters-->GenericSharingLogic
  GenericSharingLogic-->FacilityConsentLogic
  FacilityConsentLogic-->FacilityCategories
  FacilityCategories-->FacilityConfigService
  FacilityConfigService-->threadLocalSession
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
### Actual, Bad Categorizer Behavior
{{<mermaid>}}
sequenceDiagram
  participant CDI as CdiRequest
  participant CC as CDICategorizer
  participant STAT as StaticRequest
  participant CAT as StaticCategorizer

  activate CAT

  STAT->>CAT: Categorize this
  activate STAT
  CAT-->>STAT: Categorized
  deactivate STAT

  Note over CC: Never Used!
  CDI->>CAT: Categorize this
  activate CDI
  CAT-->>CDI: Categorized
  deactivate CAT
  deactivate CDI
  Note right of CAT: CDI kills instance

  STAT->>CAT: Categorize this
  activate STAT
  CAT-->>STAT: NPE
  deactivate STAT

  CDI->>CAT: Categorize this
  activate CDI
  CAT-->>CDI: NPE
  deactivate CDI
{{</mermaid>}}

---
### Static Categorizer Handling
{{<mermaid>}}
sequenceDiagram
  participant CDI as CdiRequest
  participant STAT as StaticRequest
  participant CAT as StaticCategorizer
  STAT->>CAT: Categorize this
  activate CAT
  CAT-->>STAT: Categorized
  CDI->>CAT: Categorize this
  CAT-->>CDI: Categorized
  deactivate CAT
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

