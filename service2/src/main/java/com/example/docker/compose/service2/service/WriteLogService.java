package com.example.docker.compose.service2.service;

import com.example.docker.compose.service2.controller.dto.MsgDto;
import com.example.docker.compose.service2.controller.dto.MsgRespDto;
import jakarta.servlet.http.HttpServletRequest;

public interface WriteLogService {

    MsgRespDto writeLogs(MsgDto msgDto, HttpServletRequest request);

}
