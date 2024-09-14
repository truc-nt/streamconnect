package com.hcmut.gateway.configuration;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class ChannelConfiguration {
//    @Value( "${coreServiceServerHost}" )
//    private String coreServiceServerHost;
//
//    @Value( "${coreServiceServerPort}" )
//    private Integer coreServiceServerPort;
//
//    @Bean
//    @Qualifier("coreServiceClient")
//    public ExternalServiceClient coreServiceClient() {
//        return new ExternalServiceClient(coreServiceServerHost, coreServiceServerPort);
//    }

    @Value("${ecommerceServiceServerHost}")
    private String ecommerceServiceServerHost;

    @Value("${ecommerceServiceServerPort}")
    private Integer ecommerceServiceServerPort;

    @Bean
    @Qualifier("ecommerceServiceClient")
    public ExternalServiceClient ecommerceServiceClient() {
        return new ExternalServiceClient(ecommerceServiceServerHost, ecommerceServiceServerPort);
    }
}
