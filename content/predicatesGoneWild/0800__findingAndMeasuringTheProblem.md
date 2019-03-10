+++
+++

# The Actual Problem

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
#### Without my code, logs reduced, but were still very active.
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
### How bad could it be in practice?
---
#### 10 Visits
### 2,000 Queries

---
#### 100 Visits
### 2,000,000 Queries

---
#### 5000 Visits
### 250,000,000,000 Queries

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

