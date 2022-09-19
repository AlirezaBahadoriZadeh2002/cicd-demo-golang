pipeline
{
	agent any
	/*tools
	{
        	go 'go-1.13'
        }*/
	environment 
	{
        	GO111MODULE = 'on'
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
				sh "nohup go run main.go Z>&1 &"
			}
		}

	}
}
