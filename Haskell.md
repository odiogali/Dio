2024/10/31 20:36
Status: #idea
Tags: #academic #language #concept #functional
# What is Haskell?
Haskell is a general-purpose, statically-typed, purely functional programming language with type inference and lazy evaluation. Adasfaksfdja sldkajskdajslkd jask jdaskl jask jdaks jdaks jlda sdjas kdj askdjsakjda ksjd sakj akjdaks jaksj daskj dakjs kdjask djaskj dkjsakjdskajdkasjd.
___
# About
- First implemented in 1990
- GHC stands for Glasgow Haskell Compiler
- First fully functional lazy language
# Features
- Functional syntax, pattern matching
- Evaluation by hand, lazy evaluation 
- Lists and list comprehensions
- Datatypes and higher-ordering programming
- Structural induction
- Classes and monads
- Relation to lambda calculus
# Control Flow
## If-then-else:
- For conditional logic, Haskell uses inline `if-then-else` expressions
- Syntax:
	- `if condition then expression1 else expression2`
- Example:
```Haskell
isPositive :: Integer -> String
isPositive a = if a > 0 then "Positive" else "Non-positive"

-- Test
isPositive 5 
isPositive 0
isPositive -1
```
## Guards:
- Provide an elegant way to handle multiple conditions 
- Written with a series of conditions prefixed by `|`
- Syntax:
```Haskell
functionName args
	| condition1 = expression1
	| condition2 = expression2
	| condition3 = expression3
	| otherwise = defaultExpression
```
- Example:
```Haskell
-- Determine grade based on score
grade :: Int -> String
grade score 
	| score >= 90 = "A"
	| score >= 80 = "B"
	| score >= 70 = "C"
	| score >= 60 = "D"
	| otherwise = "F"

-- Test
grade 85
grade 54
```
## Pattern Matching:
- Pattern matching simplifies function definitions by directly deconstructing values
- Syntax:
```Haskell
functionName pattern1 = expression1
functionName pattern2 = expression2
```
- Example:
```Haskell
-- Describe a tuple
describeTuple :: (Int, Int) -> String
describeTuple (0, 0) = "Both are zero"
describeTuple (x, 0) = "First is " ++ show x ++ ", second is zero"
describeTuple (0, y) = "First is zero, second is" ++ show y 
describeTuple (x, y) = "First is " ++ show x ++ ", second is" ++ show y 

-- Test
describeTuple (0, 0)
describeTuple (3, 0)
describeTuple (0, 5)
describeTuple (3, 4)
```
## Let and Where Expressions
- Both `let` and `where` are used to define local bindings
- `let` expressions are inline and can be used anywhere
	- Syntax: `let bindings in expression`
- `let` example:
```Haskell
-- Calculate area using let
circleArea :: Double -> Double
circleArea r = let piVal = 3.14159 in piVal * r * r

-- Test
circleArea 3
```
- `where` is used to define bindings at the end of a function definition
	- Syntax is as below:
```Haskell
functionName args = expression
	where bindings
```
- `where` example:
```Haskell
-- Calculate area using where
circleArea :: Double -> Double
circleArea r = piVal * r * r
	where piVal = 3.14159

-- Test
circleArea 3
```
## Case Expressions
- Allows us to perform certain actions based on the inputs without having to define an entirely different function 
- It's pattern matching without multiple function definitions
- The following two expressions are semantically equivalent:
```Haskell
-- Function defintions
f p11 .. p1k = e1
...
f pn1 .. pnk = en

-- Case expressions
f x1 x2 ... xk = case (x1, x2, ..., xk) of 
	(p11 .. p1k) -> e1
	...
	(pn1 .. pnk) -> en
```
- The following is an implementation of `take` that uses case expressions:
```Haskell
mytake m ys = case (m, ys) of
	(0, _) -> []
	(_, []) -> []
	(n, x : xs) -> x : take (n - 1) xs 
```
# Lists
- Haskell has a built-in syntax for lists
	- By default, lists in Haskell are akin to linked lists in other languages
