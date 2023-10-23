package com.example.docker.compose.service2.service;

import com.rabbitmq.client.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.nio.charset.StandardCharsets;


@Component
public class RabbitMQConsumer implements ApplicationListener<ApplicationReadyEvent>  {
    private final Logger logger = LoggerFactory.getLogger(RabbitMQConsumer.class);

    @Value("${service2.rabbitmq.listen.topic}")
    private String listenTopic;

    @Value("${service2.rabbitmq.service.name}")
    private String containerName;

    @Autowired
    private RabbitMQService rabbitMQService;

    private final static int MAX_RETRIES = 10;
    private final static long RETRY_INTERVAL = 2000; // 5 seconds

    @Override
    public void onApplicationEvent(ApplicationReadyEvent event) {

        String rabbitMQHost = "";

        try {
            InetAddress inetAddress = InetAddress.getByName(containerName);
            System.out.println("IP Address of " + containerName + ": " + inetAddress.getHostAddress());
            rabbitMQHost = inetAddress.getHostAddress();
        } catch (UnknownHostException e) {
            System.err.println("Error: " + e.getMessage());
        }

        ConnectionFactory factory = new ConnectionFactory();
        factory.setHost(rabbitMQHost);// RabbitMQ server address
        factory.setUsername("guest");
        factory.setPassword("guest");

        int retryCount = 0;
        boolean connected = false;

        while (retryCount < MAX_RETRIES && !connected) {
            try {
                Channel channel;
                Connection connection = factory.newConnection();
                channel = connection.createChannel();
                channel.queueDeclare(listenTopic, false, false, false, null);

                System.out.println("Waiting for messages. To exit, press Ctrl+C");

                DefaultConsumer consumer = new DefaultConsumer(channel) {
                    @Override
                    public void handleDelivery(String consumerTag, Envelope envelope, AMQP.BasicProperties properties, byte[] body) {
                        String message = new String(body, StandardCharsets.UTF_8);
                        System.out.println("Received: " + message);
                        message = message + " MSG";
                        try {
                            rabbitMQService.sendMessage(message);
                        } catch (Exception e) {
                            logger.error("Error Publishing Messages: " + e.getMessage());
                        }
                    }
                };

                channel.basicConsume(listenTopic, false, consumer);
                connected = true;

            } catch (Exception e) {
                retryCount++;
                System.out.println("Connection failed. Retrying in " + (RETRY_INTERVAL / 1000) + " seconds...");
                try {
                    Thread.sleep(RETRY_INTERVAL);
                } catch (InterruptedException ex) {
                    Thread.currentThread().interrupt();
                }
            }
        }
        if (retryCount >= MAX_RETRIES) {
            System.out.println("Failed to establish a connection after " + MAX_RETRIES + " retries.");
        }
    }
}
