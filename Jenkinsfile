pipeline{
    agent any

    tools {go "go"}

    stages{
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