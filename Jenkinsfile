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
				sh " sudo systemctl restart shellscript.service"
		        }
	        }
	}
}
