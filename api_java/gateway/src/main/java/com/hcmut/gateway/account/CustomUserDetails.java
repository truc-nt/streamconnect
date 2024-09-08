package com.hcmut.gateway.account;

import java.util.Collection;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.User;

public class CustomUserDetails extends User {
    private final long userId;

    public CustomUserDetails(String username, String password, boolean enabled, boolean accountNonExpired,
            boolean credentialsNonExpired, boolean accountNonLocked, Collection<? extends GrantedAuthority> authorities,
            long userId) {
        super(username, password, enabled, accountNonExpired, credentialsNonExpired, accountNonLocked, authorities);
        this.userId = userId;
    }

    public CustomUserDetails(String username, String password, Collection<? extends GrantedAuthority> authorities, long userId) {
        super(username, password, authorities);
        this.userId = userId;
    }

    public long getUserId() {
        return userId;
    }
}
