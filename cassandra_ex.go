// Example use of Cassandra from: https://github.com/gocql/gocql

/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(tml mesg, id UUID, mesg mesg, PRIMARY KEY(id));
create index on example.tweet(tml);
*/
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	// cluster := gocql.NewCluster("104.131.182.5")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 1
	cluster.Keyspace = "mykeyspace"
	cluster.Consistency = gocql.Quorum
	log.Println("Create")
	session, _ := cluster.CreateSession()
	log.Println("Created")
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (tml, user_id, mesg) VALUES (?, ?, ?)`,
		"me", 1, "hello world").Exec(); err != nil {
		log.Fatal(err)
	}
	// if err := session.Query(`INSERT INTO tweet (tml, mesg) VALUES (?, ?)`,
	// 	"me", "hello world").Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Inserted")

	// var id gocql.UUID
	var mesg string

	/* Search for a specific set of records whose 'tml' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	if err := session.Query(`SELECT mesg FROM tweet WHERE user_id = ? LIMIT 1`,
		1).Consistency(gocql.One).Scan(&mesg); err != nil {
		log.Println(err)
	}
	fmt.Println("Tweet:", mesg)

	// list all tweets
	iter := session.Query(`SELECT mesg FROM tweet WHERE user_id = ?`, 1).Iter()
	for iter.Scan(&mesg) {
		fmt.Println("Tweet:", mesg)
	}
	if err := iter.Close(); err != nil {
		log.Println(err)
	}
}
