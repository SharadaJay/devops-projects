package com.example.docker.compose.service2.service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;

@Component
public class LogfileTruncator implements ApplicationListener<ApplicationReadyEvent> {

    private final Logger logger = LoggerFactory.getLogger(LogfileTruncator.class);

    @Value("${service2.log.file.path}")
    private String logFilePath;

    @Override
    public void onApplicationEvent(ApplicationReadyEvent event) {
        try {
            Path logPath = Path.of(logFilePath);
            if (Files.exists(logPath)) {
                Files.write(logPath, new byte[0], StandardOpenOption.TRUNCATE_EXISTING);
                logger.info("Log file truncated: " + logFilePath);
            } else {
                Files.createFile(logPath);
                logger.info("Log file created: " + logFilePath);
            }
        } catch (IOException e) {
            logger.error("Error truncating log file: " + e.getMessage());
        }
    }
}
