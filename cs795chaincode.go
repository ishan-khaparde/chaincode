package main 

import (
	"errors"
	
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	)

type Submission struct{
	HomeworkSubmissionHash string `json:"hash"`
	StudentId string `json:"student-id"`
}
type SimpleChaincode struct{

}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string)([]byte,error){
	return nil,nil;
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string)([]byte,error){
	return nil,nil;
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string)([]byte,error){
	
	if function =="postSubmittedHomework"{
		result,err := postSubmittedHomework(stub,args)

		if err != nil{
			fmt.Println("Error posting homework")
			return nil,err
		}
		fmt.Println("Homework posted to blockchain!")
		return result,nil
	}
	return nil,nil;
}

func postSubmittedHomework(stub shim.ChaincodeStubInterface, args []string)([]byte,error){

	if len(args) < 2{
		fmt.Println("Error! Illegal number of arguments recevied.")
		return nil , errors.New("Missing homework file hash or public key of candidate.")
	}
	 txHash := args[0];
	 student_id := args[1];

	var homework Submission = Submission{txHash,student_id};
	bytes,err := json.Marshal(&homework);
	if err!=nil{
		fmt.Println("Error in doing JSON marshal of homework data.")
		return nil,err;
	}

	err = stub.PutState(student_id,bytes)
	return bytes,err; 
}
func main(){

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Could not start SimpleChaincode")
	} else {
		fmt.Println("SimpleChaincode successfully started")
	}

}