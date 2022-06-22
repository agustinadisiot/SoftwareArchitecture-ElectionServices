module stats_service

go 1.18

replace own_logger => ./../own_logger

replace encrypt => ./../encrypt

replace pipes_and_filters => ./../pipes_and_filters

replace message_queue => ./../message_queue

require message_queue v0.0.0-00010101000000-000000000000

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.9.1
	own_logger v0.0.0-00010101000000-000000000000
	pipes_and_filters v0.0.0-00010101000000-000000000000
)