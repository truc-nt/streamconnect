package com.hcmut.streamconnect.account;

import static com.hcmut.streamconnect.util.CollectionUtils.mapToList;

import com.hcmut.streamconnect.model.repository.UserRepository;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import com.hcmut.streamconnect.model.entity.User;

@Service("userDetailsService")
@Transactional
public class CustomUserDetailsService implements UserDetailsService {

    private final UserRepository userRepository;

    @Autowired
    public CustomUserDetailsService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public UserDetails loadUserByUsername(final String username) throws UsernameNotFoundException {

        User user = userRepository.findByUsernameOrEmail(username)
                .orElseThrow(() -> new UsernameNotFoundException("No user found with username: " + username));

        List<GrantedAuthority> grantedAuthorities = mapToList(user.getRoles(), SimpleGrantedAuthority::new);
        return new CustomUserDetails(user.getUsername(), user.getHashedPassword(), true, true, true, true,
                grantedAuthorities, user.getId());
    }
}
