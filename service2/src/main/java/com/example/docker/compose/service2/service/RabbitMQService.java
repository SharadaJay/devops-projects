package com.example.docker.compose.service2.service;

import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;
import com.rabbitmq.client.Channel;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.net.InetAddress;
import java.net.UnknownHostException;

@Service
public class RabbitMQService {

    private final Logger logger = LoggerFactory.getLogger(RabbitMQConsumer.class);

    @Value("${service2.rabbitmq.publish.topic}")
    private String publishTopic;

    @Value("${service2.rabbitmq.service.name}")
    private String containerName;

    public void sendMessage(String message) throws Exception {

        String rabbitMQHost = "";

        try {
            InetAddress inetAddress = InetAddress.getByName(containerName);
            System.out.println("IP Address of " + containerName + ": " + inetAddress.getHostAddress());
            rabbitMQHost = inetAddress.getHostAddress();
        } catch (UnknownHostException e) {
            System.err.println("Error: " + e.getMessage());
        }

        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost(rabbitMQHost); // RabbitMQ server address
        factory.setUsername("guest");
        factory.setPassword("guest");

        try (Connection connection = factory.newConnection();
             Channel channel = connection.createChannel()) {

            channel.queueDeclare(publishTopic, false, false, false, null);

            channel.basicPublish("", publishTopic, null, message.getBytes());

            System.out.println("Sent HTTP: " + message);
        }
    }
}

