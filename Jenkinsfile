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
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && dep ensure'
            }
        }
        stage('Tests'){
            parallel{
                stage('TestSum'){
                    steps{
                        sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go test -run TestSum'
                    }
                }

                stage('TestHashToBcrypt'){
                    steps{
                        sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go test -run TestHashToBcrypt'
                    }
                }

                stage('TestRemoveNonNumeric'){
                    steps{
                        sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go test -run TestRemoveNonNumeric'
                    }
                }
            }
        }

        stage('Build'){
            steps{
                sh 'cd ${GOPATH}/src/github.com/uasouz/goTestingexercise && go build -o app main.go'
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