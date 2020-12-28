package com.proto.api;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler",
    comments = "Source: api/student_api.proto")
public final class StudentSrvGrpc {

  private StudentSrvGrpc() {}

  public static final String SERVICE_NAME = "api.StudentSrv";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.proto.model.Students.Student,
      com.proto.Common.Result> METHOD_NEW_STUDENT =
      io.grpc.MethodDescriptor.<com.proto.model.Students.Student, com.proto.Common.Result>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "api.StudentSrv", "NewStudent"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.proto.model.Students.Student.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.proto.Common.Result.getDefaultInstance()))
          .build();
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<com.proto.api.StudentApi.QueryStudent,
      com.proto.Common.Result> METHOD_STUDENT_BY_ID =
      io.grpc.MethodDescriptor.<com.proto.api.StudentApi.QueryStudent, com.proto.Common.Result>newBuilder()
          .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
          .setFullMethodName(generateFullMethodName(
              "api.StudentSrv", "StudentByID"))
          .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.proto.api.StudentApi.QueryStudent.getDefaultInstance()))
          .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
              com.proto.Common.Result.getDefaultInstance()))
          .build();

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static StudentSrvStub newStub(io.grpc.Channel channel) {
    return new StudentSrvStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static StudentSrvBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new StudentSrvBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static StudentSrvFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new StudentSrvFutureStub(channel);
  }

  /**
   */
  public static abstract class StudentSrvImplBase implements io.grpc.BindableService {

    /**
     */
    public void newStudent(com.proto.model.Students.Student request,
        io.grpc.stub.StreamObserver<com.proto.Common.Result> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_NEW_STUDENT, responseObserver);
    }

    /**
     */
    public void studentByID(com.proto.api.StudentApi.QueryStudent request,
        io.grpc.stub.StreamObserver<com.proto.Common.Result> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_STUDENT_BY_ID, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_NEW_STUDENT,
            asyncUnaryCall(
              new MethodHandlers<
                com.proto.model.Students.Student,
                com.proto.Common.Result>(
                  this, METHODID_NEW_STUDENT)))
          .addMethod(
            METHOD_STUDENT_BY_ID,
            asyncUnaryCall(
              new MethodHandlers<
                com.proto.api.StudentApi.QueryStudent,
                com.proto.Common.Result>(
                  this, METHODID_STUDENT_BY_ID)))
          .build();
    }
  }

  /**
   */
  public static final class StudentSrvStub extends io.grpc.stub.AbstractStub<StudentSrvStub> {
    private StudentSrvStub(io.grpc.Channel channel) {
      super(channel);
    }

    private StudentSrvStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StudentSrvStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new StudentSrvStub(channel, callOptions);
    }

    /**
     */
    public void newStudent(com.proto.model.Students.Student request,
        io.grpc.stub.StreamObserver<com.proto.Common.Result> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_NEW_STUDENT, getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void studentByID(com.proto.api.StudentApi.QueryStudent request,
        io.grpc.stub.StreamObserver<com.proto.Common.Result> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_STUDENT_BY_ID, getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class StudentSrvBlockingStub extends io.grpc.stub.AbstractStub<StudentSrvBlockingStub> {
    private StudentSrvBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private StudentSrvBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StudentSrvBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new StudentSrvBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.proto.Common.Result newStudent(com.proto.model.Students.Student request) {
      return blockingUnaryCall(
          getChannel(), METHOD_NEW_STUDENT, getCallOptions(), request);
    }

    /**
     */
    public com.proto.Common.Result studentByID(com.proto.api.StudentApi.QueryStudent request) {
      return blockingUnaryCall(
          getChannel(), METHOD_STUDENT_BY_ID, getCallOptions(), request);
    }
  }

  /**
   */
  public static final class StudentSrvFutureStub extends io.grpc.stub.AbstractStub<StudentSrvFutureStub> {
    private StudentSrvFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private StudentSrvFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected StudentSrvFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new StudentSrvFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.proto.Common.Result> newStudent(
        com.proto.model.Students.Student request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_NEW_STUDENT, getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.proto.Common.Result> studentByID(
        com.proto.api.StudentApi.QueryStudent request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_STUDENT_BY_ID, getCallOptions()), request);
    }
  }

  private static final int METHODID_NEW_STUDENT = 0;
  private static final int METHODID_STUDENT_BY_ID = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final StudentSrvImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(StudentSrvImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_NEW_STUDENT:
          serviceImpl.newStudent((com.proto.model.Students.Student) request,
              (io.grpc.stub.StreamObserver<com.proto.Common.Result>) responseObserver);
          break;
        case METHODID_STUDENT_BY_ID:
          serviceImpl.studentByID((com.proto.api.StudentApi.QueryStudent) request,
              (io.grpc.stub.StreamObserver<com.proto.Common.Result>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class StudentSrvDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.proto.api.StudentApi.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (StudentSrvGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new StudentSrvDescriptorSupplier())
              .addMethod(METHOD_NEW_STUDENT)
              .addMethod(METHOD_STUDENT_BY_ID)
              .build();
        }
      }
    }
    return result;
  }
}
