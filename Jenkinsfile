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

        stage('Move Results'){
            steps{
                sh 'cp ${GOPATH}/src/github.com/uasouz/goTestingexercise/app .'
            }
        }

        stage('Go to Production?'){
            steps{
                timeout(time: 30, unit: 'SECONDS') {
                    script{
                        def INPUT_PARAMS = input message: 'Por favor,tome uma acao', ok: 'Next',parameters: [ choice(name: 'TOPROD', choices: ['Sim','Não'].join('\n'), description: 'Devo ir para producao?')],
                        env.TOPROD = INPUT_PARAMS.TOPROD
                        echo "${env.TOPROD}"
                    }
                }
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