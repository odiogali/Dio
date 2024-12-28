2024/09/03 22:59
Status: #idea
Tags: #datastructure #academic
# What is a Set?
A set is a collection of distinct items. A `HashSet` is an unsorted, unordered set. Java's `HashSet` does not have all the 'set theory' type operators like disjoint, power set, etc. `HashSet<Key> set = new HashSet<Key>();` - `<Key>` is the [[Generics|generic]] type of the items used as keys. The main functions that belong to the `HashSet` class are - `add()` which cannot add duplicates, `clear()` which empties the set, `contains()`, `remove()`, and `size()`.
____
# About Sets (Theory)
- A set is an unordered collection of distinct objects
	- Order of the objects in set doesn't matter
	- Generally there are no duplicates in sets but in CPSC 351, it doesn't matter if elements are listed more than once:
		- {1,6,10,3} is the same as {6,6,10,3,1,3,1}
- To specify a set, you can list them explicitly:
	- 𝑆<sub>1</sub> = {3, 5, 18, "𝑎𝑝𝑝𝑙𝑒"}
	- 𝑆<sub>2</sub> = {1, 2, 3, 4, 5, … }
- Or define them according to a rule, called a comprehension:
	- 𝑆<sub>3</sub> = {𝑓: 𝑓 is a frog named Fred}
	- Colon (:) above means "such that"
- Sets can be infinite or finite
	- The empty set is denoted ∅ = {}.
- We write 𝑥 ∈ 𝑆 to indicate the object 𝑥 is an element (member) of the set 𝑆
	- 2 ∈ {1,2,4,8}
- Subsets:
	- 𝐴 ⊆ 𝐵 if and only if every element of 𝐴 is also an element of 𝐵.
	- We say 𝐴 is a subset of 𝐵, or 𝐵 is a superset of 𝐴.
	- Two sets are equal if both 𝐴 ⊆ 𝐵 and 𝐵 ⊆ 𝐴.
- Important Sets:
	- The natural numbers are ℕ = {0, 1, 2, 3, 4, … }
	- The integers are ℤ = {… − 2, −1, 0, 1, 2, … }
	- The real numbers are ℝ = {0, 𝜋, 1, −15.3 … } (any number that exists)
- [[Quantifiers]]:
	- The universal quantifier is denoted ∀, which means “for all”.
	- The existential quantifier is denoted ∃, which means “there exists”.
		- 𝑆<sub>1</sub> = {𝑥 ∈ ℕ: ∃𝑦 ∈ ℕ, 𝑥 = 𝑦<sup>2</sup>}
		- 𝑆<sub>2</sub> = {𝑥 ∈ ℕ: ∀𝑦 ∈ {2,3,5}, 𝑥 is divisible by 𝑦}
# Set Operations
- The union of two sets A and B, denoted 𝐴 ∪ 𝐵, is the set of all elements contained in one or the other or both.
	- If A = {1,2,3} and B = {2,4,6}, 𝐴 ∪ 𝐵 is {1,2,3,4,6}
- The intersection of two sets 𝐴 and 𝐵, denoted 𝐴 ∩ 𝐵, is the set of all elements contained in both.
	- Two sets are disjoint if their intersection is the empty set
- The difference of two sets 𝐴 and 𝐵, denoted 𝐴 − 𝐵, is the set of all elements contained in 𝐴 but not 𝐵
- The symmetric difference of two sets 𝐴 and 𝐵, denoted 𝐴 ⊕ 𝐵, is the set of all elements contained in one or the other, but not both.
- Venn diagrams give an easy way to reason about set memberships
	- 𝐴 ∪ 𝐵 will be everything in the Venn diagram
	- 𝐴 ∩ 𝐵 will be the middle of the Venn diagram
	- A - B will be just A minus the middle
	- 𝐴 ⊕ 𝐵 will be A and B but not the middle
- One of the reasons we care about sets is that they capture [[Formal Logic|logic]] operations. For example:
	- 𝐴 = {𝑥 ∶ condition 𝐶<sub>𝐴</sub>(𝑥) is true}
	- 𝐵 = {𝑥 ∶ condition 𝐶<sub>𝐵</sub>(𝑥) is true}
	- Then, 𝐴 ∪ 𝐵 = {𝑥 ∶ 𝐶<sub>𝐴</sub> (𝑥) ∨ 𝐶<sub>𝐵</sub> (𝑥) are true}
![Cartesian Products](Sequences.md#Cartesian%20Products)
# Functions and Relations
- If A and B are sets, then a function from A to B is an assignment of exactly one object b ∈ B to every object a ∈ A
- Functions from A to B are sometimes called total functions from A to B
	- A partial function from A to B is an assignment of at most one value b ∈ B to every value a ∈ A
	- Therefore, it is possible some elements of A are left without an assigned value
- The notation for functions is thus the following (it represents the fact that $f$ is a partial function from A to B): ![](Pasted%20image%2020240903231104.png)
	- For any such function $f$, if a ∈ A and there exists an object b ∈ B such that $f$ associates b with a, we say that f is defined at $a$ and we write the value associated with $a$, by $f$ as $f(a)$
	- If there is no value b ∈ B associated with a by f, we say f is undefined at a
- If f: A -> B, then A is called the domain of the function f, and B is called the range of f
- Suppose that f: A -> B
	- f is called *injective* or 1-1 if it satisfies the following property: for every b ∈ B, there exists at most one object a ∈ A such that f(a) = b so that for all objects in A, if f(a<sub>1</sub>) = f(a<sub>2</sub>) then a<sub>1</sub> = a<sub>2</sub>
		- B's are not necessarily all mapped to an a, but if they are, it's just one
	- f is called *surjective* or onto, if f satisfies the following property: for every object b ∈ B, there exists at least one object a ∈ A such that f (a) = b
		- All b's are mapped to at least one a
	- f is called *bijective* if f is both injective and surjective: so that, for every object b ∈ B, there exists exactly one object a ∈ A such that f (a) = b. A function f : A → B, that is bijective, is called a bijection from A to B.
		- All b's are mapped to exactly one a
- When f : A<sub>1</sub> × A<sub>2</sub> × · · · × A<sub>k</sub> → B for some positive integer k and for sets A<sub>1</sub>, A<sub>2</sub>, . . . , A<sub>k</sub> and B, we call f a k-ary function. This is called a “unary” function when k = 1 and it is called a “binary” function when k = 2.
- A predicate or property is a function whose range is {true, false}
	- A property whose domain is the set of k-tuples A × A × ... × A for some positive integer k and set A is called a relation 
- Infix notation is used when defining and using various familiar binary functions and relations
	- For the example, while the “integer addition function” would be represented as a total function + : Z × Z → Z, based on what is given above, we generate wrote “a + b” instead of “+(a, b)”,  as what you get when applying this function to the pair of integers a and b.
# Questions