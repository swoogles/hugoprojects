+++
title = "My presentation"
outputs = ["Reveal"]
+++

# Static Shock
or, how a single request can generate 800,000 exceptions.

---
## Troubling Symptoms
<ul>
{{% fragment %}}<li> Database Degradation </li>{{% /fragment %}}
{{% fragment %}}<li> Out of memory errors! </li>{{% /fragment %}}
{{% fragment %}}<li> The hard drives are full!? </li>{{% /fragment %}}
</ul>


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
# Attempt #2

---
{{% section %}}
# Deploying
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### No immediate failures.
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### Need to rollback for a different, unrelated issue.

---
### "Well at least it ran cleanly before the rollback!"
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### "The load on the web monoliths is skyrocketing" - Doug

{{% /section %}}

---
# The Actual Problem

---
### Number of logs
<pre>
<code>

for numberOfItems in range(1, 10, 1):
  print(str(numberOfItems) + ": "
    + str(2 * (numberOfItems * numberOfItems * numberOfItems)))

</code>
</pre>
<pre>
<code>

1: 2
2: 16
3: 54
4: 128
5: 250

</code>
</pre>

---
### The numbers take flight.
<p>
{{% fragment %}}10:   2,000{{% /fragment %}}
</p>
<p>
{{% fragment %}}100:  2,000,000{{% /fragment %}}
</p>
<p>
{{% fragment %}}1000: 2,000,000,000{{% /fragment %}}
</p>
<p>
{{% fragment %}}5000: 250,000,000,000{{% /fragment %}}
</p>

<p>
{{% fragment %}}9999: 1,999,400,059,998{{% /fragment %}}
</p>

---
{{% section %}}
### The numbers take flight.
---
### 10 Visits
### 2,000 Database Calls

---
### 100 Visits
### 2,000,000 Database Calls

---
### 1000 Visits
### 2,000,000,000 Database Calls

---
### 5000 Visits
### 250,000,000,000 Database Calls

<h3>
{{% fragment %}} Number of humans that have ever lived{{% /fragment %}}
</h3>
<h3>
{{% fragment %}} 107,000,000,000{{% /fragment %}}
</h3>


---
### 9999 Visits
### 1,999,400,059,998 Database Calls

<h3>
{{% fragment %}} Calls take 1/10th of a millisecond {{% /fragment %}}
</h3>
<h3>
{{% fragment %}} 100199940005.9998 .1*milliseconds {{% /fragment %}}
</h3>
1999400059998 * (.0001)
199940005.9998


---
{{% /section %}}

<p>
{{% fragment %}}9999: 1,999,400,059,998{{% /fragment %}}
</p>


---
# Predicate-ment
<pre>
<code>
List&ltEntity&gt unfilteredItems;

Predicate&ltEntity&gt requiresPermission =
  (entity) -> 
    logic.sensitiveFields(unfilteredItems).matches(entity);

return 
  unfilteredItems
    .filter(requiresPermission);

</code>
</pre>

---
# Predicate-ment
<pre>
<code>
Predicate&ltEntity&gt requiresPermission =
  (entity) -> 
    logic.sensitiveFields(unfilteredItems).matches(entity);

</code>
</pre>

---
# Predicate-ment
<pre>
<code>
Predicate&ltEntity&gt requiresPermission =
    logic.sensitiveFields(unfilteredItems)::matches;

</code>
</pre>



---
{{% section %}}
# Final Contributors

---
### Can't get Facility null

{{% /section %}}

---
Junk drawer from here on out.


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

