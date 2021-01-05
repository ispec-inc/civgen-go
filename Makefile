.PHONY: generate_example
generate_example:
	go run ./modelgen --name User --fields ID:int64,Name:string,Email:string,CreatedAt:time.Time,UpdateAt:time.Time --base_dir ./example
