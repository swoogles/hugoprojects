Functional Interfaces

---
- Supplier
- Predicate
- Function
- Consumer

---
- Predicate<T>
- Function<T,R> 

---
{{% section %}}
### Predicate&lt;T&gt;
"Take a T and return true or false"

---
#### Is a Person an adult?
    Predicate<Person> isAnAdult = (person) -> person.age() >= 18;

---
#### Why is that any better than this?
    boolean isAnAdult(Person person) {
        return person.age() > 18);
    }
{{% /section %}}


---
{{% section %}}
### Function&lt;T,R&gt;
"Take a T and return an R"

---
### Function&lt;UUID,User&gt;
"Take a UUID and return a User"

{{% /section %}}

