package dynamo

import (
	"context"
	"fmt"
	"time"

	"EBS/m/v2/pkg/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var svc = dynamodb.New(awsutil.Session)

const (
	MaxBatchGetInput   = 100
	MaxBatchWriteInput = 25
)

type Table struct {
	Name         string
	compositeKey bool
	ttl          time.Duration
	useRev       bool // ignores revisions if not set
}

// Deprecated in favour of NewTable
func New(name string) *Table {
	return &Table{Name: name}
}

func NewTable(name string, compositeKey, revisioned bool) *Table {
	return &Table{
		Name:         name,
		compositeKey: compositeKey,
		useRev:       revisioned,
	}
}

func NewTableWithTTL(name string, compositeKey bool, ttl time.Duration) *Table {
	return &Table{
		Name:         name,
		compositeKey: compositeKey,
		ttl:          ttl,
	}
}

func (t *Table) PutItem(ctx context.Context, item interface{}) (result *dynamodb.PutItemOutput, err error) {
	it, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return
	}

	result, err = svc.PutItemWithContext(ctx,
		&dynamodb.PutItemInput{
			Item:      it,
			TableName: aws.String(t.Name),
		},
	)

	return
}

func (t *Table) UpdateItem(ctx context.Context, q *dynamodb.UpdateItemInput) (err error) {
	_, err = svc.UpdateItem(q)
	return
}

func (t *Table) Get(ctx context.Context, id string, data interface{}) (err error) {
	result, err := svc.GetItemWithContext(ctx,
		&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": &dynamodb.AttributeValue{
					S: aws.String(id),
				}},
			TableName: aws.String(t.Name),
		},
	)
	if err != nil {
		return
	}

	if len(result.Item) == 0 {
		err = db.ErrNotFound
		return
	}

	err = UnmarshalMap(result.Item, &data)

	return
}

// Maxiumum limit of 100 items.
// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchGetItem.html
func (t *Table) BatchGetItems(
	keys []map[string]*dynamodb.AttributeValue,
) (
	result *dynamodb.BatchGetItemOutput,
	err error,
) {

	if len(keys) > MaxBatchGetInput {
		err = fmt.Errorf("Cannot retrieve more than %d items.", MaxBatchGetInput)
		return
	}

	dynamoItems := map[string]*dynamodb.KeysAndAttributes{
		t.Name: &dynamodb.KeysAndAttributes{
			Keys: keys,
		},
	}

	result, err = svc.BatchGetItem(&dynamodb.BatchGetItemInput{
		RequestItems: dynamoItems,
	})
	if len(result.UnprocessedKeys) > 0 {
		err = fmt.Errorf("Error retrieving %d keys", len(result.UnprocessedKeys))
	}

	return
}

// There cannot be duplicate items in the batch request.
// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html
func (t *Table) BatchPutItems(
	items []interface{},
) (
	result *dynamodb.BatchWriteItemOutput,
	err error,
) {

	if len(items) > MaxBatchWriteInput {
		err = fmt.Errorf("Cannot process more than %d requests.", MaxBatchWriteInput)
		return
	}

	writeReqs := []*dynamodb.WriteRequest{}
	for _, item := range items {
		var dynamoItem map[string]*dynamodb.AttributeValue

		dynamoItem, err = dynamodbattribute.MarshalMap(item)
		if err != nil {
			return
		}

		writeReqs = append(writeReqs, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: dynamoItem,
			},
		})
	}

	dynamoItems := map[string][]*dynamodb.WriteRequest{
		t.Name: writeReqs,
	}

	result, err = svc.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: dynamoItems,
	})

	return
}
