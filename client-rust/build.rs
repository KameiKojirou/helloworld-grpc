// client-rust/build.rs
fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        // turn off server code generation if you only want a client
        .build_server(false)
        // the first slice is your input files, the second is your include path(s)
        .compile_protos(
            &["../proto/helloworld.proto"],
            &["../proto"],
        )?;
    Ok(())
}