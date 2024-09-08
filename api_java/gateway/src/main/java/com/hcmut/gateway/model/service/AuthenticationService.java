package com.hcmut.gateway.model.service;

import com.hcmut.gateway.controller.AuthenticationController.LoginUserDto;
import com.hcmut.shared_lib.model.entity.User;

public interface AuthenticationService {
    User authenticate(LoginUserDto input);
    User register(User user);
}
