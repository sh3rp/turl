node {
    def root = tool name: 'Go1.11', type: 'go'
    ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/sh3rp/turl") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"
            
            stage 'Checkout'
        
            git url: 'https://github.com/sh3rp/turl.git'
        
            stage 'preTest'
            sh 'go version'
            
            stage 'Test'
            sh 'go vet'
            sh 'go test -cover'
            
            stage 'Build'
            sh 'go build .'
            
            stage 'Deploy'
            // Do nothing.
        }
    }
}
