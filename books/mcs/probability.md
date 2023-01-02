# Probability

## 16. Event and Probability Spaces

**Definition 1** : A countable _sample space_ $S$ is a nonempty countable set. An element $\omega\isin S$ is called an _outcome_. A subset of $S$ is called an _event_

**Definition 2** : A _probability function_ on a sample space $S$ is a total function $Pr:S\to R$ such that
+ $Pr[\omega]>=0$ for all $\omega \isin S$
+ $\sum_{\omega \isin S}Pr[\omega]=1$

A sample space together with a probability function is called a _probability space_. For any event $E \subseteq S$, _the probability of $E$_ is defined to be the sum of the probabilities of the outcomes in $E$
$Pr[E] \coloncolonequals \sum_{\omega \isin E}Pr[\omega]$

::: tip  Probability Rules from Set Theory
+ _**Sum Rule**_ : If $E_0,E_1,...,E_n...$ are _pairwise disjoint events_, then
$$Pr[\bigcup_{n \isin \N}E_n]=\sum_{n \isin \N}Pr[E_n]$$
+ _**Complement Rule**_ : Another consequence of the Sum Rule is that $Pr[\overline{A}]=1-Pr[A]$, which follows because $Pr[S]=1$ and $S$ is the union of the disjoint sets $A$ and $\overline{A}$
$$Pr[\overline{A}]=1-Pr[A]$$
+ _**Difference Rule**_
$$Pr[B-A]=Pr[B]-Pr[A \cap B]$$
+ **_Inclusion-Exclusion_**
$$Pr[A \cup B] = Pr[A]+Pr[B]-Pr[A \cap B]$$
+ **_Boole's Inequality_**
$$Pr[A \cup B] <= Pr[A]+Pr[B]$$
+ _**Monotonicity Rule**_
$$A \subseteq B \implies Pr[A] <= Pr[B]$$
+ _**Union Bound**_
$$Pr[E_1 \cup ... \cup E_n \cup ...] <= Pr[E_1]+...+Pr[E_n]+...$$
:::

_**Definition 3**_ : A finite probability space $S$, is said to be _uniform_ if $Pr[\omega]$ is the same for every outcome $\omega \isin S$

Uniform sample spaces are particularly easy to work with. That's because for every event $E \subseteq S$
$$Pr[E]=\frac{|E|}{|S|}$$

## 17. Conditional Probability

**Definition 4** : Let $X$ and $Y$ be events where $Y$ has nonzero probability. Then
$$Pr[X|Y] \coloncolonequals \frac{Pr[X \cap Y]}{Pr[Y]}$$
The Conditional Probability Product for $n$ events is
$$Pr[E_1 \cap E_2 \cap ... \cap E_n]=Pr[E_1] \cdot Pr[E_2 | E_1] \cdot Pr[E_3|E_1 \cap E_2] ... \cdot Pr[E_n | E_1 \cap E_2 \cap ... \cap E_{n-1}]$$

## 18. Random Variables
A _random variable_ $R$ on a probability space is a total function whose domain is the sample space

> Notice that the name random variable is a misnomer; random variables are actually functions

### 18.1 Random Variables Examples
An _indicator random variable_ is a random variable that maps every outcome to either 0 or 1. Indicator random variables are also called _Bernoulli variables_

In particular, an indicator random variable partitions the sample space into those outcomes mapped to 1 and those outcomes mapped to 0

In the same way, an event $E$ partitions the sample space into those outcomes in $E$ and those not in $E$. So $E$ is naturally associated with an indicator random variable $I_E$, where $I_{E}(w)=1$ for outcomes $w \isin E$ and $I_{E}(w)=0$ for outcomes $w \not \isin E$

There's a strong relationship between events and more general random variables as well. A random variable that takes on several values partitions the sample space into several blocks. **More generally any assertion about the values of random variables defines an event**

### 18.2 Independence

Random variables $R1$ and $R2$ are _independent_ iff for all $x_1, x_2$, the two events
$$[R_1=x_1] \text{ and } [R_2=x_2]$$
are independent



## 19. Deviation from the Mean


## 20. Random Walks

