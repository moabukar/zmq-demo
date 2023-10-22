extern crate zmq;

fn main() {
    let context = zmq::Context::new();
    let publisher = context.socket(zmq::PUB).unwrap();

    publisher.bind("tcp://*:5561").unwrap();

    loop {
        publisher.send("1 Hello from Rust!", 0).unwrap();
        std::thread::sleep(std::time::Duration::from_secs(1));
    }
}
