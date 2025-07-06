import grpc
import helloworld_pb2
import helloworld_pb2_grpc

def main():
    # 1) connect to server
    channel = grpc.insecure_channel('localhost:50051')
    stub = helloworld_pb2_grpc.GreeterStub(channel)

    # 2) make the call
    req = helloworld_pb2.HelloRequest(name='Python')
    resp = stub.SayHello(req)

    # 3) print the reply
    print('Greeting:', resp.message)

if __name__ == '__main__':
    main()
