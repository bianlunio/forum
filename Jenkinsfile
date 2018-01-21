pipeline {
  agent {
    docker {
      image 'golang:1.9-alpine'
    }
  }
  stages {
    stage('Copy project') {
      steps {
        sh '''
          if [ ! -d "${HOME}/src" ]; then
            mkdir "${HOME}/src"
          fi

          if [ -d "${HOME}/src/forum" ]; then
            rm -fr "${HOME}/src/forum"
          fi

          mkdir "${HOME}/src/forum"
          cp -R . ${HOME}/src/forum
        '''
      }
    }
  }
  environment {
      GOPATH = '/home/'
  }
}
