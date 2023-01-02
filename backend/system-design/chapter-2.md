# chapter-2

## Database and DBMS

### Components

Here're some common components found across different databases

#### Schema
The role of a schema is to define the shape of a data structure, and specify what kinds of data can go where

#### Table
Each table contains various columns just like in a spreadsheet

#### Column
A column contains a set of data values of a particular type, one value for each row of the database

### Types
![Img](./FILES/chapter-2.md/5582a41b.png)

Below are different types of database
+ [SQL](https://www.karanpratapsingh.com/courses/system-design/sql-databases): A collection of data items with pre-defined relationships between them
+ [NoSQL](https://www.karanpratapsingh.com/courses/system-design/nosql-databases): A broad category that includes any database that doesn't use SQL as its primary data access language
  + Document
  + Key-value
  + Graph
  + Timeseries
  + Wide column
  + Multi-model


### Challenges
Some common challenges faced while running databases at scale
+ **Absorbing significant increases in data volume**

## SQL databases

A SQL(or relational) database is a collection of data items with pre-defined relationships between them. Thease items are organized as a set of tables with columns and rows

Each row in a table could be marked with a unique identifier called a primary key. SQL databases usually follow the ACID consistency model


## Database Replication

Replication is a process that involves sharing information to ensure consistency between redundant resources such as multiple databases, to improve reliability, fault-tolerance, or accessibility

### Master-Slave Replication

The master serves reads and writes, replicating writes to one or more slaves, which serve only reads. Slaves can also replicate additional slaves in a tree-like fashion. If the master goes offline, the system can continue to operate in read-only mode until a slave is promoted to a master or a new master is provisioned

![Img](./FILES/chapter-2.md/4c7b0a6a.png)

#### Advantages
+ Backups of the entire database of relatively no impact on the master
+ Applications can read from the slave(s) without impacting the master
+ Slaves can be taken offline and synced back to the master without any downtime

#### Disadvantages
+ Replication adds more hardware and additional complexity
+ Downtime and possibly loss of data when a master fails
+ All writes also have to be made to the master in a master-slave architecture
+ The more read slaves, the more we have to replicate, which will increase replication log

### Master-Master Replication
Both masters serve reads/writes and coordinate with each other. If either master goes down, the system can continue to operate with both reads and writes

![Img](https://raw.githubusercontent.com/karanpratapsingh/portfolio/master/public/static/courses/system-design/chapter-II/database-replication/master-master-replication.png)

#### Advantages
+ Applications can read from both masters
+ Distributes write load across both master ndoes
+ Simple, automatic, and quick failover

#### Disadvantages
+ Not as simple as master-slave to configure and deploy
+ Either loosely consistent or have increased write latency due to synchronization
+ Conflict resolution comes into play as more write nodes are added and as latency increases

### Synchronous vs Asynchronous replication

The primary difference between synchronous and asynchronous replication is how the data is written to the replica. In synchronous replication, data is written to primary storage and the replica simultaneously. As such, the primary copy and the replica should always remain synchronized

In contrast, asynchronous replication copies the data to the replica after the data is already written to the primary storage. Although the replication process may occur in near-real-time, it's more common for replication to occur on a scheduled basis and it's more cost-effective