2024/05/14 13:18
Status: #idea
Tags: #concept #idea 
# What is Asymptotic Notation
Asymptotic notation and analysis allows us to consider the growth rate of the running time of the algorithm as a function of its input size, n. The most famous asymptotic notation tool is [[Big O Notation]].
___
# Big O Family
- Big-Omega (asymptotically >=)
	- $f(n)$ is Ω(g(n)) -> $f(n) \geq c×g(n)$ 
	- $f(n)$ is Ω(g(n)) if $g(n)$ is O($f(n)$) ![[Pasted image 20240514223619.png|600]]
- Big-Theta (asymptotically =)
	- $f(n)$ is Θ(g(n)) -> $f(n) = c×g(n)$ 
	- $f(n)$ is Θ(g(n)) if $f(n)$ is O(g(n)) and $f(n)$ is Ω(g(n))![[Pasted image 20240514223650.png|550]]

