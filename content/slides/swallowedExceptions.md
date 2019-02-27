# Exception, what Exception?

---
2 methods below my code queried the DB

---
FacilityLogic

    Facility get(String facilityId) {
      Facility facility = null;
      try {
        facility = DB.getFacility(facilityId);
      } catch (HorribleDbException ex) {
        log.error("Error getting facility");
      }
      return facility; // Everything is totally fine!
    }

---

FacilityConfigLogic

    String configByName(Facility facility, String configName) {
      String ret = null;
      try { // facility is null!
        ret = DB.getConfig(facility, configName); 
      } catch (NullPointException ex) {
        log.("Can't get value for facility", facility); 
      } 
      return ret; // Just keep swimming
    }
