+++
title = "My presentation"
outputs = ["Reveal"]
+++

# How many Exceptions can a single request generate?  

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
### Errors in FacilityConfig

---
### Can't get Facility null

{{% /section %}}


---
{{% section %}}
### The Rabbit Hole

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

---
# The Actual Problem

---
{{% section %}}
### After adding some logging...

---
### DB Activities when loading Visit page
    1 visit :   2 DB queries
    2 visits:  4 DB queries
    3 visits:  27 DB queries
    4 visits: ~100 DB queries
    5 visits: ~250 DB queries
{{% /section %}}

---
{{% section %}}
# Divergent Expectations

---
![Example image](/images/EncounterRequests1.png)

---
![Example image](/images/EncounterRequests2.png)

---
![Example image](/images/EncounterRequests3.png)

{{% /section %}}

---
{{% section %}}
### How bad could it be?
---
#### 10 Visits
### 2,000 Database Calls

---
#### 100 Visits
### 2,000,000 Database Calls

---
#### DB queries for 5000 Visits
### 250,000,000,000
<br>

<h4>
{{% fragment %}} Total number of humans ever{{% /fragment %}}
</h4>
<h3>
{{% fragment %}} 107,000,000,000{{% /fragment %}}
</h3>


---
### 10,000 Visits
### 2,000,000,000,000 Database Calls

<h3>
{{% fragment %}} 1 microsecond per stacktrace {{% /fragment %}}
</h3>
<h3>
{{% fragment %}} ~23 days{{% /fragment %}}
</h3>

{{% /section %}}

---
{{% section %}}
# n^2


---
{{%readfile file="/content/slides/predicateCodeSamples.md" markdown="true"%}}
{{% /section %}}


---
{{% section %}}
# n^3

---
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
{{% section %}}
# Is your code acting like a pig?

---
![Example image5](/images/Piggy.png)

---
# Be more like a Koala!

---
![Example image5](/images/KoalaEatingEucalyptus_small.jpg)
#### Die when you eat the wrong thing!

{{% /section %}}

