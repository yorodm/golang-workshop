# Golang y Kafka

## Enviando mensajes con `sarama`{.verbatim}

`sarama`{.verbatim} es una biblioteca desarrollada por Shopify para la
comunicación con Apache Kafka

Para utilizar `sarama`{.verbatim} solo tenemos que importar el módulo

``` go
import  "github.com/Shopify/sarama"
```

### Configuración

``` go
config := sarama.NewConfig()
config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
config.Producer.Return.Successes = true
tlsConfig := createTlsConfiguration()
if tlsConfig != nil {
    config.Net.TLS.Config = tlsConfig // de crypto/tls
    config.Net.TLS.Enable = true
}
```

### Enviando mensajes síncronos

`sarama.SyncProducer`{.verbatim} bloquea la gorutina en espera de
confirmación

``` go
brokerlist := [...]string{"host1:por1", "host2:port2"}
producer := sarama.NewSyncProducer(brokerList, config)
partition, offset, err := s.DataCollector.SendMessage(&sarama.ProducerMessage{
    Key:   sarama.StringEncoder("llave"),
    Topic: "topic",
    Value: sarama.StringEncoder("Mensaje"),
})
producer.Close() //Liberar recursos
```

### Enviando mensajes asíncronos.

`sarama.AsyncProducer`{.verbatim} envía de forma asíncrona.

``` go
brokerlist := [...]string{"host1:por1", "host2:port2"}
producer, err := sarama.NewAsyncProducer(brokerList, config)
producer.Input() <- &sarama.ProducerMessage{
    Key:   sarama.StringEncoder("llave"),
    Topic: "topic",
    Value: sarama.StringEncoder("Mensaje"),,
}
producer.AsyncClose() // Liberar recursos
```

## Recibiendo mensajes con `sarama`{.verbatim}

`sarama`{.verbatim} permite dos modos de recepción de mensajes.

1.  Obtener mensajes desde una sola partición utilizando
    `sarama.Consumer`{.verbatim}
2.  Obtener mensajes de un clúster con `sarama.ConsumerGroup`{.verbatim}

En la carpeta `day-07/21-sarama/consumergroup`{.verbatim} se incluye un
ejemplo de consumidor usando `sarama.ConsumerGroup`{.verbatim}

## Interceptors

Los *interceptors* permiten procesar un mensaje antes de que sea enviado
o recibido.

### Interceptando mensajes enviados

``` go
//
func (m *myinterp) OnSend(p *sarama.ProducerMessage) {
    // se procesa el mensaje
}
// Configuración del productor
conf.Producer.Interceptors = []sarama.ProducerInterceptor{&myinterp{}}
```

### Interceptando mensajes recibidos.

``` go
// Implementar sarama.ConsumerInterceptor
func (m *myinterp) OnConsume(p *sarama.ConsumerMessage) {
    // se procesa el mensaje
}
// Configuración del consumidor
conf.Consumer.Interceptors = []sarama.ConsumerInterceptor{&myinterp{}}
```
