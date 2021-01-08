#!/bin/bash

Green='\033[0;32m'
echo -e "${Green} Installing all apps under $HOME/install folder"
mkdir ~/install
cd install

#Homebrew
echo -e "${Green} Installing homebrew..."
mkdir homebrew && curl -L https://github.com/Homebrew/brew/tarball/master | tar xz --strip 1 -C homebrew
./homebrew/bin/brew
echo "export PATH=$HOME/install/homebrew/bin:$PATH" >> ~/.bash_profile
source ~/.bash_profile

##############Must haves
#Spectacle
echo -e "${Green} Installing Spectacle..."
curl -o spectacle.zip https://s3.amazonaws.com/spectacle/downloads/Spectacle+1.2.zip
unzip spectacle.zip

#Intellij
curl -o intellij-ultimate.dmg https://download.jetbrains.com/idea/ideaIU-2020.2.3.dmg?_ga=2.211839446.338057530.1605536848-2071424282.1605536848
sudo hdiutil attach intellij-utlimate.dmg
sudo installer -package /Volumes/intellij-ultimate/intellij-ultimate.pkg -target /
sudo hdiutil detach /Volumes/intellij-ultimate

##############Languages
#java
echo -e "${Green} Installing Java..."
brew cask install java11

#gradle
echo -e "${Green} Installing Gradle..."
brew install gradle

#scala
echo -e "${Green} Installing Scala 2.12..."
brwe install scala@2.12

#go
echo -e "${Green} Installing Go..."
brew install go
mkdir -p ~/go-workspace
vi ~/.zshrc
echo "export GOPATH=$HOME/go-workspace" >> ~/.zshrc
echo "export GOROOT=/usr/local/opt/go/libexec" >> ~/.zshrc
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.zshrc
echo "export PATH=$PATH:$GOROOT/bin" >> ~/.zshrc

############Tools
#jq and yq
echo -e "${Green} Installing jq and yq..."
brew install jq
brew install python-yq

#docker
echo -e "${Green} Installing docker..."
brew install docker

#kubectl
echo -e "${Green} Installing kubectl..."
brew install kubectl

#helm
echo -e "${Green} Installing helm..."
brew install helm

#iTerm2
echo -e "${Green} Installing iterm2..."
brew cask install iterm2

#sqlectron
echo -e "${Green} Installing sqlectron..."
brew cask install sqlectron

#ansible
echo -e "${Green} Installing ansible..."
brew install ansible

#awscli
echo -e "${Green} Installing aws cli..."
brew install awscli

echo -e "${Green} Completed!"
