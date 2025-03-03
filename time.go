package dynamodbsynchro

import (
	"errors"
	"time"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Time[T synchro.TimeZone] struct {
	synchro.Time[T]
}

func (c Time[T]) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: c.Format(time.RFC3339Nano),
	}, nil
}

func (c *Time[T]) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	if av, ok := av.(*types.AttributeValueMemberS); ok {
		t, err := synchro.Parse[T](time.RFC3339Nano, av.Value)
		if err != nil {
			return err
		}
		c.Time = t
		return nil
	}
	return errors.New("attribute value is not a string")
}

func New[T synchro.TimeZone](t synchro.Time[T]) Time[T] {
	return Time[T]{Time: t}
}

var _ attributevalue.Marshaler = Time[tz.UTC]{}
var _ attributevalue.Unmarshaler = (*Time[tz.UTC])(nil)
