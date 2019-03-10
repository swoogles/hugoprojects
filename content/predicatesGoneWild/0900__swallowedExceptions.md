+++
+++

# But
<ul>
{{% fragment %}}<li> Why was it so hard to find?</li>{{% /fragment %}}
{{% fragment %}}<li> How was it able to kill entire machines?</li>{{% /fragment %}}
{{% fragment %}}<li> Why didn't the existing code fail?</li>{{% /fragment %}}
</ul>



---
{{% section %}}
# Exception, what Exception?

---
2 methods below my code queried the DB

---
FacilityLogic

    Facility get(String facilityId) {
      try {
        return DB.getFacility(facilityId);
      } catch (HorribleDbException ex) {
        log.error("Error getting facility");
        return null; // Everything is totally fine!
      }
    }

---

FacilityConfigLogic

    String configByName(Facility facility, String configName) {
      try { // facility is null!
        return DB.getConfig(facility, configName); 
      } catch (NullPointException ex) {
        log.("Can't get value for facility", facility); 
        return null; // Just keep swimming
      } 
    }

{{% /section %}}
