package com.hcmut.gateway.configuration;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.hcmut.shared_lib.model.entity.User;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.stereotype.Service;

import java.security.Key;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

@Service
public class JwtService {
    @Value("${security.jwt.secret-key}")
    private String secretKey;

    @Value("${security.jwt.expiration-time}")
    private long jwtExpiration;

    private final ObjectMapper objectMapper = new ObjectMapper();

    public String extractUsername(String token) {
        JwtUserInfo userInfo = extractUserInfo(token);
        return userInfo.getUsername();
    }

    public Long extractUserId(String token) {
        JwtUserInfo userInfo = extractUserInfo(token);
        return userInfo.getUserId();
    }

    private JwtUserInfo extractUserInfo(String token) {
        try {
            return objectMapper.readValue(extractClaim(token, Claims::getSubject), JwtUserInfo.class);
        } catch (JsonProcessingException ex) {
            throw new RuntimeException("Cannot deserialize user info from JSON", ex);
        }
    }

    public <T> T extractClaim(String token, Function<Claims, T> claimsResolver) {
        final Claims claims = extractAllClaims(token);
        return claimsResolver.apply(claims);
    }

    public String generateToken(User user) {
        return buildToken(new HashMap<>(), user.getId(), user.getUsername(), jwtExpiration);
    }

    /*public String generateToken(UserDetails userDetails) {
        return generateToken(new HashMap<>(), userDetails);
    }

    public String generateToken(Map<String, Object> extraClaims, UserDetails userDetails) {
        return buildToken(extraClaims, userDetails.getUsername(), jwtExpiration);
    }*/

    public long getExpirationTime() {
        return jwtExpiration;
    }

    private String buildToken(Map<String, Object> extraClaims, long userId, String userName, long expiration)
    {
        JwtUserInfo userInfo = new JwtUserInfo(userId, userName);
        try {
            return Jwts.builder().setClaims(extraClaims)
                    .setSubject(objectMapper.writeValueAsString(userInfo))
                    .setIssuedAt(new Date(System.currentTimeMillis()))
                    .setExpiration(new Date(System.currentTimeMillis() + expiration))
                    .signWith(getSignInKey(), SignatureAlgorithm.HS256)
                    .compact();
        } catch (JsonProcessingException ex) {
            throw new RuntimeException("Cannot serialize user info to JSON", ex);
        }
    }

    public boolean isTokenValid(String token, UserDetails userDetails) {
        final String username = extractUsername(token);
        return (username.equals(userDetails.getUsername())) && !isTokenExpired(token);
    }

    private boolean isTokenExpired(String token) {
        return extractExpiration(token).before(new Date());
    }

    private Date extractExpiration(String token) {
        return extractClaim(token, Claims::getExpiration);
    }

    private Claims extractAllClaims(String token) {
        return Jwts.parserBuilder().setSigningKey(getSignInKey()).build().parseClaimsJws(token).getBody();
    }

    private Key getSignInKey() {
        byte[] keyBytes = Decoders.BASE64.decode(secretKey);
        return Keys.hmacShaKeyFor(keyBytes);
    }

    public static class JwtUserInfo {
        @JsonProperty("userId")
        private long userId;

        @JsonProperty("username")
        private String username;

        public JwtUserInfo() {}

        public JwtUserInfo(long userId, String username) {
            this.userId = userId;
            this.username = username;
        }

        public long getUserId() {
            return userId;
        }

        public String getUsername() {
            return username;
        }

        public void setUserId(long userId) {
            this.userId = userId;
        }

        public void setUsername(String username) {
            this.username = username;
        }
    }
}
