# Don't hide it under a bush-null

    Facility get(String facilityId) {
      Facility facility = null;
      try {
        facility = DB.getFacility(facilityId);
      } catch (HorribleDbException ex) {
      }
      return facility; // Everything is totally fine!
    }

---
### FacilityConfigLogic

    Optional<Boolean> get(String facilityId, String configName) {
      return 
        Optional.ofNullable(configByName(facilityId, configName))
          .map(Boolean::parse);
    }

    String configByName(Facility facility, String configName) {
          String ret = null;
          try {
            facility = DB.getFacility(facility); 
          } catch (Exception ex) {
            log.("Can't get value for facility", facility); 
          } return ret; // Just keep swimming
    }
