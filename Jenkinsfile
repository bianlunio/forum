pipeline {
  agent {
    docker {
      image 'go:1.9-alpine'
    }
    
  }
  stages {
    stage('') {
      steps {
        ws(dir: '/usr/src/forum') {
          git(url: 'https://github.com/bianlunio/forum.git', branch: 'develop', changelog: true)
          sh '''pwd
ls -l'''
        }
        
      }
    }
  }
  environment {
    GOPATH = '/usr/'
  }
}