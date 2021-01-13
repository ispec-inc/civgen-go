.PHONY: model_example
model_example:
	cd example && \
	go run ../model --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time

.PHONY: mockio_example
mockio_example:
	go run ./mockio --source example/pkg/domain/repository/user.go --destination example/pkg/domain/repository/mockio/user.go
