package com.hbchen.app;

import com.hbchen.app.greeter.proto.GreeterGrpc;
import com.hbchen.app.greeter.proto.HelloReply;
import com.hbchen.app.greeter.proto.HelloRequest;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

import java.util.logging.Logger;

/**
 * Server that manages startup/shutdown of a {@code Greeter} server.
 */
@GrpcService
public class GreeterService extends GreeterGrpc.GreeterImplBase {
    private static final Logger logger = Logger.getLogger(GreeterService.class.getName());

    @Override
    public void sayHello(HelloRequest req, StreamObserver<HelloReply> responseObserver) {
        logger.info("Greeter received sayHello name:" + req.getName());
        HelloReply reply = HelloReply.newBuilder().setMessage("Hello " + req.getName()).build();
        responseObserver.onNext(reply);
        responseObserver.onCompleted();
    }
}
