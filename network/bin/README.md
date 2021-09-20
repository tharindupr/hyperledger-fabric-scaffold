#### Fabric 2.0

Creating Crypto Assets, Genesis Block, Channel Configuration and Anchor Peers
-----------------------------------------------------------------------------
```sh
cd network/bin
chmod -R 755 ./*
./create_network.sh 
```

Starting the network
--------------------
This will start the ../devenv/docker-compose.base.yaml docker compose enviroment. 
```sh
./init_dev_env.sh      
```

Creating the channel and join Peers to the Channel
--------------------------------------------------
```sh
./create_channel.sh 
```



# Deploying the Chaincodes 

## 6. Getting dependecies and packaging the chain code 
./package_chaincode.sh [path_to_src] [chaincode_name] [version]

## 7. Installing chaincode on two orgs
./install_chaincode.sh [path_to_src] [chaincode_name] [version]

## 8. Aproving the chaincode on two orgs
./approve_chaincode.sh [path_to_src] [chaincode_name] [version]

## 9. Committing the chaincode on two orgs
./commit_chaincode.sh [path_to_src] [chaincode_name] [version]

### Learn more on the Chaincode lifecyle : <a href="https://hyperledger-fabric.readthedocs.io/en/release-2.0/chaincode_lifecycle.html#fabric-chaincode-lifecycle"> Here </a>

# Invoking the Chaincodes 

## 10. Invoking the init function of the installed chaincode. 
./invoke_init.sh [chaincodename]


## 11. Using Invoke function of the installed chaincode. 
./invoke_chaincode.sh [chaincodename] <br>
Make sure you set the the necesory arguments required



# Quick Network Dev Env

## Restart the HLF network
./restart_dev_env.sh <br>

## Deploy Subject Contract
./deploy_chaincode_assetmanagementcontract.sh [version]

This will start all the containers + create the channel



# Start client application

## Change Directory to api (If you are in bin)
cd ../../api

## Start the node API
node app.js


