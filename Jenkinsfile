pipeline {
  agent {
    docker {
      image 'golang:1.9-alpine'
    }
    
  }
  stages {
    stage('Copy project') {
      steps {
        sh '''if [ ! -d "${HOME}/src" ]; then
    mkdir "${HOME}/src"
fi

if [ -d "${HOME}/src/forum" ]; then
    rm -fr "${HOME}/src/forum"
fi

mkdir "${HOME}/src/forum"
cp -R . ${HOME}/src/forum

ls'''
      }
    }
    stage('Test') {
      steps {
        sh '''cd ${HOME}/src/forum

go test -cover -coverpkg=./... -covermode=count -coverprofile=coverage.out ./tests'''
      }
    }
  }
  environment {
    GOPATH = '/home/'
  }
}