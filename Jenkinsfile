pipeline{
    agent { any }

    stages{
        stage('Test'){
            steps{
                sh 'go test'
            }
        }

        stage('Build'){
            steps{
                sh 'go build main.go math.go'
            }
        }
    }
}