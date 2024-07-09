## 1. Introduction
While we are preparing this documentation, there is 21,208,535 blocks in the Cosmos blockchain.Full nodes does not download all the blocks from genesis, so they dont have the historical data. To get a historical data, users need to get the data from "Archive Nodes", but only thing you can do is trusting them.
With "projectname", we remove the need of trust, providing much better decentrilazation. "Dont trust, verify! "


## 2. Problem We Solve
Amount of Archive nodes on Cosmos is not that much. If your app needs historical data of blockchain, you either can set up your own Archive Node(which is not a good way since storage requirements of that machines are very high, hence they are very expensive) or use the RPC of an existing one. But in that case, you can not verify if the data you get is correct. you need to "trust" the node.

With integrating ZKProof mechanism, Every node will keep the proof of the block among all the other data. To be sure if a data is valid, you can simply verify the proof. We implemented that ZKProof mechanism with O1Js. Check the "Tech Stack" part to read more about it.

## 3. Why o1js?
We chose o1js for the Zero-knowledge proof generation and verification functionality because o1js is one of the few projects that can truly aggregate ZKProofs. o1js is perfect fit because:
##### Infinite Recursion
Using recursion technology of O1js, Only having the proof of the last block is enough because if you verify it passing the genesis block of the chain as the public input, you can be sure all the blocks in the middle are correct.

##### Constant size
No matter what is the amount of blocks you create proof, it has always constant size. Recursing all the proofs of blocks does not affect the size of the final proof.

##### Constant Verification Time
No matter what is the amount of blocks you create proof, it always require constant (and relativly low) time to verify it. 

## 4. Architecture
![Architecture ](https://i.imgur.com/xIOp4OG.png)
#### Cosmos SDK

The Cosmos SDK is the underlying framework used to build the blockchain application. It handles block generation, verification, and transaction verification. The zkproof module is integrated into this SDK to enhance its functionality with zero-knowledge proofs.

#### zkproof Module

The zkproof module is a dedicated component within the Cosmos SDK responsible for managing zero-knowledge proofs. It interacts with other parts of the SDK to ensure the blockchain's data integrity.

#### ProofMechanism

The ProofMechanism is the core component of the zkproof module. It is actually a node.js project that has ZkProgram written with o1js. Creating and verifying the proof process is executing in here, returning the result of verifying or new ZKP to the module. Functions that calls ZkProgram methods are accessible via CLI Scripts, that's how call these inside of the module that is written with Go.  
It consists of two main functions:

-   *Generate Genesis Proof*: This function takes block hash as public input, generates a new zero-knowledge proof for genesis block.
-   *Generate New Proof*: This function takes recent block hash as public input,  generates a new zero-knowledge proof for each block by combining the block's hash with the last generated proof.
-   *Verify Proof*: This function verifies the correctness of the zero-knowledge proof received from the state before it is used for new blocks.

## 5. Code Review
In this chapter, we will break down and demonstrate our project repo

### x/cosmicproof

#### Keeper

-   **Purpose**: The `keeper` package  contains the logic for creating and verifying zeroknowledge proofs and interacting with the module's state. It is responsible for reading from and writing to the store.
-   **Components**:
    -   **Keeper Struct**: This struct holds references to other necessary parts of the application (e.g., store keys, codec, other module keepers).
    -   **Methods**: 
    - 
	    **1. Generatecosmicproof**
		This function executes the js code which is written with "o1.js",it creates zeroknowledge proof by using latest proof and current blockhash as inputs then returns a zeroknowledge proof for this block.
		
		**2. Verifycosmicproof**
		This function executes the js code which is written with "o1.js",it verifys zeroknowledge proof by using proof and genesis blockhash as inputs then returns a boolean value for given proof.

		**3. Savecosmicproof**
			This function saves a proof associated with a specific block height into a prefixed key-value store within the Cosmos SDK's context.
			**4. Getcosmicproof**
			This function gets the proof of specific block height.
			
#### Module

-   **Purpose**: The `module` package contains the module's basic setup and configuration. It defines how the module is initialized, how it handles messages, and how it integrates with the rest of the Cosmos SDK application.
-   **Components**:
    -   **AppModuleBasic**: Implements basic non-dependent module functionalities like registration of codec.
    -   **AppModule**: Implements dependent module functionalities like genesis handling, message and query routing, and module manager integration.
#### Proofmechanism
#### **/src**
1. **cli**
This file contains command line management system to create zeroknowledge proofs by executing javascript code in go.
2. **GenerateBlockProof**
This file contains the javascript code for generating and verifying zeroknowledge proofs by using "o1.js".
1. **GenerateBlockProof.test**
This file contains tests for generating and verifying zeroknowledge proofs by using "o1.js".

#### Types

-   **Purpose**: The `types` package contains the types and constants used throughout the module. It serves as a central place for defining the module's core data structures, message types, and error handling.
-   **Components**:
    -   **Messages**: Define the messages that the module can handle, such as MsgCreateProof or MsgVerifyProof.
    -   **State Types**: Define the data structures that will be stored in the module's state such as
	  type  cosmicproof  struct {
			Proof  string  `json:"proof"`
		}
    -   **Errors**: Define custom error types for the module.
    -   **Constants and Variables**: Any module-specific constants or global variables.
#### Handler.go
-   **Purpose**: This file contains the message handling logic for the module. It defines how different messages are processed and their effects on the module's state.
-   **Components**:
    -   **Message Handlers**: Functions that define how each message type is processed and validated.
    -   **Router Setup**: Integration of the message handlers with the Cosmos SDK's message routing system.
#### Module.go
-   **Purpose**: This file is the main entry point for the module's integration with the Cosmos SDK. It ties together the various components of the module and implements the necessary interfaces.
-   **Components**:
    -   **AppModule Interface Implementation**: Functions that implement the `AppModule` interface, integrating the module with the application lifecycle (e.g., `InitGenesis`, `ExportGenesis`, `Route`, `QuerierRoute`).

## 6. What is different for a Node now?

Archive nodes verifies all the blocks from start, and store their ZKP's. But it is enough if only 1 node do it because then all of the proofs are verifiable and this situation eliminates trust factor.
Full nodes can get recent proof from another node, verify it and store it in the State. In  every validation process of a block, they generate the recent states ZKP, and update it in the state. They dont have to store every seperate blocks ZKP because having one ZKP is knowing that you can verify every block before due to recursion.

Each block's proof is generated recursively, meaning each proof includes a verification of the previous block's proof. This creates a chain of verifiable proofs that link back to the genesis block.
#### Steps:
1.⁠ ⁠*Initial Proof Generation*: For each new block, the archive node generates a proof using o1js.
2.⁠ ⁠*Recursive Proofs*: Each new proof includes a hash of the previous block's proof, ensuring a continuous chain of trust.
3.⁠ ⁠*Proof Storage*: Proofs are stored alongside block data in the blockchain.