//package org.fi;
//import org.apache.flink.api.common.functions.MapFunction;
//import org.apache.flink.streaming.api.datastream.DataStream;
//import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
//import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;
//import org.springframework.boot.SpringApplication;
//import org.springframework.boot.autoconfigure.SpringBootApplication;
//import org.springframework.kafka.annotation.KafkaListener;
//import org.springframework.kafka.core.KafkaTemplate;
//import org.springframework.stereotype.Service;
//
//@SpringBootApplication
//public class KafkaFlinkSpringBootApplication {
//	private final KafkaTemplate<String, String> kafkaTemplate;
//	public KafkaFlinkSpringBootApplication(KafkaTemplate<String, String> kafkaTemplate) {
//		this.kafkaTemplate = kafkaTemplate;
//	}
//	public static void main(String[] args) {
//		SpringApplication.run(KafkaFlinkSpringBootApplication.class, args);
//	}
//	@Service
//	public static class KafkaConsumerService {
//		private final KafkaTemplate<String, String> kafkaTemplate;
//		public KafkaConsumerService(KafkaTemplate<String, String> kafkaTemplate) {
//			this.kafkaTemplate = kafkaTemplate;
//		}
//		@KafkaListener(topics = "trans", groupId = "mygroup")
//		public void consume(String message) {
//			// Capitalize the received message
//			String capitalizedMessage = message.toUpperCase();
//			// Send the capitalized message to another Kafka topic
//			kafkaTemplate.send("output-kafka-topic", capitalizedMessage);
//			System.out.println("Received and capitalized message: " + capitalizedMessage);
//		}
//	}
//	public static class CapitalizeMapFunction implements MapFunction<String, String> {
//		@Override
//		public String map(String value) throws Exception {
//			return value.toUpperCase();
//		}
//	}
//}

package org.fi;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

@SpringBootApplication
public class KafkaFlinkSpringBootApplication {
    private final KafkaTemplate<String, String> kafkaTemplate;

    public KafkaFlinkSpringBootApplication(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public static void main(String[] args) {
        SpringApplication.run(KafkaFlinkSpringBootApplication.class, args);
    }

    @Service
    public static class KafkaConsumerService {
        private final KafkaTemplate<String, String> kafkaTemplate;

        public KafkaConsumerService(KafkaTemplate<String, String> kafkaTemplate) {
            this.kafkaTemplate = kafkaTemplate;
        }

        @KafkaListener(topics = "trans", groupId = "mygroup")
        public void consume(String message) {
            // Print the received message on the console
            System.out.println("Received message: in flink " + message);
        }
    }
}

