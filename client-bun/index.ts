import path from "path";
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";

// adjust path to your proto folder
const PROTO_PATH = path.resolve(__dirname, "../proto/helloworld.proto");

// 1) load the .proto defs
const packageDef = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

// 2) create the gRPC package
const grpcObj = grpc.loadPackageDefinition(packageDef) as any;

// 3) instantiate the client
const client = new grpcObj.helloworld.Greeter(
  "localhost:50051",
  grpc.credentials.createInsecure()
);

// 4) call the RPC
client.sayHello({ name: "Bun" }, (err: grpc.ServiceError | null, res: any) => {
  if (err) {
    console.error("RPC error:", err);
    process.exit(1);
  }
  console.log("Greetings:", res.message);
});