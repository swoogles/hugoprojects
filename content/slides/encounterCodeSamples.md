## Calling code
    List<Entity> endpoint(List<Entity> bulkItems) {

      Predicate<Entity> otherPredicate = ...

      Predicate<Entity> primaryRequirement =
        (entity) -> 
          logic.allowed(bulkItems).contains(entity);

      return 
        bulkItems
          .filter(primaryRequirement.or(otherPredicate);
    }

---
## Calling code
                 endpoint(List<Entity> bulkItems) {




        (entity) -> 
          logic.allowed(bulkItems)




    }
