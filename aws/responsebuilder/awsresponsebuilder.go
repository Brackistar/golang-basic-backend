package responsebuilder

import "github.com/aws/aws-lambda-go/events"

type AWSResponseBuilder struct {
	statusCode int
	body       string
	headers    map[string]string
}

func NewAWSResponseBuilder() *AWSResponseBuilder {
	return &AWSResponseBuilder{
		headers: make(map[string]string),
	}
}

func (r *AWSResponseBuilder) SetStatusCode(status uint) {
	r.statusCode = int(status)
}

func (r *AWSResponseBuilder) SetBody(body string) {
	r.body = body
}

func (r *AWSResponseBuilder) AddHeader(key string, val string) {
	r.headers[key] = val
}

func (r *AWSResponseBuilder) Clear() {
	r.statusCode = 0
	r.body = ""
	r.headers = make(map[string]string)
}

func (r *AWSResponseBuilder) Build() *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: r.statusCode,
		Body:       r.body,
		Headers:    r.headers,
	}
}
