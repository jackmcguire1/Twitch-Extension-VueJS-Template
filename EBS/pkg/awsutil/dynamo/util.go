package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func UnmarshalMap(
	fields map[string]*dynamodb.AttributeValue,
	obj interface{},
) (
	err error,
) {
	err = dynamodbattribute.UnmarshalMap(fields, &obj)
	return
}

func UnmarshalAttr(
	field *dynamodb.AttributeValue,
	obj interface{},
) (
	err error,
) {
	err = dynamodbattribute.Unmarshal(field, &obj)
	return
}

func UnmarshalList(
	list []*dynamodb.AttributeValue,
	obj interface{},
) (
	err error,
) {
	err = dynamodbattribute.UnmarshalList(list, &obj)
	return
}

func UnmarshalSliceOfMaps(
	items []map[string]*dynamodb.AttributeValue,
	obj interface{},
) (
	err error,
) {
	err = dynamodbattribute.UnmarshalListOfMaps(items, &obj)
	return
}
