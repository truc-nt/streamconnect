package com.hcmut.streamconnect.model.service;

import com.hcmut.streamconnect.controller.AuthenticationController.LoginUserDto;
import com.hcmut.streamconnect.model.entity.Account;

public interface AuthenticationService {
    Account authenticate(LoginUserDto input);
}
