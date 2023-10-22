extern crate zmq;

fn main() {
    let context = zmq::Context::new();
    let requester = context.socket(zmq::REQ).unwrap();

    requester.connect("tcp://localhost:5560").unwrap();

    for request_nbr in 1..=10 {
        requester.send("Hello", 0).unwrap();

        let reply = requester.recv_msg(0).unwrap();
        println!("Received reply {} of message: {:?}", request_nbr, reply.as_str().unwrap());
    }
}
