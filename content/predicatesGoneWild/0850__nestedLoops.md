+++
+++

{{% section %}}
# n^2


---
### Predicate&lt;T&gt;
Function that takes a single T and returns a boolean

---
## The Predica(te)ment
    List<Entity> unfilteredItems;

    Predicate<Entity> requiresPermission =
      (entity) -> 
        logic.sensitiveFields(unfilteredItems).matches(entity);

    List<Entity> privateData  =
        unfilteredItems
          .filter(requiresPermission);

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

{{% /section %}}


---
{{% section %}}
# n^3

---
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
{{% /section %}}


