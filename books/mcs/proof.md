#  Proofs

## 1. What is a Proof
+ **_Proofs_** : Simply put, a proof is a method of establishing truth. A _mathematical proof_ of a proposition is a chain of _logical deductions_, leading to the proposition from a base set of _axioms_
+ **_Propositions_** : A _proposition_ is a statement (communication) that is either true or false
+ **_Predicates_** : A _predicate_ can be understood as a proposition whose truth depends on the value of one or more variables. So "n is a perfect square" describes a predicate, since you can't say if it's true or false until you know what the value of the variable n happens to be
::: warning The notation for predicates
The notation for predicates is confusingly similar to ordinary function notation. If P is a predicate, then P(n) is either true or false, depending on the value of n. On the other hand, if p is an ordinary function, like $n^2+1$, then p(n) is a *numerical quantity*
:::
+ **_Theorems_** : Important true propositions are called _theorems_
+ **_Lema_** : A _lemma_ is a preliminary proposition useful for proving later propositions
+ **_Corollary_** : A _corollary_ is a proposition that follows in just a few logical steps from a theorem
+ _**Implications**_ : Propositions of the form "If P, then Q" are called _implications_
+ _**Logical Deduction**_ : Logical deduction, or _inference rules_, are used to prove new propositions using previously proved ones

Proposition of the form "If $P$ then $Q$" are called _implications_. This implication is often rephrased as $P \implies Q$

There're a couple of standard methods for proving an implication
:::: row Proving an Implication
::: col
_**Method 1**_ : In order to prove that $P \implies Q$
  1. Write "Assume $P$"
  2. Show that $Q$ logically follows
:::
::: col
_**Method 2**_ : An implication $P \implies Q$ is logically equivalent to its _contrapositive_
$$\lnot Q \implies \lnot P$$
Proving one is as good as proving the other, and proving the contrapositive is sometimes easier than proving the original statement. If so, then you can proceed as follows:
1. Write "We prove the contrapositive" and then state the contrapositive
1. Proceed as in Method 1
:::
::::

:::: row Proving an If and Only If
::: col
**Prove Each Statement Implies the Other**

The statement $P \iff Q$ is equivalent to the two statements $P \implies Q$ and $Q \implies P$. So you can prove an "iff" by proving _two_ implications
1. Write "We prove $P \implies Q$"  and vice-versa
1. Write "First, we show $P \implies Q$. Do this by one of the methods in Proving an Implication"
1. Write "Now, we show $Q \implies P$. Again, do this by one of the methods in Proving an Implication"
:::
::: col
**Construct a chain of Iffs**

In order to prove that $P \iff Q$ is true:
1. Write "We construct a chain of if-and-only-if implications"
1. Prove $P$ is equivalent to a second statement which is equivalen to a third statement and so forth until you reach $Q$
:::
::::

:::: row Proof by Cases
::: col
Breaking a complicated proof into cases and proving each case separately is a common useful strategy
:::
::::

:::: row Proof by Contradiction
::: col
In a proof by contradiction, or _indirect proof_, you show that if a proposition were false, then some false fact would be true. Since a false fact by definition can't be true, the proposition must be true

**Method** : In order to prove a proposition $P$ by contradiction:
1. Write "We use proof by contradiction"
1. Write "Suppose $P$ is false"
1. Deduce something known to be false (a logical contradiction)
1. Write "This is a contradiction. Therefore $P$ must be true"
:::
::::

Offer some general tips on writing good proofs:
+ State your game plan
+ Keep a linear flow
+ A proof is an essay, not a calculation
+ Avoid excessive symbolism
+ Revise and simplify
+ Introduce notation thoughtfully
+ Structure long proofs
+ Be wary of "obvious"
+ Finish



## 2. The Well Ordering Principle

::: tip The Well Ordering Principle
Every nonempty set of non-negative integers has a smallest element
:::

::: section Template for Well Ordering Principle
To prove that "$P(n)$ is true for all $n \isin \N$ using the Well Ordering Principle"
+ Define the set $C$, of _counterexamples_ to $P$ being true. Specifically, define
$$C \coloncolonequals \lbrace n \isin \N | \lnot P(n) \rbrace$$
> The notation $\lbrace n|Q(n) \rbrace$ means that "the set of all elements $n$ for which $Q(n)$ is true"
+ Assume for proof by contradiction that $C$ is nonempty
+ By the Well Ordering Principle, there will be a smallest element $n$ in $C$
+ Reach a contradiction somehow -- often by showing that $P(n)$ is actually true or by showing that there's another member of $C$ that is smaller than $n$. This's the open-ended part of the proof task
+ Conclude that $C$ must be empty, that is, no counterexamples exist
:::

