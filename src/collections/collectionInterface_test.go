package collections

type collectionTest struct {
	values []interface{}
}

func (collection collectionTest) IsEmpty() bool {
	if len(collection.values) == 0 {
		return true
	}
	return false
}

func (collection collectionTest) Size() uint {

}
