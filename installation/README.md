# Fabric 2.0
# September 2021

## Prerequisites
Ubuntu 18 or higher 
Docker version 17.06.2-ce or greater is required.
Go version 1.12.x is required.


## If you are on Windows 
Make sure to install Virtual Box and Vagrant
cd into this directory.
vagrant up


## 1 Open a terminal and give execute permissions
cd installation  <br/>
chmod -R 755 ./*  <br/>

# 2 Install Pre-Requisites & Validate
This will install the minimum version required. 
./install-prereqs.sh  <br/>
source ~/.profile  <br/>
source ~/.bashrc   <br/>

# Logout and Login Again (To reflect the changes to GOPATH)
sudo ./validate-prereqs.sh


# 3 Install the Fabric binaries & images
sudo -E ./install-fabric.sh  <br/>
sudo ./validate-fabric.sh  <br/>

<!-- # 4 Install Hyperledger Explorer tool
./install-explorer.sh
sudo ./validate-explorer.sh -->

# 5 Install the Go Tools
./install-gotools.sh





