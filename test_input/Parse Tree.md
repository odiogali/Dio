2024/09/29 18:35
Status: #idea
Tags: #academic #theory #algorithm #language 
# What is a Parse Tree?
A parse tree or parsing tree is an ordered, rooted tree that represents the syntactic structure of a string according to some context-free grammar. The term "parse tree" is used primarily in computational linguistics, in theoretical syntax, the term syntax tree is more common. Parse trees can be used to represent real-world constructions like sentences or mathematical expressions.
___
# Definition
A parse tree for a [[Regular Expressions|regular expression]] $\omega$ over $\Sigma$ is a rooted tree that can be defined as follows, for a given regular expression $\omega \in \Sigma*$:
- **(a)** If $\omega$ is a symbol $\sigma\in \Sigma$ then the parse tree for $\omega$ has a single node with label $\sigma$
- **(b)** If $\omega$ is the symbol "$\lambda$", then the parse tree for $\omega$ has a single node, with label "$\lambda$"
- **(c)** If $\omega$ is the symbol "$\emptyset$", then the parse tree for $\omega$ has a single node, with label "$\emptyset$"
- **(d)** If $\omega$ is the symbol "$\Sigma$", then the parse tree for $\omega$ has a single node with label "$\Sigma$"
- **(e)** If $\omega$ is a regular expression "($\mu \cup \nu$)", for regular expressions $\mu$ and $\nu$ over $\Sigma$, then the parse tree for $\omega$ has a root with label "$\cup$" with two children
	- The first child is the root of a parse tree for $\mu$ and the second child is the root of a parse tree for $\nu$
- **(f)** If $\omega$ is a regular expression "($\mu \circ\nu$)", for regular expressions $\mu$ and $\nu$ over $\Sigma$, then the parse tree for $\omega$ has a root with label "$\circ$" with two children
	- The first child is the root of a parse tree for $\mu$ and the second child is the root of a parse tree for $\nu$
- **(g)** If $\omega$ is a regular expression "($\mu$)", for a regular expression $\mu$ over $\Sigma$, then the parse tree for $\omega$ has a root with label "\*" with one child
	- The child is the root of a parse tree for $\mu$
For example, if $\Sigma=\{0,1\}$, then the parse tree for the regular expression (((((Σ⋆) ◦ 1) ◦ (λ ∪ 0)) ◦ 1) ◦ (Σ)⋆) is the following:![[Pasted image 20240929191049.png]]
# Parsing
Consider, now, the following computational problem.
## The "Parsing" Problem:
- *Precondition*: A string $\omega\in \hat{\Sigma}*$ is given as input
- *Postcondition*: If $\omega$ is a regular expression over $\Sigma$ then a parse tree for $\omega$ is returned as output. An empty tree is returned as output otherwise
It will be useful to have a solution for the following problem when developing a solution for the above one.
## The "Splitting" Problem:
- *Precondition*: A string $\omega=\sigma_{1}\sigma_{2}\dots \sigma_{n}\in \hat{\Sigma}*$ with length $n\geq{3}$ such that $\sigma_{1}=$ "(" and $\sigma_{n}$ = ")" is given as input 
- *Postcondition*: An integer $k$ such that $1 \leq k \leq n$ is returned as output. If $\omega$ is a regular expression over $\Sigma$ then $n \geq 5$ and the following properties are satisfied.
	- **(a)** $3 \leq k \leq n - 2$ and either $\sigma_k = \text{"} \cup \text{"}$ or $\sigma_k = \text{"}\circ\text{"}$.
	- **(b)** The following strings are regular expressions over $\Sigma$
  $$
  \mu = \sigma_2 \sigma_3 \dots \sigma_{k-1} \quad \text{and} \quad \nu = \sigma_{k+1} \sigma_{k+2} \dots \sigma_{n-1}
  $$
![[Regular Expressions#Going From a String to a Regular Expression]]

- Lemma 2 can be used to develop and prove the correctness of an algorithm for the "splitting" proble: it suffices to sweep over the input string, $\omega$ in order to keep track of the difference between the number of copies of "(" and ")" that have been seen so far
	- If an integer $k$ is found such that $1 ≤ k ≤ n$, $\omega_{k}$ is either "$\cup$" or "$\circ$" and there is exactly one more copy of "(" than there is of ")", then this integer $k$ should be returned as output
	- If no such integer $k$ with these properties exists, then $\omega$ cannot be a regular expression at all and it suffices to return the value $k=1$
- Similarly, since $\omega$ begins with "(" and ends with ")" - $\omega$ cannot be a regular expression (so that 1 can be returned) if $|\omega|\leq{4}$
```Psuedo
splitting ( ω: ̂ Σ⋆ ) {  
	1. integer n := |ω|  
	2. if (n ≥ 5) {  
		// Suppose that ω = σ1σ2 . . . σn  
		3. integer k := 0  
		4 integer diff := 0  
		5. while (k < n)  
			6. k := k + 1  
			7. if ( σk == “ ( ” ) {  
				8. diff := diff + 1  
			9. } else if ( σk == “ ) ” ) {  
				10. diff := diff − 1  
			}  
			11. if ((diff == 1) and ((σk == “ ∪ ”) or (σk == “ ◦ ”))) {  
				12. return k  
			}  
		}  
		13. return 1  
	} else {  // If n < 5
		14. return 1  
	}  
}
```
- Once we have an algorithm for splitting, an algorithm that recursively solves the "parsing" problem is easy to describe
```Psuedo
parsing ( ω : ̂ Σ⋆ ) {  
	1. integer n := |ω|  
	// Suppose ω = σ1σ2 . . . σn  
	2. if (n == 0) {  
		3. Return an empty tree.  
	4. } else if ( n == 1 ) {  
		5. if ((σ1 ∈ Σ) or (σ1 ∈ {“ λ ”, “ ∅ ”, “ Σ ”})) {  
		6. Return a parse tree with size one, whose root has label σ1.  
		} else {  
			7. Return an empty tree.  
		}  
	8. } else if ((n≥4) and (σ1 == “(”) and (σn−1 == “)”) and (σn == “⋆”)) {  
		9. Set μ to be the string σ1σ2 . . . σn−1, so that ω = (μ)⋆  
		10. Set ̂ T to be the tree parsing(μ)  
		11. if ( ̂ T is not an empty tree ) {  
			12. Return a parse tree whose root has label “⋆” and a single child — which is the root of the parse tree ̂ T.  
		} else {  
			13. Return an empty tree.  
		}
	14. } else if ((n≥5) and (σ1 == “(”) and (σn == “)”)) {  
		15. integer k := splitting(ω)  
		16. if ((σk == “∪”)) or (σk == “◦”)) {  
			17. Set μ to be the string σ2σ3 . . . σk−1 and set ν to be the string  
			σk+1σk+2 . . . σn−1, so that ω = ( μ σk ν )  
			18. Set TL to be the parse tree parsing(μ)  
			19. Set TR to be the parse tree parsing(ν)  
			20. if ((TL is not an empty tree) and (TR is not an empty tree)) {  
				21. Return a parse tree whose root has label σk, with two children: The left child is the root of the parse tree TL, and the right child is the root of the parse tree TR.  
			} else {  
				22. Return an empty tree.  
			}  
		} else {  
			23. Return an empty tree.  
		}  
	} else {  
		24. Return an empty tree.  
	}  
}
```
- Any execution of the parsing algorithm is at most quadratic in the length of the input string
	- Thus, this is a "polynomial-time" algorithm
# Questions