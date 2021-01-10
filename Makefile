.PHONY: example
example:
	cd example && \
	go run ../model --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time
