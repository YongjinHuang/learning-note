# Structures

The default convention in this chapter is that variables range over the set of integers, $\Z$

## 8. Number Theory

### 8.1 Divisibility
The nature of number theory emerges as soon as we consider the _divides_ relation

**_$a$ divides $b$_ (notation $a | b$) iff there's an integer $k$ such that**
$$ak = b$$

Some immediate consequences of the definition above are that
$$\forall n \isin \Z. n|0; n|n; \underline{+}1|n$$
$$0|n \implies n=0$$

The following lemma collects some basic facts about divisibility
1. $a|b \land b|c \implies a|c$
1. $a|b \land a|c \implies a|sb+tc$
1. $\forall c \not = 0. a|b \iff ca|cb$

A number of the form $sb+tc$ is called an _integer linear combination_ of b and c

**An integer is a linear combination of numbers $b_0,...b_k$, iff**
$$n=s_0b_0+s_1b_1...+s_kb_k$$
**for some integers $s_0,...,s_k$**

**Division Theorem: Let $n$ and $d$ be integers such that $d>0$. Then there exists a unique pair of integers $q$ and $r$, such that**
$$n=q\cdot d+r \land 0 <= r < d$$
The number $q$ is called the _quotient_ ($qcnt(n,d)$) and the number $r$ is called the _remainder_ of $n$ divided by $d$ ($rem(n,d)$)

### 8.2 The Greatest Common Divisor
A _common divisor_ of $a$ and $b$ is a number that divides them both. The _greatest common divisor_ of $a$ and $b$ is written $gcd(a,b)$

Some immediate consequences of the definition of gcd are that for $n > 0$
$$gcd(n,n)=n, gcd(n,1)=1,gcd(n,0)=n$$
where the last equality follows from the fact that everything is a divisor of 0

**Euclid's Algorithm** : For $b \not = 0$
$$gcd(a,b)=gcd(b,rem(a,b))$$

An integer is a linear combination of $a$ and $b$ iff it's a multiple of $gcd(a,b)$

**The greatest common divisor of $a$ and $b$ is a linear combination of $a$ and $b$. That is**
$$gcd(a,b)=sa+tb$$
for some integers $s$ and $t$


### 8.3 Prime Mysteries
**A _prime_ is number greater than 1 that is divisible only by itself and 1. A number other than 0, 1, and -1 that is not a prime is called _composite_**

Let $\pi(n)$ denote the number of primes up to $n$
$$\pi(n) \coloncolonequals |\lbrace p \isin [2..n] | p \text{ is prime} \rbrace|$$

**Prime Number Theorem**
$$\lim_{n \to \infty}\frac{\pi(n)}{\frac{n}{\ln n}} = 1$$
Thus, primes gradually taper off

(Chebyshev's Theorem on Prime Density). For $n > 1$
$$\pi(n) > \frac{n}{3\ln n}$$

### 8.4 The Fundamental Theorem of Arithmetic
A sequence of numbers is _weakly decreasing_ when each number in the sequence is at least as big the numbers after it

::: tip Fundamental Theorem of Arithmetic
Every postive integer is a product of a unique weakly decreasing sequence of primes

The Fundamental Theorem is also called the _Unique Factorization Theorem_
:::

If $p$ is a prime and $p|ab$, then $p | a$ or $p|b$

Let $p$ be a prime. If $p | a_1a_2...a_n$, then p divides some $a_i$

<!-- ### 8.5 Alan Turing -->

### 8.6 Modular Arithmetic
**$a$ is congruent to b modulo $n$ iff $n | (a-b)$**
$$a \equiv b \pmod n$$

$$a \equiv b \pmod n \iff rem(a,n) = rem(b,n)$$

$$a \equiv rem(a,n) \pmod n$$


::: tip Equality Relation
A relation is an _equivalence relation_ if it's reflexive, symmetric and transitive
:::

The congurence relation has properties like an equality relation
+ $$a \equiv a \pmod n$$
+ $$a \equiv b \iff b \equiv a \pmod n$$
+ $$a \equiv b \land b \equiv c \implies a \equiv c \pmod n$$

The next most useful fact about congruences is that they are preserved by addition and multiplication

If $a \equiv b \pmod n$ and $c \equiv d \pmod n$, then
$$a + c \equiv b + d \pmod n \\
ac \equiv bd \pmod n
$$

Integers that have no prime factor in common are called _relatively prime_(Other texts call them _coprime_). This's the same as having no common divisor (prime or not) greater than 1. It's also equivalent to saying that $gcd(a,b)=1$

**If $k \isin \lbrack0..n)$ is relatively prime to $n$, then $k$ has an inverse in $\Z_n$**

