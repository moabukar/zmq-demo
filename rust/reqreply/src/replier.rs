extern crate zmq;

fn main() {
    let context = zmq::Context::new();
    let responder = context.socket(zmq::REP).unwrap();

    responder.bind("tcp://*:5560").unwrap();

    loop {
        let message = responder.recv_msg(0).unwrap();
        println!("Received request: {:?}", message.as_str().unwrap());

        std::thread::sleep(std::time::Duration::from_millis(1000));

        responder.send("World", 0).unwrap();
    }
}
