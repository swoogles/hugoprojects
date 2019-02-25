# Predica(te)ment
    List<Entity> allowed(List<Entity> unfilteredItems) {

      Predicate<Entity> requiresPermission =
        (entity) -> 
          logic.sensitiveFields(unfilteredItems).matches(entity);

      return 
        unfilteredItems
          .filter(requiresPermission);
    }


---
### What I said
    Predicate<Entity> requiresPermission =
      (entity) -> 
        logic.sensitiveFields(unfilteredItems).matches(entity);

### What I meant
    Predicate<Entity> requiresPermission =
        logic.sensitiveFields(unfilteredItems)::matches;


