# goTestingExercise - Jenkins Basic

##Iniciando Jenkins via Docker

`docker run \
-u root \
-rm \
-p 8080:8080 \
-v jenkins-data:/var/jenkins_home \
-v /var/run/docker.sock:/var/run/docker.sock \
--name jenkins \
jenkinsci/blueocean`


## Instalando Jenkins via bash (Ubuntu)

`
sudo apt update
sudo apt install openjdk-8-jdk

wget -q -O - https://pkg.jenkins.io/debian/jenkins.io.key | sudo apt-key add -

sudo sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'

sudo apt update
sudo apt install jenkins
`
