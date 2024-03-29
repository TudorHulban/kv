package kvnuts

import (
	"testing"

	"github.com/TudorHulban/kv"
	"github.com/TudorHulban/kv/definition"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory_SetUpdateDelete(t *testing.T) {
	store, errNewStore := NewStoreInMemory(_segmentSizeTests)
	require.NoError(t, errNewStore)

	defer func() {
		assert.NoError(t,
			store.Close())
	}()

	definition.SetUpdateDeleteTest(t, store)
}

func TestDiskSetUpdateDelete(t *testing.T) {
	store, err := NewStore(_segmentSizeTests)
	require.NoError(t, err)

	defer func() {
		assert.NoError(t,
			store.Close())
	}()

	key := []byte("x")
	value := []byte("y")
	bucket := "A"

	assert.NoError(t,
		store.Set(
			bucket,
			kv.KV{
				Key:   key,
				Value: value,
			},
		),
	)

	fetchedValue1, errGetExistentValue := store.GetValueFor(bucket, key)
	assert.NoError(t, errGetExistentValue)
	assert.Equal(t, value, fetchedValue1)

	updateValue := []byte("z")

	assert.NoError(t,
		store.Set(
			bucket,
			kv.KV{
				Key:   key,
				Value: updateValue,
			},
		),
	)

	fetchedValue2, errGetUpdatedValue := store.GetValueFor(bucket, key)

	t.Logf("updated value: %s", updateValue)
	t.Logf("fetched: %s", fetchedValue2)

	assert.NoError(t, errGetUpdatedValue)
	assert.Equal(t, updateValue, fetchedValue2)

	require.NoError(t,
		store.DeleteKVBy(bucket, key))

	fetchedValue3, errGetDeletedKey := store.GetValueFor(bucket, key)
	assert.Error(t, errGetDeletedKey)
	assert.Empty(t, fetchedValue3)
}
