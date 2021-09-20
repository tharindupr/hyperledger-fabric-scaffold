package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	//"reflect"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"
	//"github.com/hyperledger/fabric/common/util"
	//"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

type Parameters struct{
	Gamma float64 `json:Gamma`
}


type Prediction struct{
	ModelID string `json:ModelID`
	NodeID string `json:NodeID`
	Timestamp string `json:Timestamp`
	OutPut bool `json:OutPut`

}

type Model struct{
	ModelID string `json:ModelID`
	ModelDescription string `json:ModelDescription`
	NodeID string `json:NodeID`
	MalciousPrecision float64 `json:MalciousPrecision`
	MalciousRecall float64 `json:MalciousRecall`
	BenignPrecision float64 `json:BenignPrecision`
	BenignRecall float64 `json:BenignRecall`
	Hash string `json:Hash`
	PositiveTrustScore float64 `json:PositiveTrustScore`
	NegativeTrustScore float64 `json:NegativeTrustScore`
}


var logger = flogging.MustGetLogger("models_cc")


// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	logger.Infof("Chiancode : datacollection initiated")

	var parameters Parameters
	parameters.Beta = 3
	parameters.AlphaPositive = 10
	parameters.AlphaNegative = -10

	parametersAsBytes, _ := json.Marshal(parameters)
	APIstub.PutState("parameters", parametersAsBytes)
	return shim.Success(nil)
}


// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))
	logger.Errorf("Function name is:  %d", function)
	switch function {
		case "createModel":
			return s.createModel(APIstub, args)
		case "traceModel":
			return s.traceModel(APIstub, args)
		case "getModel":
			return s.getModel(APIstub, args)
	
	}
	return shim.Error("Invoke Function Not Success.")
}


//add a new detection
func (s *SmartContract) reportPrediction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var prediction Prediction
	json.Unmarshal([]byte(args[0]), &prediction)


	//getting the model trust parameters from Chain
	parametersAsBytes, _ := APIstub.GetState("parameters")
	if parametersAsBytes == nil {
		return shim.Error("Error in retrieving model trust parameters")
	}
	parameters := Parameters{}
	json.Unmarshal(parametersAsBytes, &parameters)


	var PostiveWeight = 2 * (model.BenignPrecision * model.BenignRecall)/(model.BenignPrecision + model.BenignRecall)
	var NegativeWeight = 2 * (model.MalciousPrecision * model.MalciousRecall)/(model.MalciousPrecision + model.MalciousRecall)


	model.PositiveTrustScore = PostiveWeight * parameters.AlphaPositive
	model.NegativeTrustScore = NegativeWeight * parameters.AlphaNegative

	modeldAsBytes, _ := json.Marshal(model)
	APIstub.PutState(model.ModelID, modeldAsBytes)

	
	logger.Infof("Model Successfully Added")
	return shim.Success(modeldAsBytes)

}

//retrieve a new model 
func (s *SmartContract) getModel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	tAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(tAsBytes)
}

//update Paramters
func (s *SmartContract) updateParameters(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	var parameters Parameters
	json.Unmarshal([]byte(args[0]), &parameters)


	parametersAsBytes, _ := json.Marshal(parameters)
	APIstub.PutState("parameters", parametersAsBytes)

	
	logger.Infof("Patameters updated successfully")
	return shim.Success(parametersAsBytes)

}


func (t *SmartContract) traceModel(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ID := args[0]
	logger.Infof("searching for the Perf Contract with id %s", args[0])
	resultsIterator, err := stub.GetHistoryForKey(ID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()

		logger.Infof(string(response.TxId))
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	logger.Infof("- getHistoryforDEC returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}




// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}