### 8.7 Remainder Arithmetic
::: tip General Principle of Remainder
To find the remainder on division by $n$ of the result of a series of additions and multiplications, applied to some integers:
+ replace each integer operated by its remainder on division by $n$
+ keep each result of an addition or multiplication in the range $[0..n)$ by immediately replacing any result outside that range by its remainder on division by $n$
:::
Let's introduce the notation $+_n$ for doing an addition and then immediately taking a remainder on division by $n$, as specified by the general principle; likewise for multiplying:
$$
i+_nj \coloncolonequals rem(i+j,n)\\
i \cdot_nj \coloncolonequals rem(ij,n)
$$
$$
rem(i+j,n)=rem(i,n)+_nrem(j,n)\\
rem(ij,n)=rem(i,n)\cdot_nrem(j,n)
$$
**The set of integers in the range $[0..n)$ together with the operations $+_n$ and $\cdot_n$ is refered to as $\Z_n$, _the ring of integers modulo $n$_**

### 8.9 Multiplicative Inverses and Cancelling
The multiplicative inverse of a number $x$ is another number $x^{-1}$ such that
$$x^{-1} \cdot x = 1$$

Integers that have no prime factor in common are called _relatively prime_. This is the same as having no common divisor(prime or not) greather than 1. It's also equivalent to saying that $gcd(a,b)=1$

If $k \isin [0..n)$ is relatively prime to $n$, then $k$ has an inverse in ${\Z}_n$

If $i$ and $j$ are both inverses of $k$ in ${\Z}_n$, then $i = j$

A number $k$ is _cancellable_ in $Z_n$ iff
$$k \cdot a = k \cdot b \implies a = b (\Z_n)$$
for all $a,b \isin [0..n)$

if $k$ has an inverse in $\Z_n$, then it's cancellable

The following are equivalent for $k \isin [0..n)$
+ $gcd(k,n)=1$
+ $k$ has an inverse in $\Z_n$
+ $k$ is cancellable in $\Z_n$

### 8.10 Euler's Theorem
For $n>0$, define
$$\phi(n) \coloncolonequals \text{ the number of integers in (0..n) that are relatively prime to n }$$
This function $\phi$ is known as Euler's $\phi$ function. More generally, if $p$ is a prime, then $\phi(p)=p-1$ since every positive number in $(0..p)$ is relatively prime to $p$

::: section Euler's Theorem
if $n$ and $k$ are relatively prime, then
$$k^{\phi(n)}\equiv 1 \pmod n$$
:::

Let $\Z^{*}_n$ be the integers in $(0..n)$, that are relatively prime to $n$
$$\Z^{*}_n \coloncolonequals \lbrace k \isin (0..n) | gcd(k, n) = 1 \rbrace$$

Consequently,
$$\phi(n) = |\Z^{*}_n|$$

::: section Euler's Theorem
for all $k \isin \Z^{*}_n$
$$k^{\phi(n)}=1(\Z_n)$$
:::

::: section Fermat's Little Theorem
Suppose $p$ is a prime and $k$ is not a multiple of $p$. Then
$$k^{p-1}\equiv 1 \pmod p$$
:::

$$\phi(pq)=(p-1)(q-1)$$
for primes $p\not=q$
1. If $p$ is a prime, then $\phi(p^k) = p^k-p^{k-1}$ for $k>=1$
2. If $a$ and $b$ are relatively prime, then $\phi(ab)=\phi(a)\phi(b)$


### 8.11 RSA Public Key Encryption
::: section The RSA Cryptosystem
TODO
:::



## 9. Directed graphs & Partial Orders
_Directed graphs_, called _digraphs_ for short, provide a handy way to represent how things are connected together and how to get from one thing to another by following those connections. They're usually pictured as a bunch of dots or circles with arrows between some of the dots. The dots are called _nodes_ or _vertices_ and the lines are called _directed edges_ or _arrows_

**A directed graph $G$, consists of a nonempty set, $V(G)$, called the vertices of $G$, and a set $E(G)$, called the edges of $G$. An element of $V(G)$ is called a vertex(node). An element of $E(G)$ is called a _directed edge_(arrow). A directed edge starts at some vertex, $u$, called the tail of the edge, and ends at some vertex, $v$, called the head of the edge. Such an edge can be represented by the ordered pair $(u,v)$. The notation $u \to v$ denotes the edge**

### 9.1 Vertex Degrees
The _in-degree_ of a vertex in a digraph is the number of arrows comming into it, and similarily its _out-degree_ is the number of arrows out of it. More precisely

::: section Vertex Degrees
If $G$ is a digraph and $v \isin V(G)$, then
$$
indeg(v) \coloncolonequals |\lbrace e \isin E(G) | head(e)=v \rbrace| \\
outdeg(v) \coloncolonequals |\lbrace e \isin E(G) | tail(e)=v \rbrace|
$$
An immediate consequence of this definition is
$$\sum_{v \isin V(G)}indeg(v) = \sum_{v \isin V(G)}outdeg(v) = |E(G)|$$
:::

