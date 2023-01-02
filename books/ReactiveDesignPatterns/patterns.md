# patterns

## Fault tolerance and recovery patterns
### Simple Component
::: tip Simple Component
A component shall do only one thing, but do it in full
:::
This pattern applies wherever a system performs multiple functions or the functions it performs are so complex that they need to be broken into different components. Applied to object-oriented software design, it's usually stated as follows: A class should have only one reason to change

[Drawio](./drawio/simple-pattern.drawio){link-type="drawio"}

The goal of the Simple Component pattern is to implement the single responsibility principle

### Error Kernel

::: tip Error Kernel
In a supervision hierachy, keep important application state or functionality near the root while delegating risky operations towards the leaves
:::

This pattern builds on the Simple Component pattern and is applicable wherever components with different failure probability and reliability requirements are combined into a larger system or application -- some functions of the system must never go down, whereas others are necessarily exposed to failure. Applying the Simple Component pattern will frequently leave you in this position, so it pays to familiarize yourself well with the Error Kernel pattern.

[Drawio](./drawio/error-kernel.drawio){link-type="drawio"}

### Let-it-crash

::: tip Let-it-crash
Prefer a full component restart to internal failure handling
:::
In chapter 7 of the book we discuss principled failure handling, noting that the internal recovery mechanisms of each component are limited because they are not sufficiently separated from the failing parts -- every thing within a component can be affected by a failure. This's especially true for hardware failures that take down the component as a whole, but it's also true for corrupted state taht is the result of some programming error only observable in rare circumstances. For this reason, **it is necessary to delegate failure handling to a supervisor instead of attempting to solve it within the component**

[Drawio](./drawio/let-it-crash.drawio){link-type="drawio"}

This principle is also called crash-only software: the idea is that transient but rare failures are often costly to diagnose and fix, making it preferable to recover a working system by rebooting parts of it. This hierarchical restart-based failure handling makes it possible to greatly simplify the failure model and at the same time leads to a more robust system that even has a chance to survive failures that were entirely unforeseen

### Circuit Breaker
::: tip Circuit Breaker
Protect services by breaking the connection to their users during prolonged failure conditions
:::
This pattern describes how to safely connect different parts of the system so that failures do not spread uncontrollably across them. Its origin lies in electrical engineering: in order to protect electrical circuits from each other and introduce decoupled failure domains, a technique was established of breaking the connection when the transmitted power exceeds a given threshold

Translated to a Reactive application, this means the flow of requests from one component to the next may be broken up deliberately when the recipient is overloaded or otherwise failing. Doing so serves two purposes:
1. The recipient gets some breathing room to recover from possible load-induced failures
1. The sender decides that requests will fail instead of waisting time with waiting for negative replies

## Replication patterns

### Active-passive replication
::: tip Active-passive
Content
:::


## Resource management patterns
### Resource encapsulation
::: tip Resource encapsulation
A resource and its lifecycle are responsibilities that must be owned by one component
:::

### Resource loan
::: tip Resource loan
Give a client exclusive transient access to scarce resource without transfering ownership
:::

### Complex command
::: tip Complex command
Send compound instructions to the resource to avoid excessive network usage
:::

### Resource pool
::: tip Resource pool
Hide an elastic pool of resources behind their owner
:::

### Managed blocking
::: tip Managed blocking
Blocking a resource requires consideration and ownership
:::


## Message flow patterns

### Request-response

### Self-contained message

### Ask

### Forward Flow

### Aggregator

### Saga

### Business Handshake

## Flow control patterns
### Pull
::: tip Pull
Have the consumer ask the producer for batches of data
:::
### Managed queue
::: tip Managed queue
Manage an explicit input queue, and react to its fill level
:::

### Drop
::: tip Drop
Dropping requests is preferable to failing uncontrollably
:::
Imagine a system that is exposed to uncontrollably user input. Any deployment will be finite in both its processing capability and its buffer capability, and if user input exceeds the former for long enough, the latter will be used up and something will need to fail. If this's not foreseen explicitly, then an automatic out-of-memory killer will make a decision that is likely to be less satisfactory than a planned load-shedding mechanisms -- and shedding load means dropping requests

This's more of a philosophy than an implementation pattern. Network protocols, operating system, programming platforms, and libraries will all drop packets, messages or requests when overloaded; they do so in order to protect the system so that it can recover when load decreases. In the same spirit, authors of Reactive systems need to be comfortable with the notion of sometimes deliberately losing messages


### Throttling
::: tip Throttling
Throttle your own output rate according to contracts with other services
:::



## State management and persistence patterns

### Domain object
::: tip Domain object
Separate the business domain logic from communication and state management
:::
The Domain Object pattern describes how to matain a clear boundary and separation between the different concerns of business logic, state management and 


### Sharding

::: tip Sharding
Scale out the management of a large number of domain objects by grouping them into shards based on unique and stable object properties
:::
The Sharding pattern places an upper bound on the size of the directory by grouping the domain objects into a configurable number of shards -- the domain is fractured algorithmically into pieces of manageable size. The term algorithmically means the association between objects and shards is determined by a fixed formula that can be evaluated whenever an object needs to be located

### Event sourcing

::: tip Event sourcing
Perform state changes only by applying events. Make them durable by storing the events in a log
:::



### Event stream

::: tip Event stream
Publish the events emitted by a component so that the rest of the system can derive knowledge from them
:::
The events that a component stores in its log represent the sum of all the knowledge it has ever possessed. This's a treasure trove for the rest of the system to delve into