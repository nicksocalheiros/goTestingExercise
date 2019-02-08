#!/usr/bin/env groovy
pipeline{

    agent { docker { image 'golang' } }

    stages{

        stage('Dependencies'){

            steps{
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/github.com/uasouz/goTestingexercise'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/uasouz/goTestingexercise'
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && dep ensure'
            }
        }

        stage('Test'){
    	    parallel{
        		stage('ParallelTest'){
        			steps{
        				sh 'echo "testing"'
        		     }
        		}
        		stage('TestExercise'){
        		
                    steps{
                        sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go test'
                    }
        		}
    	    }
        }

        stage('Build'){
            steps{
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go build -o app main.go '
            }
        }

        stage('Move Results'){
            steps{
                sh 'cp ${GOPATH}/src/github.com/uasouz/goTestingexercise/app .'
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
