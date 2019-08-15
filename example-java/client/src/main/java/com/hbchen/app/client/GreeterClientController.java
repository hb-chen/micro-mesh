package com.hbchen.app.client;

import com.hbchen.app.greeter.proto.GreeterGrpc;
import com.hbchen.app.greeter.proto.HelloRequest;
import io.grpc.StatusRuntimeException;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.logging.Logger;

@RestController
public class GreeterClientController {
    private static final Logger logger = Logger.getLogger(GreeterClientController.class.getName());

    @GrpcClient("mm-example")
    private GreeterGrpc.GreeterBlockingStub greeterStub;

    @RequestMapping("/greeting")
    public String printMessage(@RequestParam(defaultValue = "Michael") String name) {
        try {
            HelloRequest request = HelloRequest.newBuilder().setName(name).build();
            return greeterStub.sayHello(request).getMessage();
        } catch (final StatusRuntimeException e) {
            return "FAILED with " + e.getStatus().getCode().name();
        }
    }
}
