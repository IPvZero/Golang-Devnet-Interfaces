## Scrapligo & TextFSM
This is a simple example of using the [Scrapligo](https://github.com/scrapli/scrapligo) library developed by [Carl Montanari](https://github.com/carlmontanari/) to retrieve and parse interface information from the Cisco Always-On DevNet Sandbox.

***

### INSTALL GO
First, you must have Golang installed. You can find information on how to do this at [https://go.dev/doc/install](https://go.dev/doc/install).

Alternatively, you can follow these steps to download Go version go1.17.5:

```
cd /usr/local/

sudo wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz

sudo nano $HOME/.profile
```

Add the following line to your $HOME/.profile file:

```export PATH=$PATH:/usr/local/go/bin```

Save the changes and then issue the command:

```source $HOME/.profile```

***


### RUNNING THE SCRIPT
With Golang now installed, change back into your home directory by issuing the command:

```cd $HOME```

Git clone this repository:

```git clone https://github.com/IPvZero/Golang-Devnet-Interfaces.git```

Change into the Golang-Devnet-Interfaces directory:

```cd Golang-Devnet-Interfaces/```

Run with script with:
```go run main.go```

If successful, you can also compile the file:

```go build main.go```

Once compiled to execute the code by simple typing:

```./main.go```

