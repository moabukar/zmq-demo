#include <zmq.hpp>
#include <string>
#include <iostream>
#include <sstream>
#include <unistd.h>

int main () {
  const char * protocol =
    "tcp://*:5555";
    //  Prepare our context and socket
  zmq::context_t context (1);
  zmq::socket_t socket (context, ZMQ_PUB);
  //    socket.setsockopt (ZMQ_RA`TE, &rate, sizeof (rate));
  //socket.connect ("epgm://eth0;239.192.1.1:5556");
  //socket.bind("epgm://eth0;239.192.1.1:5556");
  socket.bind(protocol);
  
  int i = 0;
  while (true) {
    //  Wait for next request from client
    //        socket.recv (&request);
    
    //  Do some 'work'
    usleep (100000);
    
    //  Send reply back to client
    // zmq::message_t reply (5);
    // memcpy ((void *) reply.data (), "World", 5);
    std::stringstream ss;
    ss << "xy.z|Hello: " << i++;
    zmq::message_t request((void*)ss.str().c_str(), ss.str().size()+1, NULL);
    std::cout << "sending: \"" << (const char*)request.data();
    socket.send(request);
    std::cout << "\"... done." << std::endl;
    
  }
  return 0;
}
