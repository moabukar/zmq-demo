#include <zmq.hpp>
#include <string.h>
#include <stdio.h>
#include <unistd.h>

int main (void)
{
  zmq::context_t context(1);
  const char * protocol =
    "tcp://localhost:5555";
  //  Socket to talk to server
  printf ("Connecting to hello world server...");
  zmq::socket_t sock (context, ZMQ_SUB);
  //  sock.bind("epgm://eth0;239.192.1.1:5556");
  sock.connect(protocol);
  sock.setsockopt (ZMQ_SUBSCRIBE, "", 0);
  printf ("done. \n");
  
  int request_nbr;
  while(true){
    //zmq::message_t request((void*)"Hello", 5, NULL);
    //        zmq::msg_init_size (&request, 5);
    //        memcpy (zmq::msg_data (&request), "Hello", 5);
    //printf ("Sending Hello %d…\n", request_nbr);
    //sock.send ( &request, 0);
    //zmq::msg_close (&request);
    zmq::message_t reply;
    sock.recv (&reply, 0);
    printf ("Received Word %d bytes: \"%s\"\n", reply.size(), reply.data());
  }
  sock.close();
  return 0;
}
