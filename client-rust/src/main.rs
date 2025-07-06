use tonic::transport::Channel;
use tonic::Request;

// this module is generated at compile time by build.rs
pub mod helloworld {
    tonic::include_proto!("helloworld");
}

use helloworld::greeter_client::GreeterClient;
use helloworld::HelloRequest;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // connect to our Go server
    let mut client: GreeterClient<Channel> =
        GreeterClient::connect("http://127.0.0.1:50051").await?; 

    // build the request
    let req = Request::new(HelloRequest {
        name: "Rust".into(),
    });

    // perform the RPC
    let res = client.say_hello(req).await?;

    println!("RESPONSE: {}", res.into_inner().message);
    Ok(())
}