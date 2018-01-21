pipeline {
  agent {
    docker {
      image 'go:1.9-alpine'
    }
    
  }
  stages {
    stage('Test') {
      steps {
        sh 'ls'
      }
    }
  }
  environment {
    GOPATH = '/usr/'
  }
}