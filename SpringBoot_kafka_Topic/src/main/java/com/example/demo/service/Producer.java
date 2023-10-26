package com.example.demo.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

import com.example.demo.dto.TransactionDTO;

@Service
public class Producer {

	public static final String topic = "trans";

//  @Autowired 
//  private KafkaTemplate<String, String> kafkaTemp;
//  
//  public void publishToTopic(String message) {
//	  System.out.println("Publishing to topic "+topic);
//	  this.kafkaTemp.send(topic, message);

	@Autowired
	private KafkaTemplate<String, TransactionDTO> kafkaTemplate; // Update the value type to TransactionDTO

	public void publishToTopic(TransactionDTO transaction) {
		System.out.println("Publishing to topic " + topic);
		this.kafkaTemplate.send(topic, transaction);

	}
}
