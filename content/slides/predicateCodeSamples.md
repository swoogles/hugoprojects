### Predicate&lt;T&gt;
Function that takes a single T and returns a boolean

---
## The Predica(te)ment
    List<Entity> privateData(List<Entity> unfilteredItems) {

      Predicate<Entity> requiresPermission =
        (entity) -> 
          logic.sensitiveFields(unfilteredItems).matches(entity);

      return 
        unfilteredItems
          .filter(requiresPermission);
    }

---
## The Predica(te)ment
                 privateData(List<Entity> unfilteredItems) {


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

