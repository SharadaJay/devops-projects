package com.example.docker.compose.service2.controller;

import com.example.docker.compose.service2.controller.dto.MsgDto;
import com.example.docker.compose.service2.controller.dto.MsgRespDto;
import com.example.docker.compose.service2.service.RabbitMQService;
import jakarta.servlet.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/orchestration/svc2")
public class BaseController {

    private final Logger logger = LoggerFactory.getLogger(this.getClass());

    private final RabbitMQService rabbitMQService;

    public BaseController(RabbitMQService rabbitMQService) {
        this.rabbitMQService = rabbitMQService;
    }

    @PostMapping(path = "/write", produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<MsgRespDto> writeLog(
            @RequestBody
            MsgDto msgDto, HttpServletRequest request) throws Exception {
        String message = msgDto.getMessage() + " " + request.getRemoteAddr() + ":" + request.getRemotePort();
        rabbitMQService.sendMessage(message);
        MsgRespDto msgRespDto = new MsgRespDto();
        return new ResponseEntity<>(msgRespDto, HttpStatus.OK);

    }
}
