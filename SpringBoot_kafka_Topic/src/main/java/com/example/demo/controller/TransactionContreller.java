package com.example.demo.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.example.demo.dto.TransactionDTO;
import com.example.demo.service.Producer;

@RestController
@RequestMapping("/transactions")
public class TransactionContreller {

//	@Autowired 
//	Producer producer;
//	
//	@PostMapping(value="/post")
//	public void sendMessage(@RequestBody List<TransactionDTO> list) {
//		//producer.publishToTopic(list);
//	}
	
	@RestController
	@RequestMapping("/kafkaapp")
	public class KafkaController {

		@Autowired 
		Producer producer;
		
	    public static final String topic = "trans";

	    @Autowired
	    private KafkaTemplate<String, String> kafkaTemplate;

	    @PostMapping("/post")
	    public void publishToTopic(@RequestBody String message) {
	        System.out.println("Publishing to topic " + topic);
	        System.out.println(message);
	        kafkaTemplate.send(topic, message);
	    }
	}
	
	
	
}

