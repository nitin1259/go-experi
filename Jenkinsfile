pipeline{
  agent any
  tools {
    go 'go-1.18'
  }
  environment{
    GO111MODULE='on'
  }

  stages{
    stage("Test"){
      steps{
        git 'https://github.com/kodekloudhub/go-webapp-sample.git'
        sh 'go test ./...'
      }
    }
    stage('Build docker image'){
      steps{
        script{
          app = docker.build("kodekloudhub/go-webapp-sample")
        }
      }
    }
    stage('Build'){
      steps{
        git 'https://github.com/kodekloudhub/go-webapp-sample.git'
        sh 'go build .'
      }
    }
    stage('Run'){
      steps{
        sh 'cd /var/lib/jenkins/workspace/go-full-pipeline && go-webapp-sample &'
      }
    }
    
  }
}