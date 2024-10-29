package persistence

import "github.com/google/uuid"

//Tadis is Inherently in-memory but f it is to be truly useful,
//it has to have a mechanism to move and push objects into persistent storage
//this attempt tries to achieve this using Write ahead Logs.
//
//Write Ahead Logs are used in some databases (like postgres) to manage states across
//distributed database systems and In Recovery.


//schema is a generic type
//at invocation the Fields to be used allocated and set
//to ensure robustness, i would support struct embedding
//ie you can define a small schema and place that schema in a larger schema
//pros
//relatively easy to implement
//

type Schema struct {
	id uuid.UUID
	//string fields
	Field1 string
	Field2 string
	Field3 string
	//field 4 is skipped
	Field5 int
	Field6 int
	//field 7 is skipped
	Field7 []string
	Field8 []string
	//field 9 is skipped
	Field10 map[string]string
}


//A Transaction is a safe way to ensure a process is run to completion

//Begin starts a transaction
func Begin() {
	
}

func Commit() {

}
