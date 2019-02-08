#!/usr/bin/env groovy
pipeline{
    agent {
        docker {
         image 'golang'
        }
    }

    stages{

        stage('Dependencies'){
            steps{
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/goTestingexercise'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/goTestingexercise'
                sh 'pwd'
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'cd ${GOPATH}/src/goTestingexercise && dep init'
            }
        }

        stage('Test'){
            steps{
                sh 'go test'
            }
        }

        stage('Build'){
            steps{
                sh 'go build main.go math.go -o app'
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