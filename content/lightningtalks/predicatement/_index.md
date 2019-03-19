+++
title = "Predicates Gone Wild"
outputs = ["Reveal"]
+++
## How many Exceptions can a single request generate?

---
## Goal:
Add a new filter to hospital visits returned by an API.

---
# Code is deployed

---
# Everything is fine!

---
## Troubling Symptoms
<ul>
{{% fragment %}}<li> Database degradation </li>{{% /fragment %}}
{{% fragment %}}<li> Out of memory errors! </li>{{% /fragment %}}
{{% fragment %}}<li> The hard drives are full!? </li>{{% /fragment %}}
</ul>

---
# Tracking down the the Problem!

---
{{% section %}}
### After adding some logging...

---
### DB Activity on Census page
    1 visit :  2 DB queries

    2 visits:  4 DB queries

    3 visits:  27 DB queries

    4 visits: ~100 DB queries

    5 visits: ~250 DB queries

---
# 2n^3

{{% /section %}}

---
{{% section %}}
# Divergent Expectations

---
### Desired
![Example image](/images/EncounterRequests1.png)

---
### Deployed/Tolerated
![Example image](/images/EncounterRequests2.png)

---
### Death
![Example image](/images/EncounterRequests3.png)

{{% /section %}}

---
{{% section %}}
### How bad could it be in practice?
---
#### 10 Visits
### 2,000 Queries

---
#### 100 Visits
### 2,000,000 Queries

---
### 10,000 Visits
### 2,000,000,000,000 Queries

<hr>
<h3>
{{% fragment %}} 1 microsecond per stacktrace {{% /fragment %}}
</h3>
<h3>
{{% fragment %}} ~23 days{{% /fragment %}}
</h3>

{{% /section %}}

---
{{% section %}}
### Predicate&lt;T&gt;
"Take a T and return true or false"

---

### What I said
Determine sensitive Fields for each item. Use once.

<br>

### What I said
    Predicate<HospitalVisit> requiresPermission =
      (visit) ->
        logic.sensitiveFields(unfilteredItems).matches(visit);

---
### What I meant
Determine sensitive Fields once. Use for all entities.

<br>

### What I meant
    Predicate<HospitalVisit> requiresPermission =
        logic.sensitiveFields(unfilteredItems)::matches;





{{% /section %}}

---
{{% section %}}
# Is your code acting like a pig?

---
![Example image5](/images/Piggy.png)

---
#### Be more like a Koala
![Example image5](/images/KoalaEatingEucalyptus_small.jpg)

---
#### Die when you eat the wrong thing!

{{% /section %}}


