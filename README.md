# install
aws-cli<br>
sam-cli

# deploy

### 環境変数
```bash
export AppName="fix-ip-lambda"
exoprt FixIPEnabled="true"
```

### S3
```bash
aws cloudformation deploy \
		--stack-name "${AppName}-s3" \
		--template-file template-s3.yaml \
        --parameter-overrides \
					BucketName="${AppName}" \
```

### VPC
```bash
aws cloudformation deploy \
		--stack-name "${AppName}-vpc" \
		--template-file template-vpc.yaml \
		--capabilities CAPABILITY_IAM
```

### Lambda
```bash
sam build
    --template-file template-lambda.yaml
	sam deploy \
		--stack-name  "${AppName}-exec" \
		--s3-bucket "${AppName}" \
		--capabilities CAPABILITY_IAM \
		--parameter-overrides \
                    AppName="${AppName}" \
                    URL="${URL}" \
					FixIPEnabled="${FixIPEnabled}"
sam deploy
```