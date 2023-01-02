# System Deisng Interview

System desin is the process of defining the architecture, interfaces, and data for a system that satisfies specific requirements

System design helps us define a solution that meets the business requirements. It's one of the earliest decisions we can make when building a system. Often it's essential to think from a high level as these decisions are very difficult to correct later. It also makes it easier to reason about and manage architectural changes as the system evolves

System design is a very extensive topic and system design interviews are designed to evaluate your capability to produce technical solutions to abstract problems, as such, they're not designed for a specific answer


## Requirements clarifications

Usually, requirements are divided into 3 parts
1. Functional
1. Non-functional
1. Extended

### Functional requirements
Requirements that the end user specifically demands as basic functionalities that the system should offer
+ What are the features that we need to design for this system?
+ What are the edge cases we need to consider, if any, in our design?

### Non-functional requirements
The quality constraints that the system must satisfy according to the project contract. They're also called Non-behavioral requirements
+ portability
+ maintainability
+ reliability
+ scalability
+ security

### Extended requirements
Basically "nice to have" requirements that might be out of the scope of the system
+ Our system should record metrics and analytics
+ Service health and performance monitoring

## Estimation and Constraints
Estimate the scale of the system we are going to design
+ What's the desired scale that this system will need to handle
+ What's the read/write ratio of our system
+ How many requests per second
+ How much storage will be needed

## Data model design
In this step, we basically define all the entities and relationships between them
+ What are the different entities in this system
+ What are the relationships between these entities
+ How many tables do we need
+ Is NoSQL a better choice here


## API design
APIs will help us define the expectations from the system explicitly. We don't have to write any code, just a simple interface defining the API requirements such as parameters, functions, classes, types, entities, etc
```ts
createUser(name: string, email: string): User
```

It's advised to keep the interface as simple as possible and come back to it later when convering extended requirements

## High-level component design
It's time to identify system components that are needed to solve our problem and draft the first design of our system
+ Is it best to design a monolithic or a microservices architecture
+ What type of a database should we use

Once we have a basic diagram, we can start discussing how the system will work from the client's perspective

## Detailed design
Now it's time to go into detail about the major components of the system we designed

+ How should we partition our data
+ What about load distribution
+ Should we use cache
+ How will we handle a sudden spike in traffic

Try not to be opinionated about certain technologies

## Identify and resolve bottlenecks

Finally it's time to discuss bottlenecks and approaches to mitigate them
+ Do we have enough database replicas
+ Is there any single point of failure
+ Is database sharding required
+ How can we make our system more robust
+ How to improve the availability of our cache

Make sure to read the engineering blog of the company you're interviewing with. This will help you get a sense of what technology stack they're using and which problems are important to them

