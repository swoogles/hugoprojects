# Exception, what Exception?

---
FacilityLogic

    Facility get(String facilityId) {
      Facility facility = null;
      try {
        facility = DB.getFacility(facilityId);
      } catch (HorribleDbException ex) {
      }
      return facility; // Everything is totally fine!
    }

---

FacilityConfigLogic

    String configByName(Facility facility, String configName) {
      String ret = null;
      try {
        ret = DB.getConfig(facility, configName); 
      } catch (Exception ex) {
        log.("Can't get value for facility", facility); 
      } 
      return ret; // Just keep swimming
    }
