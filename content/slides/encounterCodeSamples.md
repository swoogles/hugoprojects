# Calling code
    List<Entity> endpoint(List<Entity> bulkItems) {

      Predicate<Entity> otherPredicate = ...

      Predicate<Entity> allowed =
        (entity) -> 
          logic.allowed(bulkItems).contains(entity);

      return 
        bulkItems
          .filter(allowed.or(otherPredicate);
    }
