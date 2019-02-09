+++
title = "My presentation"
outputs = ["Reveal"]
+++

# Hello world!

This is my first slide.

---
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
- {{% fragment %}}{{% /fragment %}}
- {{% fragment %}}{{% /fragment %}}
- {{% fragment %}}{{% /fragment %}}

---
#

---
#

---
- {{% fragment %}}One{{% /fragment %}}
- {{% fragment %}}Two{{% /fragment %}}
- {{% fragment %}}Three{{% /fragment %}}
- {{% fragment %}}Four{{% /fragment %}}

---
{{<mermaid>}}
  graph TD
  Start --> Stop
{{</mermaid>}}

{{ readFile "mermaid/mermaid.js" }}

{ {- renderFile "mermaid/mermaid.js" -} }

<style>{ {- renderFile "/static/css/site.css" -} }</style>

<link href="{{"mermaid/mermaid.css" | relURL}}{{ if not .Site.Params.disableAssetsBusting }}?{{ now.Unix }}{{ end }}" type="text/css" rel="stylesheet" />


<script>{ {- renderFile "mermaid/mermaid.js" -} }</script>
<script src="{{"mermaid/mermaid.js" | relURL}}{{ if not .Site.Params.disableAssetsBusting }}?{{ now.Unix }}{{ end }}"></script>
