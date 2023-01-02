# Counting
## 13. Sums and Asymptotics

Expressions like $\frac{n(n+1)}{2}$ that do not make subscripted summations or products, or those handy but sometimes troublesome sequences of three dots, are called _closed form_

Another example is the closed form for a *geometric sum*

$$\forall n \isin \N \land x \not = 1.  1+x+x^2+x^3+...+x^n = \frac{1-x^{n+1}}{1-x}$$

::: tip Infinite Geometric Series
If $|x| < 1$, then
$$\sum^{\infty}_{i = 0}{x^i}=\frac{1}{1-x}$$

_Proof_.
$$
\begin{split}
\sum^{\infty}_{i=0}{x^i} &\coloncolonequals \lim_{n \to \infty}{\sum^n_{i=0}x^i}\\
&= \lim_{n \to \infty}{\frac{1-x^{n+1}}{1-x}}\\
&= \frac{1}{1-x}
\end{split}
$$
The final line follows from the fact that $\lim_{n \to \infty}x^{n+1} = 0$ when $|x| < 1$

:::

## 14. Cardinality Rules

### 14.2 Counting Sequences

The _Product Rule_ gives the size of a product of sets. Recall that if $P_1,P_2,..,P_n$ are sets, then
$$P_1 \times P_2 \times ... \times P_n$$
is the set of all sequences whose first term is drawn from $P_1$, second term is drawn from $P_2$, and so forth

_**Product Rule**_: If $P_1 \times P_2 \times ... \times P_n$ are finite sets, then:
$$|P_1 \times P_2 \times ... \times P_n| = |P_1|\cdot|P_2|...|P_n|$$

We can write the set of all $n$-bit 0/1 sequences as a product of sets
$$\lbrace 0,1 \rbrace^n \coloncolonequals \underbrace{\lbrace 0,1 \rbrace \times \lbrace 0,1 \rbrace \times ... \times \lbrace 0,1 \rbrace}_{n \text{ terms}}$$
Then Product Rule gives the answer
$$|\lbrace 0,1 \rbrace^n| = |\lbrace 0,1 \rbrace|^n = 2^n$$

_**Sum Rule**_: If $A_1, A_2, ..., A_n$ are disjoint sets, then:
$$|A_1 \cup A_2 \cup ... \cup A_n| = |A_1|+|A_2| + ... + |A_n|$$

::: section Generalized Product Rule
Let $S$ be a set of length-$k$ sequences. If there're:
+ $n_1$ possible first entries
+ $n_2$ possible second entries for each first entry,
...
+ $n_k$ possible kth entries for each sequence of first $k-1$ entries
then
$$|S| = n_1 \cdot n_2 \cdot ... n_k $$
:::