### 9.2 Walks and Paths
::: section Walks and Paths
A walk in a digraph $G$, is an alternating sequence of vertices and edges that begins with a vertex, ends with a vertex, and such that for every edge $(u \to v)$ in the walk, vertex $u$ is the element just before the edge, and vertex $v$ is the next element after the edge

So a walk $\bold v$ is a sequence of the form
$$\bold v \coloncolonequals v_0 (v_0 \to v_1) v_1 (v_1 \to v_2) v_2 ...(v_{k-1}, v_k) v_k$$
where $(v_i \to v_{i+1}) \isin E(G)$ for $i \isin [0..k)$. The walk is said to start at $v_0$, to end at $v_k$, and the length, $|\bold v|$, of the walk is defined to be $k$

The walk is a path iff all the $v_i$'s are different, that is , if $i \not= j$, then $v_i \not= v_j$
:::
The _closed walk_ is a walk that begins and ends at the same vertex

A _cycle_ is a postive length closed walk whose vertices are distinct except for the beginning and end vertices

::: warning Zero Path
Note that a single vertex counts as a length zero path that begins and ends at itself. It is also a closed walk, but does not count as a cycle, since cycles by definition must have positive length

Length one cycles are sometimes called _self-loops_
:::

In general, if a walk $\bold f$ ends with a vertex $v$, and a walk $\bold r$ starts with the same vertex $v$, we'll say that their _merge_ $\bold f\text{\textasciicircum}\bold r$ is the walk that starts with $\bold f$ and continues with $\bold r$

Sometimes it's useful to name the node $v$ where the walks merge; we'll use the notation $\bold f \hat{v} \bold r$ to describe the merge of a walk $\bold f$ that ends at $v$ with a walk $\bold r$ that begins at $v$

A consequence of this definition is that 

$$|\bold f\text{\textasciicircum}\bold r| = |\bold f| + |\bold r|$$

::: section Finding a Path
**The shortest walk from one vertex to another is a path**

The _distance_, $dist(u,v)$, in a graph from vertex $u$ to vertex $v$ is the length of a shortest path from $u$ to $v$

$$\text{The Triangle Inequality } dist(u,v)<= dist(u,w)+dist(w,v)$$
for all vertices $u,v,w$ with equality holding iff $w$ is on a shortest path from $u$ to $v$ 
:::

**The shortest postive length _closed walk_ through a vertex is a _cycle_ through that vertex**

### 9.3 Adjacency Matrices
If a graph $G$ has n vertices $v_0, v_1, ..., v_{n-1}$, a useful way to represent it is with an $n \times n$ matrix of zeroes and ones called its _adjacency matrix_, $A_G$. The $ij$ th entry of the adjacency matrix $(A_G)_{ij}$ is $1$ if there's an edge from vertex $v_i$ to vertex $v_j$ and $0$ otherwise. That is:
$$
(A_G)_{ij} \coloncolonequals \begin{cases}
   1 &\text{if } (v_i \to v_j) \isin E(G) \\
   0 &\text{otherwise}
\end{cases}
$$

::: section The length-$k$ walk counting matrix
The length-$k$ walk counting matrix for an $n$-vertex graph $G$ is the $n \times n$ matrix $C$ such that
$$C_{uv} \coloncolonequals \text{the number of length-k walks from u to v}$$
:::

### 9.4 Walk Relations
For any digraph $G$, we're interested in a binary relation $G^*$, called the _walk relation_ on $V(G)$ where
$$uG^*v \coloncolonequals \text{ there's a walk in G from u to v}$$
Similarly, there's a _positive walk relation_
$$uG^+v \coloncolonequals \text{ there's a positive length walk in G from u to v}$$
When there's a walk from vertex $v$ to vertex $w$, we say that $w$ is _reachable_ from $v$, or equivalently, that $v$ is connected to $w$

Let $G^n$ denote the composition of $G$ with itself $n$ times, it's easy to check that $G^n$ is the length-$n$ walk relation:
$$aG^nb \iff \text{there's a length n walk in G from a to b}$$


### 9.5 Directed Acyclic Graphs & Scheduling
::: section Directed Acyclic Graph
A _directed asyclic graph_(DAG) is a directed graph with no cycles
:::

A _topological sort_ of a finite DAG is a list of all the vertices such that each vertex $v$ appears earlier in the list than every other vertex reachable from $v$

A vertex $v$ of a DAG, $D$, is _minimum_ iff every other vertex is reachable from $v$

A vertex $v$ is _minimal_ iff $v$ is not reachable from any other vertex

A DAG may have no minimum element but lots of minimal elements

***Every finite DAG has a topological sort***


### 9.10 Equivalence Relations
A relation is an _equivalence relation_ if it's reflexive, symmetric and transitive

If $f: A \to B$ is a total function, define a relation $\equiv_f$ by the rule:
$$a \equiv_f a' \iff f(a) = f(a')$$


## 10. Communication Networks


## 11. Simple Graphs


## 12. Planar Graphs

