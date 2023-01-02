# Recurrences

A _recurrence_ describes a sequence of numbers. Early terms are specified explicitly, and later terms are expressed as a function of their predecessors

As a trivial example, here is a recurrence describing the sequence 1,2,3,...

$T_1=1$
$T_n=T_{n-1}+1$

Here the first term is defined to be 1 and each subsequent term is one more than its predecessors

$T_n=2^n-1$ satifies the recurrence:

$T_1=1$
$T_n=2T_{n-1}+1 (\forall n >=2)$

The maximum number of comparisons needed to **Merge Sort** n items is given by this recurrence:

$T_1=0$
$T_n=2T_{n/2}+n-1$ for $n>=2$ and a power of 2

A _**closed-form expression**_ would be much more helpful

::: tip Solving the recurrence -- Plug and Chug
#### Step 1: Plug and Chug Until a Pattern Appears
First, we expand the recurrence equation by alternately plugging and chugging until a pattern appears
$$
\begin{gather*}
\begin{split}
T_n &=2T_{n/2}+n-1\\
&=2(2T_{\frac{n}{4}}+n/2-1)+(n-1)\\
&=4T_{\frac{n}{4}}+(n-2)+(n-1)\\
&=4(2T_{\frac{n}{8}}+n-1)+(n-2)+(n-1)\\
&=8T_{\frac{n}{8}}+(n-4)+(n-2)+(n-1)
\end{split}
\end{gather*}
$$
A pattern is emerging. In particular, the formula seems holds:
$$
\begin{gather*}
\begin{split}
T_n &=2^kT_{\frac{n}{2^k}}+(n-2^{k-1})+(n-2^{k-2})+...+(n-2^{0})\\
&=2^kT_{\frac{n}{2^k}}+kn-(2^k-1)\\
&=2^kT_{\frac{n}{2^k}}+kn-2^k+1
\end{split}
\end{gather*}
$$

#### Step 2: Verify the Pattern
Next, we verify the pattern with one additional round of plug-and-chug. If we guessed the wrong pattern, then this is where we’ll discover the mistake

$$
\begin{equation*}
\begin{split}
T_n &=2^kT_{\frac{n}{2^k}}+kn-2^k+1\\
&=2^k(2T_{\frac{n}{2^{k+1}}}+\frac{n}{2^k}-1)+kn-2^k+1\\
&=2^{k+1}T_{\frac{n}{2^{k+1}}}+(k+1)n-2^{k+1}+1
\end{split}
\end{equation*}
$$
The formula is unchanged except that k is replaced by $k+1$. This amounts to the inducton step in a proof that the formua holds for all $k>=1$

#### Step 3: Write $T_n$ Using Early Terms with Known Values
Finally we express $T_n$ using early terms whose values are known. Specifically, if we let $k=\log{n}$, then $T_{\frac{n}{2^k}}=T_1$, which we know is 0:
$$
\begin{equation*}
\begin{split}
T_n &=2^kT_{\frac{n}{2^k}}+kn-2^k+1\\
&=2^{\log{n}}T_{\frac{n}{2^{\log{n}}}}+n\log{n}-2^{\log{n}}+1\\
&=nT_1+n\log{n}-n+1\\
&=n\log{n}-n+1
\end{split}
\end{equation*}
$$
:::

$T_n=n\log{n}-n+1$
