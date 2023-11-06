package com.macro.mall;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.web.server.LocalServerPort;
import org.springframework.http.ResponseEntity;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class MallGatewayApplicationTests {

    @LocalServerPort
    private int port;

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    public void performanceTestForApiEndpoint() {
        String url = "http://localhost:" + port + "/mall-gateway"; // replace with your actual endpoint

        long startTime = System.currentTimeMillis();

        ResponseEntity<String> response = restTemplate.getForEntity(url, String.class);

        long elapsedTime = System.currentTimeMillis() - startTime;

        assertThat(response.getStatusCodeValue()).isEqualTo(200);
        System.out.println("API call took: " + elapsedTime + " ms");
    }

    @Test
    public void contextLoads() {
    }
}
