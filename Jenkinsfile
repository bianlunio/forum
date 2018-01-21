pipeline {
  agent any
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
    GOPATH = '/home/'
  }
}