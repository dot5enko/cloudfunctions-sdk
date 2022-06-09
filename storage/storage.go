package storage

var StoreHandler StoreHandlerFunction

type StorageHandlers struct {
	QueryHandler QueryBuilderHandler
	StoreHandler StoreHandlerFunction
}
