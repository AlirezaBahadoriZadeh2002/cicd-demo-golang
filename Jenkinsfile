pipeline
{
	agent any
	tools
	{
		go 'ttt'
	}
	stages
	{
		stage('Build')
		{
			steps
			{
				sh "echo slalam"
				sh "go build main.go"
			}
		}
		stage('Test')
		{
			steps
			{
				sh "go test"
			}
		}
		stage('Deploy')
		{
			steps
			{
				sh "cp /var/lib/jenkins/workspace/alireza-go/main /var/lib/alireza/"
		        }
	        }
	}
}
