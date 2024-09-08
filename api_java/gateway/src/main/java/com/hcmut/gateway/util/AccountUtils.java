package com.hcmut.gateway.util;

import com.hcmut.gateway.account.CurrentUserDetails;
import com.hcmut.gateway.account.CustomUserDetails;
import com.hcmut.gateway.account.Role;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class AccountUtils {
    private static final Logger logger = LoggerFactory.getLogger(AccountUtils.class);

    private AccountUtils() { }

    public static CurrentUserDetails getCurrentUserDetails() {
        CurrentUserDetails response = new CurrentUserDetails();
        Object principal = SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        if (principal instanceof UserDetails) {
            if (principal instanceof CustomUserDetails) {
                response.setId(((CustomUserDetails) principal).getUserId());
            } else {
                logger.warn("Using UserDetails");
            }
            String username = ((UserDetails) principal).getUsername();
            response.setUsername(username);
            List<String> auths = ((UserDetails) principal).getAuthorities()
                    .stream()
                    .map(GrantedAuthority::getAuthority)
                    .collect(Collectors.toList());
            List<Role> roles = auths.stream()
                    .map(Role::getRoleById)
                    .filter(Objects::nonNull)
                    .collect(Collectors.toList());
            response.setRoles(roles);
            response.setAuthorities(auths);
        } else {
            logger.warn("Principal should be an instance of PalexyUserDetails or UserDetails");
        }
        return response;
    }
}
