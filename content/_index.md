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
{{% section %}}
### Bad Categorizer Behavior
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
### Desired Categorizer Behavior
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
{{% /section %}}


---
# Attempt #2

---
{{% section %}}
# Deploying
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### No immediate failures
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### Need to rollback for a different, unrelated issue

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
    1: 2
    2: 16
    3: 54
    4: 128
    5: 250

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
### DB queries for 5000 Visits
### 250,000,000,000
<br>

<h3>
{{% fragment %}} Total number of humans ever{{% /fragment %}}
</h3>
<h3>
{{% fragment %}} 107,000,000,000{{% /fragment %}}
</h3>


---
### 9999 Visits
### 1,999,400,059,998 Database Calls

<h3>
{{% fragment %}} 1 microsecond per stacktrace {{% /fragment %}}
</h3>
<h3>
{{% fragment %}} ~23 days{{% /fragment %}}
</h3>

{{% /section %}}

---
{{% section %}}
{{%readfile file="/content/slides/predicateCodeSamples.md" markdown="true"%}}
{{% /section %}}



---
{{% section %}}
{{%readfile file="/content/slides/encounterCodeSamples.md" markdown="true"%}}
{{% /section %}}


---
# But
<ul>
{{% fragment %}}<li> Why was it so hard to find?</li>{{% /fragment %}}
{{% fragment %}}<li> How was it able to kill entire machines?</li>{{% /fragment %}}
{{% fragment %}}<li> Why didn't the existing code fail?</li>{{% /fragment %}}
</ul>

---
{{% section %}}
{{%readfile file="/content/slides/swallowedExceptions.md" markdown="true"%}}
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

