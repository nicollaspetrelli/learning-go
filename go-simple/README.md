# Repository to learn GoLang




## Installing GO:

wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
sudo tar -xvf go1.17.1.linux-amd64.tar.gz
sudo mv go /usr/local

## Adding GOPATH:
FOR BASH USE .BASHRC file
for ZSH use .zshrc file
then put tree lines of code in your file and save

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH


### Fonts:

Installing GoLang in WSL2
https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374

### Installing in VSCODE:
https://github.com/golang/vscode-go