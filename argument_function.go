package main

/**
 * tokenv4
 * Shows the use of ChaincodeStub API for getting the arguments sent from client
 * set-chain-env.sh -i '{"Args":["FunctionName","Arg-1", "Arg-2"]}'
 **/
import (
	"fmt"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("============Invoke Executed============")

	// V4
	// Get the args in a 2D array of byte
	argsArray := stub.GetArgs()
	fmt.Println("============GetArgs()")
	for ndx, arg := range argsArray {
		// Convert the byte[] to string
		argStr := string(arg)
		fmt.Printf("[%d]=%s  \n", ndx, argStr)
	}

	// V4
	// Get the Args[] sent by the client
	fmt.Println("==========GetStringArgs()")
	fmt.Println(stub.GetStringArgs())

	// V4
	// Get the function & parameters
	fmt.Println("=========GetFunctionAndParameters()")
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Printf("Function=%s  Args=%s\n", funcName, args)

	// V4
	fmt.Println("==============GetArgsSlice()")
	argsSlice, _ := stub.GetArgsSlice()
	length := len(argsSlice)
	fmt.Println(length)
	fmt.Println(argsSlice)
	s := string(argsSlice)
	fmt.Println(s)

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Println("Started Chaincode. token/v4")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