- Each of the following lists are equal:
	- `[1,2,3]`
	- `1:[2, 3]`
		- Colon is the list constructor
	- `1:2:[3]`
		- `[3]` is called a singleton list
	- `1:2:3:[]`
		- `[]` is called an empty list or the list constructor
- In prelude, lists are declared
	- As below with 'a' being the type
```Haskell
data[a] = []
		| (:) a [a]
```
```Haskell
zip :: [a] -> [b] -> [(a, b)]
zip (a : as) (b : bs) = (a, b) : (zip as bs)
zip _ _ = []
```
- Types of lists:
	- `[]` is the empty list
	- `x:xs` is the list with first element x and "tail" of the list xs
	- `[a]` is (besides the one element list), the type of list of a's
## List Comprehensions:
- List builder notation
	- Write description of a list using the elements of another list
		- `[ x | x <- list]`
		- `[ x * 2 | x <- list]`
- Basically, using the description of a list 
- Haskell translates list comprehensions into "core" Haskell
- Examples:
	- List of pairs of elements from the list "as" and list "bs"
	```Haskell
	pairs :: [a] ->  [b] -> [(a,b)]
	pairs as bs = [(a, b)| a <- as, b <- bs]
	```
	- Element of list satisfying predicate "p":
	```Haskell
	filter :: (a -> Bool) -> [a] -> [a]
	filter p as = [a | a <- as, p a]
	```
	- Flatten (or `concat` in the Prelude):
	```Haskell
	flatten :: [[a]] -> [a]
	flatten ass = [a | as <- ass, a <- as]
	```
	- Apply function to all elements in list
	```Haskell
	mapList :: (a-> b) -> [a] -> [b]
	mapList f as = [f a | a <- as]
	```
	- Quick Sort
	```Haskell
	qsort :: Ord a => [a] -> [a]
	qsort [] = []
	qsort (a:as) = qsort [y | y <- as, y < a] 
					++ [a] ++ qsort [y | y <- as, y >= a]
	```
	- Pythagorean triples: $x,y,z \in \mathbb{N}$ such that $x^2+y^{2}=z^2$
		- Give all Pythagorean triples where $x,y,z \leq n$
	```Haskell
	-- Very inefficient
	pythag :: Integer -> [(Integer, Integer, Integer)]
	pythag n = [(x, y, z) | x <- [1..n]
						    y <- [1..n]
						    z <- [1..n]
						    pyth x y z] where
		pyth x y z = x * x + y * y == z * z
	```
	- Sieve of Eratosthenes (old guy who was a polymath)
		- Program to work out the sequence of prime numbers
	```Haskell
	primes = pfilter [2..] where
		pfilter [] = []
		pfilter (p:xs) = p : pfilter [x | x <- xs, not (x `mod` p == 0)]
	```
## List Comprehensions:
- Haskell evaluates list comprehensions according to three rules:
	1. `{ [e | x <- xs, r] }` = `concat (map (\x -> { [e | r] }) xs)` where `r` is a list of conditions
	2. `{ [e | p x : r] }` = `if (p x) then { [e | r] } else []`
	3. `{ [e | ] }` = `[e]`
- The above is called a translation of list comprehensions
	- How Haskell itself interprets list comprehensions
- An example of a translation:
```Psuedocode
{ [m * m | m <- [1 .. 10], m * m < 50] }
 == concat (map (\x -> { [m * m | m * m < 50] } ) [1..10]) -- using rule 1
 == concat (map (\x -> if (m * m < 50) then {[m*m | ]} else []) [1..10]) -- using rule 2
 == concat (map (\x -> if (m * m < 50) then [m*m] else []) [1..10]) -- using rule 3
```
- It leads to simpler reasoning about program equality
	- Done through structural induction
## Folds:
- Fold is a family of higher order functions that process a data structure in some order and build a return value
	- This is compared to the family of *unfold* functions which take a starting value and apply it to a function to generate a data structure
	- Typically, fold deals with a combining function and a data structure
		- Fold then proceeds to combine elements of the data structure using the function in some systematic way
