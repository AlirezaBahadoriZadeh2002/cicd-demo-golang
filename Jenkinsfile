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
				sh "go run main.go  &"
			}
		}

	}
}
