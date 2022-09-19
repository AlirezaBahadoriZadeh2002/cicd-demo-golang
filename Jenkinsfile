pipeline
{
	agent any
	
	stages
	{
		stage('Build')
		{
			steps
			{
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
				sh "nohup go run main.go Z>&1 &"
			}
		}

	}
}