A set of numbers is _well ordered_ when each of its nonempty subsets has a minimum element

**Well ordering commonly comes up in computer science as a method for proving that computations won't run forever**

For any nonnegative integer $n$, the set of integers greater than or equal to $-n$ is well ordered

A _lower bound_(respectively, _upper bound_) for a set $S$, of real number is a number $b$ such that $b <= s$ (respectively, $b >= s$) for every $s \isin S$

> Note that a lower or upper bound of set $S$ is not required to be in the set

Any set of integers with a lower bound is well ordered

Any nonempty set of integers with an upper bound has a maximum element


## 3. Logical Formulas
Mathematicians use the words _NOT_, _AND_, and _OR_ for operations that change or combine propositions

> Proposition variables are also called _Boolean variables_

In general, a _truth table_ indicates the true/false value of a proposition for each possible set of truth values for the variables

| $P$ | $\lnot P$ |
| :-- | :-- |
| $T$ | $F$ |
| $F$ | $T$ |

The truth table for implications can be summarized in words as follows:

> An implication is true exactly when the if-part is false or the then-part is true

::: details False Hypotheses
It often bothers people when they first learn that implications which have false hypotheses are considered to be true

> if $C_i$: the system sensors are in condition $i$, then $A_i$: the system takes actiion $i$

Or more concisely:
$[C_1 \implies A_1]\land[C_2\implies A_2]\land ... \land [C_{12} \implies A_{12}]$

:::


The proposition "P if and only if Q" asserts that P and Q have the same truth value. Either both are true or both are false. For example, the following if-and-only-if statement is true for every real number $x$:
$x^2-4>=0 \iff |x|>=2$

| $P$ | $Q$ | $P \land Q$  | $P \lor Q$  | $P \oplus Q$ | $P \implies Q$  | $P \iff Q$ |
| -- | -- | -- | -- | -- | -- | -- | -- |
| $T$ | $T$ | $T$ | $T$ | $F$ | $T$ | $T$ |
| $T$ | $F$ | $F$ | $T$ | $T$ | $F$ | $F$ |
| $F$ | $T$ | $F$ | $T$ | $T$ | $T$ | $F$ |
| $F$ | $F$ | $F$ | $F$ | $F$ | $T$ | $T$ |

The mathematical notation is concise but cryptic

| English | Symbolic Notation |
| -- | -- |
| NOT(P) | $\lnot P$ |
| P AND Q | $P \land Q$ |
| P OR Q | $P \lor Q$ |
| P IMPLIES Q | $P \implies Q$ |
| if P then Q | $P \implies Q$ |
| P IFF Q | $P \iff Q$ |
| P XOR Q | $P \oplus Q$ |

A statement of the form "NOT(Q) IMPLIES NOT(P)" is called the _contrapositive_ of the implication "P IMPLIES Q". **The truth table shows that an implication and its contrapositive are _equivalent_**

$P \implies Q$ is equivalent with $\lnot Q \implies \lnot P$

$(P \implies Q)\land(Q \implies P)$ is the same as $P \iff Q$

**A _valid_ formula is one which is always true, no matter what truth values its variables may have**. The simplest example is $P \lor \lnot(P)$

Equivalence of formulas is really a special case of validity. Namely, statements $F$ and $G$ are equivalent precisely when the statement $F \iff G$ is valid

 A formula is valid iff it's equivalent to true

**A _satisfiable_ formula is one which can sometimes be true**. A statement P is satisfiable iff its negative NOT(P) is not valid


Every propositional formula is equivalent to a "sum-of-products" or _disjunctive form_. More precisely, a disjunctive form is simply an OR of AND-terms, where each AND-term is an AND of variables or negations of variables

$$(A \land B) \lor (A \land \lnot C)$$

An expression of this form is called _disjunctive normal form_ (_**DNF**_)

:::: row Propositions in Normal Form
::: col
**Distributive Law of AND orver OR**
$A \land (B \lor C)$ is equivalent to $(A \land B) \lor (A \land C)$
:::
::: col
**Distributive Law of OR orver AND**
$A \lor (B \land C)$ is equivalent to $(A \lor B) \land (A \lor C)$
:::
::::

