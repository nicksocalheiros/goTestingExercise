#!/usr/bin/env groovy
pipeline{

    agent { docker { image 'golang' } }

    stages{

        stage('Dependencies'){

            steps{
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github.com/uasouz/goTestingexercise'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/uasouz/goTestingexercise'
                sh 'pwd'
                sh 'go get -u github.com/jstemmer/go-junit-report'
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && dep ensure'
            }
        }

        stage('Test'){
            steps{
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go test -v 2>&1 | go-junit-report > report.xml'
            }
        }

        stage('Build'){
            steps{
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go build -o app main.go '
            }
        }
    }
    post {
        always {
            sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && pwd && ls'
            archiveArtifacts artifacts: '/go/src/github.com/uasouz/goTestingexercise/app', onlyIfSuccessful: true
            junit '/go/src/github.com/uasouz/goTestingexercise/report.xml'
            deleteDir()
        }
    }
}