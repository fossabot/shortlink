package kafka

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/mq/v1/query"
)

type Config struct {
	URI           []string // addresses of available kafka brokers
	ConsumerGroup string
}

type Kafka struct { // nolint:decorder
	*Config
	client   sarama.Client
	producer sarama.SyncProducer
	consumer sarama.ConsumerGroup
}

func (mq *Kafka) Init(ctx context.Context) error {
	var err error

	// Set configuration
	config := mq.setConfig()

	if mq.client, err = sarama.NewClient(mq.Config.URI, config); err != nil {
		return err
	}

	// Create new producer
	if mq.producer, err = sarama.NewSyncProducerFromClient(mq.client); err != nil {
		return err
	}

	// Create new consumer
	if mq.consumer, err = sarama.NewConsumerGroupFromClient(mq.ConsumerGroup, mq.client); err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) Close() error {
	var err error

	if mq.client != nil {
		err = mq.client.Close()
	}

	if mq.producer != nil {
		err = mq.producer.Close()
	}

	if mq.consumer != nil {
		err = mq.consumer.Close()
	}

	return err
}

func (k *Kafka) Publish(ctx context.Context, target string, message query.Message) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic:     "shortlink",
		Key:       sarama.StringEncoder(message.Key),
		Value:     sarama.ByteEncoder(message.Payload),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
	})

	return err
}

func (mq *Kafka) Subscribe(target string, message query.Response) error {
	consumer := Consumer{
		ch: message,
	}

	if err := mq.consumer.Consume(context.Background(), []string{"shortlink"}, &consumer); err != nil {
		return err
	}

	return nil
}

func (mq *Kafka) UnSubscribe(target string) error {
	panic("implement me!")
}

// setConfig - Construct a new Sarama configuration.
func (mq *Kafka) setConfig() *sarama.Config {
	viper.AutomaticEnv()
	viper.SetDefault("MQ_KAFKA_URI", "localhost:9092")       // Kafka URI
	viper.SetDefault("MQ_KAFKA_CONSUMER_GROUP", "shortlink") // Kafka consumer group

	mq.Config = &Config{
		URI: []string{
			viper.GetString("MQ_KAFKA_URI"),
		},
		ConsumerGroup: viper.GetString("MQ_KAFKA_CONSUMER_GROUP"),
	}

	// sarama config
	config := sarama.NewConfig()

	config.ClientID = viper.GetString("SERVICE_NAME")

	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Compression = sarama.CompressionSnappy
	config.Version = sarama.V2_5_0_0

	config.Consumer.Return.Errors = true

	return config
}