_Conjunctive normal form_ (_**CNF**_): An AND of OR-terms in which the OR-terms are OR's only of variables or their negations

Every propositional formula is equivalent to both a disjunctive normal form and a conjunctive normal form

We List below a bunch of equivalence axioms with the symbol $\longleftrightarrow$

::: section Proving Equivalences
Commutativity of AND
$$A \land B \longleftrightarrow B \land A$$
Associativity of AND
$$(A \land B) \land C \longleftrightarrow A \land (B \land C)$$
Identity for AND
$$\bold T \land A \longleftrightarrow A$$
Zero for AND
$$\bold F \land A \longleftrightarrow \bold F$$
Idempotance for AND
$$A \land A \longleftrightarrow A$$
Contradiction for AND
$$A \land \lnot A \longleftrightarrow \bold F$$
Double negation
$$\lnot(\overline{A}) \longleftrightarrow A$$
Validity for OR
$$A \lor \overline{A} \longleftrightarrow \bold T$$
DeMorgan for AND
$$\lnot (A \land B) \longleftrightarrow \overline{A} \lor \overline{B}$$
DeMorgan for OR
$$\lnot (A \lor B) \longleftrightarrow \overline{A} \land \overline{B}$$

Any propositional formula can be transformed into a disjunctive normal form or a conjunctive normal form from using the equivalence listed above

(Completeness of the propositional equivalence axioms). Two propositional formula are equivalent iff they can be proved equivalent using the equivalence axioms listed above
:::

**The general problem of deciding whether a proposition is satisfiable is called _SAT_**

The problem of determining whether or not SAT has a polynomial time solution is known as the $\bold P$ vs. $\bold NP$ problem


:::: row Quantifiers
::: col
**Always True** $\forall$
+ For all $x \isin D$, $P(x)$ is true
+ $P(x)$ is true for every $x$ in the set $D$
:::
::: col
**Sometimes True** $\exists$
+ There's an $x \isin D$ such that $P(x)$ is true
+ $P(x)$ is true for some $x$ in the set $D$
+ $P(x)$ is true for at least one $x \isin D$
:::
::::
All these sentences "quantify" how often the predicate is true. Specifically, an assertion that a predicate is always true is called a _univeral quantification_, and an assertion that a predicate is sometimes true is an _existential quantification_

$\lnot(\forall x, P(x))$ is equivalent to $\exists x, \lnot P(x)$
$\lnot(\exists x, P(x))$ is equivalent to $\forall x, \lnot P(x)$

The general principle is that _moving a NOT across a quantifier changes the kind of quantifier_

$$\lnot (\forall x, P(x)) \iff \exists x, \lnot P(x)$$

A useful example of a valid assertion is 
$$\exists x \forall y, P(x,y) \implies \forall y \exists x, P(x,y)$$

However, the formula
$$\forall y \exists x, P(x,y) \implies \exists x \forall y, P(x,y)$$
is not valid

## 4. Mathematical Data types
### 4.1 Sets
Informally, a _set_ is a bunch of objects, which are called the _elements_ of the set. For example, here are some sets:
+ $A=\lbrace Alex,Tippy \rbrace$
+ $B=\lbrace red,yellow,blue \rbrace$

Bigger sets or infinite sets might be defined by indicating how to generate a list of them:

$D \coloncolonequals \lbrace 1,2,4,8,16,... \rbrace \text{the powers of 2}$

The order of elements is not significant. There's no notion of an element appearing more than once in a set

The expression $e \isin B$ asserts that $e$ is an element of set $S$

| symbol | set | elements |
| -- | -- | -- |
| $\emptyset$ | the empty set | none |
| $\N$ | nonnegative integers | $\lbrace 0,1,2,3,.. \rbrace$ |
| $\Z$ | integers | $\lbrace ...,-3,-2,-1,0,1,2,3,.. \rbrace$ |
| $\mathbf{Q}$ | rational numbers | $\frac{1}{2}, 16, ...$ |
| $\reals$ | real numbers | $\pi,-9,\sqrt{2}...$ |
| $\Complex$ | complex numbers | $i, \frac{19}{2}, \sqrt{2}-2i, ...$ |

