package com.hcmut.shared_lib.common_util;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.http.HttpRequest;
import org.springframework.http.client.ClientHttpResponse;

import java.io.IOException;

public class ExternalRequestUtils {
    public static void handleErrorRequest(HttpRequest request, ClientHttpResponse httpResponse) {
        try {
            ErrorResponse errorResponse = new ObjectMapper().readValue(httpResponse.getBody(), ErrorResponse.class);
            if (httpResponse.getStatusCode().is4xxClientError()) {
                throw new IllegalArgumentException(errorResponse.getError());
            }
            throw new RuntimeException(errorResponse.getError());
        } catch (IOException exception) {
            throw new RuntimeException("Failed to read error response from external service");
        }

    }

    private static class ErrorResponse {
        private String error;

        public ErrorResponse() {}

        public ErrorResponse(String error) {
            this.error = error;
        }

        public String getError() {
            return error;
        }
    }
}
