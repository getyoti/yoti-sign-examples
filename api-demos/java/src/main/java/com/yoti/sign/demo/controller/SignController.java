package com.yoti.sign.demo.controller;

import com.yoti.sign.demo.config.Config;
import org.apache.commons.io.FileUtils;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.ContentType;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.ResourceLoader;
import org.springframework.http.MediaType;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;


@Controller
@Configuration
@EnableConfigurationProperties(Config.class)
public class SignController {

    private static final Logger LOG = LoggerFactory.getLogger(SignController.class);
    private Config properties;
    private ResourceLoader resourceLoader;

    @Autowired
    public SignController(Config properties, ResourceLoader resourceLoader) {
        this.properties = properties;
        this.resourceLoader = resourceLoader;
    }

    @RequestMapping(value = "/", method = RequestMethod.GET,
            produces = MediaType.APPLICATION_JSON_VALUE)
    public @ResponseBody String getIndex() {

        try {

            //Read options.json
            File optionsFile = new File(String.valueOf(getClass().getClassLoader().getResource("options.json").getFile()));
            String options = FileUtils.readFileToString(optionsFile, "UTF-8");

            //Read test.pdf
            File file = new File(String.valueOf(getClass().getClassLoader().getResource("test.pdf").getFile()));

            //Build and Post payload
            CloseableHttpClient httpClient = HttpClients.createDefault();
            HttpPost httpPost = new HttpPost(this.properties.getBaseUrl());
            MultipartEntityBuilder builder = MultipartEntityBuilder.create();
            httpPost.setHeader("Authorization","Bearer " + this.properties.getAuthenticationToken());

            builder.addTextBody("options", options, ContentType.TEXT_PLAIN);

            builder.addBinaryBody(
                    "file",
                    new FileInputStream(file),
                    ContentType.APPLICATION_OCTET_STREAM,
                    file.getName()
            );

            HttpEntity multipart = builder.build();
            httpPost.setEntity(multipart);
            CloseableHttpResponse response = httpClient.execute(httpPost);

            return EntityUtils.toString(response.getEntity());

        } catch (IOException e) {
            LOG.error(e.getMessage());
        }
        return "error";
    }
}