A superscript $^+$ restricts a set to its positive elements; for example, $\reals ^+$ denotes the set of positive real numbers. Similarly $\Z ^-$ denotes the set of negative integers

The expression $S \subseteq T$ indicates that set $S$ is a *subset* of set $T$, which means the every elements of $S$ is also an element of $T$. $S \subset T$ means that $S$ is a subset of $T$, but the two are not equal. $S \subset T$ is a _strict subset of $T$_

The _union_ of sets $A$ and $B$, denoted $A \cup B$, includes exactly the elements appearing in $A$ or $B$ or both. That is:
$$x \isin A \cup B \iff x\isin A \lor x \isin B$$

The _intersections_ of $A$ and $B$, denoted $A \cap B$, consists of all elements that appear in _both $A$ and $B$_. That is:
$$x \isin A \cap B \iff x\isin A \land x \isin B$$

The _set difference_ of $A$ and $B$, denoted $A - B$, consists of all elements that are in A, but not in B. That is:
$$x \isin A-B \iff x\isin A \land x \notin B$$

Often all the sets being considered are subsets of a known domain discourse $D$. Then for any subset $A \subseteq D$, we define $\overline{A}$ to be the set of all elements of $D$ *not* in $A$. That is: 
$$
\overline{A} \coloncolonequals D-A
$$
The set $\overline{A}$ is called the *complement* of A. So:
$$\overline{A} = \empty \iff A = D$$

For example, if the domain we're working with is the integers, the complement of the nonnegative integers is the set of negative integers:
$$\N = \Z^-$$
$$A \subseteq B \iff A \cap \overline{B} = \empty$$

The set of all the subsets of a set, $A$, is called the *power set*, $pow(A)$. So:
$$B \isin power(A) \iff B \subseteq A$$

More generally, if $A$ has $n$ lements then there're $2^n$ sets in $pow(A)$

An important use of predicates is in *set builder notation*. The idea is to define a set using a *predicate*; in particular, the set consists of all values that make the predicate true

$$
A \coloncolonequals \lbrace n \isin \N | \text{n is a prime and n = 4k+1 for some integer k} \rbrace \\
B \coloncolonequals \lbrace x \isin \R | x^3-3x+1>0  \rbrace \\
C \coloncolonequals \lbrace a+bi \isin \Complex | a^2+2b^2<=1 \rbrace
$$

Two sets are defined to be equal if they have exactly the same elements. That is $X = Y$ means that $z \isin X$ if and only if $z \isin Y$, for all elements $z$

### 4.2 Sequences
Sets provide one way to group a collection of objects. Another way is in a _sequence_, which is a list of objects called _terms_ or _components_
There're several differences between Sets and sequences
+ The elements of a set are required to be distinct, but terms in a sequence can be the same
+ The terms in a sequence have a special order, but the element of a set do not
+ Texts differ on notation for the _empty sequence_: $\lambda$

Length two sequences are called _pairs_

### 4.3 Functions
A _function_ assigns an element of one set, called the _domain_, to an element of another set, called the _codomain_. The notation:
$$f:A \to B$$
indicates that $f$ is a function with domain $A$ and codomain $B$

In general, functions may be _partial functions_, meaning that there may be domain elements for which the function is not defined

**If a function is defined on every element of its domain, it's called a _total function_**

If $f:A \to B$ and $S \subseteq A$, we define $f(S)$ to be the set of all the values that $f$ takes when it it applied to elements of $S$. That is
$$f(S) \coloncolonequals \lbrace b\isin B | f(s)=b \text{  for some  } s \isin S \rbrace$$

**Applying $f$ to a set $S$ of arguments is referred to as applying $f$ pointwise to $S$, and the set $f(S)$ is referred to as the image of $S$ under $f$**. The set of values that arise from applying $f$ to all possible arguments is called the _range_ of $f$. That is,
$$range(f) \coloncolonequals f(domain(f))$$

For functions $f:A \to B$ and $g: B\to C$, the _composition_, $g \circ f$, of $g$ with $f$ is defined to be the function from $A$ to $C$ defined by the rule:
$$(g \circ f)(x) \coloncolonequals g(f(x))$$
for all $x \isin A$

### 4.4 Binary Relations
**A _binary relation_, $R$, consists of a set $A$ called the domain of $R$, a set called the codomain of the $R$, and a subset of $A \times B$ called the graph of $R$**

When the domain and codomain are the same set, we simply say the relation is on A. It's common to use $a R b$ to mean that the pair $(a,b)$ is in the graph of $R$

