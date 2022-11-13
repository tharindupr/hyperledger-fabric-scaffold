npm init-y
sudo npm install --only=prod @hyperledger/caliper-cli@0.4.2 --unsafe-perm
sudo npx caliper bind --caliper-bind-sut fabric:2.1 --caliper-bind-args="--unsafe-perm"

make sure node version is 8.X or 10.x


sudo npx caliper launch manager  --caliper-benchconfig scenarios/config-trustupdate.yaml --caliper-networkconfig networks/fabric/network-config5.yaml --caliper-fabric-gateway-enabled 
