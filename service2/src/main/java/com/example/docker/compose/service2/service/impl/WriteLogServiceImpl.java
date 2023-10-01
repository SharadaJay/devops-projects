package com.example.docker.compose.service2.service.impl;

import com.example.docker.compose.service2.controller.dto.MsgDto;
import com.example.docker.compose.service2.controller.dto.MsgRespDto;
import com.example.docker.compose.service2.service.WriteLogService;
import jakarta.servlet.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.StandardOpenOption;

@Service
public class WriteLogServiceImpl implements WriteLogService {

    @Autowired
    private ApplicationContext context;

    @Value("${service2.log.file.path}")
    private String logFilePath;

    private final Logger logger = LoggerFactory.getLogger(WriteLogServiceImpl.class);

    private static final String STOP_STR = "STOP";


    @Override
    public MsgRespDto writeLogs(MsgDto msgDto, HttpServletRequest request) {

        MsgRespDto msgRespDto = new MsgRespDto();

        if (STOP_STR.equals(msgDto.getMessage())) {
            int exitCode = SpringApplication.exit(context, () -> 0);
            System.exit(exitCode);
        } else {
            writeLogMessage(msgDto.getMessage() + " " + request.getRemoteAddr() + ":" + request.getRemotePort());
            msgRespDto.setMessage("Received Successfully");
        }
        return msgRespDto;
    }

    private void writeLogMessage(String logMsg) {
        try {
            Path logFilePath = Path.of(this.logFilePath);
            Files.writeString(logFilePath, logMsg + "\n", StandardCharsets.UTF_8,
                    StandardOpenOption.WRITE, StandardOpenOption.CREATE, StandardOpenOption.APPEND);
        } catch (IOException e) {
            logger.error("Error writing to log file: " + e.getMessage());
        }
    }
}