- On lists, you can either recursively combine the first element with the results of combining the rest (`foldr`) or by recursively combining the results of combining all but the last element with the last one (`foldl`)
	- The following code shows how folds are defined in Haskell
```Haskell
-- if the list is empty, the result is the initial value z; else
-- apply f to the first element and the result of folding the rest
foldr f z []     = z 
foldr f z (x:xs) = f x (foldr f z xs) 

-- if the list is empty, the result is the initial value; else
-- we recurse immediately, making the new initial value the result
-- of combining the old initial value with the first element.
foldl f z []     = z                  
foldl f z (x:xs) = foldl f (f z x) xs
```
- `foldr` versus `foldl`:![[Pasted image 20250130193734.png|600]]![[Pasted image 20250130193744.png|600]]
- How to fold on lists:
	- In the prelude is called "foldr"
	- Type definition: `foldr :: c -> (a -> c -> c) -> [a] -> c`
	- What it does:![[Pasted image 20250129113704.png]]
	- Example: ![[Pasted image 20250129113640.png]]
```Haskell
foldr c f [] = c
foldr c f (a : as) = f a (foldr c f as)
```
- Foldr captures a common form of recursion - it has good properties
	- `fold c f as` terminates whenever `c` and `f` terminate (assuming that `as` is a finite list)
