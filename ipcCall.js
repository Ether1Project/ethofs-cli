const net = require('net');
const os = require('os');
const client = net.createConnection({
  path:'ethofs.ipc'
}, () => {
  // 'connect' listener
  console.log('Sending ethoFS Request...');
});
var obj = {};
obj.jsonrpc = "2.0";
obj.method = "ethofs_list"; //"ethofs_register" "ethofs_add" "ethofs_remove" "ethofs_extend" "ethofs_list"
obj.params = [""];  // Enter private key to list hosting contracts
obj.id = 01;
client.write(JSON.stringify(obj),() => {
  // 'connect' listener
  console.log('ethoFS Request Complete');
});
client.on('data', (data) => {
  console.log('ethoFS Response Received: \n'+ data.toString());
  client.end();
});
client.on('end', () => {
  console.log('Disconnecting From ethoFS');
});
