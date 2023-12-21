pipeline {
    environment {
        ECR = '432086256981.dkr.ecr.us-west-2.amazonaws.com/gusher-api-2'
        CONTAINER = 'gusher-api-2'
        CTX = 'default'
    }
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'DOCKER_CONTEXT=${CTX} docker build -t ${CONTAINER} .'
            }
        }
        stage('Push') {
            steps {
                sh 'aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 432086256981.dkr.ecr.us-west-2.amazonaws.com'
                sh 'DOCKER_CONTEXT=${CTX} docker tag ${CONTAINER} ${ECR}'
                sh 'DOCKER_CONTEXT=${CTX} docker push ${ECR}'
            }
        }
    }

}