- Suppose you want the average of a list of floats:
```Haskell
average :: [Float] -> Float
average xs = mydiv (foldr (\n (sum, l) -> (sum + n, l + 1)) (0, 0) xs)
	where
		mydiv (s, n) = s/n
```
- Bubble sort
```Haskell
bsort :: Ord a => [a] -> [a]
bsort xs = foldr push [] xs where
	push a [] = [a]
	push a (b: bs) 
		| a <= b = a : b : bs
		| otherwise = b : (push a bs)
```
- Splitting a list into two lists, the list of odd indexed elements and even indexed elements
```Haskell
split :: [a] -> ([a], [a])
split = foldr (\a (odd, even) -> (a : even, odd)) ([], []) 
```
- Given a list, return the tails of the list
```Haskell
tails :: [a] -> [[a]]
tails xs = (\(current, tails) -> current : tails) (foldr (\a (current, tails) -> (a : current, current : tails)) ([], []) xs )
```
- Reversing lists (naive way):
```Haskell
nrev :: [a] -> [a]
nrev [] = []
nrev (a:as) = (nrev as) ++ [a] -- this recursion is in fold form

-- Can be rewritten
nrev xs = foldr (\a as -> as ++ [a]) [] xs
```
- Optimized reversing of lists
```Haskell
srev :: [a] -> [a]
srev xs = shunt xs [] where
	shunt [] ys = ys
	shunt (x:xs) ys = shunt xs (x:ys)
```
- Higher order reverse:
```Haskell
hrev :: [a] -> [a]
hrev as = foldr (\x f -> f . ((:) x)) id as []
-- fold builds a function
```
# Datatypes
- Lists are defined as such:
```Haskell
data [a] = (:) a [a] -- Both [a]'s are recursive (infinitely many elements)
		  | []
```
- Booleans can either be true or false:
```Haskell
data Bool = True | False
```
- The success or fail datatype which is also called `Maybe` in the Prelude:
```Haskell
data SF a = SS a | FF 
-- SF is referred to as type/data constructor, FF or SS are term constructors
```
- We can also write a datatype for the set of natural numbers:
```Haskell
data Nat = Zero
		 | Succ Nat -- Unary representation - n is represented by applying Succ 
					-- constructor n times to Zero
	(deriving Eq, Show)

eq :: Nat -> Nat -> Bool -- Programming the equality test
eq Zero Zero = True
eq (Succ n) (Succ m) = eq n m
eq _ _ = False

pred :: Nat -> Nat
pred (Succ n) = n
pred Zero = Zero

add :: Nat -> Nat -> Nat
add Zero m = m
add (Succ n) m = Succ (add n m)

foldNat :: c -> (c -> c) -> Nat -> c
foldNat c f Zero = c
foldNat c f (Succ n) = f (foldNat c f n)

-- Add as a fold
add n m = foldNat m Succ n

-- Mutliply as a fold
mult n m = foldNat 0 (add n) m

monas n m = foldNat n pred m -- 3 - 6 = 0
```
- A basic tree where the internal nodes don't have values looks like the following:
```Haskell
data Tree a = Leaf a 
			| Node (Tree a) (Tree a) -- folds replace constructors (Node/Leaf)
```
- For instance, a Tree datatype would look like this:![[NodeTrees.excalidraw|700]]
- To sum the elements of the tree: `sum t = foldTree id (+) t`
```Haskell
foldTree :: (a -> c) -> (c -> c -> c) -> (Tree a) -> c
foldTree leaf node (Leaf a) = Leaf a
foldTree leaf node (Node t1 t2) = Node (foldTree leaf node t1) 
									   (foldTree leaf node t2)
```
- To collect the leaves of the tree: `collect t = foldtree (\a -> [a]) (++) t`
- To get the height of a tree
```Haskell
hgt (Leaf _) = 1
hgt (Node t1 t2) = 1 + max (hgt t1) (hgt t2)

-- Written as a fold
hgt = foldTree (\_ -> 1) (\n1 n2 -> 1 + max n1 n2)
```
- Search trees:
```Haskell
data STree a = SNode (STree a) a (STree a) | Tip
	deriving (Show, Eq)

foldSTree :: (c -> a -> c -> c) -> c -> (STree a) -> c
foldSTree g co Tip = co
foldSTree g co (SNode t1 a t2) = g (foldSTree g co t1) a (foldSTree g co t2)

mapSTree :: (a -> b) -> (STree a) -> (STree b)
mapSTree f Tip = Tip
mapSTree f (SNode t1 a t2) = SNode (mapSTree f t1) (f a) (mapSTree f t2)

hgtSTree t = foldSTree (\h1 a h2 -> 1 + max h1 h2) 0 t

-- Working out all the paths of a search tree
pathSTree :: STree a -> [[a]]
pathSTree = foldSTree (\p1 a p2 -> [a] : map ((:)a) (p1 ++ p2)) []
```
- Rose trees:
```Haskell
-- Tree with arbitrary branching and information at the nodes
data Rose a = Rs a [Rose a]
	deriving (Show, Eq)

mapRose :: (a -> b) -> (Rose a) -> (Rose b)
mapRose f (rs a ts) = rs (f a) (map (mapRose f) ts)

foldRose (a -> [c] -> c) -> (Rose a) -> c
foldRose g (rs a ts) = g a (map (foldRose g) ts)

mapRose = foldRose (\a ts -> rs (f a) ts)

sumRose :: Rose Int -> Int
sumRose t = foldRose (\n ns -> n + (foldr (+) 0 ns)) t

hgtRose :: Rose a -> Int
hgtRose = foldRose (\a hs -> 1 + (foldr max 0 hs))

pathRose = foldRose (\a ps -> [a] (map ((:) a) (concat ps)))
```
# Hand Evaluation
- The reason we hand evaluate is to understand what happens when you run a program particularly when debugging![[Pasted image 20250120121621.png]]![[Pasted image 20250120121636.png]]
- In the Prelude, the boolean datatype is defined like this where:
	- The data declaration bit is `data Bool`
	- The values `False` and `True` are the constructors
	- `Show` and `Eq` are typeclasses:
		- `Show` allows you to see the results of the computation in the interpreter
		- `Eq` refers to equality and allows the datatype to be compared
```Haskell
data Bool = False | True
	(deriving Show, Eq) -- Asking Haskell to automatically provide equality 
						-- tests and showing functions
```
- Haskell is lazy (it only does the computations it needs to do)
```Haskell
-- This is how (&&) is programmed in the Prelude
(&&) :: Bool -> Bool -> Bool
(&&) True True = True
(&&) _ _ = False

-- Scott's Bottom
bottom :: a -- a is a type variable
bottom = bottom -- bottomless recursion

False && bottom === (&&) False bottom
```
- In the above example, Scott's bottom is an "element" of each function that never terminates
	- It is thus an undefined or divergent element
	- Any program that cause bottom to be evaluated will also not terminate
	- Thus, one can use bottom to to determine whether a program touches an arguement with this function
