pipeline {
  agent {
    docker {
      image 'golang:1.9-alpine'
    }
    
  }
  stages {
    stage('Test') {
      steps {
        sh '''echo $HOME
cd $HOME
pwd
ls

go version
which go'''
      }
    }
  }
  environment {
    GOPATH = '/usr/'
  }
}