---
# Predica(te)ment
    List<Entity> unfilteredItems;

    Predicate<Entity> requiresPermission =
      (entity) -> 
        logic.sensitiveFields(unfilteredItems).matches(entity);

    return 
      unfilteredItems
        .filter(requiresPermission);


---
# Predicate-ment
    Predicate<Entity> requiresPermission =
      (entity) -> 
        logic.sensitiveFields(unfilteredItems).matches(entity);

---
# Predicate-ment
    Predicate<Entity> requiresPermission =
        logic.sensitiveFields(unfilteredItems)::matches;


