pipeline {
  agent {
    docker {
      image 'golang:1.9-alpine'
      args '''-v "$PWD":/usr/src/forum
-w /usr/src/forum'''
    }
    
  }
  stages {
    stage('Test') {
      steps {
        sh '''pwd
ls'''
        ws(dir: '/usr/src/forum') {
          sh 'ls -l'
        }
        
      }
    }
  }
  environment {
    GOPATH = '/usr/'
  }
}