# Matching
- In the evaluation of a Haskell program, there are a number of steps that require *matching*
- Matching is an algorithm which takes in two terms (ordered trees)
	- One of these is the *pattern* (or a template) and can have variables at the leaves, all of which must be distinct
	- The other is sometimes called the *subject* term: the purpose is to determine whether the subject term matches the pattern term
- Explicitly, the matching algorithm determines whether there is a substitution of the variables of the pattern term which will make the two terms equal
	- For instance, in `fib 42 =: fib n`, `fib 42` is the subject term and `fib n` is the pattern term (which has the variable `n`)
	- The matching algorithm in this case succeeds and produces the substituion `[n := 42]`
- Matching can fail
	- For example, the matching algorithm for `fib 42 =: fib 0` fails: here the pattern term has no variable so the subject term must equal the pattern if the matching is to succeed
- The matching algorithm is simple and can be described recursively by:
	- If the pattern is a variable, return the substitution of that variable by the subject: `t =: x = [x := t]`
	- If the pattern starts with a term constructor (a function symbol) with a number of arguments the subject term must start in exactly the same way otherwise matching fails. The arguments are then recursively matched and the union of the substitutions generated for each argument returned. Thus $F(p_{1},\dots,p_{n}):=F(t_{1},\dots,t_{n})=\bigcup_{i=1\dots n}p_{i} := t_{i}$
# Currying
- Currying is a way of turning a function which takes multiple arguments into a series of functions that each take a single argument and return a single value
- Currying is used for:
	- Defining functions with multiple arguments
	- Creating useful functions from partial applications of curried functions
- Standard approach of defining function with multiple functions vs currying approach:
	- Standard: `add :: (x, y) -> Int`
		 - Functions can only take one argument and return 1 value
	- Curried: `c_add :: Int -> (Int -> Int)`
		- The function takes the x, returns another function that takes the y and the function that took the y, returns the result
		- Returns a function - thus curried functions are high-order functions
- Currying shorthand :
	- The following two type definitions are the same
		- `sum :: Int -> (Int -> (Int -> (Int -> Int)))`
		- `sum :: Int -> Int -> Int -> Int -> Int`
	- The following two functions are actually the same:
		- `sum a b c d`
		- `(((sum a) b) c) d`
- Arrow types are right-associative
	- `a -> b -> c -> d == a -> (b -> (c -> d))`
- Function application is left associative
	- `f :: a -> (b -> (c -> d))`
		- Is applied in this manner `((f x1) x2) x3`
	- `f x1 x2 g x3 x4` where `f :: a -> b -> c -> d` and `g :: e -> f -> c`
		- Is applied in this manner `((((f x1) x2) g) x3) x4 => TYPE ERROR`
		- Should actually be `f x1 x2 (g x3 x4)`
- Function application binds more tightly than any operator `(+, -, ...)`
	- `f n + 1` => `(f n) + 1`
- Function application binds more tightly than the composition operator
	- `f . g x => f . (g x) ~> TYPE ERROR`
	- `(f . g) x`
- Haskell functions are curried by default
- Here's an example of a partially applied function:
```Haskell
multiply :: Int -> Int -> Int
multiply x y = x * y

double :: Int -> Int
double = multiply 2
```
# Higher-Order Functions
- Take in a function as a parameter and returns either a function or value itself
- Examples:
	- `mapList :: (a -> b)  -> [a] -> [b]`
		- Applies the given function to all elements of the given list
	- `filter :: (a -> Bool) -> [a] -> [a]` where the function is called a predicate
		- Only keeps in the list the things that satisfy a predicate
	- `(.) :: (b -> c) -> (a -> b) -> (a -> c)`: Composition operator
		- Applies one function, then another
		- E.g. `(f . g) x = f (g x)`
# Questions
