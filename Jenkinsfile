#!/usr/bin/env groovy
pipeline{
    agent any

    tools { go }

    stages{

        stage('Dependencies'){
            steps{
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/goTestingexercise'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/goTestingexercise'
                sh 'pwd'
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'cd ${GOPATH}/src/goTestingexercise && dep ensure'
            }
        }

        stage('Test'){
            steps{
                sh 'cd ${GOPATH}/src/goTestingexercise && go test'
            }
        }

        stage('Build'){
            steps{
                sh 'cd ${GOPATH}/src/goTestingexercise && go build main.go math.go -o app'
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'app', onlyIfSuccessful: true
            deleteDir()
        }
    }
}