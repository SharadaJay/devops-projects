package com.example.docker.compose.monitor.controller;

import com.example.docker.compose.monitor.service.RabbitMQConsumer;
import com.example.docker.compose.monitor.service.RabbitMQRunLogConsumer;
import jakarta.servlet.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("")
public class BaseController {

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    @Autowired
    private RabbitMQConsumer rabbitMQConsumer;

    @Autowired
    private RabbitMQRunLogConsumer rabbitMQRunLogConsumer;

    @GetMapping(path = "/", produces = MediaType.TEXT_PLAIN_VALUE)
    public String getMessages(
            HttpServletRequest request) throws Exception {
        List <String> values = getListOfValues();
        return String.join("\n", values);
    }

    @GetMapping(path = "/run-log", produces = MediaType.TEXT_PLAIN_VALUE)
    public String getRunLog(
            HttpServletRequest request) throws Exception {
        List <String> values = getListOfValuesRunLog();
        return String.join("\n", values);
    }

    private List<String> getListOfValues() {
        return rabbitMQConsumer.messageList;
    }

    private List<String> getListOfValuesRunLog() {
        return rabbitMQRunLogConsumer.messageList;
    }
}
