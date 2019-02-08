#!/usr/bin/env groovy
pipeline{
    agent {
    docker {
     image 'golang'
     customWorkSpace '/go/src/goTestingExercise'
      }
      }

    stages{

        stage('Dependencies'){
            steps{
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'dep init'
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
        }
    }
}