/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the Healchain structure

type Heal struct {
	Name string `json:"name"`
	Bgroup string `json:"bgroup"`
	HP string `json:"hp"`
	SP string `json:"sp"`
	DP string `json:"dp"`
	OG string `json:"og"`
	Allergic string `json:"allergic"`
	Generalinfo string `json:"generalinfo"`
	MedicalHistory string `json:"medicalhistory"`
	MedicalHistoryview string `json:"medicalhistoryview"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createAcc" {
		return s.createAcc(APIstub, args)
	} else if function == "getProfile" {
		return s.getProfile(APIstub, args)
	} else if function == "editProfile" {
		return s.editProfile(APIstub, args)
	}

	
	
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	
	h := Heal{Name: "Charmander", Bgroup: "AB+"}
	hAsBytes, _ := json.Marshal(h)
	APIstub.PutState("0x31028", hAsBytes)
	
	return shim.Success(nil)
}

//Create new account for users
func (s *SmartContract) createAcc(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var heal = Heal{Name: args[1], Bgroup: args[2], HP: args[3], SP: args[4],OG: args[5],DP: args[6], MedicalHistory: args[7], MedicalHistoryview: args[8], Allergic: args[9], Generalinfo: args[10]}
	healerAsBytes, _ := json.Marshal(heal)
	APIstub.PutState(args[0], healerAsBytes)

	return shim.Success(nil)
}


//Fetch profile data
func (s *SmartContract) getProfile(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	healerAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(healerAsBytes)
}

//Edit existing profile
func (s *SmartContract) editProfile(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments.")
	}
	
	fetchprofile, _ := APIstub.GetState(args[0])
	heal := Heal{}
	json.Unmarshal(fetchprofile, &heal)
	heal.Name = args[1]
	heal.Bgroup = args[2]
	heal.HP = args[3]
	heal.SP = args[4]
	heal.OG = args[5]
	heal.DP = args[6]
	heal.MedicalHistory = args[7]
	heal.MedicalHistoryview = args[8]
	heal.Allergic = args[9]
	heal.Generalinfo = args[10]
	
	fetchprofile, _ = json.Marshal(heal)
	APIstub.PutState(args[0], fetchprofile)

	return shim.Success(nil)
}


// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
