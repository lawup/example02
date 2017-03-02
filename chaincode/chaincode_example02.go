/*
Hyperledger Hackathon 2017 @ShangHai
Team ：HyperTerminator
Works: Consensus of Justice A.I. 
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	err1 := stub.PutState("hello_world1", []byte(args[0]))
	if err1 != nil {
		return nil, err1
	}

	err2 := stub.PutState("hello_world2", []byte(args[1]))
	if err2 != nil {
		return nil, err2
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key1, value1, key2, value2 string
	var err1, err2 error
	fmt.Println("running write()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4. name of the key and value to set")
	}

	key1 = args[0] //rename for funsies
	value1 = args[1]
	key2 = args[2] //rename for funsies
	value2 = args[3]
	
	err1 = stub.PutState(key1, []byte(value1)) //write the variable into the chaincode state
	if err1 != nil {
		return nil, err1
	}
	err2 = stub.PutState(key2, []byte(value2)) //write the variable into the chaincode state
	if err2 != nil {
		return nil, err2
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key1,key2,jsonResp1,jsonResp2 string
	var err1,err2 error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key1 = args[0]
	valAsbytes1, err1 := stub.GetState(key1)
	if err1 != nil {
		jsonResp1 = "{\"Error\":\"Failed to get state for " + key1  +" \"}"
		return nil, errors.New(jsonResp1)
	}

	key2 = args[0]
	valAsbytes2, err2 := stub.GetState(key2)
	if err2 != nil {
		jsonResp2 = "{\"Error\":\"Failed to get state for " + key2  +" \"}"
		return nil, errors.New(jsonResp2)
	}
	
	valAsbytes1 += valAsbytes2
	
	return valAsbytes, nil
}

