const EventEmitter = require('events');


function main() {
    const eventEmitter = new EventEmitter();

    eventEmitter.on('example1', (msg) => {
        console.log("Received message  on `example1` event", msg);
    });

    eventEmitter.on('example2', (msg) => {
        console.log("Received message  on `example2` event", msg);
    });

    eventEmitter.emit('example1', 'send a');
   
    eventEmitter.emit('example2', 'send b');
    eventEmitter.emit('example1', 'send c');
}

main()




