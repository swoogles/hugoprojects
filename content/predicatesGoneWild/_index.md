+++
title = "My presentation"
outputs = ["Reveal"]
+++

## How many Exceptions can a single request generate?  

---
## Goal:
Filter facility visits returned by Census endpoint.

---
# Attempt #1

---
## Troubling Symptoms
<ul>
{{% fragment %}}<li> Database degradation </li>{{% /fragment %}}
{{% fragment %}}<li> Out of memory errors! </li>{{% /fragment %}}
{{% fragment %}}<li> The hard drives are full!? </li>{{% /fragment %}}
</ul>


---
{{% section %}}
# Analyzing

---
### "I see something about Consent" - Austin

---
### Errors in Facility Config

---
### Can't get Facility null

---
### This connection is closed.

---
### Transaction cannot proceed: STATUS ROLLBACK

---
### Abort of action invoked while multiple threads active within it.

{{% /section %}}


---
{{% section %}}
### The CDI Rabbit Hole

---
### Hypothesis
#### Something is being improperly shared between CDI and Static classes!

---
### We wanted it to work like this:

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
### But instead it worked like this:


---
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
### Because the dependencies looked like this:

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
### But forget all that!
{{% /section %}}



---
# Attempt #2

---
{{% section %}}
# Deploying
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### No immediate failures!

---
### Need to rollback for a different, unrelated issue

---
### "Well at least it ran cleanly before the rollback!"
{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}{{% fragment %}}.{{% /fragment %}}

---
### "The load on the web monoliths is skyrocketing" - Doug

{{% /section %}}