> Writing the relation or operation symbol between its arguments is called _infix notation_

::: danger Notice
The definition of a _binary relation_ is exactly the same as the definition of a _function_, except that **it doesn't require the functional condition that, for each domain element, $a$, there's at most one pair in the graph whose first coordinate is $a$**

As we said, a function is a special case of a binary relation
:::

::: section Binary Relation
A binary relation $R$ is
+ a _function_ when it has the $[<=1 \text{ arrow out}]$ property
+ _surjective_ when it has the $[>= 1 \text{ arrows in}]$ property. That is, every point in the righthand, codomain column has at least one arrow pointing to it
+ _total_ when it has the $[>=1 \text{ arrows out}]$ property
+ _injective_ when it has the $[<=1 \text{ arrow in}]$ property
+ _bijective_ when it has both the $[=1 \text{ arrow out}]$ and the $[=1 \text{ arrow in}]$ property
:::
::: section Relational Images
**The image of a set $Y$, under a relation $R$, written $R(Y)$, is the set of elements of the codomain $B$ of $R$ that are related to some element in $Y$. In terms of the relation diagram, $R(Y)$ is the set of points with an arrow comming in that starts from some point in $Y$**

The inverse $R^{-1}$ of a relation $R:A \to B$ is the relation from $B$ to $A$ defined by the rule
$$b R^{-1} a \iff a R b$$

In other words, $R^{-1}$ is the relation you get by reversing the direction of the arrows in the diagram of $R$

The image of a set under the relation $R^{-1}$ is called the _inverse image_ of the set. That is, the inverse image of a set $X$ under the relation $R$ is defined to be $R^{-1}(X)$
:::
### 4.5 Finite Cardinality
A finite set is one that has only a finite number of elements. this number of elements is the size or _cardinality_ of the set

::: section Cardinality
If $A$ is a finite set, the _cardinality_ of $A$, written $|A|$, is the number of elements in A
:::

