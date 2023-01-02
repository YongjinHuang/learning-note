# [InnoDB Lock](https://jahfer.com/posts/innodb-locks/)

## Introduction

InnoDB only has a handful of locking concepts, but their use and behavior depnd greatly on the [transaction isolation level](https://dev.mysql.com/doc/refman/5.7/en/innodb-transaction-isolation-levels.html) that is active for the connection

::: tip Transaction Isolation Level
The isolation level is the setting that fine-tunes the balance between performance and reliablity, consistency, and reproducibility of results when multiple transactions are making changes and performing queries at the same time
:::

There're four transaction isolation levels for InnoDB (in order of most-to-least strict):
+ SERIALIZABLE
+ REPEATABLE READ (default)
+ READ COMMITTED
+ READ UNCOMMITTED

```sql
CREATE TABLE `books` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author_id` bigint(20) NOT NULL,
  `title` varchar(255) NOT NULL,
  `borrowed` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_books_on_author_id` (`author_id`)
);

INSERT INTO `books` (`author_id`, `title`)
VALUES
  (101, "The Pragmatic Programmer"),
  (102, "Clean Code"),
  (102, "The Clean Coder"),
  (104, "Ruby Under a Microscope");
```

| id | author_id | title | borrowed |
| -- | -- | -- | -- |
| 1 | 101 | The Pragmatic Programmer | False |
| 2 | 102 | Clean Code | False |
| 3 | 103 | The Clean Coder | False |
| 4 | 104 | Ruby Under a Microscope | False |


## InnoDB Lock Types

### Shared or Exclusive Locks

InnoDB locks can either be [shared or exclusive](https://dev.mysql.com/doc/refman/5.7/en/innodb-locking.html#innodb-shared-exclusive-locks)


::: tip Shared or Exclusive
A shared(S) lock permits the transaction that holds the lock to read a row

An exclusive(X) lock permits the transaction that holds the lock to update or delete a row

If transaction *T1* holds an S lock on row *r*, then the requests from some distinct transaction *T2* for a lock on row *r* are handled as follows:
+ A request by *T2* for an S lock can be granted immediately. As a result, both *T1* and *T2* hold an S lock on *r*
+ A request by *T2* for an X lock can't be granted immediately

> This's how InnoDB can guarantee the results of a read performed under a S lock won't change while being read


If a transaction *T1* holds an X lock on row *r*, a request for some distinct transaction *T2* for a lock of either type on *r* can't be granted immediately. Instead, transaction *T2* has to wait for transaction *T1* to release its lock on row *r*

:::

S locks will only lock rows in the index used to perform the read. This is because the read needs to be reproducible, but only in the context of how the records were found

On the opposite end, X locks are taken against the primary key of the record(s) which may have a larger impact since all secondary indexes are affected

### Intention Locks
InnoDB supports *multiple granularity locking* which permits coexistence of row locks and table locks. To make locking at multiple granularity levels practible, InnoDB uses intention locks. Intention locks are table-level locks that indicate which type of lock (shared or exclusive) a transaction requires later for a row in a table. There're two types of intention locks:
+ An intention shared lock (IS) indicates that a transaction intends to set a shared lock on individual rows in a table
+ An intention exclusive lock (IX) indicates that a transaction intends to set an exclusive lock on individual rows in a table

For example, `SELECT ... LOCK IN SHARE MODE` sets an IS lock, and `SELECT ... FOR UPDATE` sets an IX lock. The main purpose of intention locks is to show that someone is locking a row, or going to lock a row in a table

The intention locking protocol is as follows:
+ Before a transaction can accquire an S lock on a row in a table, it must first acquire an IS lock or stronger on the table 
+ Before a transaction can accquire an X lock on a row in a table, it must accquire an IX lock on the table

A lock is granted to a requesting transaction if it's compatible(✅) with existing locks, but not if it conflicts(❌) with existing locks. If a lock request conflicts with an existing lock and can't be granted because it would cause deadlock, an error occurs

|  | X | IX | S | IS |
| -- | -- | -- | -- | -- |
| X | ❌ | ❌ | ❌ | ❌ |
| IX | ❌ | ✅ | ❌ | ✅ |
| S | ❌ | ❌ | ✅ | ✅ |
| IS | ❌ | ✅ | ✅ | ✅ |


### Record Locks
A record lock is a lock on an index record. Record locks always lock index records, even if a table is defined with no indexes. For such cases, InnoDB creates a hidden clustered index and uses this index for record locking
<video width="520" height="160" controls>
  <source src="./mp4/record-lock-basic.mp4" type="video/mp4">
</video>

### Gap Locks
A gap lock is a lock on a gap between index records, or a lock on the gap before the first or after the last index record. InnoDB uses this lock type to ensure a set of selected records and the surrounding records maintain their relationship. A gap might span a signle index value, multiple index values, or even be empty

**If a lock is held on a gap, not other statement is allowed to INSERT, UPDATE or DELETE a record that falls within that gap**

A gap lock may be taken on any index for the table, including the table's clustered index

<video width="520" height="160" controls>
  <source src="./mp4/gap-lock-basic.mp4" type="video/mp4">
</video>

> **It's worth noting here that conflicting locks can be held on a gap by different transactions.** For example, transaction A can hold a shared gap lock on a gap while transaction B holds an exclusive gap lock on the same gap. The reason conflicting gap locks are allowed is that if a record is purged from an index, the gap locks held on the record by different transactions must be merged 

### Next-key Locks
A next-key lock is a combination of a record lock on the index record and a gap lock on the gap before the index record. Thus the nex-key locks enables you to lock the nonexistence of something in your table

> Note: When InnoDB scans an index, it can also lock the gap after the last record in the index

By default InnoDB operates in REPEATABLE READ transaction isolation level. In this case, InnoDB uses next-key locks for searches and index scans, which prevents [phantom rows](https://dev.mysql.com/doc/refman/5.7/en/innodb-next-key-locking.html)

::: tip Phantom Rows
The so-called phantom problem occurs within a transaction when the same query produces different sets of rows at different times. For example, if a SELECT is executed twice, but returns a row the second time that was not returned the first time, the row is a phantom row
:::

<video width="520" height="160" controls>
  <source src="./mp4/multi-lock.mp4" type="video/mp4">
</video>

## Transaction Isolation Level: SERIALIZABLE
Every record read by the query must be locked to ensure no other connection can modify the records being viewed for statements performed within a SERIALIZABLE transaction isolation level. This's the strictest isolation level and its locking ensure that concurrent transactions can be reordered safely without impacting one another. The search sets shared next-key locks on the index records it encounters

If the SELECT query is using a unique index (such as the PRIMARY key), [it doesn't need to use a gap lock since it can guarantee it is only affecting unique records](https://dev.mysql.com/doc/refman/5.7/en/innodb-locks-set.html)


## Transaction Isolation Level: REPEATABLE READ
The default transaction isolation level InnoDB uses is REPEATABLE READ. A transaction using REPEATABLE READ will perform reads as if they were run at the same point in time(snapshotted) as the very first read in the transaction. This allows a consistent view of the database across queries without running into phantom rows: records that appear or disappear on subsequent reads

### Locks on Primary Keys & Unique Indexes
The simplest case is when a locking query selects records using the table's primary key

Under REPEATABLE READ InnoDB ensures that the update is reproducible: It will always affect the same records because it blocks any other statements from impacting the results

When using a primary key or a unique index, InnoDB doesn't need to take a gap lock on the left or right of the selected records since it is sure that the values matched by the query are distinct from all others


### Locks on Non-Unique Indexes
If the statement is selecting records that can not be guaranteed as unique, it must use a gap lock to ensure the read query performed can be repeated while returning the same result (hence REPEATABLE READ). This's where the semantics of the next-key locking come into play

> If one session has a shared or exclusive lock on record R in an index, another session can't insert a new index record in the gap immediately before R in the index order

::: warning A particularly surprising case
When accquiring a next-key lock on the first or last record in an index and the search doesn't meet the criteria for uniqueness, InnoDB must lock all values towards positive or negative infinity to make sure no other record overlaps this record's position in the table

This can be a particularly nasty problem since it's not uncommon for insertions happen at the tail end of an index
:::
