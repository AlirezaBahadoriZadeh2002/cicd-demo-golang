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
				script
				{
                			withEnv(['BUILD_ID=dontkill']) 
					{
                    				sh "nohup go run main.go &"
                			}
				}
		         }
	        }
	}
}
