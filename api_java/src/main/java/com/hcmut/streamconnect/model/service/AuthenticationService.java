package com.hcmut.streamconnect.model.service;

import com.hcmut.streamconnect.controller.AuthenticationController.LoginUserDto;
import com.hcmut.streamconnect.model.entity.User;

public interface AuthenticationService {
    User authenticate(LoginUserDto input);
    User register(User user);
}