Now suppose $R: A \to B$ is a function. This means that every element of $A$ contribute at most one arrow to the diagram for $R$, so the number of arrows is at most the number of elements in $A$. That is, if $R$ is a function, then
$$|A|>=\text{\#arrows}$$

If $R$ is also surjective, then every element of $B$ has at least an arrow into it, so there must be at least as many arrows in the diagram as the size of $B$. That is
$$\text{\#arrows} >= |B|$$

Combining these inequalities implies that if $R$ is a surjective function, then
$$|A|>=|B|$$

::: section Domain and codomain size to relational properties
Let $A, B$ be (not necessarily finite) sets. Then
1. $A \text{ surj } B$ iff there is a _surjective function_ from $A$ to $B$
1. $A \text{ inj } B$ iff there is an _injective total_ relation from $A$ to $B$
1. $A \text{ bij } B$ iff there is a _bijection_ from $A$ to $B$

For _finite sets_ $A, B$
1. If $A \text{ surj } B$, then $|A| >= |B|$
1. If $A \text{ inj } B$, then $|A| <= |B|$
1. If $A \text{ bij } B$, then $|A| = |B|$
:::

::: section Mapping Rules
For _finite sets_ $A, B$
1. $A \text{ surj } B \iff |A| >= |B|$
1. $A \text{ inj } B \iff |A| <= |B|$
1. $A \text{ bij } B \iff |A| = |B|$
:::

There're $2^n$ subsets of an $n$-element set. that is
$$|A|=n \implies |pow(A)| = 2^n$$



## 5. Induction
::: tip The Induction Principle
Let $P$ be a predicate on nonnegative integers. If
+ $P(0)$ is true, and
+ $P(n) \implies P(n+1) \forall n \isin \N$

then $P(m)$ is true for all nonnegative integers $m$

Induction Rule
$$\frac{P(0), \forall n \isin \N. P(n) \implies P(n+1)}{\forall m \isin \N. P(m)}$$
:::

Even the most complicated induction proof follows exactly the same template. There're five components:
1. **State that the proof uses induction**
1. **Define an appropriate predicate $P(n)$**. The predicate $P(n)$ is called _induction hypothesis_
1. **Prove that $P(0)$ is true**. This part of the proof is called the _base case_ or _basis step_.
1. **Prove that $P(n)$ implies $P(n+1)$ for every nonnegative integer $n$**. This is called the _induction step_
1. **Invoke induction**

The induction proof is both a weakness and a strength. It's a weakness when a proof doesn't provide insight. But it's a strength that a proof can provide a reader with a reliable guarantee of correctness without requiring insight

The only change from the ordinary induction principle is that strong induction allows you make more assumptions in the inductive step of your proof. In an ordinary induction argument, you assume that $P(n)$ is true and try to prove that $P(n+1)$ is also true. In a strong induction argument, you may assume that $P(0),P(1),...$ and $P(n)$ are all true when you trying to prove $P(n+1)$. **So you can assume a stronger set of hypotheses which make your proof easier**

::: tip Principle of Strong Induction
Let $P$ be a predicate on nonnegative integers. If
+ $P(0)$ is true, and
+ $P(0) \land P(1) \land... \land P(n) \implies P(n+1) \forall n \isin \N$

then $P(m)$ is true for all nonnegative integers $m$
:::

Strong Induction Rule
$$\frac{P(0). \forall n \isin \N, P(0) \land P(1) \land... \land P(n) \implies P(n+1)}{\forall m \isin \N. P(m)}$$

::: tip State Machines
Formally a _state machine_ is nothing more than a binary relation on set, except that the elements of the set are called _states_, the relation is called the transition relation, and an arrow in the graph of the transition relation is called a _transition_. A transition from state $q$ to state $r$ will be written $q \to r$. The transition relation is also called the *state graph* of the machine. A state machine also comes equipped with a designated _start state_
:::



A property that is preserved through a series of operations or steps is known as a _preserved invariant_

An _execution_ of the state machine is a (possibly infinite) sequence of states with the property that:
+ it begins with the start state, and
+ if $q$ and $r$ are consecutive states in the sequence, then $q \to r$

A state is called _reachable_ if it appears in some execution

A _preserved invariant_ of a state machine is a predicate, $P$, on states, such that whenever $P(q)$ is true of a state $q$, and $q \to r$ for some state $r$, then $P(r)$ holds

::: tip The Invariant Principle
If a preserved invariant of a state machine is true for the start state, then it's true for all reachable states

:::
The Invariant Principle is nothing more than the Induction Principle reformulated in a convenient form for state machines. Showing that a predicate is true in the start state is the base case of the induction, and showing that a predicate is a preserved invariant corresponds to the inductive step

## 6. Recursive Data Types

Definition of recursive data types have two parts:
1. **Base case(s)** specifying that some known mathematical elements are in the data type, and
1. **Constructor case(s)** that specifies how to construct new data elements from previously constructed elements or from base elements

_Structural Induction_ is a method for proving that all the elements of a recursively defined data type have some property. A structural induction proof has two parts corresponding to the recursive definition:
+ Prove that each base case element has the property
+ Prove that each constructor case element has the property, when the constructor is applied to elements that have the property

::: tip The Principle of Structural Induction
Let $P$ be a predicate on a recursively defined data type $R$. If
+ $P(b)$ is true for each base case element, $b \isin R$, and
+ for all two-argument constructor(and likewise for all constructors taking other numbers of arguments), $c$,
$$\forall r, s \isin R, [P(r) \land P(s)] \implies P(c(r, s))$$

then
$$\forall r \isin R. P(r) \text{is true}$$

:::


## 7. Inifite Sets
An infinite set $A$ should be considered "as big as" a set $B$ when $A \small \text{ surj } B$. So we could consider $A$ to be "strictly smaller" than $B$ which we abbreviate as $A \small \text{ strict } B$, when $A$ is not "as big as" $B$
$$A \small \text{ strict } B \iff \lnot (A \small \text{ surj } B)$$

For finite sets $A, B$
$$A \small \text{ strict } B \iff |A| < |B|$$

Don't assume that surj has any particular "as big as" property on _infinite sets_ until it's been proved

For any sets $A,B,C$
1. $$A \text{ surj } B \iff B \text{ inj } A$$
2. $$A \text{ surj } B \land B \text{ surj } C \implies A \text{ surj } C$$
3. $$A \text{ bij } B \land B \text{ bij } C \implies A \text{ bij } C$$
4. $$A \text{ bij } B \iff B \text{ bij } A$$
5. $$A \text{ surj } B \land B \text{ surj } A \implies A \text { bij } B$$

Part 1 follows from the fact that $R$ has the $<= 1 \text{ out },>=1\text{ in }$ surjective function property iff $R^{-1}$ has the $>= 1 \text{ out },<=1\text{ in }$ total injective property

Part 2 follows from the fact that compositions of surjections are surjections

Part 3 and 4 follow from the first two parts because $R$ is a bijection iff $R$ and $R^{-1}$ are surjective functions

Part 5 is the Schro ̈der-Bernstein Theorem: If $A$ is at least as big as $B$ and conversel, $B$ is at least as big as $A$, then $A$ is the same size as $B$

For all sets $A,B$
$$A \text{ surj } B \text{ OR } B \text{ surj } A$$

For all sets $A,B,C$
$$A \text{ strict } B \land B \text{ strict } C \implies A \text{ strict } C$$

A basic property of finite sets that does not carry over to infinite sets is that adding something new makes a set bigger. That is, if $A$ is a finite set and $b \not \isin A$, then $|A \cup \lbrace b \rbrace| = |A| + 1$, and so $A$ and $A \cup \lbrace b \rbrace$ are not the same size. But if A is infinite, then these two sets are the same size

Let $A$ be a set and $b \not \isin A$. Then $A$ is infinite iff $A \text{ bij } A \cup \lbrace b \rbrace$

A set $C$ is _countable_ iff its elements can be listed in order, that is, the elements in $C$ are precisely the elements in the sequence
$$c_0,c_1,...,c_n,...$$
Assuming no repeats in the list, saying that $C$ can be listed in this way is formally the same as saying that the function $f: \N \to C$ defined by the rule that $f(i) \coloncolonequals  c_i$ is a bijection

A set $C$ is _countably infinite_ iff $\N \text{ bij } C$. A set is _countable_ iff it's finite or countably infite

A set $C$ is _countable_ iff $\N \text{ surj } C$. In fact a nonempty set $C$ is countable iff there's a total surjective function $g:\N \to C$

::: tip Power sets are strictly bigger
For any set $A$
$$A \text{ strict } pow(A)$$

Proof. To show that $A$ is strictly smaller than $pow(A)$, we have to show that if $g$ is a function from $A$ to $pow(A)$, then $g$ is not a surjection. To do this, we'll simply find a subset, $A_g \subseteq A$ that is not in the range of $g$. The idea is for any element $a \isin A$ to look at the set $g(a) \subseteq A$ and ask whether or not $a$ happens to be in $g(a)$

First define
$$A_g \coloncolonequals \lbrace a \isin A | a \not \isin g(a) \rbrace$$

$A_g$ is now a well-defined subset of $A$, which means it's a member of $pow(A)$. But $A_g$ can'be in the range of $g$, because if it were, we would have
$$A_g = g(a_0)$$
for some $a_0 \isin A$, so by definition of $A_g$

$$a \isin g(a_0) \iff a \isin A_g \iff a \not \isin g(a)$$
for all $a \isin A$. Now letting $a = a_0$ yields the contradiction
$$a_0 \isin g(a_0) \iff a_0 \not \isin g(a_0)$$
So $g$ is not a surjection, because there's an element in the power set of $A$, specifically the set $A_g$, that is not in the range of $g$

$pow(\N)$ is uncountable

$$pow(\N) \text{ bij } \lbrace 0,1 \rbrace ^{w}$$

$\lbrace 0,1 \rbrace ^{w}$ is uncountable
:::

+ If $U$ is an uncountable set and $A \text{ surj } U$, then $A$ is uncountable
+ if $C$ is a countable set and $C \text{ surj } A$, then $A$ is countable











::: danger Russell's Paradox
Let $S$ be a variable ranging over all sets, and define
$$W \coloncolonequals \lbrace S|S \not \isin S \rbrace$$
So by definition,
$$S \isin W \iff S \not \isin S$$
for every set $S$. In particular we can let $S$ be $W$, and obtain the contradictory result that
$$W \isin W \iff W \not \isin W$$
:::

A _formula of set theory_(Technically this's called a _first-order predicate formula_ of set theory) is a predicate formula that only uses the predicates $x=y$ and $x \isin y$

It's generally agreed that using some simple logical deduction rules, essentially all of mathematics can be derived from some formulas of set theory called the Axioms of Zermelo-Fraenkel Set Theory with Choice (ZFC)

$$(x \subseteq y) \coloncolonequals \forall z.(z \isin x \implies z \isin y)$$