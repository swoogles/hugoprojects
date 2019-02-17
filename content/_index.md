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
### Actual, Bad Categorizer Behavior
{{<mermaid>}}
sequenceDiagram
  participant StaticRequest
  participant StaticCategorizer
  participant CDICategorizer
  participant CdiRequest

  activate StaticCategorizer

  StaticRequest->>StaticCategorizer: Categorize this
  activate StaticRequest
  StaticCategorizer-->>StaticRequest: Categorized
  deactivate StaticRequest

  Note over CDICategorizer: Never Used!
  CdiRequest->>StaticCategorizer: Categorize this
  activate CdiRequest
  StaticCategorizer-->>CdiRequest: Categorized
  deactivate StaticCategorizer
  deactivate CdiRequest
  Note left of StaticCategorizer: CDI cleanup

  StaticRequest->>StaticCategorizer: Categorize this
  activate StaticRequest
  StaticCategorizer-->>StaticRequest: NPE
  deactivate StaticRequest

  CdiRequest->>StaticCategorizer: Categorize this
  activate CdiRequest
  StaticCategorizer-->>CdiRequest: NPE
  deactivate CdiRequest
{{</mermaid>}}

---
### Desired Categorizer Behavior(FIX!!)
{{<mermaid>}}
sequenceDiagram
  participant StaticRequest
  participant StaticCategorizer
  participant CDICategorizer
  participant CdiRequest

  activate StaticCategorizer
  activate CDICategorizer

  StaticRequest->>StaticCategorizer: Categorize this
  activate StaticRequest
  StaticCategorizer-->>StaticRequest: Categorized
  deactivate StaticRequest

  CdiRequest->>CDICategorizer: Categorize this
  activate CdiRequest
  CDICategorizer-->>CdiRequest: Categorized
  deactivate CdiRequest

  StaticRequest->>StaticCategorizer: Categorize this
  activate StaticRequest
  StaticCategorizer-->>StaticRequest: Categorized
  deactivate StaticRequest

  CdiRequest->>CDICategorizer: Categorize this
  activate CdiRequest
  CDICategorizer-->>CdiRequest: Categorized
  deactivate CdiRequest

  deactivate StaticCategorizer
  deactivate CDICategorizer
{{</mermaid>}}


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
# Attempt #2

---
{{% section %}}
# Deploying

---
### No immediate failures.

---
### Need to rollback for a different, unrelated issue.

---
### "Well at least it ran cleanly before the rollback!"

---
### "The load on the web monoliths is skyrocketing" - Doug

{{% /section %}}

---
# Predicate-ment
<pre>
<code>
List<Categorized> unfilteredItems;

Predicate<Categorized> requiresPermission =
  (entity) -> 
    logic.sensitiveFields(unfilteredItems).matches(entity);

return 
  unfilteredItems
    .filter(requiresPermission);


</code>
</pre>

---
{{% section %}}
# Final Contributors

---
### Can't get Facility null

{{% /section %}}


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

