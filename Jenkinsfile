pipeline {
  agent {
    docker {
      image 'golang:1.9-alpine'
    }
    
  }
  stages {
    stage('Test') {
      steps {
        sh 'go test -cover -coverpkg=./... -covermode=count -coverprofile=coverage.out ./tests'
      }
    }
  }
  environment {
    GOPATH = '/usr/'
  }
}