package dynamodbsynchro

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type NullTime[T synchro.TimeZone] struct {
	Time  Time[T]
	Valid bool
}

func (c NullTime[T]) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	if !c.Valid {
		return &types.AttributeValueMemberNULL{}, nil
	}
	return c.Time.MarshalDynamoDBAttributeValue()
}

func (c *NullTime[T]) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	if _, ok := av.(*types.AttributeValueMemberNULL); ok {
		c.Valid = false
		return nil
	}
	c.Valid = true
	return c.Time.UnmarshalDynamoDBAttributeValue(av)
}

func (c NullTime[T]) ToSynchro() synchro.NullTime[T] {
	return synchro.NullTime[T]{Time: c.Time.Time, Valid: c.Valid}
}

func NewNullTime[T synchro.TimeZone](t synchro.NullTime[T]) NullTime[T] {
	return NullTime[T]{Time: New(t.Time), Valid: t.Valid}
}

var _ attributevalue.Marshaler = NullTime[tz.UTC]{}
var _ attributevalue.Unmarshaler = (*NullTime[tz.UTC])(nil)
