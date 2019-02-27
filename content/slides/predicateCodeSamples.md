### Predicate
Function that takes a single argument and returns a boolean

---
# The Predica(te)ment
    List<Entity> allowed(List<Entity> unfilteredItems) {

      Predicate<Entity> requiresPermission =
        (entity) -> 
          logic.sensitiveFields(unfilteredItems).matches(entity);

      return 
        unfilteredItems
          .filter(requiresPermission);
    }

---
# The Predica(te)ment
    List<Entity> allowed(List<Entity> unfilteredItems) {


        (entity) -> 
          logic.sensitiveFields(unfilteredItems)




    }

---
### What I said
Call sensitiveFields for each item. Use it once.

<br>


### What I meant
Call sensitiveFields once. Use it for all entities.


---
### What I said
    Predicate<Entity> requiresPermission =
      (entity) -> 
        logic.sensitiveFields(unfilteredItems).matches(entity);



### What I meant
    Predicate<Entity> requiresPermission =
        logic.sensitiveFields(unfilteredItems)::matches;

---
### Could also be:
    Validator validator = 
      logic.sensitiveFields(unfilteredItems)

    Predicate<Entity> requiresPermission =
      (entity) -> 
        validator.matches(entity);

