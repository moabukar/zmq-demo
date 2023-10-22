extern crate zmq;

fn main() {
    let context = zmq::Context::new();
    let subscriber = context.socket(zmq::SUB).unwrap();

    subscriber.connect("tcp://localhost:5561").unwrap();
    subscriber.set_subscribe("1".as_bytes()).unwrap();

    loop {
        let msg = subscriber.recv_string(0).unwrap().unwrap();
        println!("Received: {}", msg);
    }
}
